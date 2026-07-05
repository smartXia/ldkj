import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const apiSource = readFileSync(new URL('../src/services/adminApi.js', import.meta.url), 'utf8')
const routesSource = readFileSync(new URL('../src/admin/index.js', import.meta.url), 'utf8')
const layoutSource = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')

assert.match(apiSource, /const SESSION_KEY = 'wsd_admin_session'/)
assert.match(apiSource, /function getSession\(\)/)
assert.match(apiSource, /function hasPermission\(permission\)/)
assert.match(apiSource, /window\.localStorage\.setItem\(SESSION_KEY/)

assert.match(routesSource, /meta:\s*\{\s*permission:\s*'content:read'\s*\}/)
assert.match(routesSource, /path:\s*'users'[\s\S]*meta:\s*\{\s*permission:\s*'user:manage'\s*\}/)

assert.match(layoutSource, /permission:\s*'user:manage'/)
assert.match(layoutSource, /const visibleNavItems = computed/)
assert.match(layoutSource, /hasPermission\(item\.permission\)/)
