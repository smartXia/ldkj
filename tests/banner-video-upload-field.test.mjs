import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'

const uploadField = readFileSync(new URL('../src/admin/components/UploadField.vue', import.meta.url), 'utf8')

assert.match(uploadField, /const\s+isBannerUpload\s*=/)
assert.match(uploadField, /const\s+acceptTypes\s*=/)
assert.match(uploadField, /video\/mp4/)
assert.match(uploadField, /if\s*\(isBannerVideo\(file\)\)\s*\{[\s\S]*await uploadFile\(file\)/)
assert.match(uploadField, /:accept="acceptTypes"/)
assert.match(uploadField, /v-if="isVideoValue"/)
