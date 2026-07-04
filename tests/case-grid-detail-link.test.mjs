import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const gridSource = readFileSync(new URL('../src/components/case/CaseGrid.vue', import.meta.url), 'utf8')
const viewSource = readFileSync(new URL('../src/views/CaseView.vue', import.meta.url), 'utf8')

assert.match(gridSource, /<RouterLink[\s\S]*class="case-link"[\s\S]*:to="`\/case\/\$\{item\.slug \|\| item\.id\}`"[\s\S]*>/)
assert.doesNotMatch(gridSource, /@click="emit\('select', item\)"/)
assert.doesNotMatch(viewSource, /CaseDetailModal/)
