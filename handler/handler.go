package handler

import (
	_ "embed"
	"fmt"
	"html/template"
	"math/rand"
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
			value = getEmail(cfg.Prefix, externalParty, cfg.SuffixRandomSet, cfg.Separator)
		}
		tmpl.Execute(w, templateData{Value: value})
	}, nil
}

func getEmail(prefix, externalParty string, suffixRandomSet, separator string) string {
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
	// Note: This uses an insecure random source. This is by design as suffix is not security sensitive.
	ranomSetLength := len(suffixRandomSet)
	for i := 0; i < SuffixLength; i++ {
		_ = output.WriteByte(suffixRandomSet[rand.Intn(ranomSetLength)])
	}
	return output.String()
}
