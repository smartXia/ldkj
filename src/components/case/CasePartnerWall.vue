<script setup>
defineProps({
  logos: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section class="partner-wall">
    <div class="container">
      <div class="partner-heading">
        <span>服务客户</span>
        <h2>跨行业品牌的社交营销伙伴</h2>
        <p>从种草、转化到长期运营，沉淀品牌在社交生态中的增长资产。</p>
      </div>
      <div class="logo-grid">
        <div v-for="logo in logos" :key="logo.name" class="logo-card">
          <img :src="logo.image" :alt="logo.name" loading="lazy" />
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.partner-wall {
  padding: 92px 0 104px;
  background: #fff;
}

.partner-heading {
  margin-bottom: 42px;
  text-align: center;
}

.partner-heading span {
  color: var(--color-brand);
  font-size: 15px;
  font-weight: 900;
}

.partner-heading h2 {
  margin: 10px 0 14px;
  color: var(--color-ink);
  font-size: clamp(30px, 3.2vw, 44px);
  line-height: 1.2;
}

.partner-heading p {
  margin: 0;
  color: var(--color-ink-muted);
  font-size: 18px;
}

.logo-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 14px;
}

.logo-card {
  height: 92px;
  border: 1px solid var(--color-border);
  border-radius: 8px;
  display: grid;
  place-items: center;
  background: #fff;
  transition: transform var(--motion-fast), box-shadow var(--motion-fast);
}

.logo-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 14px 32px rgba(30, 42, 70, 0.08);
}

.logo-card img {
  max-width: 76%;
  max-height: 52px;
  object-fit: contain;
}

@media (max-width: 920px) {
  .logo-grid { grid-template-columns: repeat(3, 1fr); }
}

@media (max-width: 520px) {
  .partner-wall { padding: 66px 0 76px; }
  .logo-grid { grid-template-columns: repeat(2, 1fr); }
}
</style>
