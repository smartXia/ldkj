import assert from 'node:assert/strict'

async function loadConfig(suffix) {
  const { default: config } = await import(`../vite.config.js?${suffix}`)
  return typeof config === 'function' ? config({ mode: 'development' }) : config
}

let resolvedConfig = await loadConfig('default')
let proxy = resolvedConfig.server?.proxy

assert.ok(proxy?.['/api'], 'Vite dev server should proxy /api requests to the backend')
assert.equal(proxy['/api'].target, 'http://localhost:8080')
assert.equal(proxy['/api'].changeOrigin, true)

assert.ok(proxy?.['/oss'], 'Vite dev server should proxy uploaded /oss assets to the backend')
assert.equal(proxy['/oss'].target, 'http://localhost:8080')
assert.equal(proxy['/oss'].changeOrigin, true)

process.env.VITE_BACKEND_PROXY_TARGET = 'http://127.0.0.1:9000'
resolvedConfig = await loadConfig('override')
proxy = resolvedConfig.server?.proxy

assert.equal(proxy['/api'].target, 'http://127.0.0.1:9000')
assert.equal(proxy['/oss'].target, 'http://127.0.0.1:9000')

delete process.env.VITE_BACKEND_PROXY_TARGET
