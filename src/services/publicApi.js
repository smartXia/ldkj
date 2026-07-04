import { caseCards, caseFilterGroups, caseHero } from '../data/caseContent'
import { messageArticles, messageCategories, messageHero } from '../data/messageContent'
import { heroSlides, logoList } from '../data/siteContent'
import { partnerFaqs, publicServices } from '../data/publicContent'

const API_BASE = import.meta.env.VITE_PUBLIC_API_BASE || ''

function endpoint(path) {
  return `${API_BASE}${path}`
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
    async () => unwrapItem(await request('/api/public/site'), { heroSlides, logoList }),
    { heroSlides, logoList }
  )
}

export function getServices() {
  return withFallback(
    async () => unwrapList(await request('/api/services'), publicServices),
    publicServices
  )
}

export function getServiceById(id) {
  const fallback = publicServices.find((item) => item.id === id) || publicServices[0]

  return withFallback(
    async () => unwrapItem(await request(`/api/services/${encodeURIComponent(id)}`), fallback),
    fallback
  )
}

export function getCases() {
  const fallback = { hero: caseHero, filters: caseFilterGroups, items: caseCards }

  return withFallback(async () => {
    const payload = await request('/api/public/cases')
    return {
      hero: payload?.hero || payload?.data?.hero || caseHero,
      filters: payload?.filters || payload?.data?.filters || caseFilterGroups,
      items: unwrapList(payload, caseCards),
    }
  }, fallback)
}

export function getCaseById(id) {
  const fallback = caseCards.find((item) => item.id === id) || caseCards[0]

  return withFallback(
    async () => unwrapItem(await request(`/api/public/cases/${encodeURIComponent(id)}`), fallback),
    fallback
  )
}

export function getArticles() {
  const fallback = { hero: messageHero, categories: messageCategories, items: messageArticles }

  return withFallback(async () => {
    const payload = await request('/api/public/news')
    return {
      hero: payload?.hero || payload?.data?.hero || messageHero,
      categories: payload?.categories || payload?.data?.categories || messageCategories,
      items: unwrapList(payload, messageArticles),
    }
  }, fallback)
}

export function getArticleById(id) {
  const fallback = messageArticles.find((item) => item.id === id) || messageArticles[0]

  return withFallback(
    async () => unwrapItem(await request(`/api/public/news/${encodeURIComponent(id)}`), fallback),
    fallback
  )
}

export function getPartnerContent() {
  return withFallback(
    async () => unwrapItem(await request('/api/partner'), { faqs: partnerFaqs }),
    { faqs: partnerFaqs }
  )
}

export async function submitPublicForm(payload) {
  return request('/api/public/forms', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}
