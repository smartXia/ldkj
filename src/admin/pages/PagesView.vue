<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'pageKey', label: '页面标识' },
  { key: 'title', label: '标题' },
  { key: 'status', label: '状态', type: 'status' },
  { key: 'updatedAt', label: '更新时间', type: 'date' },
]

const fields = [
  { key: 'pageKey', label: '页面标识', placeholder: '如 about / service / partner' },
  { key: 'title', label: '标题', placeholder: '页面标题' },
  { key: 'extra_data', label: '扩展数据', type: 'textarea', placeholder: '可填写 JSON，用于页面扩展配置' },
  {
    key: 'status',
    label: '状态',
    type: 'select',
    default: 'draft',
    options: [
      { label: '草稿', value: 'draft' },
      { label: '已发布', value: 'published' },
      { label: '归档', value: 'archived' },
    ],
  },
  { key: 'content', label: '页面内容', type: 'richtext', placeholder: '静态页面正文' },
]
</script>

<template>
  <ContentManager title="页面管理" resource="pages" :columns="columns" :fields="fields" upload-type="page" />
</template>
