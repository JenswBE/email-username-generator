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
	Value         string
	ExternalParty string
	Domain        string
}

func NewHandler(cfg config.Config) (http.HandlerFunc, error) {
	tmpl, err := template.New("response").Parse(responsePage)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response template: %w", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		externalParty := r.URL.Query().Get("p")
		domain := r.URL.Query().Get("d")
		var value string
		if externalParty != "" {
			if value, err = getEmail(cfg.Prefix, externalParty, cfg.SuffixRandomSet, cfg.Separator, domain); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to generate email: %v", err)
				return
			}
		}
		data := templateData{
			Value:         value,
			ExternalParty: externalParty,
			Domain:        domain,
		}
		if err := tmpl.Execute(w, data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to execute template: %v", err)
		}
	}, nil
}

func getEmail(prefix, externalParty string, suffixRandomSet, separator, domain string) (string, error) {
	// Prepare
	cleanedParty := strings.ReplaceAll(strings.TrimSpace(externalParty), " ", separator)
	output := strings.Builder{}
	output.Grow(len(prefix) + len(cleanedParty) + SuffixLength + 2 + 1 + len(domain)) // Max 2 separators + @ before domain

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

	// Add domain
	if domain != "" {
		output.WriteRune('@')
		output.WriteString(domain)
	}
	return output.String(), nil
}
