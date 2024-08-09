package auth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input     map[string][]string
		wantKey   string
		wantError error
	}{
		"simple": {input: http.Header{"Authorization": {"ApiKey your-api-key"}}, wantKey: "your-api-key", wantError: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil {
				fmt.Println(err)
			}
			diff := cmp.Diff(tc.wantKey, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
