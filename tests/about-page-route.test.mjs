import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const router = readFileSync(new URL('../src/router/index.js', import.meta.url), 'utf8')
const header = readFileSync(new URL('../src/components/SiteHeader.vue', import.meta.url), 'utf8')

assert.match(router, /path:\s*'\/about'[\s\S]*name:\s*'about'[\s\S]*AboutView\.vue/)
assert.match(router, /path:\s*'\/zh-cn\/about'[\s\S]*redirect:\s*'\/about'/)
assert.match(header, /label:\s*'关于我们'[\s\S]*path:\s*'\/about'/)
