package main

import (
	"context"
	"strconv"
	"strings"
)

type memoryStore struct {
	site       SiteConfig
	banner     Banner
	seo        map[string]SEOSetting
	services   map[int64]Service
	cases      map[int64]Case
	news       map[int64]News
	forms      map[int64]LeadForm
	pages      map[int64]StaticPage
	faqs       map[int64]PartnerFAQ
	logs       []OperationLog
	nextServiceID int64
	nextCaseID int64
	nextNewsID int64
	nextFormID int64
	nextPageID int64
	nextFAQID  int64
	nextLogID  int64
}

func newMemoryStore() *memoryStore {
	return &memoryStore{
		site:       SiteConfig{ID: 1, SiteTitle: "Nanjing Lingdong Information Technology Website"},
		banner:     Banner{ID: 1, Title: "Lingdong Information", IsPublished: true},
		seo:        map[string]SEOSetting{"home": {ID: 1, PageKey: "home", Title: "Nanjing Lingdong Information Technology"}},
		services:   make(map[int64]Service),
		cases:      make(map[int64]Case),
		news:       make(map[int64]News),
		forms:      make(map[int64]LeadForm),
		pages:      make(map[int64]StaticPage),
		faqs:       make(map[int64]PartnerFAQ),
		logs:       []OperationLog{},
		nextServiceID: 1,
		nextCaseID: 1,
		nextNewsID: 1,
		nextFormID: 1,
		nextPageID: 1,
		nextFAQID:  1,
		nextLogID:  1,
	}
}

func (m *memoryStore) GetSiteConfig(context.Context) (SiteConfig, error) {
	return m.site, nil
}

func (m *memoryStore) SaveSiteConfig(_ context.Context, site SiteConfig) (SiteConfig, error) {
	site.ID = 1
	site.UpdatedAt = nowString()
	m.site = site
	return site, nil
}

func (m *memoryStore) GetBanner(context.Context) (Banner, error) {
	return normalizeBanner(m.banner), nil
}

func (m *memoryStore) SaveBanner(_ context.Context, banner Banner) (Banner, error) {
	banner.ID = 1
	banner.UpdatedAt = nowString()
	m.banner = normalizeBanner(banner)
	return m.banner, nil
}

func (m *memoryStore) GetSEO(_ context.Context, page string) (SEOSetting, error) {
	seo, ok := m.seo[page]
	if !ok {
		return SEOSetting{}, notFound("SEO setting not found")
	}
	return seo, nil
}

func (m *memoryStore) SaveSEO(_ context.Context, seo SEOSetting) (SEOSetting, error) {
	if seo.ID == 0 {
		seo.ID = int64(len(m.seo) + 1)
	}
	seo.UpdatedAt = nowString()
	m.seo[seo.PageKey] = seo
	return seo, nil
}

func (m *memoryStore) ListSEO(context.Context) ([]SEOSetting, error) {
	items := make([]SEOSetting, 0, len(m.seo))
	for _, item := range m.seo {
		items = append(items, item)
	}
	return items, nil
}

func (m *memoryStore) ListServices(_ context.Context, opts ListOptions) (ListResult[Service], error) {
	items := make([]Service, 0, len(m.services))
	for _, item := range m.services {
		if matchesService(item, opts) {
			items = append(items, normalizeService(item))
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) GetService(_ context.Context, key string, includeDraft bool) (Service, error) {
	for _, item := range m.services {
		if strconv.FormatInt(item.ID, 10) == key || item.Slug == key {
			if !includeDraft && item.Status != "published" {
				return Service{}, notFound("service not found")
			}
			return normalizeService(item), nil
		}
	}
	return Service{}, notFound("service not found")
}

func (m *memoryStore) CreateService(_ context.Context, item Service) (Service, error) {
	item.ID = m.nextServiceID
	m.nextServiceID++
	item.CreatedAt = nowString()
	item.UpdatedAt = item.CreatedAt
	item = normalizeService(item)
	m.services[item.ID] = item
	return item, nil
}

func (m *memoryStore) UpdateService(_ context.Context, id int64, item Service) (Service, error) {
	old, ok := m.services[id]
	if !ok {
		return Service{}, notFound("service not found")
	}
	item.ID = id
	item.CreatedAt = old.CreatedAt
	item.UpdatedAt = nowString()
	item = normalizeService(item)
	m.services[id] = item
	return item, nil
}

func (m *memoryStore) DeleteService(_ context.Context, id int64) error {
	if _, ok := m.services[id]; !ok {
		return notFound("service not found")
	}
	delete(m.services, id)
	return nil
}

func (m *memoryStore) ListCases(_ context.Context, opts ListOptions) (ListResult[Case], error) {
	items := make([]Case, 0, len(m.cases))
	for _, item := range m.cases {
		if matchesCase(item, opts) {
			items = append(items, item)
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) GetCase(_ context.Context, key string, includeDraft bool) (Case, error) {
	for _, item := range m.cases {
		if strconv.FormatInt(item.ID, 10) == key || item.Slug == key {
			if !includeDraft && item.Status != "published" {
				return Case{}, notFound("case not found")
			}
			return item, nil
		}
	}
	return Case{}, notFound("case not found")
}

func (m *memoryStore) CreateCase(_ context.Context, item Case) (Case, error) {
	item.ID = m.nextCaseID
	m.nextCaseID++
	item.CreatedAt = nowString()
	item.UpdatedAt = item.CreatedAt
	item = normalizeCase(item)
	m.cases[item.ID] = item
	return item, nil
}

func (m *memoryStore) UpdateCase(_ context.Context, id int64, item Case) (Case, error) {
	old, ok := m.cases[id]
	if !ok {
		return Case{}, notFound("case not found")
	}
	item.ID = id
	item.CreatedAt = old.CreatedAt
	item.UpdatedAt = nowString()
	item = normalizeCase(item)
	m.cases[id] = item
	return item, nil
}

func (m *memoryStore) DeleteCase(_ context.Context, id int64) error {
	if _, ok := m.cases[id]; !ok {
		return notFound("case not found")
	}
	delete(m.cases, id)
	return nil
}

func (m *memoryStore) ListNews(_ context.Context, opts ListOptions) (ListResult[News], error) {
	items := make([]News, 0, len(m.news))
	for _, item := range m.news {
		if matchesNews(item, opts) {
			items = append(items, item)
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) GetNews(_ context.Context, key string, includeDraft bool) (News, error) {
	for _, item := range m.news {
		if strconv.FormatInt(item.ID, 10) == key || item.Slug == key {
			if !includeDraft && item.Status != "published" {
				return News{}, notFound("news not found")
			}
			return item, nil
		}
	}
	return News{}, notFound("news not found")
}

func (m *memoryStore) CreateNews(_ context.Context, item News) (News, error) {
	item.ID = m.nextNewsID
	m.nextNewsID++
	item.CreatedAt = nowString()
	item.UpdatedAt = item.CreatedAt
	item = normalizeNews(item)
	m.news[item.ID] = item
	return item, nil
}

func (m *memoryStore) UpdateNews(_ context.Context, id int64, item News) (News, error) {
	old, ok := m.news[id]
	if !ok {
		return News{}, notFound("news not found")
	}
	item.ID = id
	item.CreatedAt = old.CreatedAt
	item.UpdatedAt = nowString()
	item = normalizeNews(item)
	m.news[id] = item
	return item, nil
}

func (m *memoryStore) DeleteNews(_ context.Context, id int64) error {
	if _, ok := m.news[id]; !ok {
		return notFound("news not found")
	}
	delete(m.news, id)
	return nil
}

func (m *memoryStore) CreateForm(_ context.Context, form LeadForm) (LeadForm, error) {
	form.ID = m.nextFormID
	m.nextFormID++
	form.CreatedAt = nowString()
	form.UpdatedAt = form.CreatedAt
	form = normalizeForm(form)
	m.forms[form.ID] = form
	return form, nil
}

func (m *memoryStore) ListForms(_ context.Context, opts ListOptions) (ListResult[LeadForm], error) {
	items := make([]LeadForm, 0, len(m.forms))
	for _, item := range m.forms {
		if opts.Status == "" || item.Status == opts.Status {
			items = append(items, item)
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) UpdateFormStatus(_ context.Context, id int64, status string) (LeadForm, error) {
	item, ok := m.forms[id]
	if !ok {
		return LeadForm{}, notFound("form not found")
	}
	item.Status = status
	item.UpdatedAt = nowString()
	m.forms[id] = normalizeForm(item)
	return item, nil
}

func (m *memoryStore) ListStaticPages(_ context.Context, opts ListOptions) (ListResult[StaticPage], error) {
	items := make([]StaticPage, 0, len(m.pages))
	for _, item := range m.pages {
		if matchesPage(item, opts) {
			items = append(items, normalizePage(item))
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) GetStaticPage(_ context.Context, key string, includeDraft bool) (StaticPage, error) {
	for _, item := range m.pages {
		if strconv.FormatInt(item.ID, 10) == key || item.PageKey == key || item.Page == key {
			if !includeDraft && item.Status != "published" {
				return StaticPage{}, notFound("page not found")
			}
			return normalizePage(item), nil
		}
	}
	return StaticPage{}, notFound("page not found")
}

func (m *memoryStore) CreateStaticPage(_ context.Context, item StaticPage) (StaticPage, error) {
	item.ID = m.nextPageID
	m.nextPageID++
	item.CreatedAt = nowString()
	item.UpdatedAt = item.CreatedAt
	item = normalizePage(item)
	m.pages[item.ID] = item
	return item, nil
}

func (m *memoryStore) UpdateStaticPage(_ context.Context, id int64, item StaticPage) (StaticPage, error) {
	old, ok := m.pages[id]
	if !ok {
		return StaticPage{}, notFound("page not found")
	}
	item.ID = id
	item.CreatedAt = old.CreatedAt
	item.UpdatedAt = nowString()
	item = normalizePage(item)
	m.pages[id] = item
	return item, nil
}

func (m *memoryStore) DeleteStaticPage(_ context.Context, id int64) error {
	if _, ok := m.pages[id]; !ok {
		return notFound("page not found")
	}
	delete(m.pages, id)
	return nil
}

func (m *memoryStore) ListFAQs(_ context.Context, opts ListOptions) (ListResult[PartnerFAQ], error) {
	items := make([]PartnerFAQ, 0, len(m.faqs))
	for _, item := range m.faqs {
		if opts.Status == "" || (opts.Status == "published" && item.IsPublished) || (opts.Status == "draft" && !item.IsPublished) {
			items = append(items, normalizeFAQ(item))
		}
	}
	return paginate(items, opts), nil
}

func (m *memoryStore) GetFAQ(_ context.Context, id int64) (PartnerFAQ, error) {
	item, ok := m.faqs[id]
	if !ok {
		return PartnerFAQ{}, notFound("faq not found")
	}
	return normalizeFAQ(item), nil
}

func (m *memoryStore) CreateFAQ(_ context.Context, item PartnerFAQ) (PartnerFAQ, error) {
	item.ID = m.nextFAQID
	m.nextFAQID++
	item.CreatedAt = nowString()
	item.UpdatedAt = item.CreatedAt
	item = normalizeFAQ(item)
	m.faqs[item.ID] = item
	return item, nil
}

func (m *memoryStore) UpdateFAQ(_ context.Context, id int64, item PartnerFAQ) (PartnerFAQ, error) {
	old, ok := m.faqs[id]
	if !ok {
		return PartnerFAQ{}, notFound("faq not found")
	}
	item.ID = id
	item.CreatedAt = old.CreatedAt
	item.UpdatedAt = nowString()
	item = normalizeFAQ(item)
	m.faqs[id] = item
	return item, nil
}

func (m *memoryStore) DeleteFAQ(_ context.Context, id int64) error {
	if _, ok := m.faqs[id]; !ok {
		return notFound("faq not found")
	}
	delete(m.faqs, id)
	return nil
}

func (m *memoryStore) Dashboard(context.Context) (DashboardStats, []OperationLog, error) {
	pending := int64(0)
	for _, form := range m.forms {
		if form.Status == "new" || form.Status == "pending" {
			pending++
		}
	}
	stats := DashboardStats{
		Cases:        int64(len(m.cases)),
		News:         int64(len(m.news)),
		Forms:        int64(len(m.forms)),
		PendingForms: pending,
		Services:     int64(len(m.services)),
	}
	logs := make([]OperationLog, 0, len(m.logs))
	for _, item := range m.logs {
		logs = append(logs, normalizeLog(item))
	}
	return stats, logs, nil
}

func (m *memoryStore) LogOperation(_ context.Context, log OperationLog) error {
	log.ID = m.nextLogID
	m.nextLogID++
	if log.CreatedAt == "" {
		log.CreatedAt = nowString()
	}
	m.logs = append([]OperationLog{normalizeLog(log)}, m.logs...)
	if len(m.logs) > 100 {
		m.logs = m.logs[:100]
	}
	return nil
}

func matchesCase(item Case, opts ListOptions) bool {
	return (opts.Status == "" || item.Status == opts.Status) &&
		(opts.Industry == "" || item.Industry == opts.Industry) &&
		(opts.Platform == "" || item.Platform == opts.Platform) &&
		(opts.Strategy == "" || item.Strategy == opts.Strategy) &&
		(opts.Keyword == "" || strings.Contains(item.Title, opts.Keyword))
}

func matchesNews(item News, opts ListOptions) bool {
	return (opts.Status == "" || item.Status == opts.Status) &&
		(opts.Category == "" || item.Category == opts.Category) &&
		(opts.Keyword == "" || strings.Contains(item.Title, opts.Keyword))
}

func matchesService(item Service, opts ListOptions) bool {
	return (opts.Status == "" || item.Status == opts.Status) &&
		(opts.Keyword == "" || strings.Contains(item.Title, opts.Keyword) || strings.Contains(item.Summary, opts.Keyword))
}

func matchesPage(item StaticPage, opts ListOptions) bool {
	return (opts.Status == "" || item.Status == opts.Status) &&
		(opts.Keyword == "" || strings.Contains(item.Title, opts.Keyword) || strings.Contains(item.PageKey, opts.Keyword))
}

func paginate[T any](items []T, opts ListOptions) ListResult[T] {
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 {
		opts.PageSize = 10
	}
	total := len(items)
	start := (opts.Page - 1) * opts.PageSize
	if start > total {
		start = total
	}
	end := start + opts.PageSize
	if end > total {
		end = total
	}
	return ListResult[T]{Items: items[start:end], Total: int64(total), Page: opts.Page, PageSize: opts.PageSize}
}
