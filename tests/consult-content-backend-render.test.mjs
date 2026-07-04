import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const publicApi = readFileSync(new URL('../src/services/publicApi.js', import.meta.url), 'utf8')
const dialog = readFileSync(new URL('../src/components/consulting/MarketingConsultDialog.vue', import.meta.url), 'utf8')
const contactCta = readFileSync(new URL('../src/components/home/ContactCta.vue', import.meta.url), 'utf8')

assert.match(publicApi, /export function getConsultContent\(\)/)
assert.match(publicApi, /JSON\.parse/)
assert.match(publicApi, /extra_data/)

assert.match(dialog, /getConsultContent/)
assert.match(dialog, /onMounted\(async/)
assert.match(dialog, /consultContent/)
assert.match(dialog, /demandOptions/)

assert.match(contactCta, /getConsultContent/)
assert.match(contactCta, /onMounted\(async/)
assert.match(contactCta, /consultContent/)
assert.match(contactCta, /demandOptions/)
