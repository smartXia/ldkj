import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/views/CaseDetailView.vue', import.meta.url), 'utf8')

assert.match(source, /caseItem\.content/)
assert.match(source, /v-html="caseItem\.content"/)
assert.doesNotMatch(source, /项目复盘[\s\S]*本案例围绕品牌阶段目标/)
