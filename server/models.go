package main

import "time"

type SiteConfig struct {
	ID          int64  `json:"id"`
	SiteTitle   string `json:"site_title"`
	LogoURL     string `json:"logo_url"`
	Contact     string `json:"contact"`
	Copyright   string `json:"copyright"`
	FooterLinks string `json:"footer_links"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type Banner struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	ImageURL    string `json:"image_url"`
	LinkURL     string `json:"link_url"`
	ButtonText  string `json:"button_text"`
	IsPublished bool   `json:"is_published"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type SEOSetting struct {
	ID          int64  `json:"id"`
	PageKey     string `json:"page_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type Case struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Industry    string `json:"industry"`
	Platform    string `json:"platform"`
	Strategy    string `json:"strategy"`
	CoverURL    string `json:"cover_url"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	CoreMetrics string `json:"core_metrics"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type News struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Category    string `json:"category"`
	CoverURL    string `json:"cover_url"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type LeadForm struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	Email       string `json:"email"`
	Requirement string `json:"requirement"`
	Interest    string `json:"interest"`
	Source      string `json:"source"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type ListOptions struct {
	Page     int
	PageSize int
	Status   string
	Industry string
	Platform string
	Strategy string
	Category string
	Keyword  string
}

type ListResult[T any] struct {
	Items    []T   `json:"items"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func nowString() string {
	return time.Now().Format(time.RFC3339)
}
