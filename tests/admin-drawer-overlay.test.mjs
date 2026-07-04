import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const styles = readFileSync(new URL('../src/admin/styles.css', import.meta.url), 'utf8')

assert.match(styles, /\.admin-drawer\s*\{[\s\S]*position:\s*fixed/)
assert.match(styles, /\.admin-drawer\s*\{[\s\S]*inset:\s*0/)
assert.match(styles, /\.admin-drawer::before\s*\{[\s\S]*background:\s*rgba/)
assert.match(styles, /\.admin-drawer\s*\{[\s\S]*grid-template-columns:\s*minmax\(0,\s*1fr\)\s*min/)
assert.match(styles, /\.admin-form-actions\s*\{[\s\S]*position:\s*sticky/)
