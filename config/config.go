package config

import (
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
		// Env var is explicitly set to an empty string
		return Config{}, EnvVarIsEmptyError{envVar: "EPG_SUFFIX_RANDOM_SET"}
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
