package main

import (
	"encoding/json"
	"strings"
)

func normalizeBanner(b Banner) Banner {
	if b.Image != "" {
		b.ImageURL = b.Image
	} else if b.ImageURL == "" {
		b.ImageURL = firstNonEmpty(b.Image, b.ImageURL)
	}
	if b.Image == "" {
		b.Image = b.ImageURL
	}
	if b.LinkURL == "" {
		b.LinkURL = b.Link
	}
	if b.Link == "" {
		b.Link = b.LinkURL
	}
	if b.Page == "" {
		b.Page = "home"
	}
	if b.Status == "" {
		if b.IsPublished {
			b.Status = "enabled"
		} else {
			b.Status = "disabled"
		}
	}
	b.IsPublished = b.Status != "disabled"
	b.UpdatedAtUI = b.UpdatedAt
	return b
}

func normalizeCase(item Case) Case {
	if item.Slug == "" {
		item.Slug = slugFromTitle(item.Title)
	}
	if item.Strategy == "" {
		item.Strategy = firstNonEmpty(item.Playbook, item.Method)
	}
	if item.Playbook == "" {
		item.Playbook = item.Strategy
	}
	if item.Method == "" {
		item.Method = item.Strategy
	}
	if item.CoverURL == "" {
		item.CoverURL = firstNonEmpty(item.Cover, item.Image)
	}
	if item.Cover == "" {
		item.Cover = item.CoverURL
	}
	if item.Image == "" {
		item.Image = item.CoverURL
	}
	if item.CoreMetrics == "" {
		item.CoreMetrics = item.Metrics
	}
	if item.Metrics == "" {
		item.Metrics = item.CoreMetrics
	}
	if item.Status == "" {
		item.Status = "draft"
	}
	item.CreatedAtUI = item.CreatedAt
	item.UpdatedAtUI = item.UpdatedAt
	return item
}

func normalizeNews(item News) News {
	if item.Slug == "" {
		item.Slug = slugFromTitle(item.Title)
	}
	if item.CoverURL == "" {
		item.CoverURL = firstNonEmpty(item.Cover, item.Image)
	}
	if item.Cover == "" {
		item.Cover = item.CoverURL
	}
	if item.Image == "" {
		item.Image = item.CoverURL
	}
	if item.Summary == "" {
		item.Summary = item.Desc
	}
	if item.Desc == "" {
		item.Desc = item.Summary
	}
	if item.PublishedAt == "" {
		item.PublishedAt = firstNonEmpty(item.PublishedAtUI, item.Date)
	}
	if item.PublishedAtUI == "" {
		item.PublishedAtUI = item.PublishedAt
	}
	if item.Date == "" {
		item.Date = item.PublishedAt
	}
	if item.Status == "" {
		item.Status = "draft"
	}
	item.CreatedAtUI = item.CreatedAt
	item.UpdatedAtUI = item.UpdatedAt
	return item
}

func normalizeForm(form LeadForm) LeadForm {
	if form.Status == "pending" {
		form.Status = "new"
	}
	form.CreatedAtUI = form.CreatedAt
	form.UpdatedAtUI = form.UpdatedAt
	return form
}

func normalizeService(item Service) Service {
	if item.Slug == "" {
		item.Slug = slugFromTitle(item.Title)
	}
	if item.CoverURL == "" {
		item.CoverURL = item.Image
	}
	if item.Image == "" {
		item.Image = item.CoverURL
	}
	if len(item.Highlights) == 0 {
		item.Highlights = splitList(item.HighlightText)
	}
	if item.HighlightText == "" {
		item.HighlightText = joinList(item.Highlights)
	}
	if len(item.Process) == 0 {
		item.Process = splitList(item.ProcessText)
	}
	if item.ProcessText == "" {
		item.ProcessText = joinList(item.Process)
	}
	if item.Status == "" {
		item.Status = "draft"
	}
	item.CreatedAtUI = item.CreatedAt
	item.UpdatedAtUI = item.UpdatedAt
	return item
}

func normalizePage(item StaticPage) StaticPage {
	if item.PageKey == "" {
		item.PageKey = item.Page
	}
	if item.Page == "" {
		item.Page = item.PageKey
	}
	if item.Status == "" {
		item.Status = "draft"
	}
	return item
}

func normalizeFAQ(item PartnerFAQ) PartnerFAQ {
	if item.SortOrder == 0 {
		item.SortOrder = item.Sort
	}
	if item.Sort == 0 {
		item.Sort = item.SortOrder
	}
	if item.Status == "" {
		if item.IsPublished {
			item.Status = "published"
		} else {
			item.Status = "draft"
		}
	}
	item.IsPublished = item.Status != "draft" && item.Status != "disabled"
	return item
}

func normalizeLog(item OperationLog) OperationLog {
	if item.User == "" {
		item.User = item.Username
	}
	if item.Target == "" {
		item.Target = item.Module
	}
	item.CreatedAtUI = item.CreatedAt
	return item
}

func splitList(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	var items []string
	if json.Unmarshal([]byte(raw), &items) == nil {
		return items
	}
	parts := strings.FieldsFunc(raw, func(r rune) bool {
		return r == '\n' || r == ',' || r == '，'
	})
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			items = append(items, part)
		}
	}
	return items
}

func joinList(items []string) string {
	if len(items) == 0 {
		return ""
	}
	data, err := json.Marshal(items)
	if err != nil {
		return strings.Join(items, "\n")
	}
	return string(data)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func slugFromTitle(title string) string {
	slug := strings.TrimSpace(strings.ToLower(title))
	slug = strings.ReplaceAll(slug, " ", "-")
	if slug == "" {
		return ""
	}
	return slug
}
