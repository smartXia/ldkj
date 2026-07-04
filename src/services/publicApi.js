const API_BASE = import.meta.env.VITE_PUBLIC_API_BASE || ''
const FRONTEND_ASSET_PREFIX = '/assets/'
const NEWS_CATEGORY_LABELS = {
  news: '灵动动态',
  insight: '营销洞察',
  industry: '行业资讯',
}
function endpoint(path) {
  return `${API_BASE}${path}`
}

function isFrontendAsset(value) {
  return typeof value === 'string' && value.trim().startsWith(FRONTEND_ASSET_PREFIX)
}

function backendImage(value) {
  const image = typeof value === 'string' ? value.trim() : ''
  if (!image) return ''
  if (/^https?:\/\//i.test(image) || image.startsWith('data:') || image.startsWith('blob:')) return image
  if (isFrontendAsset(image)) return image
  if (image.startsWith('/')) return API_BASE ? `${API_BASE}${image}` : image
  return image
}

function firstNonEmpty(...values) {
  return values.find((value) => typeof value === 'string' && value.trim())?.trim() || ''
}

function normalizeImageItem(item) {
  if (!item || typeof item !== 'object') return item
  const image = backendImage(item.image || item.image_url || item.cover_url || item.cover)
  const detailImage = backendImage(item.detailImage || item.detail_image || item.detail_url)
  const logo = backendImage(item.logo || item.logo_url)
  return {
    ...item,
    image,
    image_url: image,
    cover: image,
    cover_url: image,
    detailImage,
    logo,
    logo_url: logo,
  }
}

function normalizeImageList(items) {
  return Array.isArray(items) ? items.map(normalizeImageItem) : []
}

function normalizeHomeContent(content) {
  const data = content || {}
  const banner = normalizeImageItem(data.banner)
  const site = normalizeImageItem(data.site || data)
  const heroItems = data.heroSlides || data.banners || data.slides || (banner?.image ? [banner] : [])
  return {
    ...data,
    site,
    logo: site.logo,
    logo_url: site.logo,
    heroSlides: normalizeImageList(heroItems),
    marketingModules: normalizeImageList(data.marketingModules || data.marketing_modules),
    logoList: normalizeImageList(data.logoList || data.logos),
  }
}

function normalizeCaseListContent(content) {
  const data = content || {}
  return {
    ...data,
    hero: normalizeImageItem(data.hero),
    filters: data.filters || [],
    items: normalizeImageList(data.items),
  }
}

function normalizeArticleListContent(content) {
  const data = content || {}
  const items = normalizeArticleList(data.items)
  const categories = Array.isArray(data.categories) && data.categories.length
    ? data.categories
    : ['全部资讯', ...new Set(items.map((item) => item.category).filter(Boolean))]

  return {
    ...data,
    hero: normalizeImageItem(data.hero),
    categories,
    items,
  }
}

function normalizeArticleItem(item) {
  const normalized = normalizeImageItem(item)
  if (!normalized || typeof normalized !== 'object') return normalized

  const routeId = firstNonEmpty(normalized.slug, String(normalized.id || ''))
  const category = NEWS_CATEGORY_LABELS[normalized.category] || normalized.category
  const desc = firstNonEmpty(normalized.desc, normalized.summary)
  const date = firstNonEmpty(normalized.date, normalized.published_at, normalized.publishedAt)

  return {
    ...normalized,
    category,
    desc,
    summary: firstNonEmpty(normalized.summary, desc),
    date,
    published_at: date,
    publishedAt: date,
    href: routeId ? `/message/${routeId}` : '/message',
  }
}

function normalizeArticleList(items) {
  return Array.isArray(items) ? items.map(normalizeArticleItem) : []
}

const defaultConsultContent = {
  title: '预约营销咨询',
  description: '留下您的联系方式，获取一对一营销诊断',
  demandPlaceholder: '请选择营销诉求',
  demandOptions: ['整合营销', '新品推广', '电商大促', '社交种草', '账号运营', '口碑管理', '其他'],
  submitText: '立即预约',
  submittingText: '提交中...',
  successText: '预约信息已提交，我们会尽快与您联系。',
  errorText: '提交失败，请稍后再试或直接联系我们。',
}
function parseJSON(value) {
  if (!value) return {}
  if (typeof value === 'object') return value
  try {
    return JSON.parse(value)
  } catch {
    return {}
  }
}

function plainText(value) {
  return typeof value === 'string' ? value.replace(/<[^>]*>/g, '').trim() : ''
}

function normalizeConsultContent(payload) {
  const data = unwrapItem(payload, {})
  const page = data.page || data
  const extra = parseJSON(page.extra_data || page.extraData)
  const config = extra.consult || extra.marketingConsult || extra.marketing_consult || extra
  const demandOptions = Array.isArray(config.demandOptions)
    ? config.demandOptions
    : Array.isArray(config.demand_options)
      ? config.demand_options
      : defaultConsultContent.demandOptions

  return {
    ...defaultConsultContent,
    ...config,
    title: config.title || page.title || defaultConsultContent.title,
    description: config.description || config.desc || plainText(page.content) || defaultConsultContent.description,
    demandPlaceholder: config.demandPlaceholder || config.demand_placeholder || defaultConsultContent.demandPlaceholder,
    demandOptions,
    submitText: config.submitText || config.submit_text || defaultConsultContent.submitText,
    submittingText: config.submittingText || config.submitting_text || defaultConsultContent.submittingText,
    successText: config.successText || config.success_text || defaultConsultContent.successText,
    errorText: config.errorText || config.error_text || defaultConsultContent.errorText,
  }
}

async function request(path, options = {}) {
  const response = await fetch(endpoint(path), {
    headers: {
      Accept: 'application/json',
      ...(options.body ? { 'Content-Type': 'application/json' } : {}),
      ...options.headers,
    },
    ...options,
  })

  if (!response.ok) {
    throw new Error(`Request failed: ${response.status}`)
  }

  return response.status === 204 ? null : response.json()
}

function unwrapList(payload, fallback) {
  if (Array.isArray(payload)) return payload
  if (Array.isArray(payload?.data)) return payload.data
  if (Array.isArray(payload?.data?.items)) return payload.data.items
  if (Array.isArray(payload?.items)) return payload.items
  return fallback
}

function unwrapItem(payload, fallback) {
  if (payload?.data && !Array.isArray(payload.data)) return payload.data
  if (payload?.item) return payload.item
  return payload || fallback
}

function withFallback(promiseFactory, fallback) {
  return promiseFactory().catch(() => fallback)
}

export function getHomeContent() {
  return withFallback(
    async () => {
      const [sitePayload, bannerPayload] = await Promise.all([
        request('/api/public/site').catch(() => null),
        request('/api/public/banner').catch(() => null),
      ])
      const bannerItems = unwrapList(bannerPayload, [])
      return normalizeHomeContent({
        site: unwrapItem(sitePayload, {}),
        banner: bannerItems.length ? {} : unwrapItem(bannerPayload, {}),
        banners: bannerItems,
      })
    },
    normalizeHomeContent({})
  )
}

export function getServices() {
  return withFallback(
    async () => normalizeImageList(unwrapList(await request('/api/services'), [])),
    []
  )
}

export function getServiceById(id) {
  return withFallback(
    async () => normalizeImageItem(unwrapItem(await request(`/api/services/${encodeURIComponent(id)}`), {})),
    {}
  )
}

export function getCases() {
  const fallback = { hero: {}, filters: [], items: [] }

  return withFallback(async () => {
    const payload = await request('/api/public/cases')
    return normalizeCaseListContent({
      hero: payload?.hero || payload?.data?.hero || {},
      filters: payload?.filters || payload?.data?.filters || [],
      items: unwrapList(payload, []),
    })
  }, normalizeCaseListContent(fallback))
}

export function getCaseById(id) {
  return withFallback(
    async () => normalizeImageItem(unwrapItem(await request(`/api/public/cases/${encodeURIComponent(id)}`), {})),
    {}
  )
}

export function getArticles() {
  const fallback = { hero: {}, categories: [], items: [] }

  return withFallback(async () => {
    const payload = await request('/api/public/news')
    return normalizeArticleListContent({
      hero: payload?.hero || payload?.data?.hero || {},
      categories: payload?.categories || payload?.data?.categories || [],
      items: unwrapList(payload, []),
    })
  }, normalizeArticleListContent(fallback))
}

export function getArticleById(id) {
  return withFallback(
    async () => normalizeArticleItem(unwrapItem(await request(`/api/public/news/${encodeURIComponent(id)}`), {})),
    {}
  )
}

export function getPartnerContent() {
  return withFallback(
    async () => unwrapItem(await request('/api/partner'), { faqs: [] }),
    { faqs: [] }
  )
}

export function getConsultContent() {
  return withFallback(
    async () => normalizeConsultContent(await request('/api/partner')),
    defaultConsultContent
  )
}

export async function submitPublicForm(payload) {
  return request('/api/public/forms', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}
