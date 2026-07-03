<script setup>
import { onMounted, onUnmounted, shallowRef } from 'vue'
import { marketingModules } from '../../data/siteContent'

const sectionRef = shallowRef(null)
const visible = shallowRef(false)
let observer

onMounted(() => {
  observer = new IntersectionObserver(([entry]) => {
    if (entry.isIntersecting) visible.value = true
  }, { threshold: 0.28 })

  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => observer?.disconnect())
</script>

<template>
  <section ref="sectionRef" class="marketing-model" :class="{ visible }">
    <h1>基于社交媒体的数字化营销服务</h1>
    <div class="module-flow">
      <template v-for="(item, index) in marketingModules" :key="item.title">
          <article class="module-card" :style="{ transitionDelay: `${220 + index * 360}ms` }">
          <img :src="item.image" :alt="item.title" loading="lazy" />
          <p>{{ item.title }}</p>
        </article>
        <span
          v-if="index < marketingModules.length - 1"
          class="flow-arrow"
          :style="{ transitionDelay: `${520 + index * 360}ms` }"
          aria-hidden="true"
        >
          <img src="/assets/wsd/marketing-arrow.png" alt="" loading="lazy" />
        </span>
      </template>
    </div>
  </section>
</template>

<style scoped>
.marketing-model {
  padding: 76px 0 102px;
  background: #fff;
}

.marketing-model h1 {
  margin: 0 0 64px;
  text-align: center;
  font-size: 40px;
  line-height: 48px;
  font-weight: 800;
  color: var(--color-ink);
  opacity: 0;
  transform: translateY(34px);
  transition: opacity 960ms ease, transform 960ms cubic-bezier(.16, 1, .3, 1);
}

.module-flow {
  width: 1040px;
  max-width: calc(100% - 48px);
  margin: 0 auto;
  display: grid;
  grid-template-columns: 327px 48px 327px 48px 327px;
  justify-content: center;
  align-items: center;
}

.module-card {
  height: 134px;
  border-radius: 8px;
  background: linear-gradient(90deg, #f8f8fb, #f2f5fa);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  box-shadow: 0 10px 26px rgba(24, 44, 76, 0.045);
  opacity: 0;
  transform: translateY(54px);
  transition: opacity 1100ms ease, transform 1100ms cubic-bezier(.16, 1, .3, 1), box-shadow var(--motion);
}

.module-card img {
  width: 74px;
  height: 74px;
  object-fit: contain;
}

.module-card p {
  margin: 0;
  font-size: 20px;
  line-height: 28px;
  font-weight: 800;
  color: var(--color-ink);
}

.flow-arrow {
  width: 42px;
  height: 42px;
  margin: 0 auto;
  opacity: 0;
  transform: translateX(-18px) scale(.82);
  transition: opacity 900ms ease, transform 900ms cubic-bezier(.16, 1, .3, 1);
}

.flow-arrow img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: contain;
  filter: brightness(0) saturate(100%) invert(36%) sepia(93%) saturate(3113%) hue-rotate(337deg) brightness(103%) contrast(102%);
}

.marketing-model.visible h1,
.marketing-model.visible .module-card {
  opacity: 1;
  transform: translateY(0);
}

.marketing-model.visible .flow-arrow {
  opacity: .88;
  transform: translateX(0) scale(1);
}

@media (max-width: 1100px) {
  .module-flow {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .flow-arrow {
    display: none;
  }
}

@media (max-width: 768px) {
  .marketing-model {
    padding: 25px 0 26px;
  }

  .marketing-model h1 {
    width: calc(100% - 32px);
    margin: 0 auto 17px;
    font-size: 16px;
    line-height: 24px;
  }

  .module-flow {
    width: calc(100% - 32px);
    max-width: none;
    gap: 8px;
  }

  .module-card {
    height: 72px;
    border-radius: 10px;
    gap: 22px;
    justify-content: flex-start;
    padding: 0 52px;
    background: #f4f6fb;
    box-shadow: 0 8px 22px rgba(24, 44, 76, 0.055);
  }

  .module-card img {
    width: 44px;
    height: 44px;
  }

  .module-card p {
    font-size: 14px;
    line-height: 20px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .marketing-model h1,
  .module-card,
  .flow-arrow {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
</style>
