<script setup>
import { computed } from 'vue'

const props = defineProps({
  logos: {
    type: Array,
    default: () => [],
  },
})

const backendLogos = computed(() => props.logos.filter((brand) => brand.image))
</script>

<template>
  <section v-if="backendLogos.length" class="brand-wall">
    <h2>合作品牌</h2>
    <p>品牌 Logo 来自后端配置，未配置图片时不展示本模块。</p>

    <ul class="brand-list">
      <li v-for="brand in backendLogos" :key="brand.name || brand.image">
        <img :src="brand.image" :alt="brand.name || '合作品牌'" loading="lazy" />
      </li>
    </ul>
  </section>
</template>

<style scoped>
.brand-wall {
  padding: 80px max(24px, calc((100vw - 1200px) / 2));
  background: #f4f4f4;
  text-align: center;
}

.brand-wall h2 {
  margin: 0 0 16px;
  color: #111;
  font-size: 40px;
  line-height: 1.2;
}

.brand-wall p {
  max-width: 760px;
  margin: 0 auto 36px;
  color: #666;
  font-size: 18px;
  line-height: 1.7;
}

.brand-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(112px, 1fr));
  gap: 20px;
}

.brand-list li {
  height: 92px;
  border-radius: 8px;
  display: grid;
  place-items: center;
  background: #fff;
  box-shadow: 0 4px 14px rgba(0, 0, 0, .04);
}

.brand-list img {
  max-width: 78px;
  max-height: 62px;
  object-fit: contain;
}

@media (max-width: 767px) {
  .brand-wall {
    padding: 32px 16px;
  }

  .brand-wall h2 {
    font-size: 22px;
  }

  .brand-wall p {
    font-size: 12px;
  }

  .brand-list {
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }

  .brand-list li {
    height: 58px;
  }

  .brand-list img {
    max-width: 52px;
    max-height: 40px;
  }
}
</style>
