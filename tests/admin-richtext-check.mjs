import { readFile } from 'node:fs/promises'
import assert from 'node:assert/strict'

const newsView = await readFile(new URL('../src/admin/pages/NewsView.vue', import.meta.url), 'utf8')
const manager = await readFile(new URL('../src/admin/components/ContentManager.vue', import.meta.url), 'utf8')
const editor = await readFile(new URL('../src/admin/components/RichTextEditor.vue', import.meta.url), 'utf8')

assert.match(newsView, /key:\s*'content'[\s\S]*type:\s*'richtext'/)
assert.match(manager, /RichTextEditor/)
assert.match(manager, /field\.type === 'richtext'/)
assert.match(editor, /contenteditable="true"/)
assert.match(editor, /defineModel/)
assert.match(editor, /execCommand/)
assert.match(editor, /源码/)
