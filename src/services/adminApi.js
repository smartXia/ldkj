const TOKEN_KEY = 'wsd_admin_token'

export function getAdminToken() {
  return sessionStorage.getItem(TOKEN_KEY)
}

export function setAdminToken(token) {
  if (token) {
    sessionStorage.setItem(TOKEN_KEY, token)
  } else {
    sessionStorage.removeItem(TOKEN_KEY)
  }
}

export function isAdminAuthenticated() {
  return Boolean(getAdminToken())
}

async function parseResponse(response) {
  const contentType = response.headers.get('content-type') || ''
  const payload = contentType.includes('application/json') ? await response.json() : await response.text()

  if (!response.ok) {
    const message = payload?.message || payload?.code || response.statusText || '请求失败'
    const error = new Error(message)
    error.status = response.status
    throw error
  }

  return payload
}

export async function adminRequest(path, options = {}) {
  const headers = new Headers(options.headers || {})
  const controller = new AbortController()
  const timeout = window.setTimeout(() => controller.abort(), options.timeoutMs || 10000)

  if (!(options.body instanceof FormData) && !headers.has('Content-Type')) {
    headers.set('Content-Type', 'application/json')
  }

  const token = getAdminToken()
  if (token) {
    headers.set('Authorization', `Bearer ${token}`)
  }

  try {
    const response = await fetch(path, {
      ...options,
      headers,
      signal: controller.signal,
    })

    return parseResponse(response)
  } catch (error) {
    if (error.name === 'AbortError') {
      throw new Error('请求超时，请检查后端服务或数据库连接')
    }
    throw error
  } finally {
    window.clearTimeout(timeout)
  }
}

export async function loginAdmin(credentials) {
  const result = await adminRequest('/api/admin/login', {
    method: 'POST',
    body: JSON.stringify(credentials),
  })
  setAdminToken(result.token)
  return result
}

export function logoutAdmin() {
  setAdminToken('')
}

export function listResource(resource, params = {}) {
  const query = new URLSearchParams()
  for (const [key, value] of Object.entries(params)) {
    if (value !== undefined && value !== null && value !== '') {
      query.set(key, value)
    }
  }
  const suffix = query.toString() ? `?${query}` : ''
  return adminRequest(`/api/admin/${resource}${suffix}`)
}

export function createResource(resource, payload) {
  return adminRequest(`/api/admin/${resource}`, {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export function updateResource(resource, id, payload) {
  return adminRequest(`/api/admin/${resource}/${id}`, {
    method: 'PUT',
    body: JSON.stringify(payload),
  })
}

export function deleteResource(resource, id) {
  return adminRequest(`/api/admin/${resource}/${id}`, {
    method: 'DELETE',
  })
}

export function updateFormStatus(id, status) {
  return adminRequest(`/api/admin/forms/${id}/status`, {
    method: 'PATCH',
    body: JSON.stringify({ status }),
  })
}

export function uploadFile(file) {
  const body = new FormData()
  body.append('file', file)
  return adminRequest('/api/admin/upload', {
    method: 'POST',
    body,
  })
}
