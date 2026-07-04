<script setup>
import { computed, nextTick, reactive, ref } from 'vue'
import { upload } from '../../services/adminApi'
import { getCenteredCropArea, getImageSpec, hasExactDimensions } from '../imageSpecs'

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'common',
  },
})

const emit = defineEmits(['update:modelValue'])

const fileInput = ref(null)
const cropViewport = ref(null)
const uploading = ref(false)
const error = ref('')
const imageSpec = computed(() => getImageSpec(props.type))
const specText = computed(() => {
  if (!imageSpec.value) return ''
  return `${imageSpec.value.label}要求 ${imageSpec.value.width} x ${imageSpec.value.height}px`
})

const crop = reactive({
  visible: false,
  file: null,
  url: '',
  naturalWidth: 0,
  naturalHeight: 0,
  viewportWidth: 0,
  viewportHeight: 0,
  imageWidth: 0,
  imageHeight: 0,
  offsetX: 0,
  offsetY: 0,
  zoom: 1,
  dragging: false,
  startX: 0,
  startY: 0,
  startOffsetX: 0,
  startOffsetY: 0,
})

const cropViewportStyle = computed(() => {
  if (!imageSpec.value) return {}
  return { aspectRatio: `${imageSpec.value.width} / ${imageSpec.value.height}` }
})

const cropImageStyle = computed(() => ({
  width: `${crop.imageWidth}px`,
  height: `${crop.imageHeight}px`,
  transform: `translate(${crop.offsetX}px, ${crop.offsetY}px)`,
}))

function readImage(file) {
  return new Promise((resolve, reject) => {
    const url = URL.createObjectURL(file)
    const image = new Image()
    image.onload = () => resolve({ url, width: image.naturalWidth, height: image.naturalHeight })
    image.onerror = () => {
      URL.revokeObjectURL(url)
      reject(new Error('图片读取失败，请更换图片后重试'))
    }
    image.src = url
  })
}

function clampCropOffset() {
  const minX = Math.min(0, crop.viewportWidth - crop.imageWidth)
  const minY = Math.min(0, crop.viewportHeight - crop.imageHeight)
  crop.offsetX = Math.min(0, Math.max(minX, crop.offsetX))
  crop.offsetY = Math.min(0, Math.max(minY, crop.offsetY))
}

function syncCropImageSize() {
  if (!crop.viewportWidth || !crop.viewportHeight) return
  const coverScale = Math.max(crop.viewportWidth / crop.naturalWidth, crop.viewportHeight / crop.naturalHeight)
  crop.imageWidth = Math.round(crop.naturalWidth * coverScale * crop.zoom)
  crop.imageHeight = Math.round(crop.naturalHeight * coverScale * crop.zoom)
  clampCropOffset()
}

function resetCropState() {
  if (crop.url) URL.revokeObjectURL(crop.url)
  crop.visible = false
  crop.file = null
  crop.url = ''
  crop.naturalWidth = 0
  crop.naturalHeight = 0
  crop.viewportWidth = 0
  crop.viewportHeight = 0
  crop.imageWidth = 0
  crop.imageHeight = 0
  crop.offsetX = 0
  crop.offsetY = 0
  crop.zoom = 1
  crop.dragging = false
}

async function openCropper(file, imageInfo) {
  resetCropState()
  crop.visible = true
  crop.file = file
  crop.url = imageInfo.url
  crop.naturalWidth = imageInfo.width
  crop.naturalHeight = imageInfo.height
  crop.zoom = 1

  await nextTick()

  const rect = cropViewport.value.getBoundingClientRect()
  crop.viewportWidth = Math.round(rect.width)
  crop.viewportHeight = Math.round(rect.height)
  syncCropImageSize()
  crop.offsetX = Math.round((crop.viewportWidth - crop.imageWidth) / 2)
  crop.offsetY = Math.round((crop.viewportHeight - crop.imageHeight) / 2)
  clampCropOffset()
}

async function uploadFile(file) {
  uploading.value = true
  error.value = ''

  try {
    const payload = await upload(file, props.type)
    const url = payload.url || payload.data?.url || payload.path || ''
    emit('update:modelValue', url)
  } catch (err) {
    error.value = err.message || '上传失败'
  } finally {
    uploading.value = false
  }
}

async function handleFileChange(event) {
  const file = event.target.files?.[0]
  if (!file) return

  error.value = ''

  try {
    if (!imageSpec.value) {
      await uploadFile(file)
      return
    }

    const imageInfo = await readImage(file)
    if (hasExactDimensions({ width: imageInfo.width, height: imageInfo.height }, imageSpec.value)) {
      URL.revokeObjectURL(imageInfo.url)
      await uploadFile(file)
      return
    }

    await openCropper(file, imageInfo)
  } catch (err) {
    error.value = err.message || '上传失败'
  } finally {
    event.target.value = ''
  }
}

function startDrag(event) {
  crop.dragging = true
  crop.startX = event.clientX
  crop.startY = event.clientY
  crop.startOffsetX = crop.offsetX
  crop.startOffsetY = crop.offsetY
  event.currentTarget.setPointerCapture(event.pointerId)
}

function dragImage(event) {
  if (!crop.dragging) return
  crop.offsetX = crop.startOffsetX + event.clientX - crop.startX
  crop.offsetY = crop.startOffsetY + event.clientY - crop.startY
  clampCropOffset()
}

function stopDrag(event) {
  crop.dragging = false
  event.currentTarget.releasePointerCapture?.(event.pointerId)
}

function handleZoomChange(event) {
  const nextZoom = Number(event.target.value)
  const centerX = crop.viewportWidth / 2 - crop.offsetX
  const centerY = crop.viewportHeight / 2 - crop.offsetY
  const ratioX = centerX / crop.imageWidth
  const ratioY = centerY / crop.imageHeight

  crop.zoom = nextZoom
  syncCropImageSize()
  crop.offsetX = Math.round(crop.viewportWidth / 2 - crop.imageWidth * ratioX)
  crop.offsetY = Math.round(crop.viewportHeight / 2 - crop.imageHeight * ratioY)
  clampCropOffset()
}

function canvasToBlob(canvas, type) {
  return new Promise((resolve, reject) => {
    canvas.toBlob((blob) => {
      if (blob) resolve(blob)
      else reject(new Error('裁剪图片生成失败'))
    }, type)
  })
}

async function confirmCrop() {
  if (!crop.file || !imageSpec.value) return

  uploading.value = true
  error.value = ''

  try {
    const image = new Image()
    image.src = crop.url
    await image.decode()

    const scaleX = crop.naturalWidth / crop.imageWidth
    const scaleY = crop.naturalHeight / crop.imageHeight
    const sourceX = Math.max(0, Math.round(-crop.offsetX * scaleX))
    const sourceY = Math.max(0, Math.round(-crop.offsetY * scaleY))
    const sourceWidth = Math.min(crop.naturalWidth - sourceX, Math.round(crop.viewportWidth * scaleX))
    const sourceHeight = Math.min(crop.naturalHeight - sourceY, Math.round(crop.viewportHeight * scaleY))
    const fallbackArea = getCenteredCropArea(crop.naturalWidth, crop.naturalHeight, imageSpec.value)
    const drawArea = sourceWidth > 0 && sourceHeight > 0
      ? { x: sourceX, y: sourceY, width: sourceWidth, height: sourceHeight }
      : fallbackArea

    const canvas = document.createElement('canvas')
    canvas.width = imageSpec.value.width
    canvas.height = imageSpec.value.height
    const context = canvas.getContext('2d')
    context.drawImage(
      image,
      drawArea.x,
      drawArea.y,
      drawArea.width,
      drawArea.height,
      0,
      0,
      imageSpec.value.width,
      imageSpec.value.height,
    )

    const outputType = ['image/jpeg', 'image/png', 'image/webp'].includes(crop.file.type) ? crop.file.type : 'image/jpeg'
    const blob = await canvasToBlob(canvas, outputType)
    const croppedFile = new File([blob], `cropped-${crop.file.name}`, { type: outputType })
    resetCropState()
    await uploadFile(croppedFile)
  } catch (err) {
    uploading.value = false
    error.value = err.message || '裁剪失败'
  }
}
</script>

<template>
  <div class="admin-upload">
    <p v-if="specText" class="admin-upload-hint">{{ specText }}，尺寸不符时需先裁剪。</p>
    <div class="admin-upload-row">
      <input
        class="admin-input"
        :value="modelValue"
        placeholder="上传后自动回填，或粘贴图片 URL"
        @input="emit('update:modelValue', $event.target.value)"
      />
      <button class="admin-button secondary" type="button" :disabled="uploading" @click="fileInput.click()">
        {{ uploading ? '上传中' : '上传' }}
      </button>
      <input ref="fileInput" type="file" accept="image/*" hidden @change="handleFileChange" />
    </div>
    <div v-if="modelValue" class="admin-upload-preview">
      <img :src="modelValue" alt="" />
      <span class="admin-muted">{{ modelValue }}</span>
    </div>
    <p v-if="error" class="admin-muted">{{ error }}</p>

    <div v-if="crop.visible" class="admin-crop-modal" role="dialog" aria-modal="true" aria-label="裁剪图片">
      <div class="admin-crop-panel">
        <div class="admin-crop-header">
          <div>
            <strong>裁剪{{ imageSpec.label }}</strong>
            <span>输出 {{ imageSpec.width }} x {{ imageSpec.height }}px</span>
          </div>
          <button class="admin-button ghost" type="button" :disabled="uploading" @click="resetCropState">取消</button>
        </div>

        <div
          ref="cropViewport"
          class="admin-crop-viewport"
          :style="cropViewportStyle"
          @pointerdown="startDrag"
          @pointermove="dragImage"
          @pointerup="stopDrag"
          @pointercancel="stopDrag"
        >
          <img :src="crop.url" alt="" draggable="false" :style="cropImageStyle" />
          <span class="admin-crop-grid"></span>
        </div>

        <div class="admin-crop-tools">
          <label>
            缩放
            <input type="range" min="1" max="3" step="0.01" :value="crop.zoom" @input="handleZoomChange" />
          </label>
          <span>原图 {{ crop.naturalWidth }} x {{ crop.naturalHeight }}px</span>
        </div>

        <div class="admin-crop-actions">
          <button class="admin-button secondary" type="button" :disabled="uploading" @click="resetCropState">取消</button>
          <button class="admin-button" type="button" :disabled="uploading" @click="confirmCrop">
            {{ uploading ? '处理中' : '确认裁剪并上传' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
