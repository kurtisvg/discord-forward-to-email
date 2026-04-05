package cmd

import (
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

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		opts    options
		wantErr bool
	}{
		{
			name: "all set",
			opts: options{
				discordToken:     "tok",
				discordAppID:     "app",
				discordPublicKey: "key",
				gmailUser:        "user@gmail.com",
				gmailAppPassword: "pass",
			},
			wantErr: false,
		},
		{"missing token", options{discordAppID: "app", discordPublicKey: "key", gmailUser: "u", gmailAppPassword: "p"}, true},
		{"missing app id", options{discordToken: "tok", discordPublicKey: "key", gmailUser: "u", gmailAppPassword: "p"}, true},
		{"missing public key", options{discordToken: "tok", discordAppID: "app", gmailUser: "u", gmailAppPassword: "p"}, true},
		{"missing gmail user", options{discordToken: "tok", discordAppID: "app", discordPublicKey: "key", gmailAppPassword: "p"}, true},
		{"missing gmail password", options{discordToken: "tok", discordAppID: "app", discordPublicKey: "key", gmailUser: "u"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.opts.validate()
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
