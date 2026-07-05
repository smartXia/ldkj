package main

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"
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

func TestPublicAboutReturnsPublishedStaticPage(t *testing.T) {
	store := newMemoryStore()
	store.pages[1] = StaticPage{
		ID:        1,
		PageKey:   "about",
		Title:     "关于我们",
		ExtraData: `{"about":{"hero":{"title":"可编辑关于我们"}}}`,
		Status:    "published",
	}
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/public/about", nil)
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", resp.Code, resp.Body.String())
	}
	if !strings.Contains(resp.Body.String(), `"page_key":"about"`) {
		t.Fatalf("expected about page key, got %s", resp.Body.String())
	}
	if !strings.Contains(resp.Body.String(), `可编辑关于我们`) {
		t.Fatalf("expected editable about payload, got %s", resp.Body.String())
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

func TestPublicServicesPartnerAndAdminDashboard(t *testing.T) {
	store := newMemoryStore()
	store.services[1] = Service{
		ID:         1,
		Title:      "Social Service",
		Slug:       "social-service",
		Summary:    "service summary",
		CoverURL:   "/oss/service.png",
		Status:     "published",
		Highlights: []string{"strategy", "content"},
		Process:    []string{"diagnose", "deliver"},
	}
	store.faqs[1] = PartnerFAQ{ID: 1, Question: "How soon?", Answer: "Within 24 hours.", IsPublished: true}
	store.logs = append(store.logs, OperationLog{ID: 1, Username: "admin", Action: "login", Module: "auth", CreatedAt: nowString()})
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/services", nil)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"slug":"social-service"`) {
		t.Fatalf("expected public services, got %d: %s", resp.Code, resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/partner", nil)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"faqs"`) {
		t.Fatalf("expected partner FAQ payload, got %d: %s", resp.Code, resp.Body.String())
	}

	token := createToken("admin", "test", time.Hour)
	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/admin/dashboard", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"services":1`) {
		t.Fatalf("expected dashboard stats, got %d: %s", resp.Code, resp.Body.String())
	}
}

func TestAdminAliasesAndSettingsRoutes(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)
	token := createToken("admin", "test", time.Hour)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/cases", strings.NewReader(`{"title":"Alias Case","slug":"alias-case","playbook":"seed","cover":"/oss/case.png","metrics":"GMV +20%","content":"body","status":"published"}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected alias case create success, got %d: %s", resp.Code, resp.Body.String())
	}
	created := store.cases[1]
	if created.Strategy != "seed" || created.CoverURL != "/oss/case.png" || created.CoreMetrics != "GMV +20%" {
		t.Fatalf("expected case aliases to be normalized, got %#v", created)
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/admin/settings/seo", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"items"`) {
		t.Fatalf("expected SEO collection, got %d: %s", resp.Code, resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPut, "/api/admin/settings/site", strings.NewReader(`{"site_title":"Site","logo_url":"/logo.png","contact":"Contact","copyright":"Copyright","footer_links":"[]"}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || store.site.SiteTitle != "Site" {
		t.Fatalf("expected site settings update, got %d: %s", resp.Code, resp.Body.String())
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

func TestSanitizeHTMLRemovesDangerousMarkup(t *testing.T) {
	input := `<p onclick="alert(1)">ok</p><img src="/oss/a.png" onerror="alert(1)"><a href="javascript:alert(1)">bad</a><script>alert(1)</script>`
	clean := sanitizeHTML(input)

	for _, forbidden := range []string{"onclick", "onerror", "javascript:", "<script"} {
		if strings.Contains(strings.ToLower(clean), forbidden) {
			t.Fatalf("expected dangerous HTML to be removed, got %s", clean)
		}
	}
	if !strings.Contains(clean, `<p>ok</p>`) || !strings.Contains(clean, `<img src="/oss/a.png">`) {
		t.Fatalf("expected safe tags to be preserved, got %s", clean)
	}
}

func TestAdminLoginUsesStoreUserRolesAndPermissions(t *testing.T) {
	store := newMemoryStore()
	store.admins["editor"] = AdminUser{
		ID:           7,
		Username:     "editor",
		PasswordHash: hashPassword("secret"),
		Status:       "active",
		Roles:        []string{"editor"},
		Permissions:  []string{"content:read", "content:write"},
	}
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "env-secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/login", strings.NewReader(`{"username":"editor","password":"secret"}`))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected store user login success, got %d: %s", resp.Code, resp.Body.String())
	}
	if !strings.Contains(resp.Body.String(), `"roles":["editor"]`) || !strings.Contains(resp.Body.String(), `"permissions":["content:read","content:write"]`) {
		t.Fatalf("expected roles and permissions in login response, got %s", resp.Body.String())
	}
}

func TestAdminBannersSupportMultipleItems(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)
	token := createToken("admin", "test", time.Hour)

	for _, body := range []string{
		`{"title":"Hero A","image_url":"/oss/a.png","status":"enabled","sort":2}`,
		`{"title":"Hero B","image_url":"/oss/b.png","status":"enabled","sort":1}`,
	} {
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/admin/banners", strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		app.ServeHTTP(resp, req)
		if resp.Code != http.StatusCreated {
			t.Fatalf("expected banner create success, got %d: %s", resp.Code, resp.Body.String())
		}
	}

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/admin/banners", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected banner list success, got %d: %s", resp.Code, resp.Body.String())
	}
	if !strings.Contains(resp.Body.String(), `"total":2`) || !strings.Contains(resp.Body.String(), `"title":"Hero B"`) {
		t.Fatalf("expected multiple banners sorted by sort order, got %s", resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPut, "/api/admin/banners/2", strings.NewReader(`{"title":"Hero B Updated","image_url":"/oss/b2.png","status":"enabled","sort":3}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"title":"Hero B Updated"`) {
		t.Fatalf("expected banner update success, got %d: %s", resp.Code, resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPut, "/api/admin/banners/2", strings.NewReader(`{"title":"Hero B Updated","image_url":"/oss/old.png","image":"/oss/new-upload.png","status":"enabled","sort":3}`))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || !strings.Contains(resp.Body.String(), `"image_url":"/oss/new-upload.png"`) {
		t.Fatalf("expected banner image alias to update image_url, got %d: %s", resp.Code, resp.Body.String())
	}

	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodDelete, "/api/admin/banners/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	app.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK || len(store.banners) != 1 {
		t.Fatalf("expected banner delete to remove item, got %d: %s", resp.Code, resp.Body.String())
	}
}

func TestUploadPersistsMediaAsset(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test", UploadDir: t.TempDir()}, store)
	token := createToken("admin", "test", time.Hour)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	_ = writer.WriteField("type", "banner")
	part, err := writer.CreateFormFile("file", "banner.png")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = part.Write([]byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00,
	})
	_ = writer.Close()

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/upload", &body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected upload success, got %d: %s", resp.Code, resp.Body.String())
	}
	if len(store.media) != 1 || store.media[1].BizType != "banner" || store.media[1].OriginalName != "banner.png" {
		t.Fatalf("expected media asset to be persisted, got %#v", store.media)
	}
}

func TestBannerUploadAllowsMP4OnlyForBanner(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test", UploadDir: t.TempDir()}, store)
	token := createToken("admin", "test", time.Hour)
	mp4Data := []byte{
		0x00, 0x00, 0x00, 0x18,
		'f', 't', 'y', 'p',
		'i', 's', 'o', 'm',
		0x00, 0x00, 0x02, 0x00,
		'i', 's', 'o', 'm',
		'i', 's', 'o', '2',
	}

	uploadMP4 := func(uploadType string) *httptest.ResponseRecorder {
		var body bytes.Buffer
		writer := multipart.NewWriter(&body)
		_ = writer.WriteField("type", uploadType)
		part, err := writer.CreateFormFile("file", "hero.mp4")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = part.Write(mp4Data)
		_ = writer.Close()

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/admin/upload", &body)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		app.ServeHTTP(resp, req)
		return resp
	}

	bannerResp := uploadMP4("banner")
	if bannerResp.Code != http.StatusCreated || !strings.Contains(bannerResp.Body.String(), ".mp4") {
		t.Fatalf("expected banner mp4 upload success, got %d: %s", bannerResp.Code, bannerResp.Body.String())
	}

	newsResp := uploadMP4("news")
	if newsResp.Code != http.StatusBadRequest {
		t.Fatalf("expected non-banner mp4 upload to be rejected, got %d: %s", newsResp.Code, newsResp.Body.String())
	}
}

func TestFormSubmissionCreatesEmailNotificationRecordWhenEnabled(t *testing.T) {
	store := newMemoryStore()
	store.emailSettings.IsEnabled = true
	store.emailSettings.Recipients = "ops@example.com"
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/public/forms", strings.NewReader(`{"name":"Zhang San","phone":"13800138000","company":"Acme","requirement":"Need service"}`))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected form create success, got %d: %s", resp.Code, resp.Body.String())
	}
	if len(store.notifications) != 1 {
		t.Fatalf("expected one notification record, got %#v", store.notifications)
	}
	if store.notifications[1].LeadFormID != 1 || store.notifications[1].Status == "" {
		t.Fatalf("expected notification to reference form and have status, got %#v", store.notifications[1])
	}
}

func TestCaseAndNewsWritesOperationLogs(t *testing.T) {
	store := newMemoryStore()
	app := NewApp(Config{AdminUser: "admin", AdminPassword: "secret", TokenSecret: "test"}, store)
	token := createToken("admin", "test", time.Hour)

	actions := []struct {
		path string
		body string
	}{
		{"/api/admin/cases", `{"title":"Case","slug":"case","content":"body","status":"published"}`},
		{"/api/admin/news", `{"title":"News","slug":"news","content":"body","status":"published"}`},
	}
	for _, action := range actions {
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, action.path, strings.NewReader(action.body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		app.ServeHTTP(resp, req)
		if resp.Code != http.StatusCreated {
			t.Fatalf("expected create success for %s, got %d: %s", action.path, resp.Code, resp.Body.String())
		}
	}

	data, _ := json.Marshal(store.logs)
	if !strings.Contains(string(data), `"module":"cases"`) || !strings.Contains(string(data), `"module":"news"`) {
		t.Fatalf("expected case and news operation logs, got %s", data)
	}
}

func TestSanitizeNewsUsesPublishedAtAlias(t *testing.T) {
	item := sanitizeNews(News{
		Title:         "News",
		PublishedAtUI: "2026-07-05T10:30",
	})

	if item.PublishedAt != "2026-07-05 10:30:00" || item.PublishedAtUI != "2026-07-05 10:30:00" {
		t.Fatalf("expected publishedAt alias to be preserved, got published_at=%q publishedAt=%q", item.PublishedAt, item.PublishedAtUI)
	}
}
