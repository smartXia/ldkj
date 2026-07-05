const API_BASE = import.meta.env.VITE_ADMIN_API_BASE || '/api/admin'
const TOKEN_KEY = 'wsd_admin_token'
const SESSION_KEY = 'wsd_admin_session'

class AdminApiError extends Error {
  constructor(message, status, payload) {
    super(message)
    this.name = 'AdminApiError'
    this.status = status
    this.payload = payload
  }
}

function getToken() {
  return window.localStorage.getItem(TOKEN_KEY)
}

function setToken(token) {
  if (token) {
    window.localStorage.setItem(TOKEN_KEY, token)
    return
  }
  window.localStorage.removeItem(TOKEN_KEY)
}

function getSession() {
  const raw = window.localStorage.getItem(SESSION_KEY)
  if (!raw) return null
  try {
    return JSON.parse(raw)
  } catch {
    window.localStorage.removeItem(SESSION_KEY)
    return null
  }
}

function setSession(payload) {
  if (!payload) {
    window.localStorage.removeItem(SESSION_KEY)
    return
  }
  const user = payload.user || payload.data?.user || {}
  const session = {
    user,
    roles: payload.roles || user.roles || payload.data?.roles || [],
    permissions: payload.permissions || user.permissions || payload.data?.permissions || [],
  }
  window.localStorage.setItem(SESSION_KEY, JSON.stringify(session))
}

function clearSession() {
  setToken('')
  setSession(null)
}

function hasPermission(permission) {
  if (!permission) return true
  const session = getSession()
  const permissions = session?.permissions || []
  return permissions.includes('*') || permissions.includes(permission)
}

function buildUrl(path, params) {
  const url = new URL(`${API_BASE}${path}`, window.location.origin)
  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined && value !== null && value !== '') {
      url.searchParams.set(key, value)
    }
  })
  return url.toString()
}

async function parseResponse(response, expectJson = true) {
  const contentType = response.headers.get('content-type') || ''

  if (contentType.includes('application/json')) {
    const payload = await response.json()
    if (!response.ok) {
      throw new AdminApiError(payload.message || '请求失败', response.status, payload)
    }
    return payload
  }

  if (expectJson) {
    throw new AdminApiError('服务返回格式异常，请确认后端 API 已启动', response.status)
  }

  if (!response.ok) {
    throw new AdminApiError(response.statusText || '请求失败', response.status)
  }

  return response
}

async function request(path, options = {}) {
  const headers = new Headers(options.headers || {})
  const token = getToken()

  if (token) headers.set('Authorization', `Bearer ${token}`)
  if (options.body && !(options.body instanceof FormData)) {
    headers.set('Content-Type', 'application/json')
  }

  const response = await fetch(buildUrl(path, options.params), {
    ...options,
    headers,
    body: options.body instanceof FormData ? options.body : options.body ? JSON.stringify(options.body) : undefined,
  })

  if (response.status === 401) clearSession()
  return parseResponse(response, options.expectJson !== false)
}

function normalizeList(payload) {
  if (Array.isArray(payload)) {
    return { items: payload, total: payload.length }
  }
  return {
    items: payload.items || payload.list || payload.data || [],
    total: payload.total || payload.count || (payload.items || payload.list || payload.data || []).length,
  }
}

async function login(credentials) {
  const payload = await request('/login', {
    method: 'POST',
    body: credentials,
  })
  const token = payload.token || payload.data?.token
  setToken(token)
  setSession(payload)
  return payload
}

function logout() {
  clearSession()
}

function list(resource, params) {
  return request(`/${resource}`, { params }).then(normalizeList)
}

function detail(resource, id) {
  return request(`/${resource}/${id}`)
}

function create(resource, data) {
  return request(`/${resource}`, {
    method: 'POST',
    body: data,
  })
}

function update(resource, id, data) {
  return request(`/${resource}/${id}`, {
    method: 'PUT',
    body: data,
  })
}

function remove(resource, id) {
  return request(`/${resource}/${id}`, {
    method: 'DELETE',
  })
}

async function upload(file, type = 'common') {
  const formData = new FormData()
  formData.append('file', file)
  formData.append('type', type)
  return request('/upload', {
    method: 'POST',
    body: formData,
  })
}

async function exportFile(resource, params = {}, format = 'csv') {
  const response = await request(`/${resource}/export`, {
    params: { ...params, format },
    expectJson: false,
  })
  const blob = await response.blob()
  const downloadUrl = window.URL.createObjectURL(blob)
  const anchor = document.createElement('a')
  anchor.href = downloadUrl
  anchor.download = `${resource}-${new Date().toISOString().slice(0, 10)}.${format === 'excel' ? 'xlsx' : 'csv'}`
  document.body.appendChild(anchor)
  anchor.click()
  anchor.remove()
  window.URL.revokeObjectURL(downloadUrl)
}

export {
  AdminApiError,
  create,
  detail,
  exportFile,
  getSession,
  getToken,
  hasPermission,
  list,
  login,
  logout,
  remove,
  request,
  setSession,
  setToken,
  update,
  upload,
}

export default {
  create,
  detail,
  exportFile,
  getSession,
  getToken,
  hasPermission,
  list,
  login,
  logout,
  remove,
  request,
  setSession,
  setToken,
  update,
  upload,
}
