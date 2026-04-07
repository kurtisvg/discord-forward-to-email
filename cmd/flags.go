package cmd

import (
	"flag"
	"fmt"
	"os"
)

type options struct {
	version bool
	host    string
	port    string
	gateway bool

	discordToken     string
	discordAppID     string
	discordPublicKey string

	resendAPIKey string
	fromEmail    string
	toEmail      string
}

func parseFlags(args []string) options {
	var opts options
	fs := flag.NewFlagSet("fwd2email", flag.ExitOnError)
	fs.BoolVar(&opts.version, "version", false, "Print version and exit")
	fs.StringVar(&opts.host, "host", envOrDefault("HOST", ""), "HTTP server host")
	fs.StringVar(&opts.port, "port", envOrDefault("PORT", "8080"), "HTTP server port")
	fs.BoolVar(&opts.gateway, "gateway", false, "Use gateway (websocket) mode instead of webhook HTTP server")
	fs.StringVar(&opts.discordToken, "discord-token", os.Getenv("DISCORD_TOKEN"), "Discord bot token")
	fs.StringVar(&opts.discordAppID, "discord-app-id", os.Getenv("DISCORD_APP_ID"), "Discord application ID")
	fs.StringVar(&opts.discordPublicKey, "discord-public-key", os.Getenv("DISCORD_PUBLIC_KEY"), "Discord public key for signature verification")
	fs.StringVar(&opts.resendAPIKey, "resend-api-key", os.Getenv("RESEND_API_KEY"), "Resend API key")
	fs.StringVar(&opts.fromEmail, "from-email", os.Getenv("FROM_EMAIL"), "Sender email address")
	fs.StringVar(&opts.toEmail, "to-email", os.Getenv("TO_EMAIL"), "Recipient email address")
	_ = fs.Parse(args)
	return opts
}

func (o options) validate() error {
	if o.discordToken == "" {
		return fmt.Errorf("required config is not set: discord-token")
	}
	if o.discordAppID == "" {
		return fmt.Errorf("required config is not set: discord-app-id")
	}
	if !o.gateway && o.discordPublicKey == "" {
		return fmt.Errorf("required config is not set: discord-public-key")
	}
	if o.resendAPIKey == "" {
		return fmt.Errorf("required config is not set: resend-api-key")
	}
	if o.fromEmail == "" {
		return fmt.Errorf("required config is not set: from-email")
	}
	if o.toEmail == "" {
		return fmt.Errorf("required config is not set: to-email")
	}
	return nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
