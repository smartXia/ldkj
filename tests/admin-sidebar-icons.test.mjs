import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')

const navItemMatches = source.match(/\{\s*label:[\s\S]*?path:[\s\S]*?\}/g) || []
assert.ok(navItemMatches.length >= 10, 'admin sidebar should define all management nav items')

navItemMatches.forEach((item) => {
  assert.match(item, /icon:\s*'/, `nav item should declare icon: ${item}`)
})

assert.match(source, /class="admin-nav-icon"/)
assert.match(source, /aria-hidden="true"/)
assert.match(source, /:d="iconPaths\[item\.icon\]"/)
