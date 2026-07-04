import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/components/SiteHeader.vue', import.meta.url), 'utf8')

assert.match(source, /<RouterLink[\s\S]*class="btn ghost"[\s\S]*to="\/admin\/login"[\s\S]*>/)
assert.doesNotMatch(source, /href="#login"/)
assert.match(source, /language-switch/)
assert.match(source, /language-menu/)
assert.match(source, /languageOpen/)
assert.match(source, /currentLanguage/)
assert.match(source, /setLanguage/)
