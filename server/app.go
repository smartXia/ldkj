package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type App struct {
	config Config
	store  Store
	mux    *http.ServeMux
}

func NewApp(config Config, store Store) *App {
	if config.TokenSecret == "" {
		config.TokenSecret = "change-me-token-secret"
	}
	if config.UploadDir == "" {
		config.UploadDir = "oss"
	}
	if config.MaxUploadBytes == 0 {
		config.MaxUploadBytes = 5 << 20
	}
	app := &App{config: config, store: store, mux: http.NewServeMux()}
	app.routes()
	return app
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.withCORS(a.mux).ServeHTTP(w, r)
}

func (a *App) routes() {
	a.mux.HandleFunc("/api/health", a.handleHealth)
	a.mux.HandleFunc("/api/public/site", a.publicSite)
	a.mux.HandleFunc("/api/public/banner", a.publicBanner)
	a.mux.HandleFunc("/api/public/cases", a.publicCases)
	a.mux.HandleFunc("/api/public/cases/", a.publicCaseDetail)
	a.mux.HandleFunc("/api/public/news", a.publicNews)
	a.mux.HandleFunc("/api/public/news/", a.publicNewsDetail)
	a.mux.HandleFunc("/api/public/forms", a.publicForms)

	a.mux.HandleFunc("/api/admin/login", a.adminLogin)
	a.mux.HandleFunc("/api/admin/cases", a.auth(a.adminCases))
	a.mux.HandleFunc("/api/admin/cases/", a.auth(a.adminCaseDetail))
	a.mux.HandleFunc("/api/admin/news", a.auth(a.adminNews))
	a.mux.HandleFunc("/api/admin/news/", a.auth(a.adminNewsItem))
	a.mux.HandleFunc("/api/admin/banner", a.auth(a.adminBanner))
	a.mux.HandleFunc("/api/admin/settings", a.auth(a.adminSettings))
	a.mux.HandleFunc("/api/admin/seo", a.auth(a.adminSEO))
	a.mux.HandleFunc("/api/admin/forms/export", a.auth(a.adminFormsExport))
	a.mux.HandleFunc("/api/admin/forms", a.auth(a.adminForms))
	a.mux.HandleFunc("/api/admin/forms/", a.auth(a.adminFormItem))
	a.mux.HandleFunc("/api/admin/upload", a.auth(a.adminUpload))
	a.mux.Handle("/oss/", http.StripPrefix("/oss/", http.FileServer(http.Dir(a.config.UploadDir))))
}

func (a *App) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *App) publicSite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	respond(w, must(a.store.GetSiteConfig(r.Context())))
}

func (a *App) publicBanner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	respond(w, must(a.store.GetBanner(r.Context())))
}

func (a *App) publicCases(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	opts, err := listOptions(r, true)
	if err != nil {
		writeError(w, err)
		return
	}
	opts.Status = "published"
	respond(w, must(a.store.ListCases(r.Context(), opts)))
}

func (a *App) publicCaseDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	respond(w, must(a.store.GetCase(r.Context(), pathID(r.URL.Path, "/api/public/cases/"), false)))
}

func (a *App) publicNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	opts, err := listOptions(r, true)
	if err != nil {
		writeError(w, err)
		return
	}
	opts.Status = "published"
	respond(w, must(a.store.ListNews(r.Context(), opts)))
}

func (a *App) publicNewsDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	respond(w, must(a.store.GetNews(r.Context(), pathID(r.URL.Path, "/api/public/news/"), false)))
}

func (a *App) publicForms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	var form LeadForm
	if err := readJSON(r, &form); err != nil {
		writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
		return
	}
	form = sanitizeForm(form)
	if form.Name == "" || form.Phone == "" || form.Company == "" {
		writeError(w, badRequest("invalid_form", "name, phone and company are required"))
		return
	}
	if form.Status == "" {
		form.Status = "new"
	}
	respondCreated(w, must(a.store.CreateForm(r.Context(), form)))
}

func (a *App) adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := readJSON(r, &input); err != nil {
		writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
		return
	}
	if !sameSecret(input.Username, a.config.AdminUser) || !sameSecret(input.Password, a.config.AdminPassword) {
		writeError(w, AppError{Status: http.StatusUnauthorized, Code: "invalid_credentials", Message: "invalid username or password"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"token": createToken(input.Username, a.config.TokenSecret, a.config.TokenTTL())})
}

func (a *App) adminCases(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListCases(r.Context(), opts)))
	case http.MethodPost:
		var item Case
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		item = sanitizeCase(item)
		if item.Title == "" || item.Slug == "" {
			writeError(w, badRequest("invalid_case", "title and slug are required"))
			return
		}
		respondCreated(w, must(a.store.CreateCase(r.Context(), item)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminCaseDetail(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/cases/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetCase(r.Context(), strconv.FormatInt(id, 10), true)))
	case http.MethodPut:
		var item Case
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		respond(w, must(a.store.UpdateCase(r.Context(), id, sanitizeCase(item))))
	case http.MethodDelete:
		respond(w, result(map[string]bool{"deleted": true}, a.store.DeleteCase(r.Context(), id)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminNews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListNews(r.Context(), opts)))
	case http.MethodPost:
		var item News
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		item = sanitizeNews(item)
		if item.Title == "" || item.Slug == "" {
			writeError(w, badRequest("invalid_news", "title and slug are required"))
			return
		}
		respondCreated(w, must(a.store.CreateNews(r.Context(), item)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminNewsItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/news/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetNews(r.Context(), strconv.FormatInt(id, 10), true)))
	case http.MethodPut:
		var item News
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		respond(w, must(a.store.UpdateNews(r.Context(), id, sanitizeNews(item))))
	case http.MethodDelete:
		respond(w, result(map[string]bool{"deleted": true}, a.store.DeleteNews(r.Context(), id)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminBanner(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetBanner(r.Context())))
	case http.MethodPut:
		var banner Banner
		if err := readJSON(r, &banner); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		banner.Title = sanitizeText(banner.Title)
		banner.Subtitle = sanitizeText(banner.Subtitle)
		banner.ImageURL = sanitizeText(banner.ImageURL)
		banner.LinkURL = sanitizeText(banner.LinkURL)
		banner.ButtonText = sanitizeText(banner.ButtonText)
		respond(w, must(a.store.SaveBanner(r.Context(), banner)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminSettings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetSiteConfig(r.Context())))
	case http.MethodPut:
		var site SiteConfig
		if err := readJSON(r, &site); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		site.SiteTitle = sanitizeText(site.SiteTitle)
		site.LogoURL = sanitizeText(site.LogoURL)
		site.Contact = sanitizeHTML(site.Contact)
		site.Copyright = sanitizeText(site.Copyright)
		site.FooterLinks = sanitizeHTML(site.FooterLinks)
		respond(w, must(a.store.SaveSiteConfig(r.Context(), site)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminSEO(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page := sanitizeText(r.URL.Query().Get("page_key"))
		if page == "" {
			page = "home"
		}
		respond(w, must(a.store.GetSEO(r.Context(), page)))
	case http.MethodPut:
		var seo SEOSetting
		if err := readJSON(r, &seo); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		seo.PageKey = sanitizeText(seo.PageKey)
		seo.Title = sanitizeText(seo.Title)
		seo.Description = sanitizeText(seo.Description)
		seo.Keywords = sanitizeText(seo.Keywords)
		if seo.PageKey == "" {
			writeError(w, badRequest("invalid_seo", "page_key is required"))
			return
		}
		respond(w, must(a.store.SaveSEO(r.Context(), seo)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminForms(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListForms(r.Context(), opts)))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminFormsExport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	a.exportForms(w, r)
}

func (a *App) adminFormItem(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/status") && r.Method == http.MethodPatch {
		idPart := strings.TrimSuffix(pathID(r.URL.Path, "/api/admin/forms/"), "/status")
		id, ok := parseIntID(idPart)
		if !ok {
			writeError(w, badRequest("invalid_id", "invalid ID"))
			return
		}
		var input struct {
			Status string `json:"status"`
		}
		if err := readJSON(r, &input); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		if input.Status != "new" && input.Status != "processed" {
			writeError(w, badRequest("invalid_status", "status must be new or processed"))
			return
		}
		respond(w, must(a.store.UpdateFormStatus(r.Context(), id, input.Status)))
		return
	}
	methodNotAllowed(w)
}

func (a *App) exportForms(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.ListForms(r.Context(), ListOptions{Page: 1, PageSize: 10000, Status: sanitizeText(r.URL.Query().Get("status"))})
	if err != nil {
		writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="forms.csv"`)
	w.WriteHeader(http.StatusOK)
	writer := csv.NewWriter(w)
	_ = writer.Write([]string{"id", "name", "phone", "company", "position", "email", "requirement", "interest", "source", "status", "created_at"})
	for _, item := range result.Items {
		_ = writer.Write([]string{
			strconv.FormatInt(item.ID, 10), item.Name, item.Phone, item.Company, item.Position, item.Email,
			item.Requirement, item.Interest, item.Source, item.Status, item.CreatedAt,
		})
	}
	writer.Flush()
}

func (a *App) adminUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, a.config.MaxUploadBytes)
	if err := r.ParseMultipartForm(a.config.MaxUploadBytes); err != nil {
		writeError(w, badRequest("invalid_upload", "file is too large or multipart form is invalid"))
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, badRequest("missing_file", "missing file field"))
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedUploadExt(ext) {
		writeError(w, badRequest("invalid_file_type", "only jpg, jpeg, png, gif, webp and pdf files are allowed"))
		return
	}
	data, err := readUpload(file, a.config.MaxUploadBytes)
	if err != nil {
		writeError(w, badRequest("invalid_upload", "file is too large or cannot be read"))
		return
	}
	if !allowedUploadContent(ext, data) {
		writeError(w, badRequest("invalid_file_type", "file content does not match its extension"))
		return
	}

	day := time.Now().Format("20060102")
	dir := filepath.Join(a.config.UploadDir, day)
	if err := os.MkdirAll(dir, 0755); err != nil {
		writeError(w, internal(err))
		return
	}
	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(dir, name)
	dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		writeError(w, internal(err))
		return
	}
	defer dst.Close()
	if _, err := dst.Write(data); err != nil {
		writeError(w, internal(err))
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"url": "/oss/" + day + "/" + name})
}

func (a *App) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if !validateToken(bearerToken(r), a.config.TokenSecret) {
			writeError(w, AppError{Status: http.StatusUnauthorized, Code: "unauthorized", Message: "login required"})
			return
		}
		next(w, r)
	}
}

func (a *App) withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" || containsOrigin(a.config.CORSOrigins, "*") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else if containsOrigin(a.config.CORSOrigins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func listOptions(r *http.Request, public bool) (ListOptions, error) {
	query := r.URL.Query()
	page, err := positiveInt(query.Get("page"), 1, 1, 100000)
	if err != nil {
		return ListOptions{}, err
	}
	pageSize, err := positiveInt(query.Get("page_size"), 10, 1, 100)
	if err != nil {
		return ListOptions{}, err
	}
	opts := ListOptions{
		Page:     page,
		PageSize: pageSize,
		Status:   sanitizeText(query.Get("status")),
		Industry: sanitizeText(query.Get("industry")),
		Platform: sanitizeText(query.Get("platform")),
		Strategy: sanitizeText(query.Get("strategy")),
		Category: sanitizeText(query.Get("category")),
		Keyword:  sanitizeText(query.Get("keyword")),
	}
	if public {
		opts.Status = "published"
	}
	return opts, nil
}

func positiveInt(raw string, fallback, min, max int) (int, error) {
	if raw == "" {
		return fallback, nil
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v < min || v > max {
		return 0, badRequest("invalid_pagination", "invalid pagination parameters")
	}
	return v, nil
}

func readJSON(r *http.Request, out any) error {
	decoder := json.NewDecoder(http.MaxBytesReader(nilWriter{}, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	return decoder.Decode(out)
}

type nilWriter struct{}

func (nilWriter) Header() http.Header       { return http.Header{} }
func (nilWriter) Write([]byte) (int, error) { return 0, nil }
func (nilWriter) WriteHeader(int)           {}

type responseResult struct {
	data any
	err  error
}

func respond(w http.ResponseWriter, result responseResult) {
	if result.err != nil {
		writeError(w, result.err)
		return
	}
	writeJSON(w, http.StatusOK, result.data)
}

func respondCreated(w http.ResponseWriter, result responseResult) {
	if result.err != nil {
		writeError(w, result.err)
		return
	}
	writeJSON(w, http.StatusCreated, result.data)
}

func must[T any](data T, err error) responseResult {
	return responseResult{data: data, err: err}
}

func result(data any, err error) responseResult {
	return responseResult{data: data, err: err}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	var appErr AppError
	if errors.As(err, &appErr) {
		writeJSON(w, appErr.Status, appErr)
		return
	}
	writeJSON(w, http.StatusInternalServerError, AppError{Code: "internal_error", Message: "internal server error"})
}

func badRequest(code, message string) AppError {
	return AppError{Status: http.StatusBadRequest, Code: code, Message: message}
}

func notFound(message string) AppError {
	return AppError{Status: http.StatusNotFound, Code: "not_found", Message: message}
}

func internal(err error) AppError {
	return AppError{Status: http.StatusInternalServerError, Code: "internal_error", Message: err.Error()}
}

func methodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, AppError{Code: "method_not_allowed", Message: "method not allowed"})
}

func pathID(path, prefix string) string {
	return strings.Trim(strings.TrimPrefix(path, prefix), "/")
}

func parseIntID(raw string) (int64, bool) {
	id, err := strconv.ParseInt(raw, 10, 64)
	return id, err == nil && id > 0
}

func sanitizeCase(item Case) Case {
	item.Title = sanitizeText(item.Title)
	item.Slug = sanitizeText(item.Slug)
	item.Industry = sanitizeText(item.Industry)
	item.Platform = sanitizeText(item.Platform)
	item.Strategy = sanitizeText(item.Strategy)
	item.CoverURL = sanitizeText(item.CoverURL)
	item.Summary = sanitizeText(item.Summary)
	item.Content = sanitizeHTML(item.Content)
	item.CoreMetrics = sanitizeText(item.CoreMetrics)
	item.Status = validateStatus(item.Status)
	return item
}

func sanitizeNews(item News) News {
	item.Title = sanitizeText(item.Title)
	item.Slug = sanitizeText(item.Slug)
	item.Category = sanitizeText(item.Category)
	item.CoverURL = sanitizeText(item.CoverURL)
	item.Summary = sanitizeText(item.Summary)
	item.Content = sanitizeHTML(item.Content)
	item.Status = validateStatus(item.Status)
	item.PublishedAt = sanitizeText(item.PublishedAt)
	if item.PublishedAt == "" {
		item.PublishedAt = time.Now().Format("2006-01-02 15:04:05")
	}
	return item
}

func sanitizeForm(form LeadForm) LeadForm {
	form.Name = sanitizeText(form.Name)
	form.Phone = sanitizeText(form.Phone)
	form.Company = sanitizeText(form.Company)
	form.Position = sanitizeText(form.Position)
	form.Email = sanitizeText(form.Email)
	form.Requirement = sanitizeText(form.Requirement)
	form.Interest = sanitizeText(form.Interest)
	form.Source = sanitizeText(form.Source)
	if form.Status == "" {
		form.Status = "new"
	}
	return form
}

func readUpload(r io.Reader, maxBytes int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(r, maxBytes+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > maxBytes {
		return nil, errors.New("upload exceeds max size")
	}
	return data, nil
}

func allowedUploadExt(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".pdf":
		return true
	default:
		return false
	}
}

func allowedUploadContent(ext string, data []byte) bool {
	if len(data) == 0 {
		return false
	}
	contentType := http.DetectContentType(data)
	switch ext {
	case ".jpg", ".jpeg":
		return contentType == "image/jpeg"
	case ".png":
		return contentType == "image/png"
	case ".gif":
		return contentType == "image/gif"
	case ".webp":
		return contentType == "image/webp" || bytes.HasPrefix(data, []byte("RIFF")) && bytes.Contains(data[:min(len(data), 16)], []byte("WEBP"))
	case ".pdf":
		return contentType == "application/pdf" || bytes.HasPrefix(data, []byte("%PDF-"))
	default:
		return false
	}
}

func containsOrigin(list []string, origin string) bool {
	for _, item := range list {
		if item == origin {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
