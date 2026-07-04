import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const apiSource = readFileSync(new URL('../src/services/publicApi.js', import.meta.url), 'utf8')
const heroSource = readFileSync(new URL('../src/components/home/HeroCarousel.vue', import.meta.url), 'utf8')

assert.match(apiSource, /const bannerItems = unwrapList\(bannerPayload, \[\]\)/)
assert.match(apiSource, /banners: bannerItems/)
assert.match(apiSource, /banner: bannerItems\.length \? \{\} : unwrapItem\(bannerPayload, \{\}\)/)
assert.doesNotMatch(heroSource, /matchMedia\('\(max-width:\s*768px\)'\)\.matches/)
assert.match(heroSource, /visibleSlides\.value\.length <= 1/)
assert.match(heroSource, /window\.setInterval/)
