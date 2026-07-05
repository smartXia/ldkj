import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const routes = readFileSync(new URL('../src/admin/index.js', import.meta.url), 'utf8')
const layout = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')
const editor = readFileSync(new URL('../src/admin/pages/AboutPageView.vue', import.meta.url), 'utf8')

assert.match(routes, /path:\s*'about'[\s\S]*AboutPageView\.vue/)
assert.match(layout, /path:\s*'\/admin\/about'/)
assert.match(layout, /label:\s*'关于我们'/)

assert.match(editor, /list\('pages'[\s\S]*keyword:\s*'about'/)
assert.match(editor, /page_key:\s*'about'/)
assert.match(editor, /extra_data/)
assert.match(editor, /hero\.title/)
assert.match(editor, /intro\.paragraphs/)
assert.match(editor, /values/)
assert.match(editor, /honors/)
assert.match(editor, /brands/)
