package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func newTestHandler(t *testing.T) (*Handler, ed25519.PrivateKey) {
	t.Helper()
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal(err)
	}
	return &Handler{publicKey: pub}, priv
}

func signRequest(t *testing.T, priv ed25519.PrivateKey, timestamp string, body []byte) (string, string) {
	t.Helper()
	msg := append([]byte(timestamp), body...)
	sig := ed25519.Sign(priv, msg)
	return hex.EncodeToString(sig), timestamp
}

func TestHandleInteraction_Ping(t *testing.T) {
	h, priv := newTestHandler(t)

	body, _ := json.Marshal(discordgo.Interaction{Type: discordgo.InteractionPing})
	sig, ts := signRequest(t, priv, "1234567890", body)

	req := httptest.NewRequest(http.MethodPost, "/interactions", bytes.NewReader(body))
	req.Header.Set("X-Signature-Ed25519", sig)
	req.Header.Set("X-Signature-Timestamp", ts)

	rec := httptest.NewRecorder()
	h.HandleInteraction(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp discordgo.InteractionResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid response JSON: %v", err)
	}
	if resp.Type != discordgo.InteractionResponsePong {
		t.Fatalf("expected pong (type 1), got type %d", resp.Type)
	}
}

func TestHandleInteraction_InvalidSignature(t *testing.T) {
	h, _ := newTestHandler(t)

	body, _ := json.Marshal(discordgo.Interaction{Type: discordgo.InteractionPing})

	req := httptest.NewRequest(http.MethodPost, "/interactions", bytes.NewReader(body))
	req.Header.Set("X-Signature-Ed25519", "bad")
	req.Header.Set("X-Signature-Timestamp", "1234567890")

	rec := httptest.NewRecorder()
	h.HandleInteraction(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestHandleInteraction_MethodNotAllowed(t *testing.T) {
	h, _ := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/interactions", nil)
	rec := httptest.NewRecorder()
	h.HandleInteraction(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}
