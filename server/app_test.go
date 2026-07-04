package main

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestPublicCasesRejectsInjectedPageSizeAndReturnsPublishedItems(t *testing.T) {
	store := newMemoryStore()
	store.cases[1] = Case{
		ID:          1,
		Title:       "增长案例",
		Slug:        "growth",
		Industry:    "零售",
		Platform:    "抖音",
		Strategy:    "内容",
		CoverURL:    "/oss/case.jpg",
		Summary:     "摘要",
		Content:     "<b>安全正文</b>",
		CoreMetrics: "GMV +20%",
		Status:      "published",
	}
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/public/cases?page=1&page_size=9%20OR%201=1", nil)
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid page_size, got %d: %s", resp.Code, resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/public/cases?page=1&page_size=9&industry=%E9%9B%B6%E5%94%AE", nil)
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", resp.Code, resp.Body.String())
	}
	if !strings.Contains(resp.Body.String(), `"title":"增长案例"`) {
		t.Fatalf("expected case title in response, got %s", resp.Body.String())
	}
}

func TestSubmitFormValidatesRequiredFieldsAndSanitizesXSS(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/public/forms", strings.NewReader(`{"name":"","phone":"123","company":"Acme"}`))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected missing name to be rejected, got %d", resp.Code)
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/api/public/forms", strings.NewReader(`{"name":"<script>alert(1)</script>张三","phone":"13800138000","company":"Acme<script>x</script>","source":"home"}`))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected valid form to be created, got %d: %s", resp.Code, resp.Body.String())
	}
	if len(store.forms) != 1 {
		t.Fatalf("expected one form, got %d", len(store.forms))
	}
	form := store.forms[1]
	if strings.Contains(form.Name, "<script") || strings.Contains(form.Company, "<script") {
		t.Fatalf("expected script tags to be sanitized, got %#v", form)
	}
}

func TestAdminAuthCrudCSVAndUpload(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test", UploadDir: t.TempDir()}, store)

	login := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/login", strings.NewReader(`{"username":"admin","password":"secret"}`))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(login, req)
	if login.Code != http.StatusOK {
		t.Fatalf("expected login success, got %d: %s", login.Code, login.Body.String())
	}
	token := extractJSONToken(t, login.Body.String())

	unauth := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/admin/cases", nil)
	app.ServeHTTP(unauth, req)
	if unauth.Code != http.StatusUnauthorized {
		t.Fatalf("expected missing token to be rejected, got %d", unauth.Code)
	}

	create := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/api/admin/cases", strings.NewReader(`{"title":"案例","slug":"case-a","industry":"零售","content":"正文","status":"published"}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(create, req)
	if create.Code != http.StatusCreated {
		t.Fatalf("expected case create success, got %d: %s", create.Code, create.Body.String())
	}

	store.forms[1] = LeadForm{ID: 1, Name: "张三", Phone: "13800138000", Company: "Acme", Status: "new"}
	csvResp := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/admin/forms/export", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	app.ServeHTTP(csvResp, req)
	if csvResp.Code != http.StatusOK || !strings.Contains(csvResp.Body.String(), "id,name,phone,company") {
		t.Fatalf("expected csv export, got %d: %s", csvResp.Code, csvResp.Body.String())
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", "banner.png")
	if err != nil {
		t.Fatal(err)
	}
	pngData := []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00,
	}
	if _, err := part.Write(pngData); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	upload := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/api/admin/upload", &body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	app.ServeHTTP(upload, req)
	if upload.Code != http.StatusCreated {
		t.Fatalf("expected upload success, got %d: %s", upload.Code, upload.Body.String())
	}
	if !strings.Contains(upload.Body.String(), "/oss/") {
		t.Fatalf("expected upload URL under /oss, got %s", upload.Body.String())
	}

	matches, err := filepath.Glob(filepath.Join(app.config.UploadDir, "*", "*.png"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("expected uploaded file on disk, got %d files", len(matches))
	}
}

func extractJSONToken(t *testing.T, body string) string {
	t.Helper()
	const marker = `"token":"`
	start := strings.Index(body, marker)
	if start < 0 {
		t.Fatalf("token missing from response: %s", body)
	}
	start += len(marker)
	end := strings.Index(body[start:], `"`)
	if end < 0 {
		t.Fatalf("unterminated token in response: %s", body)
	}
	return body[start : start+end]
}

func TestMemoryStoreSatisfiesStore(t *testing.T) {
	var _ Store = (*memoryStore)(nil)
	_, _ = context.Background(), t
}
