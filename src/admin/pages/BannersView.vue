<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'title', label: '标题' },
  { key: 'page', label: '页面' },
  { key: 'sort', label: '排序' },
  { key: 'status', label: '状态', type: 'status' },
]

const fields = [
  { key: 'title', label: '标题', placeholder: 'Banner 主标题' },
  { key: 'subtitle', label: '副标题', placeholder: 'Banner 辅助文案' },
  { key: 'page', label: '页面', placeholder: 'home / case / message / about', default: 'home' },
  { key: 'sort', label: '排序', type: 'number', default: 0 },
  { key: 'image', label: '图片', type: 'upload', wide: true },
  { key: 'link', label: '跳转链接', placeholder: '可选' },
  {
    key: 'status',
    label: '状态',
    type: 'select',
    default: 'enabled',
    options: [
      { label: '启用', value: 'enabled' },
      { label: '停用', value: 'disabled' },
    ],
  },
]
</script>

<template>
  <ContentManager title="Banner 管理" resource="banners" :columns="columns" :fields="fields" upload-type="banner" />
</template>
