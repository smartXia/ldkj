export const adminRoutes = [
  {
    path: '/admin/login',
    name: 'admin-login',
    component: () => import('./pages/LoginView.vue'),
  },
  {
    path: '/admin',
    component: () => import('./layouts/AdminLayout.vue'),
    children: [
      { path: '', name: 'admin-dashboard', component: () => import('./pages/DashboardView.vue') },
      { path: 'services', name: 'admin-services', component: () => import('./pages/ServicesView.vue'), meta: { permission: 'content:read' } },
      { path: 'cases', name: 'admin-cases', component: () => import('./pages/CasesView.vue'), meta: { permission: 'content:read' } },
      { path: 'news', name: 'admin-news', component: () => import('./pages/NewsView.vue'), meta: { permission: 'content:read' } },
      { path: 'banners', name: 'admin-banners', component: () => import('./pages/BannersView.vue'), meta: { permission: 'banner:manage' } },
      { path: 'pages', name: 'admin-pages', component: () => import('./pages/PagesView.vue'), meta: { permission: 'content:read' } },
      { path: 'faqs', name: 'admin-faqs', component: () => import('./pages/FAQsView.vue'), meta: { permission: 'content:read' } },
      { path: 'forms', name: 'admin-forms', component: () => import('./pages/FormsView.vue'), meta: { permission: 'form:read' } },
      { path: 'site', name: 'admin-site', component: () => import('./pages/SiteSettingsView.vue'), meta: { permission: 'settings:manage' } },
      { path: 'seo', name: 'admin-seo', component: () => import('./pages/SeoSettingsView.vue'), meta: { permission: 'settings:manage' } },
      { path: 'users', name: 'admin-users', component: () => import('./pages/UsersView.vue'), meta: { permission: 'user:manage' } },
    ],
  },
]
