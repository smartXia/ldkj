<script setup>
import { computed, onBeforeUnmount, onMounted, shallowRef } from 'vue'

const props = defineProps({
  brands: {
    type: Array,
    required: true,
  },
})

const scrollIndex = shallowRef(0)
const isResetting = shallowRef(false)
const cardStep = shallowRef(257)
let timer
let resetTimer

const CARD_STEP = 257
const MOBILE_CARD_STEP = 150
const LOOP_COPY_COUNT = 3
const ORBIT_LOGOS = [
  'round-wsd.png',
  'round-global-wsd.png',
  'round-hlkj.png',
  'round-sdr.png',
  'round-hys.png',
  'round-qp.png',
  'round-yy.png',
  'round-skhd.png',
  'round-jx.png',
]

const orbitSlots = [
  { x: 24, y: 358, width: 200, z: 8 },
  { x: 68, y: 214, width: 202, z: 5 },
  { x: 304, y: 88, width: 202, z: 5 },
  { x: 536, y: 166, width: 202, z: 5 },
  { x: 650, y: 290, width: 202, z: 5 },
  { x: 650, y: 438, width: 202, z: 5 },
  { x: 536, y: 578, width: 202, z: 5 },
  { x: 304, y: 654, width: 202, z: 5 },
  { x: 68, y: 520, width: 202, z: 5 },
]

const activeIndex = computed(() => scrollIndex.value % props.brands.length)

const loopedCards = computed(() =>
  Array.from({ length: LOOP_COPY_COUNT }, (_, copyIndex) =>
    props.brands.map((brand, index) => ({
      ...brand,
      key: `${copyIndex}-${brand.title}`,
      copyIndex,
      rawIndex: index,
    }))
  ).flat()
)

const listStyle = computed(() => ({
  transform: `translate3d(0, ${-1 * (props.brands.length + scrollIndex.value) * cardStep.value}px, 0)`,
  transitionDuration: isResetting.value ? '0ms' : '850ms',
}))

const orbitBrands = computed(() =>
  props.brands.map((brand, index) => ({
    ...brand,
    rawIndex: index,
    orbitLogo: `/assets/wsd/about/${ORBIT_LOGOS[index]}`,
    orbitStyle: getOrbitStyle(index),
  }))
)

function getOrbitStyle(index) {
  const total = props.brands.length
  const slotIndex = (activeIndex.value - index + total) % total
  const slot = orbitSlots[slotIndex]

  return {
    '--orbit-x': `${slot.x}px`,
    '--orbit-y': `${slot.y}px`,
    '--orbit-width': `${slot.width}px`,
    '--orbit-z': slot.z,
  }
}

function isActive(index) {
  return index === activeIndex.value
}

function isVisibleActive(brand) {
  return brand.copyIndex === 1 && brand.rawIndex === activeIndex.value
}

function next() {
  if (isResetting.value) return

  scrollIndex.value += 1

  if (scrollIndex.value === props.brands.length) {
    window.clearTimeout(resetTimer)
    resetTimer = window.setTimeout(() => {
      isResetting.value = true
      scrollIndex.value = 0

      window.requestAnimationFrame(() => {
        window.requestAnimationFrame(() => {
          isResetting.value = false
        })
      })
    }, 880)
  }
}

function start() {
  stop()
  timer = window.setInterval(next, 2600)
}

function stop() {
  if (timer) window.clearInterval(timer)
  if (resetTimer) window.clearTimeout(resetTimer)
}

function syncCardStep() {
  cardStep.value = window.matchMedia('(max-width: 760px)').matches ? MOBILE_CARD_STEP : CARD_STEP
}

onMounted(() => {
  syncCardStep()
  window.addEventListener('resize', syncCardStep)
  start()
})

onBeforeUnmount(() => {
  stop()
  window.removeEventListener('resize', syncCardStep)
})
</script>

<template>
  <section id="about-brands" class="about-brands">
    <div class="brand-shell">
      <div class="brand-card-viewport" aria-label="子品牌自动滚动列表">
        <div class="brand-card-list" :style="listStyle">
          <article
            v-for="brand in loopedCards"
            :key="brand.key"
            class="brand-card"
            :class="{ active: isVisibleActive(brand) }"
          >
            <img :src="brand.logo" :alt="brand.alt" loading="lazy" />
            <h3>{{ brand.title }}</h3>
            <p>{{ brand.desc }}</p>
          </article>
        </div>
      </div>

      <div class="brand-map" aria-label="灵动集团子品牌关系图">
        <img class="brand-map-bg" src="/assets/wsd/about/company-bg-pc.png" alt="" />
        <div class="brand-center">灵动集团</div>
        <div
          v-for="brand in orbitBrands"
          :key="brand.title"
          class="orbit-card"
          :class="{ active: isActive(brand.rawIndex) }"
          :style="brand.orbitStyle"
        >
          <img :src="brand.orbitLogo" :alt="brand.alt" loading="lazy" />
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about-brands {
  padding: 0;
  overflow: hidden;
  background: linear-gradient(90deg, #fff 0 34.6%, #fff7f7 34.6% 100%);
}

.brand-shell {
  position: relative;
  width: 1460px;
  min-height: 804px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 440px 1fr;
}

.brand-card-viewport {
  position: relative;
  height: 804px;
  overflow: hidden;
  background: #fff;
}

.brand-card-viewport::before,
.brand-card-viewport::after {
  content: "";
  position: absolute;
  left: 0;
  right: 0;
  z-index: 5;
  height: 116px;
  pointer-events: none;
}

.brand-card-viewport::before {
  top: 0;
  background: linear-gradient(#fff 0%, rgba(255, 255, 255, 0) 100%);
}

.brand-card-viewport::after {
  bottom: 0;
  background: linear-gradient(rgba(255, 255, 255, 0) 0%, #fff 100%);
}

.brand-card-list {
  padding-top: 292px;
  transition: transform 850ms cubic-bezier(0.22, 1, 0.36, 1);
}

.brand-card {
  width: 440px;
  height: 235px;
  margin-bottom: 22px;
  padding: 32px 32px 30px;
  border-radius: 0;
  background: rgba(255, 255, 255, 0.92);
  box-shadow: 0 14px 34px rgba(31, 38, 56, 0.04);
  opacity: 0.28;
  overflow: hidden;
  transition: opacity 680ms ease, box-shadow 680ms ease;
}

.brand-card.active {
  opacity: 1;
  box-shadow: 0 16px 36px rgba(31, 38, 56, 0.06);
}

.brand-card img {
  width: 128px;
  height: 55px;
  display: block;
  object-fit: contain;
  object-position: left center;
}

.brand-card h3 {
  margin: 22px 0 12px;
  color: #1f2530;
  font-size: 22px;
  line-height: 30px;
  font-weight: 700;
}

.brand-card p {
  margin: 0;
  color: #1f2530;
  font-size: 16px;
  line-height: 23px;
  display: -webkit-box;
  overflow: hidden;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
}

.brand-card:not(.active) h3,
.brand-card:not(.active) p {
  color: rgba(31, 37, 48, 0.36);
}

.brand-map {
  position: relative;
  height: 804px;
  overflow: hidden;
}

.brand-map-bg {
  position: absolute;
  left: -118px;
  top: 0;
  width: 1074px;
  height: 804px;
  object-fit: contain;
  pointer-events: none;
}

.brand-center {
  position: absolute;
  left: 304px;
  top: 318px;
  z-index: 6;
  width: 180px;
  height: 180px;
  border-radius: 50%;
  display: grid;
  place-items: center;
  color: #fff;
  background: linear-gradient(154deg, #ff9c7a, #fb5151 83%, #fb4848);
  font-size: 28px;
  line-height: 36px;
  font-weight: 700;
}

.orbit-card {
  position: absolute;
  left: 0;
  top: 0;
  z-index: var(--orbit-z);
  width: var(--orbit-width);
  height: 80px;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 8px 24px rgba(31, 38, 56, 0.08);
  display: grid;
  place-items: center;
  transform: translate(var(--orbit-x), var(--orbit-y));
  transition: transform 850ms cubic-bezier(0.22, 1, 0.36, 1), width 850ms cubic-bezier(0.22, 1, 0.36, 1), box-shadow 500ms ease, border-color 500ms ease;
}

.orbit-card.active {
  border: 3px solid #ff5a55;
  box-shadow: 0 10px 26px rgba(251, 72, 72, 0.14);
}

.orbit-card img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}

@media (max-width: 1500px) {
  .brand-shell {
    width: 1280px;
    grid-template-columns: 440px 1fr;
  }

  .brand-map {
    transform: scale(0.92);
    transform-origin: left top;
    width: 1000px;
    margin-bottom: -64px;
  }
}

@media (max-width: 1240px) {
  .brand-shell {
    width: calc(100% - 40px);
    grid-template-columns: 1fr;
  }

  .brand-card-viewport {
    height: 420px;
  }

  .brand-card-list {
    width: 440px;
    margin: 0 auto;
    padding-top: 112px;
  }

  .brand-map {
    transform: scale(0.72);
    transform-origin: top center;
    margin-bottom: -220px;
  }
}

@media (max-width: 760px) {
  .about-brands {
    background: #fff7f7;
  }

  .brand-shell {
    width: 100%;
    min-height: 0;
    padding: 28px 18px 30px;
    overflow: hidden;
  }

  .brand-card-viewport {
    height: 150px;
  }

  .brand-card-list {
    width: 100%;
    padding-top: 0;
  }

  .brand-card {
    width: 100%;
    height: 128px;
    margin-bottom: 22px;
    padding: 18px 20px;
  }

  .brand-card img {
    width: 96px;
    height: 36px;
  }

  .brand-card h3 {
    margin: 12px 0 6px;
    font-size: 16px;
    line-height: 22px;
  }

  .brand-card p {
    font-size: 12px;
    line-height: 18px;
    -webkit-line-clamp: 2;
  }

  .brand-map {
    width: 100%;
    max-width: 100%;
    height: 230px;
    margin: 0;
    overflow: visible;
    transform: none;
  }

  .brand-map-bg {
    left: 50%;
    top: 0;
    width: 304px;
    height: 230px;
    transform: translateX(-50%);
  }

  .brand-center {
    left: 50%;
    top: 87px;
    width: 64px;
    height: 64px;
    transform: translateX(-50%);
    font-size: 14px;
    line-height: 20px;
  }

  .orbit-card {
    width: calc(var(--orbit-width) * 0.3);
    height: 25px;
    border-radius: 4px;
    transform: translate(
      calc((100vw - 36px - 304px) / 2 + var(--orbit-x) * 0.3),
      calc(var(--orbit-y) * 0.282)
    );
  }

  .orbit-card.active {
    border-width: 2px;
  }
}
</style>
