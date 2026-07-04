package main

import "context"

type Store interface {
	GetSiteConfig(context.Context) (SiteConfig, error)
	SaveSiteConfig(context.Context, SiteConfig) (SiteConfig, error)
	GetBanner(context.Context) (Banner, error)
	SaveBanner(context.Context, Banner) (Banner, error)
	GetSEO(context.Context, string) (SEOSetting, error)
	SaveSEO(context.Context, SEOSetting) (SEOSetting, error)
	ListSEO(context.Context) ([]SEOSetting, error)

	ListServices(context.Context, ListOptions) (ListResult[Service], error)
	GetService(context.Context, string, bool) (Service, error)
	CreateService(context.Context, Service) (Service, error)
	UpdateService(context.Context, int64, Service) (Service, error)
	DeleteService(context.Context, int64) error

	ListCases(context.Context, ListOptions) (ListResult[Case], error)
	GetCase(context.Context, string, bool) (Case, error)
	CreateCase(context.Context, Case) (Case, error)
	UpdateCase(context.Context, int64, Case) (Case, error)
	DeleteCase(context.Context, int64) error

	ListNews(context.Context, ListOptions) (ListResult[News], error)
	GetNews(context.Context, string, bool) (News, error)
	CreateNews(context.Context, News) (News, error)
	UpdateNews(context.Context, int64, News) (News, error)
	DeleteNews(context.Context, int64) error

	CreateForm(context.Context, LeadForm) (LeadForm, error)
	ListForms(context.Context, ListOptions) (ListResult[LeadForm], error)
	UpdateFormStatus(context.Context, int64, string) (LeadForm, error)

	ListStaticPages(context.Context, ListOptions) (ListResult[StaticPage], error)
	GetStaticPage(context.Context, string, bool) (StaticPage, error)
	CreateStaticPage(context.Context, StaticPage) (StaticPage, error)
	UpdateStaticPage(context.Context, int64, StaticPage) (StaticPage, error)
	DeleteStaticPage(context.Context, int64) error

	ListFAQs(context.Context, ListOptions) (ListResult[PartnerFAQ], error)
	GetFAQ(context.Context, int64) (PartnerFAQ, error)
	CreateFAQ(context.Context, PartnerFAQ) (PartnerFAQ, error)
	UpdateFAQ(context.Context, int64, PartnerFAQ) (PartnerFAQ, error)
	DeleteFAQ(context.Context, int64) error

	Dashboard(context.Context) (DashboardStats, []OperationLog, error)
	LogOperation(context.Context, OperationLog) error
}
