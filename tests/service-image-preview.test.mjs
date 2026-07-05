import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const source = readFileSync(new URL('../src/components/home/ServiceSection.vue', import.meta.url), 'utf8')

assert.match(source, /const previewOpen = shallowRef\(false\)/)
assert.match(source, /function openPreview\(\)/)
assert.match(source, /class="service-image-trigger"/)
assert.match(source, /@click="openPreview"/)
assert.match(source, /<Teleport v-if="previewOpen" to="body">/)
assert.match(source, /class="service-preview-modal"/)
assert.match(source, /@click="closePreview"/)
assert.match(source, /event\.key === 'Escape'/)
