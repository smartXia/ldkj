<script setup>
import { computed, onMounted, onUnmounted, shallowRef } from 'vue'

const props = defineProps({
  slides: {
    type: Array,
    default: () => [],
  },
})

const active = shallowRef(0)
const visibleSlides = computed(() => props.slides.filter((item) => item.image))
const currentSlide = computed(() => visibleSlides.value[active.value] || null)
let timer

function go(delta) {
  if (!visibleSlides.value.length) return
  active.value = (active.value + delta + visibleSlides.value.length) % visibleSlides.value.length
}

onMounted(() => {
  if (window.matchMedia('(max-width: 768px)').matches) return

  timer = window.setInterval(() => {
    if (!visibleSlides.value.length) return
    active.value = (active.value + 1) % visibleSlides.value.length
  }, 5000)
})

onUnmounted(() => {
  if (timer) window.clearInterval(timer)
})
</script>

<template>
  <section v-if="currentSlide" id="home" class="hero" aria-label="首页焦点图">
    <Transition name="hero-fade" mode="out-in">
      <img :key="currentSlide.image" class="hero-image" :src="currentSlide.image" :alt="currentSlide.title" />
    </Transition>
    <button v-if="visibleSlides.length > 1" class="hero-arrow hero-arrow-left" type="button" aria-label="上一张" @click="go(-1)">‹</button>
    <button v-if="visibleSlides.length > 1" class="hero-arrow hero-arrow-right" type="button" aria-label="下一张" @click="go(1)">›</button>
    <div v-if="visibleSlides.length > 1" class="hero-dots" aria-label="轮播分页">
      <button
        v-for="(slide, index) in visibleSlides"
        :key="slide.title"
        type="button"
        :class="{ active: active === index }"
        :aria-label="`切换到${slide.title}`"
        @click="active = index"
      ></button>
    </div>
  </section>
</template>

<style scoped>
.hero {
  position: relative;
  margin-top: var(--header-height);
  width: 100%;
  height: min(800px, calc(100vw * 0.4167));
  max-height: calc(100vh - var(--header-height));
  min-height: 560px;
  overflow: hidden;
  background: #000;
}

.hero-image {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
}

.hero-arrow {
  position: absolute;
  top: 50%;
  z-index: 2;
  width: 38px;
  height: 38px;
  border: 0;
  border-radius: 50%;
  display: grid;
  place-items: center;
  background: rgba(128, 153, 176, 0.34);
  color: rgba(255,255,255,.82);
  font-size: 38px;
  line-height: 1;
  transform: translateY(-50%);
  transition: background var(--motion-fast), color var(--motion-fast);
}

.hero-arrow:hover {
  background: rgba(128, 153, 176, 0.54);
  color: #fff;
}

.hero-arrow-left { left: 14px; }
.hero-arrow-right { right: 14px; }

.hero-dots {
  position: absolute;
  left: 50%;
  bottom: 16px;
  display: flex;
  gap: 8px;
  transform: translateX(-50%);
}

.hero-dots button {
  width: 8px;
  height: 8px;
  border: 0;
  border-radius: 50%;
  background: rgba(126, 147, 162, 0.46);
  padding: 0;
  transition: width var(--motion-fast), background var(--motion-fast);
}

.hero-dots button.active {
  width: 24px;
  border-radius: 999px;
  background: var(--color-brand);
}

.hero-fade-enter-active,
.hero-fade-leave-active { transition: opacity 420ms ease; }
.hero-fade-enter-from,
.hero-fade-leave-to { opacity: 0; }

@media (max-width: 768px) {
  .hero {
    height: 64.7vw;
    min-height: 196px;
    max-height: 252px;
  }
  .hero-image {
    object-position: center top;
  }
  .hero-arrow { display: none; }
  .hero-dots {
    bottom: 8px;
    gap: 5px;
  }
  .hero-dots button {
    width: 5px;
    height: 5px;
  }
  .hero-dots button.active {
    width: 5px;
    border-radius: 50%;
  }
}

@media (prefers-reduced-motion: reduce) {
  .hero-fade-enter-active,
  .hero-fade-leave-active { transition: none; }
}
</style>
