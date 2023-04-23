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
		Expected             *regexp.Regexp
	}{
		"all provided": {
			GivenPrefix:          "ext",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "abc",
			GivenSeparator:       ".",
			Expected:             regexp.MustCompile(`^ext\.test\.[a-c]{8}$`),
		},
		"minimal provided": {
			GivenPrefix:          "",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       "",
			Expected:             regexp.MustCompile(`^test[1-3]{8}$`),
		},
		"no prefix": {
			GivenPrefix:          "",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       ".",
			Expected:             regexp.MustCompile(`^test\.[1-3]{8}$`),
		},
		"no separator": {
			GivenPrefix:          "ext",
			GivenExternalParty:   "test",
			GivenSuffixRandomSet: "123",
			GivenSeparator:       "",
			Expected:             regexp.MustCompile(`^exttest[1-3]{8}$`),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			result, err := getEmail(tc.GivenPrefix, tc.GivenExternalParty, tc.GivenSuffixRandomSet, tc.GivenSeparator)
			if err != nil {
				t.Fatalf(`getEmail returned an error: %v`, err)
			}
			if !tc.Expected.Match([]byte(result)) {
				t.Fatalf(`getEmail result "%s" does not match with regex "%s"`, result, tc.Expected)
			}
		})
	}
}
