<script setup>
import { ref } from 'vue'
import { upload } from '../../services/adminApi'

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
const uploading = ref(false)
const error = ref('')

async function handleFileChange(event) {
  const file = event.target.files?.[0]
  if (!file) return

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
    event.target.value = ''
  }
}
</script>

<template>
  <div class="admin-upload">
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
  </div>
</template>
