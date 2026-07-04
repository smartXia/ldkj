import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/admin/pages/CasesView.vue', import.meta.url), 'utf8')

assert.match(source, /key: 'title'/)
assert.match(source, /key: 'slug'/)
assert.match(source, /访问标识/)
assert.match(source, /key: 'playbook'/)
assert.match(source, /key: 'cover'/)
assert.match(source, /key: 'metrics'/)
