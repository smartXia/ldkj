<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'title', label: '标题' },
  { key: 'slug', label: '访问标识' },
  { key: 'sort', label: '排序' },
  { key: 'status', label: '状态', type: 'status' },
  { key: 'updatedAt', label: '更新时间', type: 'date' },
]

const fields = [
  { key: 'title', label: '标题', placeholder: '服务名称' },
  { key: 'slug', label: '访问标识', placeholder: '如 social-media-marketing' },
  { key: 'subtitle', label: '副标题', placeholder: '服务副标题' },
  { key: 'summary', label: '摘要', type: 'textarea', placeholder: '列表和详情页摘要' },
  { key: 'cover', label: '封面图', type: 'upload', wide: true },
  { key: 'icon', label: '图标', type: 'upload', wide: true },
  { key: 'highlights', label: '服务亮点', type: 'textarea', placeholder: '每行一条，或填写 JSON 数组' },
  { key: 'process', label: '服务流程', type: 'textarea', placeholder: '每行一步，或填写 JSON 数组' },
  { key: 'sort', label: '排序', type: 'number', default: 0 },
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
  { key: 'content', label: '详情内容', type: 'richtext', placeholder: '服务详情正文' },
]
</script>

<template>
  <ContentManager title="服务管理" resource="services" :columns="columns" :fields="fields" upload-type="service" />
</template>
