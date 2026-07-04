<script setup>
import { onMounted, shallowRef } from 'vue'
import { getServices } from '../services/publicApi'

const services = shallowRef([])

onMounted(async () => {
  services.value = await getServices()
})
</script>

<template>
  <section class="service-page">
    <div class="service-hero">
      <p>Social Marketing Services</p>
      <h1>服务模块</h1>
      <span>从内容策略、平台合作到转化经营，提供覆盖社交营销全链路的服务能力。</span>
    </div>

    <div class="service-grid">
      <article v-for="item in services" :key="item.id" class="service-card">
        <img :src="item.image" :alt="item.title" loading="lazy" />
        <div>
          <h2>{{ item.title }}</h2>
          <p>{{ item.summary }}</p>
          <a :href="`/service/${item.id}`">查看详情</a>
        </div>
      </article>
    </div>
  </section>
</template>

<style scoped>
.service-page {
  background: #fff;
  color: #111;
}

.service-hero {
  min-height: 360px;
  padding: 144px max(24px, calc((100vw - 1200px) / 2)) 80px;
  background: linear-gradient(135deg, #fff 0%, #fff4f4 100%);
}

.service-hero p {
  margin: 0 0 12px;
  color: #ff4848;
  font-weight: 700;
}

.service-hero h1 {
  margin: 0;
  font-size: 48px;
  line-height: 1.15;
}

.service-hero span {
  display: block;
  max-width: 640px;
  margin-top: 18px;
  color: #666;
  font-size: 18px;
  line-height: 1.8;
}

.service-grid {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 76px auto 120px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}

.service-card {
  overflow: hidden;
  border: 1px solid #eee;
  border-radius: 8px;
  background: #fff;
  transition: transform .2s, box-shadow .2s;
}

.service-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 18px 42px rgba(26, 38, 66, .1);
}

.service-card img {
  width: 100%;
  aspect-ratio: 376 / 220;
  display: block;
  object-fit: cover;
  background: #f7f7f7;
}

.service-card div {
  padding: 24px;
}

.service-card h2 {
  min-height: 56px;
  margin: 0 0 14px;
  font-size: 20px;
  line-height: 1.4;
}

.service-card p {
  min-height: 78px;
  margin: 0 0 18px;
  color: #666;
  line-height: 1.7;
}

.service-card a {
  color: #ff4848;
  font-weight: 700;
}

@media (max-width: 960px) {
  .service-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .service-hero {
    min-height: 260px;
    padding-top: 104px;
  }

  .service-hero h1 {
    font-size: 34px;
  }

  .service-grid {
    grid-template-columns: 1fr;
    max-width: calc(100% - 32px);
    margin-top: 42px;
  }
}
</style>
