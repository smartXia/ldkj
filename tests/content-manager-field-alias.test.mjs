import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/admin/components/ContentManager.vue', import.meta.url), 'utf8')

assert.match(source, /function fieldAliases\(/)
assert.match(source, /publishedAt[\s\S]*published_at/)
assert.match(source, /updatedAt[\s\S]*updated_at/)
assert.match(source, /cover[\s\S]*cover_url/)
assert.match(source, /function getFieldValue\(/)
assert.match(source, /const editingRow = ref\(null\)/)
assert.match(source, /function buildPayload\(/)
assert.match(source, /\.\.\.\(isEditing\.value \? editingRow\.value \|\| \{\} : \{\}\)/)
assert.match(source, /\.\.\.form/)
assert.match(source, /function toDateTimeLocalValue\(/)
assert.match(source, /form\[field\.key\] = field\.type === 'datetime-local'/)
assert.match(source, /const value = getFieldValue\(row, column\.key\)/)
