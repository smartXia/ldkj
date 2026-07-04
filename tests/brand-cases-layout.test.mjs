import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/components/home/BrandCases.vue', import.meta.url), 'utf8')

assert.match(source, /class="case-viewport"/)
assert.match(source, /class="case-rail"/)
assert.match(source, /class="case-item"/)
assert.match(source, /class="case-expanded"/)
assert.match(source, /class="case-static"/)
assert.match(source, /width:\s*max-content/)
assert.match(source, /flex:\s*0 0 208px/)
assert.match(source, /flex-basis:\s*464px/)
assert.doesNotMatch(source, /class="case-showcase"/)
assert.doesNotMatch(source, /class="case-selector"/)
assert.doesNotMatch(source, /class="case-selector-card"/)
