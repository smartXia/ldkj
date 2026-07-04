import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

const root = resolve(import.meta.dirname, '..')
const caseContent = readFileSync(resolve(root, 'src/data/caseContent.js'), 'utf8')
const readme = readFileSync(resolve(root, 'README.md'), 'utf8')
const router = readFileSync(resolve(root, 'src/router/index.js'), 'utf8')
const app = readFileSync(resolve(root, 'src/App.vue'), 'utf8')
const adminApi = readFileSync(resolve(root, 'src/services/adminApi.js'), 'utf8')
const adminShell = readFileSync(resolve(root, 'src/components/admin/AdminShell.vue'), 'utf8')
const adminDashboard = readFileSync(resolve(root, 'src/views/admin/AdminDashboardView.vue'), 'utf8')
const viteConfig = readFileSync(resolve(root, 'vite.config.js'), 'utf8')

const mojibakeMarkers = ['鍏', '钀', '鐢', '鎴', '绉', '璧', '寰', '浠']

for (const marker of mojibakeMarkers) {
  assert.equal(
    caseContent.includes(marker),
    false,
    `caseContent.js should not contain mojibake marker: ${marker}`,
  )
}

for (const expectedText of ['全部', '行业', '营销方式', '平台', '小红书', '案例']) {
  assert.ok(caseContent.includes(expectedText), `caseContent.js should include ${expectedText}`)
}

assert.equal(readme.includes('Vue 3 + Vite'), false, 'README should not keep template heading')
assert.ok(readme.includes('WSD Social Vue Clone'), 'README should describe this project')

for (const routeName of [
  '/admin/login',
  '/admin',
]) {
  assert.ok(router.includes(routeName), `router should include ${routeName}`)
}

for (const adminRouteName of [
  'admin-cases',
  'admin-news',
  'admin-banners',
  'admin-services',
  'admin-leads',
  'admin-media',
  'admin-settings',
]) {
  assert.ok(router.includes(adminRouteName), `router should include ${adminRouteName}`)
}

assert.ok(router.includes('requiresAuth'), 'admin routes should require auth metadata')
assert.ok(router.includes('beforeEach'), 'router should protect admin routes')
assert.ok(app.includes('isAdminRoute'), 'App should isolate admin layout from public layout')
assert.ok(adminApi.includes('/api/admin/login'), 'admin API should use backend login endpoint')
assert.ok(adminApi.includes('Authorization'), 'admin API should send Bearer authorization')
assert.ok(adminApi.includes('Bearer'), 'admin API should use Bearer token')
assert.ok(viteConfig.includes('http://127.0.0.1:8080'), 'Vite should proxy API requests to Go server')
assert.ok(adminShell.includes('admin-nav'), 'admin shell should include navigation')

for (const expectedText of ['待处理线索', '案例数量', '资讯数量', '媒体素材']) {
  assert.ok(adminDashboard.includes(expectedText), `dashboard should include ${expectedText}`)
}
