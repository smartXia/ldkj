import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const homeSource = readFileSync(new URL('../src/views/HomeView.vue', import.meta.url), 'utf8')
const apiSource = readFileSync(new URL('../src/services/publicApi.js', import.meta.url), 'utf8')

const expectedOrder = [
  '<HeroCarousel',
  '<MarketingModel',
  '<ServiceSection',
  '<BrandCases',
  '<IntroStats',
  '<BrandWall',
  '<ContactCta',
]

let previousIndex = -1
for (const marker of expectedOrder) {
  const index = homeSource.indexOf(marker)
  assert.notEqual(index, -1, `${marker} should be rendered on the home page`)
  assert.ok(index > previousIndex, `${marker} should follow the reference homepage order`)
  previousIndex = index
}

assert.match(homeSource, /getServices/)
assert.match(apiSource, /siteContent/)
assert.match(apiSource, /marketingModules/)
assert.match(apiSource, /services/)
assert.match(apiSource, /brandCases/)
