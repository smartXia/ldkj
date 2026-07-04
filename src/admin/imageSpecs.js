const IMAGE_SPECS = {
  banner: { width: 1920, height: 800, label: '首页 Banner' },
  case: { width: 376, height: 250, label: '案例封面' },
  news: { width: 376, height: 188, label: '资讯封面' },
}

export function getImageSpec(type) {
  return IMAGE_SPECS[type] || null
}

export function hasExactDimensions(size, spec) {
  return Boolean(spec && size?.width === spec.width && size?.height === spec.height)
}

export function getCenteredCropArea(sourceWidth, sourceHeight, spec) {
  if (!spec || sourceWidth <= 0 || sourceHeight <= 0) {
    return { x: 0, y: 0, width: 0, height: 0 }
  }

  const targetRatio = spec.width / spec.height
  const sourceRatio = sourceWidth / sourceHeight

  if (sourceRatio > targetRatio) {
    const width = Math.round(sourceHeight * targetRatio)
    return {
      x: Math.round((sourceWidth - width) / 2),
      y: 0,
      width,
      height: sourceHeight,
    }
  }

  const height = Math.round(sourceWidth / targetRatio)
  return {
    x: 0,
    y: Math.round((sourceHeight - height) / 2),
    width: sourceWidth,
    height,
  }
}
