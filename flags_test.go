package cobracurl

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestRegisterFlags(t *testing.T) {
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)

	RegisterFlags(flagSet)

	tests := []struct {
		name         string
		flagName     string
		expectedType string
	}{
		{"Append flag", "append", "bool"},
		{"Cookie flag", "cookie", "string"},
		{"Compressed flag", "compressed", "bool"},
		{"Data flag", "data", "string"},
		{"Fail flag", "fail", "bool"},
		{"Form flag", "form", "string"},
		{"Head flag", "head", "string"},
		{"Header flag", "header", "string"},
		{"Get flag", "get", "string"},
		{"Include flag", "include", "bool"},
		{"Insecure flag", "insecure", "bool"},
		{"JSON flag", "json", "string"},
		{"Method flag", "method", "string"},
		{"Output flag", "output", "string"},
		{"Location flag", "location", "bool"},
		{"Proxy flag", "proxy", "string"},
		{"Silent flag", "silent", "bool"},
		{"Referer flag", "referer", "string"},
		{"Remote-name flag", "remote-name", "string"},
		{"Upload-file flag", "upload-file", "string"},
		{"URL flag", "url", "string"},
		{"User-agent flag", "user-agent", "string"},
		{"User flag", "user", "string"},
		{"Verbose flag", "verbose", "bool"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := flagSet.Lookup(tt.flagName)
			assert.NotNil(t, flag, "Flag should be registered")
			assert.Equal(t, tt.expectedType, flag.Value.Type(), "Flag type should match")
		})
	}
}
