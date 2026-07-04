import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const routes = readFileSync(new URL('../src/admin/index.js', import.meta.url), 'utf8')
const layout = readFileSync(new URL('../src/admin/layouts/AdminLayout.vue', import.meta.url), 'utf8')

assert.match(routes, /path:\s*'services'[\s\S]*ServicesView\.vue/)
assert.match(routes, /path:\s*'faqs'[\s\S]*FAQsView\.vue/)
assert.doesNotMatch(routes, /path:\s*'pages'[\s\S]*PagesView\.vue/)
assert.doesNotMatch(routes, /name:\s*'admin-pages'/)

assert.match(layout, /path:\s*'\/admin\/services'/)
assert.match(layout, /path:\s*'\/admin\/faqs'/)
assert.doesNotMatch(layout, /path:\s*'\/admin\/pages'/)
assert.doesNotMatch(layout, /label:\s*'页面'/)

const services = readFileSync(new URL('../src/admin/pages/ServicesView.vue', import.meta.url), 'utf8')
assert.match(services, /resource="services"/)
assert.match(services, /key:\s*'slug'/)
assert.match(services, /key:\s*'cover'/)
assert.match(services, /key:\s*'icon'/)
assert.match(services, /key:\s*'highlights'/)
assert.match(services, /key:\s*'process'/)
assert.match(services, /key:\s*'content'[\s\S]*type:\s*'richtext'/)

const faqs = readFileSync(new URL('../src/admin/pages/FAQsView.vue', import.meta.url), 'utf8')
assert.match(faqs, /resource="faqs"/)
assert.match(faqs, /key:\s*'question'/)
assert.match(faqs, /key:\s*'answer'/)
assert.match(faqs, /key:\s*'sort'/)
