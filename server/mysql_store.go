package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) Store {
	return &mysqlStore{db: db}
}

func OpenMySQL(ctx context.Context, config Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (s *mysqlStore) GetSiteConfig(ctx context.Context) (SiteConfig, error) {
	var site SiteConfig
	err := s.db.QueryRowContext(ctx, `SELECT id, site_title, logo_url, contact, copyright, footer_links, updated_at FROM site_config WHERE id = 1`).Scan(
		&site.ID, &site.SiteTitle, &site.LogoURL, &site.Contact, &site.Copyright, &site.FooterLinks, &site.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return SiteConfig{}, notFound("site config not found")
	}
	return site, err
}

func (s *mysqlStore) SaveSiteConfig(ctx context.Context, site SiteConfig) (SiteConfig, error) {
	_, err := s.db.ExecContext(ctx, `INSERT INTO site_config (id, site_title, logo_url, contact, copyright, footer_links)
VALUES (1, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE site_title = VALUES(site_title), logo_url = VALUES(logo_url), contact = VALUES(contact), copyright = VALUES(copyright), footer_links = VALUES(footer_links)`,
		site.SiteTitle, site.LogoURL, site.Contact, site.Copyright, site.FooterLinks)
	if err != nil {
		return SiteConfig{}, err
	}
	return s.GetSiteConfig(ctx)
}

func (s *mysqlStore) GetBanner(ctx context.Context) (Banner, error) {
	var b Banner
	order := bannerOrderClause(s.columnExists(ctx, "banners", "sort_order"))
	err := s.db.QueryRowContext(ctx, `SELECT id, title, subtitle, image_url, link_url, button_text, is_published, updated_at FROM banners WHERE is_published = TRUE `+order+` LIMIT 1`).Scan(
		&b.ID, &b.Title, &b.Subtitle, &b.ImageURL, &b.LinkURL, &b.ButtonText, &b.IsPublished, &b.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		err = s.db.QueryRowContext(ctx, `SELECT id, title, subtitle, image_url, link_url, button_text, is_published, updated_at FROM banners `+order+` LIMIT 1`).Scan(
			&b.ID, &b.Title, &b.Subtitle, &b.ImageURL, &b.LinkURL, &b.ButtonText, &b.IsPublished, &b.UpdatedAt)
		if errors.Is(err, sql.ErrNoRows) {
			return Banner{}, notFound("banner not found")
		}
	}
	return normalizeBanner(b), err
}

func (s *mysqlStore) SaveBanner(ctx context.Context, b Banner) (Banner, error) {
	_, err := s.db.ExecContext(ctx, `INSERT INTO banners (id, title, subtitle, image_url, link_url, button_text, is_published)
VALUES (1, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE title = VALUES(title), subtitle = VALUES(subtitle), image_url = VALUES(image_url), link_url = VALUES(link_url), button_text = VALUES(button_text), is_published = VALUES(is_published)`,
		b.Title, b.Subtitle, b.ImageURL, b.LinkURL, b.ButtonText, b.IsPublished)
	if err != nil {
		return Banner{}, err
	}
	return s.GetBanner(ctx)
}

func (s *mysqlStore) ListBanners(ctx context.Context, opts ListOptions) (ListResult[Banner], error) {
	hasPosition := s.columnExists(ctx, "banners", "position_key")
	hasSort := s.columnExists(ctx, "banners", "sort_order")
	where := ""
	args := []any{}
	if opts.Status != "" {
		where = "WHERE is_published = ?"
		args = append(args, opts.Status != "disabled" && opts.Status != "draft")
	}
	total, err := s.count(ctx, "banners", where, args)
	if err != nil {
		return ListResult[Banner]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT `+bannerSelectColumns(hasPosition, hasSort)+` FROM banners `+where+` `+bannerOrderClause(hasSort)+` LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[Banner]{}, err
	}
	defer rows.Close()
	items := []Banner{}
	for rows.Next() {
		var item Banner
		if err := rows.Scan(&item.ID, &item.Title, &item.Subtitle, &item.ImageURL, &item.LinkURL, &item.ButtonText, &item.Page, &item.Sort, &item.IsPublished, &item.UpdatedAt); err != nil {
			return ListResult[Banner]{}, err
		}
		items = append(items, normalizeBanner(item))
	}
	return ListResult[Banner]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetBannerByID(ctx context.Context, id int64) (Banner, error) {
	hasPosition := s.columnExists(ctx, "banners", "position_key")
	hasSort := s.columnExists(ctx, "banners", "sort_order")
	var item Banner
	err := s.db.QueryRowContext(ctx, `SELECT `+bannerSelectColumns(hasPosition, hasSort)+` FROM banners WHERE id = ?`, id).Scan(
		&item.ID, &item.Title, &item.Subtitle, &item.ImageURL, &item.LinkURL, &item.ButtonText, &item.Page, &item.Sort, &item.IsPublished, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Banner{}, notFound("banner not found")
	}
	return normalizeBanner(item), err
}

func (s *mysqlStore) CreateBanner(ctx context.Context, item Banner) (Banner, error) {
	item = normalizeBanner(item)
	id, err := s.nextID(ctx, "banners")
	if err != nil {
		return Banner{}, err
	}
	if s.columnExists(ctx, "banners", "position_key") && s.columnExists(ctx, "banners", "sort_order") {
		_, err = s.db.ExecContext(ctx, `INSERT INTO banners (id, title, subtitle, image_url, link_url, button_text, position_key, sort_order, is_published) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			id, item.Title, item.Subtitle, item.ImageURL, item.LinkURL, item.ButtonText, item.Page, item.Sort, item.IsPublished)
	} else {
		_, err = s.db.ExecContext(ctx, `INSERT INTO banners (id, title, subtitle, image_url, link_url, button_text, is_published) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			id, item.Title, item.Subtitle, item.ImageURL, item.LinkURL, item.ButtonText, item.IsPublished)
	}
	if err != nil {
		return Banner{}, err
	}
	return s.GetBannerByID(ctx, id)
}

func (s *mysqlStore) UpdateBanner(ctx context.Context, id int64, item Banner) (Banner, error) {
	item = normalizeBanner(item)
	var res sql.Result
	var err error
	if s.columnExists(ctx, "banners", "position_key") && s.columnExists(ctx, "banners", "sort_order") {
		res, err = s.db.ExecContext(ctx, `UPDATE banners SET title=?, subtitle=?, image_url=?, link_url=?, button_text=?, position_key=?, sort_order=?, is_published=? WHERE id=?`,
			item.Title, item.Subtitle, item.ImageURL, item.LinkURL, item.ButtonText, item.Page, item.Sort, item.IsPublished, id)
	} else {
		res, err = s.db.ExecContext(ctx, `UPDATE banners SET title=?, subtitle=?, image_url=?, link_url=?, button_text=?, is_published=? WHERE id=?`,
			item.Title, item.Subtitle, item.ImageURL, item.LinkURL, item.ButtonText, item.IsPublished, id)
	}
	if err := s.checkUpdateRows(ctx, res, err, "banners", id, "banner not found"); err != nil {
		return Banner{}, err
	}
	return s.GetBannerByID(ctx, id)
}

func (s *mysqlStore) DeleteBanner(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM banners WHERE id = ?`, id)
	return checkRows(res, err, "banner not found")
}

func (s *mysqlStore) GetSEO(ctx context.Context, page string) (SEOSetting, error) {
	var seo SEOSetting
	err := s.db.QueryRowContext(ctx, `SELECT id, page_key, title, description, keywords, updated_at FROM seo_settings WHERE page_key = ?`, page).Scan(
		&seo.ID, &seo.PageKey, &seo.Title, &seo.Description, &seo.Keywords, &seo.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return SEOSetting{}, notFound("SEO setting not found")
	}
	seo.Page = seo.PageKey
	return seo, err
}

func (s *mysqlStore) SaveSEO(ctx context.Context, seo SEOSetting) (SEOSetting, error) {
	_, err := s.db.ExecContext(ctx, `INSERT INTO seo_settings (page_key, title, description, keywords)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE title = VALUES(title), description = VALUES(description), keywords = VALUES(keywords)`,
		seo.PageKey, seo.Title, seo.Description, seo.Keywords)
	if err != nil {
		return SEOSetting{}, err
	}
	return s.GetSEO(ctx, seo.PageKey)
}

func (s *mysqlStore) ListSEO(ctx context.Context) ([]SEOSetting, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, page_key, title, description, keywords, updated_at FROM seo_settings ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SEOSetting{}
	for rows.Next() {
		var item SEOSetting
		if err := rows.Scan(&item.ID, &item.PageKey, &item.Title, &item.Description, &item.Keywords, &item.UpdatedAt); err != nil {
			return nil, err
		}
		item.Page = item.PageKey
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *mysqlStore) GetAdminUser(ctx context.Context, username string) (AdminUser, error) {
	var user AdminUser
	err := s.db.QueryRowContext(ctx, `SELECT id, username, password_hash, real_name, status FROM admin_users WHERE username = ? AND status = 'active'`, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.RealName, &user.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return AdminUser{}, notFound("admin user not found")
	}
	if err != nil {
		return AdminUser{}, err
	}
	roleRows, err := s.db.QueryContext(ctx, `SELECT r.code FROM roles r JOIN admin_user_roles aur ON aur.role_id = r.id WHERE aur.user_id = ? ORDER BY r.id`, user.ID)
	if err != nil {
		return AdminUser{}, err
	}
	defer roleRows.Close()
	for roleRows.Next() {
		var role string
		if err := roleRows.Scan(&role); err != nil {
			return AdminUser{}, err
		}
		user.Roles = append(user.Roles, role)
	}
	if err := roleRows.Err(); err != nil {
		return AdminUser{}, err
	}
	permRows, err := s.db.QueryContext(ctx, `SELECT DISTINCT p.code FROM permissions p JOIN role_permissions rp ON rp.permission_id = p.id JOIN admin_user_roles aur ON aur.role_id = rp.role_id WHERE aur.user_id = ? ORDER BY p.code`, user.ID)
	if err != nil {
		return AdminUser{}, err
	}
	defer permRows.Close()
	for permRows.Next() {
		var permission string
		if err := permRows.Scan(&permission); err != nil {
			return AdminUser{}, err
		}
		user.Permissions = append(user.Permissions, permission)
	}
	return user, permRows.Err()
}

func (s *mysqlStore) ListServices(ctx context.Context, opts ListOptions) (ListResult[Service], error) {
	where, args := serviceWhere(opts)
	total, err := s.count(ctx, "services", where, args)
	if err != nil {
		return ListResult[Service]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT `+serviceSelectColumns(s.columnExists(ctx, "services", "process_steps"))+` FROM services `+where+` ORDER BY sort_order ASC, id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[Service]{}, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var item Service
		if err := rows.Scan(&item.ID, &item.Title, &item.Slug, &item.Subtitle, &item.Summary, &item.CoverURL, &item.IconURL, &item.Content, &item.HighlightText, &item.ProcessText, &item.SortOrder, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[Service]{}, err
		}
		items = append(items, normalizeService(item))
	}
	return ListResult[Service]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetService(ctx context.Context, key string, includeDraft bool) (Service, error) {
	where := `(id = ? OR slug = ?)`
	args := []any{numericKey(key), key}
	if !includeDraft {
		where += ` AND status = ?`
		args = append(args, "published")
	}
	var item Service
	err := s.db.QueryRowContext(ctx, `SELECT `+serviceSelectColumns(s.columnExists(ctx, "services", "process_steps"))+` FROM services WHERE `+where+` LIMIT 1`, args...).Scan(
		&item.ID, &item.Title, &item.Slug, &item.Subtitle, &item.Summary, &item.CoverURL, &item.IconURL, &item.Content, &item.HighlightText, &item.ProcessText, &item.SortOrder, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Service{}, notFound("service not found")
	}
	return normalizeService(item), err
}

func (s *mysqlStore) CreateService(ctx context.Context, item Service) (Service, error) {
	item = normalizeService(item)
	res, err := s.db.ExecContext(ctx, `INSERT INTO services (title, slug, subtitle, summary, cover_url, icon_url, content, highlights, process_steps, sort_order, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		item.Title, item.Slug, item.Subtitle, item.Summary, item.CoverURL, item.IconURL, item.Content, item.HighlightText, item.ProcessText, item.SortOrder, item.Status)
	if err != nil {
		return Service{}, err
	}
	id, _ := res.LastInsertId()
	return s.GetService(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) UpdateService(ctx context.Context, id int64, item Service) (Service, error) {
	item = normalizeService(item)
	res, err := s.db.ExecContext(ctx, `UPDATE services SET title=?, slug=?, subtitle=?, summary=?, cover_url=?, icon_url=?, content=?, highlights=?, process_steps=?, sort_order=?, status=? WHERE id=?`,
		item.Title, item.Slug, item.Subtitle, item.Summary, item.CoverURL, item.IconURL, item.Content, item.HighlightText, item.ProcessText, item.SortOrder, item.Status, id)
	if err := s.checkUpdateRows(ctx, res, err, "services", id, "service not found"); err != nil {
		return Service{}, err
	}
	return s.GetService(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) DeleteService(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM services WHERE id = ?`, id)
	return checkRows(res, err, "service not found")
}

func (s *mysqlStore) ListCases(ctx context.Context, opts ListOptions) (ListResult[Case], error) {
	where, args := caseWhere(opts)
	total, err := s.count(ctx, "cases", where, args)
	if err != nil {
		return ListResult[Case]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT id, title, slug, industry, platform, strategy, cover_url, summary, content, core_metrics, status, created_at, updated_at FROM cases `+where+` ORDER BY sort_order ASC, id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[Case]{}, err
	}
	defer rows.Close()
	items := []Case{}
	for rows.Next() {
		var item Case
		if err := rows.Scan(&item.ID, &item.Title, &item.Slug, &item.Industry, &item.Platform, &item.Strategy, &item.CoverURL, &item.Summary, &item.Content, &item.CoreMetrics, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[Case]{}, err
		}
		items = append(items, item)
	}
	return ListResult[Case]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetCase(ctx context.Context, key string, includeDraft bool) (Case, error) {
	where := `(id = ? OR slug = ?)`
	args := []any{numericKey(key), key}
	if !includeDraft {
		where += ` AND status = ?`
		args = append(args, "published")
	}
	var item Case
	err := s.db.QueryRowContext(ctx, `SELECT id, title, slug, industry, platform, strategy, cover_url, summary, content, core_metrics, status, created_at, updated_at FROM cases WHERE `+where+` LIMIT 1`, args...).Scan(
		&item.ID, &item.Title, &item.Slug, &item.Industry, &item.Platform, &item.Strategy, &item.CoverURL, &item.Summary, &item.Content, &item.CoreMetrics, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Case{}, notFound("case not found")
	}
	return normalizeCase(item), err
}

func (s *mysqlStore) CreateCase(ctx context.Context, item Case) (Case, error) {
	res, err := s.db.ExecContext(ctx, `INSERT INTO cases (title, slug, industry, platform, strategy, cover_url, summary, content, core_metrics, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		item.Title, item.Slug, item.Industry, item.Platform, item.Strategy, item.CoverURL, item.Summary, item.Content, item.CoreMetrics, item.Status)
	if err != nil {
		return Case{}, err
	}
	id, _ := res.LastInsertId()
	return s.GetCase(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) UpdateCase(ctx context.Context, id int64, item Case) (Case, error) {
	res, err := s.db.ExecContext(ctx, `UPDATE cases SET title=?, slug=?, industry=?, platform=?, strategy=?, cover_url=?, summary=?, content=?, core_metrics=?, status=? WHERE id=?`,
		item.Title, item.Slug, item.Industry, item.Platform, item.Strategy, item.CoverURL, item.Summary, item.Content, item.CoreMetrics, item.Status, id)
	if err := s.checkUpdateRows(ctx, res, err, "cases", id, "case not found"); err != nil {
		return Case{}, err
	}
	return s.GetCase(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) DeleteCase(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM cases WHERE id = ?`, id)
	return checkRows(res, err, "case not found")
}

func (s *mysqlStore) ListNews(ctx context.Context, opts ListOptions) (ListResult[News], error) {
	where, args := newsWhere(opts)
	total, err := s.count(ctx, "news", where, args)
	if err != nil {
		return ListResult[News]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT id, title, slug, category, cover_url, summary, content, status, published_at, created_at, updated_at FROM news `+where+` ORDER BY published_at DESC, id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[News]{}, err
	}
	defer rows.Close()
	items := []News{}
	for rows.Next() {
		var item News
		if err := rows.Scan(&item.ID, &item.Title, &item.Slug, &item.Category, &item.CoverURL, &item.Summary, &item.Content, &item.Status, &item.PublishedAt, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[News]{}, err
		}
		items = append(items, normalizeNews(item))
	}
	return ListResult[News]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetNews(ctx context.Context, key string, includeDraft bool) (News, error) {
	where := `(id = ? OR slug = ?)`
	args := []any{numericKey(key), key}
	if !includeDraft {
		where += ` AND status = ?`
		args = append(args, "published")
	}
	var item News
	err := s.db.QueryRowContext(ctx, `SELECT id, title, slug, category, cover_url, summary, content, status, published_at, created_at, updated_at FROM news WHERE `+where+` LIMIT 1`, args...).Scan(
		&item.ID, &item.Title, &item.Slug, &item.Category, &item.CoverURL, &item.Summary, &item.Content, &item.Status, &item.PublishedAt, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return News{}, notFound("news not found")
	}
	return normalizeNews(item), err
}

func (s *mysqlStore) CreateNews(ctx context.Context, item News) (News, error) {
	res, err := s.db.ExecContext(ctx, `INSERT INTO news (title, slug, category, cover_url, summary, content, status, published_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		item.Title, item.Slug, item.Category, item.CoverURL, item.Summary, item.Content, item.Status, item.PublishedAt)
	if err != nil {
		return News{}, err
	}
	id, _ := res.LastInsertId()
	return s.GetNews(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) UpdateNews(ctx context.Context, id int64, item News) (News, error) {
	res, err := s.db.ExecContext(ctx, `UPDATE news SET title=?, slug=?, category=?, cover_url=?, summary=?, content=?, status=?, published_at=? WHERE id=?`,
		item.Title, item.Slug, item.Category, item.CoverURL, item.Summary, item.Content, item.Status, item.PublishedAt, id)
	if err := s.checkUpdateRows(ctx, res, err, "news", id, "news not found"); err != nil {
		return News{}, err
	}
	return s.GetNews(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) DeleteNews(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM news WHERE id = ?`, id)
	return checkRows(res, err, "news not found")
}

func (s *mysqlStore) CreateForm(ctx context.Context, form LeadForm) (LeadForm, error) {
	res, err := s.db.ExecContext(ctx, `INSERT INTO lead_forms (name, phone, company, position, email, requirement, interest, source, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		form.Name, form.Phone, form.Company, form.Position, form.Email, form.Requirement, form.Interest, form.Source, form.Status)
	if err != nil {
		return LeadForm{}, err
	}
	id, _ := res.LastInsertId()
	return s.getForm(ctx, id)
}

func (s *mysqlStore) ListForms(ctx context.Context, opts ListOptions) (ListResult[LeadForm], error) {
	where := ""
	args := []any{}
	if opts.Status != "" {
		where = "WHERE status = ?"
		args = append(args, opts.Status)
	}
	total, err := s.count(ctx, "lead_forms", where, args)
	if err != nil {
		return ListResult[LeadForm]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT id, name, phone, company, position, email, requirement, interest, source, status, created_at, updated_at FROM lead_forms `+where+` ORDER BY id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[LeadForm]{}, err
	}
	defer rows.Close()
	items := []LeadForm{}
	for rows.Next() {
		var item LeadForm
		if err := rows.Scan(&item.ID, &item.Name, &item.Phone, &item.Company, &item.Position, &item.Email, &item.Requirement, &item.Interest, &item.Source, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[LeadForm]{}, err
		}
		items = append(items, normalizeForm(item))
	}
	return ListResult[LeadForm]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) UpdateFormStatus(ctx context.Context, id int64, status string) (LeadForm, error) {
	res, err := s.db.ExecContext(ctx, `UPDATE lead_forms SET status = ? WHERE id = ?`, status, id)
	if err := s.checkUpdateRows(ctx, res, err, "lead_forms", id, "form not found"); err != nil {
		return LeadForm{}, err
	}
	return s.getForm(ctx, id)
}

func (s *mysqlStore) getForm(ctx context.Context, id int64) (LeadForm, error) {
	var item LeadForm
	err := s.db.QueryRowContext(ctx, `SELECT id, name, phone, company, position, email, requirement, interest, source, status, created_at, updated_at FROM lead_forms WHERE id = ?`, id).Scan(
		&item.ID, &item.Name, &item.Phone, &item.Company, &item.Position, &item.Email, &item.Requirement, &item.Interest, &item.Source, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return LeadForm{}, notFound("form not found")
	}
	return normalizeForm(item), err
}

func (s *mysqlStore) ListStaticPages(ctx context.Context, opts ListOptions) (ListResult[StaticPage], error) {
	where, args := pageWhere(opts)
	total, err := s.count(ctx, "static_pages", where, args)
	if err != nil {
		return ListResult[StaticPage]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT id, page_key, title, content, COALESCE(extra_data, ''), status, created_at, updated_at FROM static_pages `+where+` ORDER BY id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[StaticPage]{}, err
	}
	defer rows.Close()
	items := []StaticPage{}
	for rows.Next() {
		var item StaticPage
		if err := rows.Scan(&item.ID, &item.PageKey, &item.Title, &item.Content, &item.ExtraData, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[StaticPage]{}, err
		}
		items = append(items, normalizePage(item))
	}
	return ListResult[StaticPage]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetStaticPage(ctx context.Context, key string, includeDraft bool) (StaticPage, error) {
	where := `(id = ? OR page_key = ?)`
	args := []any{numericKey(key), key}
	if !includeDraft {
		where += ` AND status = ?`
		args = append(args, "published")
	}
	var item StaticPage
	err := s.db.QueryRowContext(ctx, `SELECT id, page_key, title, content, COALESCE(extra_data, ''), status, created_at, updated_at FROM static_pages WHERE `+where+` LIMIT 1`, args...).Scan(
		&item.ID, &item.PageKey, &item.Title, &item.Content, &item.ExtraData, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return StaticPage{}, notFound("page not found")
	}
	return normalizePage(item), err
}

func (s *mysqlStore) CreateStaticPage(ctx context.Context, item StaticPage) (StaticPage, error) {
	item = normalizePage(item)
	res, err := s.db.ExecContext(ctx, `INSERT INTO static_pages (page_key, title, content, extra_data, status) VALUES (?, ?, ?, ?, ?)`,
		item.PageKey, item.Title, item.Content, item.ExtraData, item.Status)
	if err != nil {
		return StaticPage{}, err
	}
	id, _ := res.LastInsertId()
	return s.GetStaticPage(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) UpdateStaticPage(ctx context.Context, id int64, item StaticPage) (StaticPage, error) {
	item = normalizePage(item)
	res, err := s.db.ExecContext(ctx, `UPDATE static_pages SET page_key=?, title=?, content=?, extra_data=?, status=? WHERE id=?`,
		item.PageKey, item.Title, item.Content, item.ExtraData, item.Status, id)
	if err := s.checkUpdateRows(ctx, res, err, "static_pages", id, "page not found"); err != nil {
		return StaticPage{}, err
	}
	return s.GetStaticPage(ctx, strconv.FormatInt(id, 10), true)
}

func (s *mysqlStore) DeleteStaticPage(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM static_pages WHERE id = ?`, id)
	return checkRows(res, err, "page not found")
}

func (s *mysqlStore) ListFAQs(ctx context.Context, opts ListOptions) (ListResult[PartnerFAQ], error) {
	where, args := faqWhere(opts)
	total, err := s.count(ctx, "partner_faqs", where, args)
	if err != nil {
		return ListResult[PartnerFAQ]{}, err
	}
	args = append(args, opts.PageSize, offset(opts))
	rows, err := s.db.QueryContext(ctx, `SELECT id, question, answer, sort_order, is_published, created_at, updated_at FROM partner_faqs `+where+` ORDER BY sort_order ASC, id DESC LIMIT ? OFFSET ?`, args...)
	if err != nil {
		return ListResult[PartnerFAQ]{}, err
	}
	defer rows.Close()
	items := []PartnerFAQ{}
	for rows.Next() {
		var item PartnerFAQ
		if err := rows.Scan(&item.ID, &item.Question, &item.Answer, &item.SortOrder, &item.IsPublished, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return ListResult[PartnerFAQ]{}, err
		}
		items = append(items, normalizeFAQ(item))
	}
	return ListResult[PartnerFAQ]{Items: items, Total: total, Page: opts.Page, PageSize: opts.PageSize}, rows.Err()
}

func (s *mysqlStore) GetFAQ(ctx context.Context, id int64) (PartnerFAQ, error) {
	var item PartnerFAQ
	err := s.db.QueryRowContext(ctx, `SELECT id, question, answer, sort_order, is_published, created_at, updated_at FROM partner_faqs WHERE id = ?`, id).Scan(
		&item.ID, &item.Question, &item.Answer, &item.SortOrder, &item.IsPublished, &item.CreatedAt, &item.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return PartnerFAQ{}, notFound("faq not found")
	}
	return normalizeFAQ(item), err
}

func (s *mysqlStore) CreateFAQ(ctx context.Context, item PartnerFAQ) (PartnerFAQ, error) {
	item = normalizeFAQ(item)
	res, err := s.db.ExecContext(ctx, `INSERT INTO partner_faqs (question, answer, sort_order, is_published) VALUES (?, ?, ?, ?)`,
		item.Question, item.Answer, item.SortOrder, item.IsPublished)
	if err != nil {
		return PartnerFAQ{}, err
	}
	id, _ := res.LastInsertId()
	return s.GetFAQ(ctx, id)
}

func (s *mysqlStore) UpdateFAQ(ctx context.Context, id int64, item PartnerFAQ) (PartnerFAQ, error) {
	item = normalizeFAQ(item)
	res, err := s.db.ExecContext(ctx, `UPDATE partner_faqs SET question=?, answer=?, sort_order=?, is_published=? WHERE id=?`,
		item.Question, item.Answer, item.SortOrder, item.IsPublished, id)
	if err := s.checkUpdateRows(ctx, res, err, "partner_faqs", id, "faq not found"); err != nil {
		return PartnerFAQ{}, err
	}
	return s.GetFAQ(ctx, id)
}

func (s *mysqlStore) DeleteFAQ(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM partner_faqs WHERE id = ?`, id)
	return checkRows(res, err, "faq not found")
}

func (s *mysqlStore) Dashboard(ctx context.Context) (DashboardStats, []OperationLog, error) {
	var stats DashboardStats
	countTable := func(table, where string, args ...any) (int64, error) {
		var total int64
		err := s.db.QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s %s", table, where), args...).Scan(&total)
		return total, err
	}
	var err error
	if stats.Cases, err = countTable("cases", ""); err != nil {
		return stats, nil, err
	}
	if stats.News, err = countTable("news", ""); err != nil {
		return stats, nil, err
	}
	if stats.Forms, err = countTable("lead_forms", ""); err != nil {
		return stats, nil, err
	}
	if stats.PendingForms, err = countTable("lead_forms", "WHERE status = ?", "new"); err != nil {
		return stats, nil, err
	}
	if stats.Services, err = countTable("services", ""); err != nil {
		return stats, nil, err
	}
	rows, err := s.db.QueryContext(ctx, `SELECT id, COALESCE(user_id, 0), username, action, module, COALESCE(target_id, 0), created_at FROM operation_logs ORDER BY id DESC LIMIT 20`)
	if err != nil {
		return stats, nil, err
	}
	defer rows.Close()
	logs := []OperationLog{}
	for rows.Next() {
		var item OperationLog
		if err := rows.Scan(&item.ID, &item.UserID, &item.Username, &item.Action, &item.Module, &item.TargetID, &item.CreatedAt); err != nil {
			return stats, nil, err
		}
		logs = append(logs, normalizeLog(item))
	}
	return stats, logs, rows.Err()
}

func (s *mysqlStore) LogOperation(ctx context.Context, log OperationLog) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO operation_logs (username, action, module, target_table, target_id, created_at) VALUES (?, ?, ?, ?, ?, NOW())`,
		log.Username, log.Action, log.Module, log.Target, log.TargetID)
	return err
}

func (s *mysqlStore) CreateMediaAsset(ctx context.Context, asset MediaAsset) (MediaAsset, error) {
	res, err := s.db.ExecContext(ctx, `INSERT INTO media_assets (original_name, file_name, file_path, file_url, mime_type, file_size, biz_type, ref_table, ref_id, uploader_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		asset.OriginalName, asset.FileName, asset.FilePath, asset.FileURL, asset.MimeType, asset.FileSize, asset.BizType, asset.RefTable, nullableID(asset.RefID), nullableID(asset.UploaderID))
	if err != nil {
		return MediaAsset{}, err
	}
	id, _ := res.LastInsertId()
	asset.ID = id
	return asset, nil
}

func (s *mysqlStore) GetEmailSetting(ctx context.Context) (EmailSetting, error) {
	var setting EmailSetting
	err := s.db.QueryRowContext(ctx, `SELECT id, smtp_host, smtp_port, smtp_user, smtp_password, sender_email, sender_name, recipients, is_ssl, is_enabled FROM email_settings ORDER BY id ASC LIMIT 1`).Scan(
		&setting.ID, &setting.SMTPHost, &setting.SMTPPort, &setting.SMTPUser, &setting.SMTPPassword, &setting.SenderEmail, &setting.SenderName, &setting.Recipients, &setting.IsSSL, &setting.IsEnabled)
	if errors.Is(err, sql.ErrNoRows) {
		return EmailSetting{}, notFound("email setting not found")
	}
	return setting, err
}

func (s *mysqlStore) CreateEmailNotification(ctx context.Context, notification EmailNotification) (EmailNotification, error) {
	res, err := s.db.ExecContext(ctx, `INSERT INTO email_notifications (lead_form_id, recipient, subject, content, status, error_message) VALUES (?, ?, ?, ?, ?, ?)`,
		nullableID(notification.LeadFormID), notification.Recipient, notification.Subject, notification.Content, notification.Status, nullableString(notification.ErrorMessage))
	if err != nil {
		return EmailNotification{}, err
	}
	id, _ := res.LastInsertId()
	notification.ID = id
	return notification, nil
}

func (s *mysqlStore) UpdateEmailNotificationStatus(ctx context.Context, id int64, status, message string) error {
	_, err := s.db.ExecContext(ctx, `UPDATE email_notifications SET status = ?, error_message = ?, sent_at = CASE WHEN ? = 'sent' THEN NOW() ELSE sent_at END WHERE id = ?`,
		status, nullableString(message), status, id)
	return err
}

func (s *mysqlStore) count(ctx context.Context, table, where string, args []any) (int64, error) {
	var total int64
	err := s.db.QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s %s", table, where), args...).Scan(&total)
	return total, err
}

func (s *mysqlStore) columnExists(ctx context.Context, table, column string) bool {
	var total int
	err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ? AND COLUMN_NAME = ?`, table, column).Scan(&total)
	return err == nil && total > 0
}

func serviceSelectColumns(hasProcessSteps bool) string {
	processSteps := "''"
	if hasProcessSteps {
		processSteps = "COALESCE(process_steps, '')"
	}
	return "id, title, slug, COALESCE(subtitle, ''), COALESCE(summary, ''), COALESCE(cover_url, ''), COALESCE(icon_url, ''), COALESCE(content, ''), COALESCE(highlights, ''), " + processSteps + ", sort_order, status, created_at, updated_at"
}

func bannerSelectColumns(hasPosition, hasSort bool) string {
	position := "''"
	if hasPosition {
		position = "position_key"
	}
	sortColumn := "0"
	if hasSort {
		sortColumn = "sort_order"
	}
	return "id, title, subtitle, image_url, link_url, button_text, " + position + ", " + sortColumn + ", is_published, updated_at"
}

func bannerOrderClause(hasSort bool) string {
	if hasSort {
		return "ORDER BY sort_order ASC, id ASC"
	}
	return "ORDER BY id ASC"
}

func (s *mysqlStore) nextID(ctx context.Context, table string) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, fmt.Sprintf("SELECT COALESCE(MAX(id), 0) + 1 FROM %s", table)).Scan(&id)
	return id, err
}

func nullableID(id int64) any {
	if id == 0 {
		return nil
	}
	return id
}

func nullableString(value string) any {
	if value == "" {
		return nil
	}
	return value
}

func caseWhere(opts ListOptions) (string, []any) {
	parts := []string{}
	args := []any{}
	addEq(&parts, &args, "status", opts.Status)
	addEq(&parts, &args, "industry", opts.Industry)
	addEq(&parts, &args, "platform", opts.Platform)
	addEq(&parts, &args, "strategy", opts.Strategy)
	if opts.Keyword != "" {
		parts = append(parts, "title LIKE ?")
		args = append(args, "%"+opts.Keyword+"%")
	}
	return whereClause(parts), args
}

func newsWhere(opts ListOptions) (string, []any) {
	parts := []string{}
	args := []any{}
	addEq(&parts, &args, "status", opts.Status)
	addEq(&parts, &args, "category", opts.Category)
	if opts.Keyword != "" {
		parts = append(parts, "title LIKE ?")
		args = append(args, "%"+opts.Keyword+"%")
	}
	return whereClause(parts), args
}

func serviceWhere(opts ListOptions) (string, []any) {
	parts := []string{}
	args := []any{}
	addEq(&parts, &args, "status", opts.Status)
	if opts.Keyword != "" {
		parts = append(parts, "(title LIKE ? OR summary LIKE ?)")
		args = append(args, "%"+opts.Keyword+"%", "%"+opts.Keyword+"%")
	}
	return whereClause(parts), args
}

func pageWhere(opts ListOptions) (string, []any) {
	parts := []string{}
	args := []any{}
	addEq(&parts, &args, "status", opts.Status)
	if opts.Keyword != "" {
		parts = append(parts, "(title LIKE ? OR page_key LIKE ?)")
		args = append(args, "%"+opts.Keyword+"%", "%"+opts.Keyword+"%")
	}
	return whereClause(parts), args
}

func faqWhere(opts ListOptions) (string, []any) {
	parts := []string{}
	args := []any{}
	switch opts.Status {
	case "published", "enabled":
		parts = append(parts, "is_published = ?")
		args = append(args, true)
	case "draft", "disabled":
		parts = append(parts, "is_published = ?")
		args = append(args, false)
	}
	if opts.Keyword != "" {
		parts = append(parts, "(question LIKE ? OR answer LIKE ?)")
		args = append(args, "%"+opts.Keyword+"%", "%"+opts.Keyword+"%")
	}
	return whereClause(parts), args
}

func addEq(parts *[]string, args *[]any, column, value string) {
	if value == "" {
		return
	}
	*parts = append(*parts, column+" = ?")
	*args = append(*args, value)
}

func whereClause(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(parts, " AND ")
}

func offset(opts ListOptions) int {
	if opts.Page <= 0 {
		opts.Page = 1
	}
	return (opts.Page - 1) * opts.PageSize
}

func numericKey(key string) int64 {
	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func checkRows(res sql.Result, err error, message string) error {
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return notFound(message)
	}
	return nil
}

func (s *mysqlStore) checkUpdateRows(ctx context.Context, res sql.Result, err error, table string, id int64, message string) error {
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}

	var total int
	if err := s.db.QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = ?", table), id).Scan(&total); err != nil {
		return err
	}
	if total == 0 {
		return notFound(message)
	}
	return nil
}
