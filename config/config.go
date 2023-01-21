package config

import (
	"errors"
	"os"
)

type Config struct {
	// Prefix for the email address
	Prefix string
	// Separator between prefix, external party and random suffix
	Separator string
	// Set of characters to use for the random suffix
	SuffixRandomSet string
}

func GetConfig() (Config, error) {
	suffixRandomSet := LookupEnvWithFallback("EPG_SUFFIX_RANDOM_SET", "abcdefghijklmnopqrstuvwxyz0123456789")
	if suffixRandomSet == "" {
		// Env var is explicetly set to an empty string
		return Config{}, errors.New("EPG_SUFFIX_RANDOM_SET cannot be an empty string")
	}

	return Config{
		Prefix:          LookupEnvWithFallback("EPG_PREFIX", "ext"),
		Separator:       LookupEnvWithFallback("EPG_SEPARATOR", "."),
		SuffixRandomSet: suffixRandomSet,
	}, nil
}

func LookupEnvWithFallback(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
