<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'title', label: '标题' },
  { key: 'category', label: '分类' },
  { key: 'status', label: '状态', type: 'status' },
  { key: 'publishedAt', label: '发布时间', type: 'date' },
]

const fields = [
  { key: 'title', label: '标题', placeholder: '资讯标题' },
  {
    key: 'category',
    label: '分类',
    type: 'select',
    default: 'news',
    options: [
      { label: '灵动动态', value: 'news' },
      { label: '营销洞察', value: 'insight' },
      { label: '行业资讯', value: 'industry' },
    ],
  },
  { key: 'cover', label: '封面图', type: 'upload', wide: true },
  { key: 'publishedAt', label: '发布时间', type: 'datetime-local' },
  {
    key: 'status',
    label: '状态',
    type: 'select',
    default: 'draft',
    options: [
      { label: '草稿', value: 'draft' },
      { label: '已发布', value: 'published' },
    ],
  },
  { key: 'summary', label: '摘要', type: 'textarea', placeholder: '列表摘要' },
  { key: 'content', label: '正文', type: 'textarea', placeholder: '资讯正文，可先填写纯文本' },
]
</script>

<template>
  <ContentManager title="资讯管理" resource="news" :columns="columns" :fields="fields" upload-type="news" />
</template>
