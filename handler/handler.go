package handler

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/JenswBE/email-prefix-generator/config"
)

//go:embed response.html
var responsePage string

const SuffixLength = 8

type templateData struct {
	Value string
}

func NewHandler(cfg config.Config) (http.HandlerFunc, error) {
	tmpl, err := template.New("response").Parse(responsePage)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response template: %w", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		externalParty := r.URL.Query().Get("p")
		var value string
		if externalParty != "" {
			if value, err = getEmail(cfg.Prefix, externalParty, cfg.SuffixRandomSet, cfg.Separator); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to generate email: %v", err)
				return
			}
		}
		if err := tmpl.Execute(w, templateData{Value: value}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to execute template: %v", err)
		}
	}, nil
}

func getEmail(prefix, externalParty string, suffixRandomSet, separator string) (string, error) {
	// Prepare
	cleanedParty := strings.ReplaceAll(strings.TrimSpace(externalParty), " ", separator)
	output := strings.Builder{}
	output.Grow(len(prefix) + len(cleanedParty) + SuffixLength + 2) // Max 2 separators

	// Add prefix if defined
	if prefix != "" {
		_, _ = output.WriteString(prefix)
		_, _ = output.WriteString(separator)
	}

	// Add external party
	_, _ = output.WriteString(externalParty)
	_, _ = output.WriteString(separator)

	// Add suffix
	randomSetLength := len(suffixRandomSet)
	randomValues := make([]byte, SuffixLength)
	if _, err := rand.Read(randomValues); err != nil {
		return "", fmt.Errorf("failed to generated random value: %w", err)
	}
	for _, randomValue := range randomValues {
		_ = output.WriteByte(suffixRandomSet[int(randomValue)%randomSetLength])
	}
	return output.String(), nil
}
