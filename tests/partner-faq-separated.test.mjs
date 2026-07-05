import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/views/PartnerView.vue', import.meta.url), 'utf8')

assert.doesNotMatch(source, /class="faq-list"/)
assert.doesNotMatch(source, /toggleFaq/)
assert.match(source, /to="\/faq"/)
