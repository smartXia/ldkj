import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

const root = resolve(import.meta.dirname, '..')
const caseContent = readFileSync(resolve(root, 'src/data/caseContent.js'), 'utf8')
const readme = readFileSync(resolve(root, 'README.md'), 'utf8')

const mojibakeMarkers = ['鍏', '钀', '鐢', '鎴', '绉', '璧', '寰', '浠']

for (const marker of mojibakeMarkers) {
  assert.equal(
    caseContent.includes(marker),
    false,
    `caseContent.js should not contain mojibake marker: ${marker}`,
  )
}

for (const expectedText of ['全部', '行业', '营销方式', '平台', '小红书', '案例']) {
  assert.ok(caseContent.includes(expectedText), `caseContent.js should include ${expectedText}`)
}

assert.equal(readme.includes('Vue 3 + Vite'), false, 'README should not keep template heading')
assert.ok(readme.includes('WSD Social Vue Clone'), 'README should describe this project')
