<script setup>
import { shallowRef } from 'vue'
import { uploadFile } from '../../services/adminApi'

const uploading = shallowRef(false)
const error = shallowRef('')
const fileUrl = shallowRef('')

async function upload(event) {
  const file = event.target.files?.[0]
  if (!file) return
  uploading.value = true
  error.value = ''
  fileUrl.value = ''
  try {
    const result = await uploadFile(file)
    fileUrl.value = result.url || result.path || result.file_url || ''
  } catch (err) {
    error.value = err.message
  } finally {
    uploading.value = false
  }
}
</script>

<template>
  <section class="admin-page">
    <div class="page-title">
      <h1>媒体上传</h1>
    </div>
    <section class="admin-card media-card">
      <label class="upload-box">
        <span>{{ uploading ? '上传中...' : '选择文件上传到后端 /api/admin/upload' }}</span>
        <input type="file" accept="image/*,video/*" @change="upload" />
      </label>
      <p v-if="error" class="admin-error">{{ error }}</p>
      <p v-if="fileUrl" class="admin-success">上传成功：{{ fileUrl }}</p>
      <img v-if="fileUrl" :src="fileUrl" alt="上传预览" />
    </section>
  </section>
</template>
*** Add File: C:\Users\Administrator\.codex\worktrees\2cd4\wsd-social-vue-clone\src\components\admin\admin.css
.admin-page {
  display: grid;
  gap: 20px;
}

.page-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
}

.page-title h1 {
  margin: 0;
  font-size: 26px;
}

.page-title p {
  margin: 4px 0 0;
  color: #7a8494;
  font-size: 14px;
}

.page-title button,
.admin-card button,
.toolbar button {
  border: 0;
  border-radius: 8px;
  padding: 9px 13px;
  background: #ff3d45;
  color: #fff;
}

.admin-grid {
  display: grid;
  grid-template-columns: minmax(300px, 380px) minmax(0, 1fr);
  gap: 18px;
  align-items: start;
}

.admin-card {
  border: 1px solid #e6ebf2;
  border-radius: 8px;
  padding: 18px;
  background: #fff;
  box-shadow: 0 12px 30px rgba(20, 31, 54, 0.06);
}

.admin-card h2 {
  margin: 0 0 16px;
  font-size: 18px;
}

.editor {
  display: grid;
  gap: 12px;
}

.editor label {
  display: grid;
  gap: 6px;
  color: #586173;
  font-size: 13px;
}

.editor input,
.editor textarea,
.toolbar input {
  width: 100%;
  border: 1px solid #d9e0ea;
  border-radius: 8px;
  padding: 9px 10px;
  color: #202838;
}

.toolbar {
  display: flex;
  gap: 10px;
  margin-bottom: 14px;
}

table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

th,
td {
  border-bottom: 1px solid #edf0f5;
  padding: 11px 8px;
  text-align: left;
  vertical-align: top;
}

th {
  color: #7b8494;
  font-weight: 600;
}

.actions {
  white-space: nowrap;
}

.actions button {
  margin-right: 6px;
  padding: 6px 9px;
  background: #f1f4f8;
  color: #313a4a;
}

.actions button:last-child {
  background: #fff0f1;
  color: #e93039;
}

.metric-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 14px;
}

.metric-grid article {
  border-radius: 8px;
  padding: 18px;
  background: #fff;
  box-shadow: 0 12px 30px rgba(20, 31, 54, 0.06);
}

.metric-grid span {
  color: #7a8494;
  font-size: 14px;
}

.metric-grid strong {
  display: block;
  margin-top: 8px;
  font-size: 32px;
}

.admin-error {
  margin: 0;
  border: 1px solid #ffc4c8;
  border-radius: 8px;
  padding: 10px 12px;
  background: #fff5f6;
  color: #d7222c;
}

.admin-success {
  margin: 0;
  border: 1px solid #bdebd0;
  border-radius: 8px;
  padding: 10px 12px;
  background: #f0fff6;
  color: #147a3d;
}

.seo-row {
  display: grid;
  gap: 10px;
  border-top: 1px solid #edf0f5;
  padding-top: 14px;
}

.upload-box {
  display: grid;
  place-items: center;
  min-height: 180px;
  border: 1px dashed #b9c3d2;
  border-radius: 8px;
  background: #f7f9fc;
  color: #596579;
}

.upload-box input {
  display: none;
}

.media-card img {
  max-width: 420px;
  border-radius: 8px;
}

@media (max-width: 980px) {
  .admin-grid,
  .metric-grid {
    grid-template-columns: 1fr;
  }
}
