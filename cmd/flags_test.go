package cmd

import (
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		wantPort string
	}{
		{"default port", []string{}, "8080"},
		{"custom port", []string{"-port", "9090"}, "9090"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := parseFlags(tt.args)
			if opts.port != tt.wantPort {
				t.Fatalf("expected port %s, got %s", tt.wantPort, opts.port)
			}
		})
	}
}

func validOpts() options {
	return options{
		discordToken:     "tok",
		discordAppID:     "app",
		discordPublicKey: "key",
		resendAPIKey:     "re_xxx",
		fromEmail:        "bot@example.com",
		toEmail:          "user@example.com",
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		opts        options
		errContains string
	}{
		{
			name: "all set",
			opts: validOpts(),
		},
		{
			name:        "missing token",
			opts:        func() options { o := validOpts(); o.discordToken = ""; return o }(),
			errContains: "discord-token",
		},
		{
			name:        "missing app id",
			opts:        func() options { o := validOpts(); o.discordAppID = ""; return o }(),
			errContains: "discord-app-id",
		},
		{
			name:        "missing public key",
			opts:        func() options { o := validOpts(); o.discordPublicKey = ""; return o }(),
			errContains: "discord-public-key",
		},
		{
			name: "gateway skips public key",
			opts: func() options { o := validOpts(); o.discordPublicKey = ""; o.gateway = true; return o }(),
		},
		{
			name:        "missing resend api key",
			opts:        func() options { o := validOpts(); o.resendAPIKey = ""; return o }(),
			errContains: "resend-api-key",
		},
		{
			name:        "missing from email",
			opts:        func() options { o := validOpts(); o.fromEmail = ""; return o }(),
			errContains: "from-email",
		},
		{
			name:        "missing to email",
			opts:        func() options { o := validOpts(); o.toEmail = ""; return o }(),
			errContains: "to-email",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.opts.validate()
			if tt.errContains == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tt.errContains)
			}
			if !strings.Contains(err.Error(), tt.errContains) {
				t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
			}
		})
	}
}
