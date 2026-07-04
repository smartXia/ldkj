import { readFileSync } from 'node:fs'
import assert from 'node:assert/strict'

const api = readFileSync(new URL('../src/services/publicApi.js', import.meta.url), 'utf8')

assert.doesNotMatch(
  api,
  /if\s*\(!image\s*\|\|\s*isFrontendAsset\(image\)\)\s*return\s*''/,
  'frontend /assets news covers must not be stripped during public API normalization'
)

assert.match(
  api,
  /function\s+normalizeArticleItem\s*\(/,
  'news items should have a dedicated normalizer for frontend card/detail fields'
)

assert.match(api, /const\s+desc\s*=\s*firstNonEmpty\([^)]*summary/, 'news summary should map to article.desc')
assert.match(api, /const\s+date\s*=\s*firstNonEmpty\([^)]*published_at/, 'news published_at should map to article.date')
assert.match(api, /href:\s*routeId\s*\?\s*`\/message\/\$\{routeId\}`\s*:\s*'\/message'/, 'news slug/id should map to article.href')
