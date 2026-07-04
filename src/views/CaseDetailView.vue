<script setup>
import { computed, onMounted, shallowRef } from 'vue'
import { useRoute } from 'vue-router'
import { getCaseById } from '../services/publicApi'

const props = defineProps({
  id: {
    type: String,
    default: '',
  },
})

const route = useRoute()
const caseItem = shallowRef(null)
const caseId = computed(() => props.id || route.params.id || route.params.caseId || '')

onMounted(async () => {
  caseItem.value = await getCaseById(caseId.value)
})
</script>

<template>
  <article v-if="caseItem" class="detail-page">
    <header class="case-detail-hero">
      <a href="/case">客户案例</a>
      <h1>{{ caseItem.title }}</h1>
      <div class="meta">
        <span>{{ caseItem.industry }}</span>
        <span>{{ caseItem.method }}</span>
        <span>{{ caseItem.platform }}</span>
      </div>
    </header>

    <img v-if="caseItem.detailImage || caseItem.image" class="detail-image" :src="caseItem.detailImage || caseItem.image" :alt="caseItem.title" />

    <section v-if="caseItem.content" class="detail-copy" v-html="caseItem.content"></section>
  </article>
</template>

<style scoped>
.detail-page {
  background: #fff;
  color: #111;
}

.case-detail-hero,
.detail-copy {
  width: 920px;
  max-width: calc(100% - 48px);
  margin: 0 auto;
}

.case-detail-hero {
  padding: 140px 0 46px;
}

.case-detail-hero a {
  color: #ff4848;
  font-weight: 700;
}

.case-detail-hero h1 {
  margin: 18px 0 22px;
  font-size: 42px;
  line-height: 1.25;
}

.meta {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.meta span {
  border-radius: 18px;
  padding: 7px 14px;
  background: #fff1f1;
  color: #ff4848;
}

.detail-image {
  width: 920px;
  max-width: calc(100% - 48px);
  margin: 0 auto 52px;
  display: block;
  border-radius: 8px;
}

.detail-copy {
  padding-bottom: 96px;
}

.detail-copy :deep(h2) {
  margin: 0 0 18px;
  font-size: 28px;
}

.detail-copy :deep(h3) {
  margin: 28px 0 14px;
  font-size: 22px;
}

.detail-copy :deep(p),
.detail-copy :deep(li) {
  color: #555;
  font-size: 17px;
  line-height: 1.9;
}

.detail-copy :deep(p),
.detail-copy :deep(ul),
.detail-copy :deep(ol),
.detail-copy :deep(blockquote) {
  margin: 0 0 18px;
}

.detail-copy :deep(ul),
.detail-copy :deep(ol) {
  padding-left: 24px;
}

.detail-copy :deep(blockquote) {
  border-left: 4px solid #ff4848;
  padding: 10px 16px;
  background: #fff7f7;
  color: #555;
}

.detail-copy :deep(img) {
  max-width: 100%;
  height: auto;
  border-radius: 8px;
  display: block;
}
</style>
