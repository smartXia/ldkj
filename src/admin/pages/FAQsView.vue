<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'question', label: '问题' },
  { key: 'sort', label: '排序' },
  { key: 'status', label: '状态', type: 'status' },
  { key: 'updatedAt', label: '更新时间', type: 'date' },
]

const fields = [
  { key: 'question', label: '问题', placeholder: '常见问题标题' },
  { key: 'answer', label: '回答', type: 'textarea', placeholder: '常见问题回答' },
  { key: 'sort', label: '排序', type: 'number', default: 0 },
  {
    key: 'status',
    label: '状态',
    type: 'select',
    default: 'draft',
    options: [
      { label: '草稿', value: 'draft' },
      { label: '已发布', value: 'published' },
      { label: '停用', value: 'disabled' },
    ],
  },
]
</script>

<template>
  <ContentManager title="FAQ 管理" resource="faqs" :columns="columns" :fields="fields" upload-type="faq" />
</template>
