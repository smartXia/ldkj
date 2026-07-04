import assert from 'node:assert/strict'

import {
  getImageSpec,
  hasExactDimensions,
  getCenteredCropArea,
} from '../src/admin/imageSpecs.js'

assert.deepEqual(getImageSpec('banner'), {
  width: 1920,
  height: 800,
  label: '首页 Banner',
})
assert.deepEqual(getImageSpec('case'), {
  width: 376,
  height: 250,
  label: '案例封面',
})
assert.deepEqual(getImageSpec('news'), {
  width: 376,
  height: 188,
  label: '资讯封面',
})
assert.equal(getImageSpec('common'), null)

assert.equal(hasExactDimensions({ width: 376, height: 250 }, getImageSpec('case')), true)
assert.equal(hasExactDimensions({ width: 752, height: 500 }, getImageSpec('case')), false)
assert.equal(hasExactDimensions({ width: 376, height: 188 }, getImageSpec('case')), false)

assert.deepEqual(getCenteredCropArea(1000, 500, getImageSpec('case')), {
  x: 124,
  y: 0,
  width: 752,
  height: 500,
})
assert.deepEqual(getCenteredCropArea(500, 1000, getImageSpec('case')), {
  x: 0,
  y: 334,
  width: 500,
  height: 332,
})
assert.deepEqual(getCenteredCropArea(376, 250, getImageSpec('case')), {
  x: 0,
  y: 0,
  width: 376,
  height: 250,
})
