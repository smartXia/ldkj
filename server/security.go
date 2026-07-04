package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var scriptTagPattern = regexp.MustCompile(`(?is)<\s*script[^>]*>.*?<\s*/\s*script\s*>`)
var tagPattern = regexp.MustCompile(`(?s)<[^>]*>`)

func sanitizeText(v string) string {
	v = strings.TrimSpace(v)
	v = scriptTagPattern.ReplaceAllString(v, "")
	v = tagPattern.ReplaceAllString(v, "")
	return html.EscapeString(v)
}

func sanitizeHTML(v string) string {
	v = strings.TrimSpace(v)
	v = scriptTagPattern.ReplaceAllString(v, "")
	return v
}

func validateStatus(status string) string {
	switch status {
	case "draft", "published", "archived":
		return status
	default:
		return "draft"
	}
}

func createToken(username, secret string, ttl time.Duration) string {
	exp := time.Now().Add(ttl).Unix()
	payload := fmt.Sprintf("%s|%d", username, exp)
	sig := sign(payload, secret)
	return base64.RawURLEncoding.EncodeToString([]byte(payload + "|" + sig))
}

func validateToken(token, secret string) bool {
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	parts := strings.Split(string(raw), "|")
	if len(parts) != 3 {
		return false
	}
	exp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || time.Now().Unix() > exp {
		return false
	}
	expected := sign(parts[0]+"|"+parts[1], secret)
	return subtle.ConstantTimeCompare([]byte(expected), []byte(parts[2])) == 1
}

func tokenSubject(token, secret string) string {
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return ""
	}
	parts := strings.Split(string(raw), "|")
	if len(parts) != 3 {
		return ""
	}
	exp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || time.Now().Unix() > exp {
		return ""
	}
	expected := sign(parts[0]+"|"+parts[1], secret)
	if subtle.ConstantTimeCompare([]byte(expected), []byte(parts[2])) != 1 {
		return ""
	}
	return parts[0]
}

func sign(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func bearerToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
}

func sameSecret(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
