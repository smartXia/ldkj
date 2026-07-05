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
	Image       string `json:"image,omitempty"`
	LinkURL     string `json:"link_url"`
	Link        string `json:"link,omitempty"`
	ButtonText  string `json:"button_text"`
	Page        string `json:"page,omitempty"`
	Sort        int    `json:"sort,omitempty"`
	Status      string `json:"status,omitempty"`
	IsPublished bool   `json:"is_published"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedAtUI string `json:"updatedAt,omitempty"`
}

type SEOSetting struct {
	ID          int64  `json:"id"`
	PageKey     string `json:"page_key"`
	Page        string `json:"page,omitempty"`
	Name        string `json:"name,omitempty"`
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
	Playbook    string `json:"playbook,omitempty"`
	Method      string `json:"method,omitempty"`
	CoverURL    string `json:"cover_url"`
	Cover       string `json:"cover,omitempty"`
	Image       string `json:"image,omitempty"`
	Summary     string `json:"summary"`
	Content     string `json:"content"`
	CoreMetrics string `json:"core_metrics"`
	Metrics     string `json:"metrics,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreatedAtUI string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedAtUI string `json:"updatedAt,omitempty"`
}

type News struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Category      string `json:"category"`
	CoverURL      string `json:"cover_url"`
	Cover         string `json:"cover,omitempty"`
	Image         string `json:"image,omitempty"`
	Summary       string `json:"summary"`
	Desc          string `json:"desc,omitempty"`
	Content       string `json:"content"`
	Status        string `json:"status"`
	PublishedAt   string `json:"published_at"`
	PublishedAtUI string `json:"publishedAt,omitempty"`
	Date          string `json:"date,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatedAtUI   string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	UpdatedAtUI   string `json:"updatedAt,omitempty"`
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
	CreatedAtUI string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedAtUI string `json:"updatedAt,omitempty"`
}

type Service struct {
	ID            int64    `json:"id"`
	Slug          string   `json:"slug"`
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle"`
	Summary       string   `json:"summary"`
	CoverURL      string   `json:"cover_url"`
	Image         string   `json:"image,omitempty"`
	IconURL       string   `json:"icon_url"`
	Content       string   `json:"content"`
	Highlights    []string `json:"highlights,omitempty"`
	HighlightText string   `json:"highlight_text,omitempty"`
	Process       []string `json:"process,omitempty"`
	ProcessText   string   `json:"process_text,omitempty"`
	SortOrder     int      `json:"sort_order"`
	Status        string   `json:"status"`
	CreatedAt     string   `json:"created_at,omitempty"`
	CreatedAtUI   string   `json:"createdAt,omitempty"`
	UpdatedAt     string   `json:"updated_at,omitempty"`
	UpdatedAtUI   string   `json:"updatedAt,omitempty"`
}

type StaticPage struct {
	ID        int64  `json:"id"`
	PageKey   string `json:"page_key"`
	Page      string `json:"page,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ExtraData string `json:"extra_data"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type PartnerFAQ struct {
	ID          int64  `json:"id"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	SortOrder   int    `json:"sort_order"`
	Sort        int    `json:"sort,omitempty"`
	IsPublished bool   `json:"is_published"`
	Status      string `json:"status,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type DashboardStats struct {
	Cases        int64 `json:"cases"`
	News         int64 `json:"news"`
	Forms        int64 `json:"forms"`
	PendingForms int64 `json:"pendingForms"`
	Services     int64 `json:"services"`
}

type OperationLog struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id,omitempty"`
	Username    string `json:"username"`
	User        string `json:"user,omitempty"`
	Action      string `json:"action"`
	Module      string `json:"module"`
	Target      string `json:"target,omitempty"`
	TargetID    int64  `json:"target_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreatedAtUI string `json:"createdAt,omitempty"`
}

type AdminUser struct {
	ID           int64    `json:"id"`
	Username     string   `json:"username"`
	PasswordHash string   `json:"-"`
	RealName     string   `json:"real_name,omitempty"`
	Email        string   `json:"email,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	Status       string   `json:"status"`
	Roles        []string `json:"roles,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	UpdatedAt    string   `json:"updated_at,omitempty"`
}

type MediaAsset struct {
	ID           int64  `json:"id"`
	OriginalName string `json:"original_name"`
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	FileURL      string `json:"file_url"`
	MimeType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
	BizType      string `json:"biz_type"`
	RefTable     string `json:"ref_table,omitempty"`
	RefID        int64  `json:"ref_id,omitempty"`
	UploaderID   int64  `json:"uploader_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type EmailSetting struct {
	ID           int64  `json:"id"`
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     int    `json:"smtp_port"`
	SMTPUser     string `json:"smtp_user"`
	SMTPPassword string `json:"-"`
	SenderEmail  string `json:"sender_email"`
	SenderName   string `json:"sender_name"`
	Recipients   string `json:"recipients"`
	IsSSL        bool   `json:"is_ssl"`
	IsEnabled    bool   `json:"is_enabled"`
}

type EmailNotification struct {
	ID           int64  `json:"id"`
	LeadFormID   int64  `json:"lead_form_id"`
	Recipient    string `json:"recipient"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
	SentAt       string `json:"sent_at,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
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
