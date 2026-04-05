package cmd

import (
	"flag"
	"os"
)

type options struct {
	port string

	discordToken     string
	discordAppID     string
	discordPublicKey string

	gmailUser        string
	gmailAppPassword string
}

func parseFlags(args []string) options {
	var opts options
	fs := flag.NewFlagSet("discord-forward-to-email", flag.ExitOnError)
	fs.StringVar(&opts.port, "port", envOrDefault("PORT", "8080"), "HTTP server port")
	fs.StringVar(&opts.discordToken, "discord-token", os.Getenv("DISCORD_TOKEN"), "Discord bot token")
	fs.StringVar(&opts.discordAppID, "discord-app-id", os.Getenv("DISCORD_APP_ID"), "Discord application ID")
	fs.StringVar(&opts.discordPublicKey, "discord-public-key", os.Getenv("DISCORD_PUBLIC_KEY"), "Discord public key for signature verification")
	fs.StringVar(&opts.gmailUser, "gmail-user", os.Getenv("GMAIL_USER"), "Gmail address")
	fs.StringVar(&opts.gmailAppPassword, "gmail-app-password", os.Getenv("GMAIL_APP_PASSWORD"), "Gmail app password")
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	return opts
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
