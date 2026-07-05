import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const router = readFileSync(new URL('../src/router/index.js', import.meta.url), 'utf8')
const header = readFileSync(new URL('../src/components/SiteHeader.vue', import.meta.url), 'utf8')

assert.match(router, /path:\s*'\/faq'[\s\S]*name:\s*'faq'[\s\S]*FAQView\.vue/)
assert.match(router, /path:\s*'\/zh-cn\/faq'[\s\S]*redirect:\s*'\/faq'/)
assert.match(header, /label:\s*'FAQ'[\s\S]*path:\s*'\/faq'/)
