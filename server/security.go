package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var scriptTagPattern = regexp.MustCompile(`(?is)<\s*script[^>]*>.*?<\s*/\s*script\s*>`)
var dangerousTagPattern = regexp.MustCompile(`(?is)<\s*(script|style|iframe|object|embed|link|meta)[^>]*>.*?<\s*/\s*(script|style|iframe|object|embed|link|meta)\s*>`)
var standaloneDangerousTagPattern = regexp.MustCompile(`(?is)<\s*(script|style|iframe|object|embed|link|meta)[^>]*\/?\s*>`)
var tagPattern = regexp.MustCompile(`(?s)<[^>]*>`)
var htmlTagPattern = regexp.MustCompile(`(?is)<\s*(/?)\s*([a-z0-9]+)([^>]*)>`)
var htmlAttrPattern = regexp.MustCompile(`(?is)([a-z_:][a-z0-9_:.-]*)\s*=\s*("[^"]*"|'[^']*'|[^\s"'=<>` + "`" + `]+)`)

func sanitizeText(v string) string {
	v = strings.TrimSpace(v)
	v = scriptTagPattern.ReplaceAllString(v, "")
	v = tagPattern.ReplaceAllString(v, "")
	return html.EscapeString(v)
}

func sanitizeHTML(v string) string {
	v = strings.TrimSpace(v)
	v = scriptTagPattern.ReplaceAllString(v, "")
	v = dangerousTagPattern.ReplaceAllString(v, "")
	v = standaloneDangerousTagPattern.ReplaceAllString(v, "")
	return htmlTagPattern.ReplaceAllStringFunc(v, sanitizeHTMLTag)
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

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func verifyPassword(password, hash string) bool {
	if strings.HasPrefix(hash, "sha256:") {
		expected := hashPassword(password)
		return sameSecret(expected, hash)
	}
	if strings.HasPrefix(hash, "{plain}") {
		return sameSecret(password, strings.TrimPrefix(hash, "{plain}"))
	}
	return false
}

func sanitizeHTMLTag(tag string) string {
	match := htmlTagPattern.FindStringSubmatch(tag)
	if len(match) != 4 {
		return ""
	}
	closing, name, attrs := match[1], strings.ToLower(match[2]), match[3]
	if !allowedHTMLTag(name) {
		return ""
	}
	if closing != "" {
		return "</" + name + ">"
	}
	cleanAttrs := sanitizeHTMLAttrs(name, attrs)
	if cleanAttrs == "" {
		return "<" + name + ">"
	}
	return "<" + name + " " + cleanAttrs + ">"
}

func allowedHTMLTag(name string) bool {
	switch name {
	case "p", "br", "strong", "b", "em", "i", "u", "s", "ul", "ol", "li", "blockquote", "pre", "code",
		"h1", "h2", "h3", "h4", "h5", "h6", "a", "img", "table", "thead", "tbody", "tr", "th", "td", "hr":
		return true
	default:
		return false
	}
}

func sanitizeHTMLAttrs(tag, attrs string) string {
	out := []string{}
	for _, match := range htmlAttrPattern.FindAllStringSubmatch(attrs, -1) {
		name := strings.ToLower(match[1])
		value := strings.Trim(match[2], `"'`)
		if strings.HasPrefix(name, "on") || name == "style" {
			continue
		}
		if !allowedHTMLAttr(tag, name, value) {
			continue
		}
		out = append(out, name+`="`+html.EscapeString(value)+`"`)
	}
	return strings.Join(out, " ")
}

func allowedHTMLAttr(tag, name, value string) bool {
	switch name {
	case "title", "alt":
		return tag == "a" || tag == "img"
	case "href":
		return tag == "a" && allowedURL(value, false)
	case "target":
		return tag == "a" && (value == "_blank" || value == "_self")
	case "rel":
		return tag == "a"
	case "src":
		return tag == "img" && allowedURL(value, true)
	case "width", "height", "colspan", "rowspan":
		for _, r := range value {
			if r < '0' || r > '9' {
				return false
			}
		}
		return value != ""
	default:
		return false
	}
}

func allowedURL(value string, allowDataImage bool) bool {
	v := strings.TrimSpace(strings.ToLower(value))
	if v == "" || strings.HasPrefix(v, "javascript:") || strings.HasPrefix(v, "vbscript:") {
		return false
	}
	if strings.HasPrefix(v, "/") || strings.HasPrefix(v, "#") || strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://") ||
		strings.HasPrefix(v, "mailto:") || strings.HasPrefix(v, "tel:") {
		return true
	}
	return allowDataImage && strings.HasPrefix(v, "data:image/")
}
