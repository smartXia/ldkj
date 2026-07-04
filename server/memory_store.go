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
	cases      map[int64]Case
	news       map[int64]News
	forms      map[int64]LeadForm
	nextCaseID int64
	nextNewsID int64
	nextFormID int64
}

func newMemoryStore() *memoryStore {
	return &memoryStore{
		site:       SiteConfig{ID: 1, SiteTitle: "Nanjing Lingdong Information Technology Website"},
		banner:     Banner{ID: 1, Title: "Lingdong Information", IsPublished: true},
		seo:        map[string]SEOSetting{"home": {ID: 1, PageKey: "home", Title: "Nanjing Lingdong Information Technology"}},
		cases:      make(map[int64]Case),
		news:       make(map[int64]News),
		forms:      make(map[int64]LeadForm),
		nextCaseID: 1,
		nextNewsID: 1,
		nextFormID: 1,
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
	return m.banner, nil
}

func (m *memoryStore) SaveBanner(_ context.Context, banner Banner) (Banner, error) {
	banner.ID = 1
	banner.UpdatedAt = nowString()
	m.banner = banner
	return banner, nil
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
	m.forms[id] = item
	return item, nil
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
