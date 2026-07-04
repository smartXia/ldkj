import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')

assert.doesNotMatch(source, /SiteHeader/, 'Admin layout should not reuse the public SiteHeader')
assert.match(source, /<RouterLink[\s\S]*class="admin-button secondary"[\s\S]*to="\/"[\s\S]*>返回前台<\/RouterLink>/)
assert.match(source, /<button[\s\S]*class="admin-button ghost"[\s\S]*@click="handleLogout"[\s\S]*>退出登录<\/button>/)
assert.match(source, /class="admin-topbar-actions"/)
