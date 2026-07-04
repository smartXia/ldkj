<script setup>
import { computed, onMounted, shallowRef } from 'vue'
import { useRoute } from 'vue-router'
import { getServiceById } from '../services/publicApi'

const props = defineProps({
  id: {
    type: String,
    default: '',
  },
})

const route = useRoute()
const service = shallowRef(null)
const serviceId = computed(() => props.id || route.params.id || route.params.serviceId || '')

onMounted(async () => {
  service.value = await getServiceById(serviceId.value)
})
</script>

<template>
  <article v-if="service" class="service-detail">
    <header class="detail-hero">
      <div>
        <a href="/service">服务模块</a>
        <h1>{{ service.title }}</h1>
        <p>{{ service.summary }}</p>
      </div>
      <img v-if="service.image" :src="service.image" :alt="service.title" />
    </header>

    <section class="detail-section">
      <h2>核心能力</h2>
      <div class="pill-grid">
        <span v-for="item in service.highlights" :key="item">{{ item }}</span>
      </div>
    </section>

    <section class="detail-section">
      <h2>服务流程</h2>
      <ol class="process-list">
        <li v-for="(item, index) in service.process" :key="item">
          <strong>{{ String(index + 1).padStart(2, '0') }}</strong>
          <span>{{ item }}</span>
        </li>
      </ol>
    </section>
  </article>
</template>

<style scoped>
.service-detail {
  background: #fff;
  color: #111;
}

.detail-hero {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 0 auto;
  padding: 140px 0 72px;
  display: grid;
  grid-template-columns: 1fr 460px;
  gap: 72px;
  align-items: center;
}

.detail-hero a {
  color: #ff4848;
  font-weight: 700;
}

.detail-hero h1 {
  margin: 18px 0;
  font-size: 46px;
  line-height: 1.18;
}

.detail-hero p {
  margin: 0;
  color: #666;
  font-size: 18px;
  line-height: 1.8;
}

.detail-hero img {
  width: 100%;
  border-radius: 8px;
}

.detail-section {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 0 auto 72px;
}

.detail-section h2 {
  margin: 0 0 24px;
  font-size: 28px;
}

.pill-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.pill-grid span,
.process-list li {
  border: 1px solid #eee;
  border-radius: 8px;
  padding: 22px;
  background: #fff;
}

.process-list {
  list-style: none;
  margin: 0;
  padding: 0 0 80px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.process-list strong {
  display: block;
  margin-bottom: 16px;
  color: #ff4848;
  font-size: 26px;
}

@media (max-width: 860px) {
  .detail-hero,
  .pill-grid,
  .process-list {
    grid-template-columns: 1fr;
  }

  .detail-hero {
    padding-top: 108px;
  }
}
</style>
