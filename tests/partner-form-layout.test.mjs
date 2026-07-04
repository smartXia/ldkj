import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/views/PartnerView.vue', import.meta.url), 'utf8')

assert.match(source, /<div class="partner-submit">/)
assert.match(source, /\.partner-submit\s*{[\s\S]*grid-column:\s*1\s*\/\s*-1/)
assert.match(source, /\.partner-submit\s*{[\s\S]*flex-wrap:\s*wrap/)
assert.match(source, /@media \(max-width:\s*520px\)[\s\S]*\.partner-form button\s*{[\s\S]*width:\s*100%/)
