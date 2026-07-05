package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
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
	a.mux.HandleFunc("/api/services", a.publicServices)
	a.mux.HandleFunc("/api/services/", a.publicServiceDetail)
	a.mux.HandleFunc("/api/partner", a.publicPartner)
	a.mux.HandleFunc("/api/public/cases", a.publicCases)
	a.mux.HandleFunc("/api/public/cases/", a.publicCaseDetail)
	a.mux.HandleFunc("/api/public/news", a.publicNews)
	a.mux.HandleFunc("/api/public/news/", a.publicNewsDetail)
	a.mux.HandleFunc("/api/public/forms", a.publicForms)

	a.mux.HandleFunc("/api/admin/login", a.adminLogin)
	a.mux.HandleFunc("/api/admin/dashboard", a.auth(a.adminDashboard))
	a.mux.HandleFunc("/api/admin/services/export", a.auth(a.adminServicesExport))
	a.mux.HandleFunc("/api/admin/services", a.auth(a.adminServices))
	a.mux.HandleFunc("/api/admin/services/", a.auth(a.adminServiceItem))
	a.mux.HandleFunc("/api/admin/pages/export", a.auth(a.adminPagesExport))
	a.mux.HandleFunc("/api/admin/pages", a.auth(a.adminPages))
	a.mux.HandleFunc("/api/admin/pages/", a.auth(a.adminPageItem))
	a.mux.HandleFunc("/api/admin/faqs/export", a.auth(a.adminFAQsExport))
	a.mux.HandleFunc("/api/admin/faqs", a.auth(a.adminFAQs))
	a.mux.HandleFunc("/api/admin/faqs/", a.auth(a.adminFAQItem))
	a.mux.HandleFunc("/api/admin/cases/export", a.auth(a.adminCasesExport))
	a.mux.HandleFunc("/api/admin/cases", a.auth(a.adminCases))
	a.mux.HandleFunc("/api/admin/cases/", a.auth(a.adminCaseDetail))
	a.mux.HandleFunc("/api/admin/news/export", a.auth(a.adminNewsExport))
	a.mux.HandleFunc("/api/admin/news", a.auth(a.adminNews))
	a.mux.HandleFunc("/api/admin/news/", a.auth(a.adminNewsItem))
	a.mux.HandleFunc("/api/admin/banners/export", a.auth(a.adminBannersExport))
	a.mux.HandleFunc("/api/admin/banners", a.auth(a.adminBanners))
	a.mux.HandleFunc("/api/admin/banners/", a.auth(a.adminBannerItem))
	a.mux.HandleFunc("/api/admin/banner", a.auth(a.adminBanner))
	a.mux.HandleFunc("/api/admin/settings/site", a.auth(a.adminSettings))
	a.mux.HandleFunc("/api/admin/settings/seo", a.auth(a.adminSEOCollection))
	a.mux.HandleFunc("/api/admin/settings", a.auth(a.adminSettings))
	a.mux.HandleFunc("/api/admin/seo", a.auth(a.adminSEO))
	a.mux.HandleFunc("/api/admin/forms/export", a.auth(a.adminFormsExport))
	a.mux.HandleFunc("/api/admin/forms", a.auth(a.adminForms))
	a.mux.HandleFunc("/api/admin/forms/", a.auth(a.adminFormItem))
	a.mux.HandleFunc("/api/admin/users", a.auth(a.requirePermission("user:manage", a.adminUsers)))
	a.mux.HandleFunc("/api/admin/users/", a.auth(a.requirePermission("user:manage", a.adminUserItem)))
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
	respond(w, must(a.store.ListBanners(r.Context(), ListOptions{Page: 1, PageSize: 20, Status: "enabled"})))
}

func (a *App) publicServices(w http.ResponseWriter, r *http.Request) {
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
	respond(w, must(a.store.ListServices(r.Context(), opts)))
}

func (a *App) publicServiceDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	respond(w, must(a.store.GetService(r.Context(), pathID(r.URL.Path, "/api/services/"), false)))
}

func (a *App) publicPartner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	faqs, err := a.store.ListFAQs(r.Context(), ListOptions{Page: 1, PageSize: 100, Status: "published"})
	if err != nil {
		writeError(w, err)
		return
	}
	page, _ := a.store.GetStaticPage(r.Context(), "partner", false)
	writeJSON(w, http.StatusOK, map[string]any{
		"page": page,
		"faqs": faqs.Items,
	})
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
	created, err := a.store.CreateForm(r.Context(), form)
	if err != nil {
		writeError(w, err)
		return
	}
	a.notifyFormSubmitted(r, created)
	writeJSON(w, http.StatusCreated, created)
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
	user, err := a.store.GetAdminUser(r.Context(), input.Username)
	if err == nil {
		if user.Status == "disabled" || !verifyPassword(input.Password, user.PasswordHash) {
			writeError(w, AppError{Status: http.StatusUnauthorized, Code: "invalid_credentials", Message: "invalid username or password"})
			return
		}
		a.logAdminWithUser(r, user.Username, "login", "auth", user.ID)
		writeJSON(w, http.StatusOK, map[string]any{
			"token":       createToken(user.Username, a.config.TokenSecret, a.config.TokenTTL()),
			"user":        user,
			"roles":       user.Roles,
			"permissions": user.Permissions,
		})
		return
	}
	var appErr AppError
	if errors.As(err, &appErr) {
		if !sameSecret(input.Username, a.config.AdminUser) || !sameSecret(input.Password, a.config.AdminPassword) {
			writeError(w, AppError{Status: http.StatusUnauthorized, Code: "invalid_credentials", Message: "invalid username or password"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"token":       createToken(input.Username, a.config.TokenSecret, a.config.TokenTTL()),
			"user":        AdminUser{Username: input.Username, Status: "active", Roles: []string{"env_admin"}, Permissions: []string{"*"}},
			"roles":       []string{"env_admin"},
			"permissions": []string{"*"},
		})
		return
	}
	writeError(w, err)
}

func (a *App) adminDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	stats, activities, err := a.store.Dashboard(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"stats": stats, "activities": activities})
}

func (a *App) adminUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListAdminUsers(r.Context(), opts)))
	case http.MethodPost:
		input, err := readAdminUserInput(r)
		if err != nil {
			writeError(w, err)
			return
		}
		if input.Password == "" {
			writeError(w, badRequest("missing_password", "password is required"))
			return
		}
		created, err := a.store.CreateAdminUser(r.Context(), input.AdminUser, input.Password)
		if err != nil {
			writeError(w, err)
			return
		}
		a.logAdmin(r, "create", "admin_users", created.ID)
		writeJSON(w, http.StatusCreated, created)
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminUserItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/users/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid user id"))
		return
	}
	switch r.Method {
	case http.MethodPut:
		input, err := readAdminUserInput(r)
		if err != nil {
			writeError(w, err)
			return
		}
		updated, err := a.store.UpdateAdminUser(r.Context(), id, input.AdminUser, input.Password)
		if err != nil {
			writeError(w, err)
			return
		}
		a.logAdmin(r, "update", "admin_users", id)
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if err := a.store.DeleteAdminUser(r.Context(), id); err != nil {
			writeError(w, err)
			return
		}
		a.logAdmin(r, "delete", "admin_users", id)
		w.WriteHeader(http.StatusNoContent)
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminServices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListServices(r.Context(), opts)))
	case http.MethodPost:
		var item Service
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		item = sanitizeService(item)
		if item.Title == "" {
			writeError(w, badRequest("invalid_service", "title is required"))
			return
		}
		created, err := a.store.CreateService(r.Context(), item)
		a.logAdmin(r, "create", "services", created.ID)
		respondCreated(w, must(created, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminServiceItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/services/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetService(r.Context(), strconv.FormatInt(id, 10), true)))
	case http.MethodPut:
		var item Service
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		updated, err := a.store.UpdateService(r.Context(), id, sanitizeService(item))
		a.logAdmin(r, "update", "services", id)
		respond(w, must(updated, err))
	case http.MethodDelete:
		err := a.store.DeleteService(r.Context(), id)
		a.logAdmin(r, "delete", "services", id)
		respond(w, result(map[string]bool{"deleted": true}, err))
	default:
		methodNotAllowed(w)
	}
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
		created, err := a.store.CreateCase(r.Context(), item)
		if err == nil {
			a.logAdmin(r, "create", "cases", created.ID)
		}
		respondCreated(w, must(created, err))
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
		updated, err := a.store.UpdateCase(r.Context(), id, sanitizeCase(item))
		if err == nil {
			a.logAdmin(r, "update", "cases", id)
		}
		respond(w, must(updated, err))
	case http.MethodDelete:
		err := a.store.DeleteCase(r.Context(), id)
		if err == nil {
			a.logAdmin(r, "delete", "cases", id)
		}
		respond(w, result(map[string]bool{"deleted": true}, err))
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
		created, err := a.store.CreateNews(r.Context(), item)
		if err == nil {
			a.logAdmin(r, "create", "news", created.ID)
		}
		respondCreated(w, must(created, err))
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
		updated, err := a.store.UpdateNews(r.Context(), id, sanitizeNews(item))
		if err == nil {
			a.logAdmin(r, "update", "news", id)
		}
		respond(w, must(updated, err))
	case http.MethodDelete:
		err := a.store.DeleteNews(r.Context(), id)
		if err == nil {
			a.logAdmin(r, "delete", "news", id)
		}
		respond(w, result(map[string]bool{"deleted": true}, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminPages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListStaticPages(r.Context(), opts)))
	case http.MethodPost:
		var item StaticPage
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		item = sanitizePage(item)
		if item.PageKey == "" || item.Title == "" {
			writeError(w, badRequest("invalid_page", "page_key and title are required"))
			return
		}
		created, err := a.store.CreateStaticPage(r.Context(), item)
		a.logAdmin(r, "create", "pages", created.ID)
		respondCreated(w, must(created, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminPageItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/pages/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetStaticPage(r.Context(), strconv.FormatInt(id, 10), true)))
	case http.MethodPut:
		var item StaticPage
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		updated, err := a.store.UpdateStaticPage(r.Context(), id, sanitizePage(item))
		a.logAdmin(r, "update", "pages", id)
		respond(w, must(updated, err))
	case http.MethodDelete:
		err := a.store.DeleteStaticPage(r.Context(), id)
		a.logAdmin(r, "delete", "pages", id)
		respond(w, result(map[string]bool{"deleted": true}, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminFAQs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListFAQs(r.Context(), opts)))
	case http.MethodPost:
		var item PartnerFAQ
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		item = sanitizeFAQ(item)
		if item.Question == "" || item.Answer == "" {
			writeError(w, badRequest("invalid_faq", "question and answer are required"))
			return
		}
		created, err := a.store.CreateFAQ(r.Context(), item)
		a.logAdmin(r, "create", "faqs", created.ID)
		respondCreated(w, must(created, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminFAQItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/faqs/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetFAQ(r.Context(), id)))
	case http.MethodPut:
		var item PartnerFAQ
		if err := readJSON(r, &item); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		updated, err := a.store.UpdateFAQ(r.Context(), id, sanitizeFAQ(item))
		a.logAdmin(r, "update", "faqs", id)
		respond(w, must(updated, err))
	case http.MethodDelete:
		err := a.store.DeleteFAQ(r.Context(), id)
		a.logAdmin(r, "delete", "faqs", id)
		respond(w, result(map[string]bool{"deleted": true}, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminBanners(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		opts, err := listOptions(r, false)
		if err != nil {
			writeError(w, err)
			return
		}
		respond(w, must(a.store.ListBanners(r.Context(), opts)))
	case http.MethodPost:
		var banner Banner
		if err := readJSON(r, &banner); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		saved, err := a.store.CreateBanner(r.Context(), sanitizeBanner(banner))
		a.logAdmin(r, "create", "banners", saved.ID)
		respondCreated(w, must(saved, err))
	default:
		methodNotAllowed(w)
	}
}

func (a *App) adminBannerItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/banners/"))
	if !ok {
		writeError(w, badRequest("invalid_id", "invalid ID"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		respond(w, must(a.store.GetBannerByID(r.Context(), id)))
	case http.MethodPut:
		var banner Banner
		if err := readJSON(r, &banner); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		banner.ID = id
		saved, err := a.store.UpdateBanner(r.Context(), id, sanitizeBanner(banner))
		a.logAdmin(r, "update", "banners", id)
		respond(w, must(saved, err))
	case http.MethodDelete:
		err := a.store.DeleteBanner(r.Context(), id)
		a.logAdmin(r, "delete", "banners", id)
		respond(w, result(map[string]bool{"deleted": true}, err))
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

func (a *App) adminSEOCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		items, err := a.store.ListSEO(r.Context())
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": items})
	case http.MethodPut:
		var input struct {
			Items []SEOSetting `json:"items"`
		}
		if err := readJSON(r, &input); err != nil {
			writeError(w, badRequest("invalid_json", "request body must be valid JSON"))
			return
		}
		for _, seo := range input.Items {
			seo.PageKey = firstNonEmpty(sanitizeText(seo.PageKey), sanitizeText(seo.Page))
			seo.Title = sanitizeText(seo.Title)
			seo.Description = sanitizeText(seo.Description)
			seo.Keywords = sanitizeText(seo.Keywords)
			if seo.PageKey == "" {
				continue
			}
			if _, err := a.store.SaveSEO(r.Context(), seo); err != nil {
				writeError(w, err)
				return
			}
		}
		a.logAdmin(r, "update", "seo_settings", 0)
		items, err := a.store.ListSEO(r.Context())
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": items})
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
		if opts.Status == "pending" {
			opts.Status = "new"
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

func (a *App) adminServicesExport(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.ListServices(r.Context(), ListOptions{Page: 1, PageSize: 10000, Status: sanitizeText(r.URL.Query().Get("status")), Keyword: sanitizeText(r.URL.Query().Get("keyword"))})
	if err != nil {
		writeError(w, err)
		return
	}
	rows := [][]string{{"id", "title", "slug", "status", "updated_at"}}
	for _, item := range result.Items {
		rows = append(rows, []string{strconv.FormatInt(item.ID, 10), item.Title, item.Slug, item.Status, item.UpdatedAt})
	}
	writeCSVRows(w, "services.csv", rows)
}

func (a *App) adminPagesExport(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.ListStaticPages(r.Context(), ListOptions{Page: 1, PageSize: 10000, Status: sanitizeText(r.URL.Query().Get("status")), Keyword: sanitizeText(r.URL.Query().Get("keyword"))})
	if err != nil {
		writeError(w, err)
		return
	}
	rows := [][]string{{"id", "page_key", "title", "status", "updated_at"}}
	for _, item := range result.Items {
		rows = append(rows, []string{strconv.FormatInt(item.ID, 10), item.PageKey, item.Title, item.Status, item.UpdatedAt})
	}
	writeCSVRows(w, "pages.csv", rows)
}

func (a *App) adminFAQsExport(w http.ResponseWriter, r *http.Request) {
	result, err := a.store.ListFAQs(r.Context(), ListOptions{Page: 1, PageSize: 10000, Status: sanitizeText(r.URL.Query().Get("status"))})
	if err != nil {
		writeError(w, err)
		return
	}
	rows := [][]string{{"id", "question", "answer", "status", "updated_at"}}
	for _, item := range result.Items {
		rows = append(rows, []string{strconv.FormatInt(item.ID, 10), item.Question, item.Answer, item.Status, item.UpdatedAt})
	}
	writeCSVRows(w, "faqs.csv", rows)
}

func (a *App) adminCasesExport(w http.ResponseWriter, r *http.Request) {
	opts, _ := listOptions(r, false)
	opts.Page, opts.PageSize = 1, 10000
	result, err := a.store.ListCases(r.Context(), opts)
	if err != nil {
		writeError(w, err)
		return
	}
	rows := [][]string{{"id", "title", "industry", "platform", "strategy", "status", "updated_at"}}
	for _, item := range result.Items {
		rows = append(rows, []string{strconv.FormatInt(item.ID, 10), item.Title, item.Industry, item.Platform, item.Strategy, item.Status, item.UpdatedAt})
	}
	writeCSVRows(w, "cases.csv", rows)
}

func (a *App) adminNewsExport(w http.ResponseWriter, r *http.Request) {
	opts, _ := listOptions(r, false)
	opts.Page, opts.PageSize = 1, 10000
	result, err := a.store.ListNews(r.Context(), opts)
	if err != nil {
		writeError(w, err)
		return
	}
	rows := [][]string{{"id", "title", "category", "status", "published_at", "updated_at"}}
	for _, item := range result.Items {
		rows = append(rows, []string{strconv.FormatInt(item.ID, 10), item.Title, item.Category, item.Status, item.PublishedAt, item.UpdatedAt})
	}
	writeCSVRows(w, "news.csv", rows)
}

func (a *App) adminBannersExport(w http.ResponseWriter, r *http.Request) {
	banner, err := a.store.GetBanner(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeCSVRows(w, "banners.csv", [][]string{
		{"id", "title", "subtitle", "image_url", "link_url", "status", "updated_at"},
		{strconv.FormatInt(banner.ID, 10), banner.Title, banner.Subtitle, banner.ImageURL, banner.LinkURL, banner.Status, banner.UpdatedAt},
	})
}

func (a *App) adminFormItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id, ok := parseIntID(pathID(r.URL.Path, "/api/admin/forms/"))
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
		if input.Status == "pending" {
			input.Status = "new"
		}
		if input.Status != "new" && input.Status != "processed" {
			writeError(w, badRequest("invalid_status", "status must be new or processed"))
			return
		}
		updated, err := a.store.UpdateFormStatus(r.Context(), id, input.Status)
		a.logAdmin(r, "update", "lead_forms", id)
		respond(w, must(updated, err))
		return
	}
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
	status := sanitizeText(r.URL.Query().Get("status"))
	if status == "pending" {
		status = "new"
	}
	result, err := a.store.ListForms(r.Context(), ListOptions{Page: 1, PageSize: 10000, Status: status})
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

func writeCSVRows(w http.ResponseWriter, filename string, rows [][]string) {
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	w.WriteHeader(http.StatusOK)
	writer := csv.NewWriter(w)
	for _, row := range rows {
		_ = writer.Write(row)
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
	bizType := sanitizeText(r.FormValue("type"))
	if !allowedUploadExt(ext, bizType) {
		writeError(w, badRequest("invalid_file_type", "only jpg, jpeg, png, gif, webp, pdf and banner mp4 files are allowed"))
		return
	}
	data, err := readUpload(file, a.config.MaxUploadBytes)
	if err != nil {
		writeError(w, badRequest("invalid_upload", "file is too large or cannot be read"))
		return
	}
	if !allowedUploadContent(ext, data, bizType) {
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
	url := "/oss/" + day + "/" + name
	asset, err := a.store.CreateMediaAsset(r.Context(), MediaAsset{
		OriginalName: sanitizeText(header.Filename),
		FileName:     name,
		FilePath:     dstPath,
		FileURL:      url,
		MimeType:     firstNonEmpty(header.Header.Get("Content-Type"), mime.TypeByExtension(ext), http.DetectContentType(data)),
		FileSize:     int64(len(data)),
		BizType:      bizType,
	})
	if err != nil {
		writeError(w, err)
		return
	}
	a.logAdmin(r, "upload", "media_assets", asset.ID)
	writeJSON(w, http.StatusCreated, map[string]any{"url": url, "asset": asset})
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

func (a *App) requirePermission(permission string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := tokenSubject(bearerToken(r), a.config.TokenSecret)
		if username == a.config.AdminUser {
			next(w, r)
			return
		}
		user, err := a.store.GetAdminUser(r.Context(), username)
		if err != nil || !adminUserHasPermission(user, permission) {
			writeError(w, AppError{Status: http.StatusForbidden, Code: "forbidden", Message: "permission denied"})
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

type adminUserInput struct {
	AdminUser
	Password string `json:"password"`
}

func readAdminUserInput(r *http.Request) (adminUserInput, error) {
	var input adminUserInput
	if err := readJSON(r, &input); err != nil {
		return input, badRequest("invalid_json", "request body must be valid JSON")
	}
	input.Username = sanitizeText(input.Username)
	input.RealName = sanitizeText(input.RealName)
	input.Email = sanitizeText(input.Email)
	input.Phone = sanitizeText(input.Phone)
	input.Status = normalizeAdminStatus(input.Status)
	cleanRoles := []string{}
	seen := map[string]bool{}
	for _, role := range input.Roles {
		role = sanitizeText(role)
		if role == "" || seen[role] {
			continue
		}
		seen[role] = true
		cleanRoles = append(cleanRoles, role)
	}
	input.Roles = cleanRoles
	input.Permissions = nil
	if input.Username == "" {
		return input, badRequest("missing_username", "username is required")
	}
	return input, nil
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
	item.Playbook = sanitizeText(item.Playbook)
	item.Method = sanitizeText(item.Method)
	item.CoverURL = sanitizeText(item.CoverURL)
	item.Cover = sanitizeText(item.Cover)
	item.Image = sanitizeText(item.Image)
	item.Summary = sanitizeText(item.Summary)
	item.Content = sanitizeHTML(item.Content)
	item.CoreMetrics = sanitizeText(item.CoreMetrics)
	item.Metrics = sanitizeText(item.Metrics)
	item.Status = validateStatus(item.Status)
	return normalizeCase(item)
}

func sanitizeNews(item News) News {
	item.Title = sanitizeText(item.Title)
	item.Slug = sanitizeText(item.Slug)
	item.Category = sanitizeText(item.Category)
	item.CoverURL = sanitizeText(item.CoverURL)
	item.Cover = sanitizeText(item.Cover)
	item.Image = sanitizeText(item.Image)
	item.Summary = sanitizeText(item.Summary)
	item.Desc = sanitizeText(item.Desc)
	item.Content = sanitizeHTML(item.Content)
	item.Status = validateStatus(item.Status)
	item.PublishedAt = sanitizeText(item.PublishedAt)
	item.PublishedAtUI = sanitizeText(item.PublishedAtUI)
	item.Date = sanitizeText(item.Date)
	item.PublishedAt = normalizeDateTimeInput(firstNonEmpty(item.PublishedAt, item.PublishedAtUI, item.Date))
	if item.PublishedAt == "" {
		item.PublishedAt = time.Now().Format("2006-01-02 15:04:05")
	}
	item.PublishedAtUI = item.PublishedAt
	item.Date = item.PublishedAt
	return normalizeNews(item)
}

func normalizeDateTimeInput(value string) string {
	value = strings.TrimSpace(value)
	if len(value) >= len("2006-01-02T15:04") && strings.Contains(value, "T") {
		value = strings.Replace(value, "T", " ", 1)
		if len(value) == len("2006-01-02 15:04") {
			value += ":00"
		}
	}
	return value
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
	if form.Status == "pending" {
		form.Status = "new"
	}
	if form.Status == "" {
		form.Status = "new"
	}
	return normalizeForm(form)
}

func sanitizeBanner(banner Banner) Banner {
	banner.Title = sanitizeText(banner.Title)
	banner.Subtitle = sanitizeText(banner.Subtitle)
	banner.ImageURL = sanitizeText(banner.ImageURL)
	banner.Image = sanitizeText(banner.Image)
	banner.LinkURL = sanitizeText(banner.LinkURL)
	banner.Link = sanitizeText(banner.Link)
	banner.ButtonText = sanitizeText(banner.ButtonText)
	banner.Page = sanitizeText(banner.Page)
	banner.Status = sanitizeText(banner.Status)
	return normalizeBanner(banner)
}

func sanitizeService(item Service) Service {
	item.Title = sanitizeText(item.Title)
	item.Slug = sanitizeText(item.Slug)
	item.Subtitle = sanitizeText(item.Subtitle)
	item.Summary = sanitizeText(item.Summary)
	item.CoverURL = sanitizeText(item.CoverURL)
	item.Image = sanitizeText(item.Image)
	item.IconURL = sanitizeText(item.IconURL)
	item.Content = sanitizeHTML(item.Content)
	item.HighlightText = sanitizeText(item.HighlightText)
	item.ProcessText = sanitizeText(item.ProcessText)
	item.Status = validateStatus(item.Status)
	return normalizeService(item)
}

func sanitizePage(item StaticPage) StaticPage {
	item.PageKey = sanitizeText(item.PageKey)
	item.Page = sanitizeText(item.Page)
	item.Title = sanitizeText(item.Title)
	item.Content = sanitizeHTML(item.Content)
	item.ExtraData = sanitizeHTML(item.ExtraData)
	item.Status = validateStatus(item.Status)
	return normalizePage(item)
}

func sanitizeFAQ(item PartnerFAQ) PartnerFAQ {
	item.Question = sanitizeText(item.Question)
	item.Answer = sanitizeHTML(item.Answer)
	item.Status = sanitizeText(item.Status)
	return normalizeFAQ(item)
}

func (a *App) logAdmin(r *http.Request, action, module string, targetID int64) {
	a.logAdminWithUser(r, tokenSubject(bearerToken(r), a.config.TokenSecret), action, module, targetID)
}

func (a *App) logAdminWithUser(r *http.Request, username, action, module string, targetID int64) {
	_ = a.store.LogOperation(r.Context(), OperationLog{
		Username:  username,
		Action:    action,
		Module:    module,
		Target:    module,
		TargetID:  targetID,
		CreatedAt: nowString(),
	})
}

func (a *App) notifyFormSubmitted(r *http.Request, form LeadForm) {
	setting, err := a.store.GetEmailSetting(r.Context())
	if err != nil || !setting.IsEnabled {
		return
	}
	recipients := splitCSV(setting.Recipients)
	if len(recipients) == 0 {
		return
	}
	subject := "New lead form submitted"
	content := fmt.Sprintf("Name: %s\nPhone: %s\nCompany: %s\nRequirement: %s\n", form.Name, form.Phone, form.Company, form.Requirement)
	for _, recipient := range recipients {
		notification, err := a.store.CreateEmailNotification(r.Context(), EmailNotification{
			LeadFormID: form.ID,
			Recipient:  recipient,
			Subject:    subject,
			Content:    content,
			Status:     "pending",
		})
		if err != nil {
			continue
		}
		if setting.SMTPHost == "" || setting.SenderEmail == "" {
			continue
		}
		if err := sendEmail(setting, recipient, subject, content); err != nil {
			_ = a.store.UpdateEmailNotificationStatus(r.Context(), notification.ID, "failed", err.Error())
			continue
		}
		_ = a.store.UpdateEmailNotificationStatus(r.Context(), notification.ID, "sent", "")
	}
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

func allowedUploadExt(ext, bizType string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".pdf":
		return true
	case ".mp4":
		return bizType == "banner"
	default:
		return false
	}
}

func allowedUploadContent(ext string, data []byte, bizType string) bool {
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
	case ".mp4":
		return bizType == "banner" && isMP4Content(data)
	default:
		return false
	}
}

func isMP4Content(data []byte) bool {
	if len(data) < 12 {
		return false
	}
	return string(data[4:8]) == "ftyp"
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
