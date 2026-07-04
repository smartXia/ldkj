import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')
const navItemsBlock = source.match(/const navItems = \[[\s\S]*?\]/)?.[0] || ''

assert.match(navItemsBlock, /label:\s*'页面'[\s\S]*path:\s*'\/admin\/pages'[\s\S]*icon:\s*'pages'/)
assert.match(navItemsBlock, /label:\s*'站点'[\s\S]*path:\s*'\/admin\/site'[\s\S]*icon:\s*'site'/)
