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
      { path: 'cases', name: 'admin-cases', component: () => import('./pages/CasesView.vue') },
      { path: 'news', name: 'admin-news', component: () => import('./pages/NewsView.vue') },
      { path: 'banners', name: 'admin-banners', component: () => import('./pages/BannersView.vue') },
      { path: 'forms', name: 'admin-forms', component: () => import('./pages/FormsView.vue') },
      { path: 'site', name: 'admin-site', component: () => import('./pages/SiteSettingsView.vue') },
      { path: 'seo', name: 'admin-seo', component: () => import('./pages/SeoSettingsView.vue') },
    ],
  },
]
