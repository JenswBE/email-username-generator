package handler

import (
	"regexp"
	"testing"
)

func Test_getEmail(t *testing.T) {
	testCases := map[string]struct {
		GivenPrefix          string
		GivenExternalParty   string
		GivenSuffixRandomSet string
		GivenSeparator       string
		GivenDomain          string
		Expected             *regexp.Regexp
	}{
		"all provided": {
			GivenPrefix:          "ext",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "abc",
			GivenSeparator:       ".",
			GivenDomain:          "example.com",
			Expected:             regexp.MustCompile(`^ext\.test\.[a-c]{8}@example.com$`),
		},
		"minimal provided": {
			GivenPrefix:          "",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       "",
			GivenDomain:          "",
			Expected:             regexp.MustCompile(`^test[1-3]{8}$`),
		},
		"no prefix": {
			GivenPrefix:          "",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       ".",
			GivenDomain:          "example.com",
			Expected:             regexp.MustCompile(`^test\.[1-3]{8}@example.com$`),
		},
		"no separator": {
			GivenPrefix:          "ext",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       "",
			GivenDomain:          "example.com",
			Expected:             regexp.MustCompile(`^exttest[1-3]{8}@example.com$`),
		},
		"no domain": {
			GivenPrefix:          "ext",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       ".",
			GivenDomain:          "",
			Expected:             regexp.MustCompile(`^ext.test.[1-3]{8}$`),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			result, err := getEmail(tc.GivenPrefix, tc.GivenExternalParty, tc.GivenSuffixRandomSet, tc.GivenSeparator, tc.GivenDomain)
			if err != nil {
				t.Fatalf(`getEmail returned an error: %v`, err)
			}
			if !tc.Expected.Match([]byte(result)) {
				t.Fatalf(`getEmail result "%s" does not match with regex "%s"`, result, tc.Expected)
			}
		})
	}
}
