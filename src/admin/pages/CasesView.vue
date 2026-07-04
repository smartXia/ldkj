<script setup>
import ContentManager from '../components/ContentManager.vue'

const columns = [
  { key: 'title', label: '标题' },
  { key: 'industry', label: '行业' },
  { key: 'platform', label: '平台' },
  { key: 'status', label: '状态', type: 'status' },
  { key: 'updatedAt', label: '更新时间', type: 'date' },
]

const fields = [
  { key: 'title', label: '标题', placeholder: '案例名称' },
  { key: 'industry', label: '行业标签', placeholder: '如 美妆 / 食品 / 教育' },
  { key: 'platform', label: '平台', placeholder: '如 抖音 / 小红书 / 视频号' },
  { key: 'playbook', label: '打法/赛道', placeholder: '筛选维度' },
  { key: 'cover', label: '封面图', type: 'upload', wide: true },
  { key: 'metrics', label: '核心数据', type: 'textarea', placeholder: '每行一条数据亮点' },
  { key: 'content', label: '详情内容', type: 'textarea', placeholder: '案例正文，可先填写纯文本' },
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
]
</script>

<template>
  <ContentManager title="案例管理" resource="cases" :columns="columns" :fields="fields" upload-type="case" />
</template>
