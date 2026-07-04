const API_BASE = import.meta.env.VITE_PUBLIC_API_BASE || ''
const FRONTEND_ASSET_PREFIX = '/assets/'

function endpoint(path) {
  return `${API_BASE}${path}`
}

function isFrontendAsset(value) {
  return typeof value === 'string' && value.trim().startsWith(FRONTEND_ASSET_PREFIX)
}

function backendImage(value) {
  const image = typeof value === 'string' ? value.trim() : ''
  if (!image || isFrontendAsset(image)) return ''
  if (/^https?:\/\//i.test(image) || image.startsWith('data:') || image.startsWith('blob:')) return image
  if (image.startsWith('/')) return API_BASE ? `${API_BASE}${image}` : image
  return image
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
  return {
    ...data,
    hero: normalizeImageItem(data.hero),
    categories: data.categories || [],
    items: normalizeImageList(data.items),
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
      return normalizeHomeContent({
        site: unwrapItem(sitePayload, {}),
        banner: unwrapItem(bannerPayload, {}),
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
    async () => normalizeImageItem(unwrapItem(await request(`/api/public/news/${encodeURIComponent(id)}`), {})),
    {}
  )
}

export function getPartnerContent() {
  return withFallback(
    async () => unwrapItem(await request('/api/partner'), { faqs: [] }),
    { faqs: [] }
  )
}

export async function submitPublicForm(payload) {
  return request('/api/public/forms', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}
