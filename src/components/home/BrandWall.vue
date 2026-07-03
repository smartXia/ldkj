<script setup>
import { computed, nextTick, onMounted, onUnmounted, shallowRef } from 'vue'

const title = '合作品牌'
const subtitle = '微思敦深耕3C家电、美妆日化、母婴宠物、食品饮料、房产家居、网服教育等行业，持续服务客户3000+，以系统标准化服务流程，让品牌获得切实可行的社交营销解决方案，为品牌增长加速带来更多的可能。'

const assetBase = '/assets/wsd/brand-wall'
const sectionRef = shallowRef(null)
const mainRef = shallowRef(null)
const visible = shallowRef(false)
const settled = shallowRef(false)
const trackItems = shallowRef([])
let observer
let settleTimer
let frameId

const otherBrands = [
  ['追觅', 'brand-other-01.png'],
  ['学而思', 'brand-other-02.jpeg'],
  ['武汉城投', 'brand-other-03.jpg'],
  ['潘婷', 'brand-other-04.jpg'],
  ['林氏家居', 'brand-other-05.png'],
  ['蔻蔻椰', 'brand-other-06.jpg'],
  ['可优比', 'brand-other-07.png'],
  ['华帝', 'brand-other-08.jpg'],
  ['海信', 'brand-other-09.png'],
  ['高露洁', 'brand-other-10.png'],
  ['锋味派', 'brand-other-11.png'],
  ['东芝', 'brand-other-12.jpg'],
  ['新东方', 'brand-other-13.png'],
  ['空刻', 'brand-other-14.png'],
  ['宝宝巴士', 'brand-other-15.jpg']
].map(([name, file]) => ({ name, image: `${assetBase}/${file}`, isOther: true }))

const industries = [
  {
    title: '3C家电',
    brands: [
      ['方太', 'brand-01.png'],
      ['小度', 'brand-02.png'],
      ['徕芬', 'brand-13.jpg'],
      ['极米', 'brand-14.jpg'],
      ['九阳', 'brand-23.png'],
      ['科沃斯', 'brand-24.jpg'],
      ['vivo', 'brand-33.png'],
      ['Haier', 'brand-34.png']
    ]
  },
  {
    title: '美妆日化',
    brands: [
      ['珀莱雅', 'brand-03.png'],
      ['卡姿兰', 'brand-04.jpg'],
      ['Olay', 'brand-15.jpg'],
      ['逐本', 'brand-16.jpg'],
      ['名创优品', 'brand-25.jpg'],
      ['维达', 'brand-26.png'],
      ['舒肤佳', 'brand-35.jpg'],
      ['ABC', 'brand-36.png']
    ]
  },
  {
    title: '母婴宠物',
    brands: [
      ['合生元', 'brand-05.png'],
      ['帮宝适', 'brand-06.png'],
      ['Babycare', 'brand-17.png'],
      ['英氏', 'brand-18.png'],
      ['素力高', 'brand-27.png'],
      ['生鲜本能', 'brand-28.jpg'],
      ['Hill’s', 'brand-37.jpg'],
      ['疯狂小狗', 'brand-pet-dog.png']
    ]
  },
  {
    title: '食品饮料',
    brands: [
      ['燕之屋', 'brand-07.png'],
      ['东阿阿胶', 'brand-08.png'],
      ['汤臣倍健', 'brand-food-01.jpg'],
      ['雀巢', 'brand-food-02.png'],
      ['隅田川', 'brand-29.png'],
      ['WonderLab', 'brand-30.jpg'],
      ['简爱', 'brand-food-03.jpg'],
      ['霸王茶姬', 'brand-38.jpg']
    ]
  },
  {
    title: '房产家居',
    brands: [
      ['中建壹品', 'brand-09.png'],
      ['武汉城建', 'brand-10.png'],
      ['米家', 'brand-19.png'],
      ['东鹏瓷砖', 'brand-20.png'],
      ['梦百合', 'brand-31.png'],
      ['住范儿', 'brand-32.png'],
      ['华润置地', 'brand-39.png'],
      ['华侨城', 'brand-40.jpg']
    ]
  },
  {
    title: '网服教育',
    brands: [
      ['美团', 'brand-11.png'],
      ['Soul', 'brand-12.png'],
      ['唯品会', 'brand-21.png'],
      ['米哈游', 'brand-22.png'],
      ['牵手APP', 'brand-app-01.png'],
      ['博为峰', 'brand-edu-01.jpg'],
      ['青藤', 'brand-edu-02.png'],
      ['自如', 'brand-edu-03.png']
    ]
  }
].map((industry) => ({
  ...industry,
  brands: industry.brands.map(([name, file]) => ({ name, image: `${assetBase}/${file}` }))
}))

const finalBrands = computed(() => industries.flatMap((industry) => industry.brands))
const finalPositionMap = computed(() => {
  const map = new Map()

  industries.forEach((industry, industryIndex) => {
    industry.brands.forEach((brand, brandIndex) => {
      const col = brandIndex % 2
      const row = Math.floor(brandIndex / 2)
      map.set(brand.name, {
        x: industryIndex * 200 + 16 + col * 92,
        y: 96 + row * 96
      })
    })
  })

  return map
})

function shuffle(list) {
  return [...list].sort((a, b) => a.name.localeCompare(b.name, 'zh-Hans-CN'))
}

function buildTrackItems() {
  const rows = [[], [], []]
  const other = shuffle(otherBrands)
  rows[0].push(...other.slice(0, 5))
  rows[1].push(...other.slice(5, 10))
  rows[2].push(...other.slice(10, 15))

  for (let groupIndex = 0; groupIndex < 4; groupIndex += 1) {
    const grouped = shuffle(finalBrands.value.slice(groupIndex * 12, groupIndex * 12 + 12))
    rows[0].push(...grouped.slice(0, 4))
    rows[1].push(...grouped.slice(4, 8))
    rows[2].push(...grouped.slice(8, 12))
  }

  trackItems.value = rows.flatMap((row, rowIndex) =>
    row.map((brand, colIndex) => ({
      ...brand,
      rowIndex,
      colIndex,
      id: `${brand.name}-${rowIndex}-${colIndex}`,
      style: {
        '--row': rowIndex,
        '--col': colIndex,
        '--start-x': `${(135 * colIndex + (rowIndex === 1 ? 70 : 0)) - 360}px`,
        '--end-x': `${(135 * colIndex + (rowIndex === 1 ? 70 : 0)) - 1160}px`,
        '--row-y': `${135 * rowIndex + 96}px`,
        '--final-x': `${finalPositionMap.value.get(brand.name)?.x ?? 0}px`,
        '--final-y': `${finalPositionMap.value.get(brand.name)?.y ?? 0}px`
      }
    }))
  )
}

async function playEntrance() {
  if (visible.value) return
  visible.value = true
  buildTrackItems()
  await nextTick()

  frameId = window.requestAnimationFrame(() => {
    sectionRef.value?.classList.add('is-running')
  })

  window.clearTimeout(settleTimer)
  settleTimer = window.setTimeout(() => {
    settled.value = true
  }, 5600)
}

onMounted(() => {
  observer = new IntersectionObserver(([entry]) => {
    if (entry.isIntersecting) playEntrance()
  }, { threshold: 0.5 })

  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => {
  observer?.disconnect()
  window.clearTimeout(settleTimer)
  window.cancelAnimationFrame(frameId)
})
</script>

<template>
  <section ref="sectionRef" class="brand-wall" :class="{ visible, settled }">
    <h2 class="brand-title">{{ title }}</h2>
    <p class="brand-desc">{{ subtitle }}</p>

    <div ref="mainRef" class="brand-main">
      <div class="industry-list">
        <div class="industry-item industry-item--other" aria-hidden="true">
          <h3 class="industry-title">其他</h3>
          <ul class="brand-list">
            <li v-for="brand in otherBrands" :key="brand.name" class="brand-item">
              <img class="brand-img" :src="brand.image" :alt="brand.name" loading="lazy" />
            </li>
          </ul>
        </div>

        <div v-for="industry in industries" :key="industry.title" class="industry-item">
          <h3 class="industry-title">{{ industry.title }}</h3>
          <ul class="brand-list">
            <li v-for="brand in industry.brands" :key="brand.name" class="brand-item">
              <img class="brand-img" :src="brand.image" :alt="brand.name" loading="lazy" />
            </li>
          </ul>
        </div>
      </div>

      <div v-if="!settled" class="motion-layer" aria-hidden="true">
        <div
          v-for="brand in trackItems"
          :key="brand.id"
          class="motion-brand"
          :class="{ 'is-other': brand.isOther }"
          :style="brand.style"
        >
          <img class="brand-img" :src="brand.image" :alt="brand.name" loading="lazy" />
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.brand-wall {
  height: 780px;
  padding-top: 80px;
  overflow: hidden;
  background: #f4f4f4;
  text-align: center;
}

.brand-title {
  margin: 0 0 20px;
  color: #000;
  font-family: PingFang SC, Microsoft YaHei, sans-serif;
  font-size: 40px;
  font-weight: 600;
  line-height: 48px;
}

.brand-desc {
  width: 1080px;
  max-width: calc(100% - 48px);
  margin: 0 auto 40px;
  color: #666;
  font-size: 20px;
  line-height: 28px;
}

.brand-main {
  position: relative;
  width: 1200px;
  height: 522px;
  margin: 0 auto;
}

.industry-list {
  position: relative;
  display: flex;
  justify-content: space-between;
}

.industry-item {
  width: 200px;
  display: flex;
  flex-direction: column;
  align-items: center;
  flex-wrap: wrap;
}

.industry-item--other {
  position: absolute;
  left: 0;
  top: 0;
}

.industry-title {
  width: 150px;
  height: 48px;
  margin: 0 0 48px;
  border-radius: 24px;
  background: #e8e8e8;
  color: #666;
  font-size: 20px;
  font-weight: 400;
  line-height: 48px;
  text-align: center;
}

.brand-list {
  width: 100%;
  padding: 0 16px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.brand-item,
.motion-brand {
  width: 76px;
  height: 76px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background: #fff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, .04);
}

.brand-item {
  margin-bottom: 20px;
  opacity: 0;
}

.settled .brand-item {
  opacity: 1;
  transition: opacity 500ms ease;
}

.industry-item--other .industry-title,
.industry-item--other .brand-item {
  opacity: 0;
}

.brand-img {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  object-fit: contain;
  display: block;
}

.motion-layer {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 2;
}

.motion-brand {
  position: absolute;
  left: 0;
  top: 0;
  opacity: 0;
  transform: translate3d(var(--start-x), var(--row-y), 0) scale(1.25);
  border-radius: 20px;
  will-change: transform, opacity, border-radius;
}

.is-running .motion-brand {
  animation: brandFly 5.5s linear forwards;
}

.is-running .motion-brand.is-other {
  animation-name: brandFlyOther;
}

@keyframes brandFly {
  0% {
    opacity: 1;
    transform: translate3d(var(--start-x), var(--row-y), 0) scale(1.25);
    border-radius: 20px;
  }
  91% {
    opacity: 1;
    transform: translate3d(var(--end-x), var(--row-y), 0) scale(1.25);
    border-radius: 20px;
  }
  100% {
    opacity: 1;
    transform: translate3d(var(--final-x), var(--final-y), 0) scale(1);
    border-radius: 8px;
  }
}

@keyframes brandFlyOther {
  0% {
    opacity: 1;
    transform: translate3d(var(--start-x), var(--row-y), 0) scale(1.25);
    border-radius: 20px;
  }
  91% {
    opacity: 1;
    transform: translate3d(var(--end-x), var(--row-y), 0) scale(1.25);
    border-radius: 20px;
  }
  100% {
    opacity: 0;
    transform: translate3d(var(--final-x), var(--final-y), 0) scale(1);
    border-radius: 8px;
  }
}

@media (min-width: 1081px) and (max-width: 1280px) {
  .brand-wall {
    height: 732px;
  }

  .brand-desc {
    width: 968px;
  }

  .brand-main {
    transform: scale(.867) translateY(-32px);
  }
}

@media (min-width: 768px) and (max-width: 1080px) {
  .brand-wall {
    height: 720px;
  }

  .brand-desc {
    width: 968px;
  }

  .brand-main {
    transform: scale(.8) translateY(-60px);
  }
}

@media (max-width: 767px) {
  .brand-wall {
    height: auto;
    padding: 29px 0 24px;
  }

  .brand-title {
    margin-bottom: 4px;
    font-size: 20px;
    line-height: 28px;
  }

  .brand-desc {
    width: calc(100% - 32px);
    margin-bottom: 24px;
    font-size: 12px;
    line-height: 16px;
  }

  .brand-main {
    width: calc(100% - 16px);
    height: auto;
  }

  .industry-list {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 10px 6px;
  }

  .industry-item:nth-child(n+5) {
    display: none;
  }

  .industry-item {
    width: auto;
  }

  .industry-title {
    width: 100px;
    height: 28px;
    margin-bottom: 18px;
    border-radius: 14px;
    font-size: 14px;
    line-height: 28px;
  }

  .brand-list {
    padding: 0 2px;
    gap: 8px 4px;
    justify-content: center;
  }

  .brand-item {
    width: 36px;
    height: 36px;
    margin-bottom: 0;
  }

  .brand-img {
    width: 31px;
    height: 31px;
  }

  .motion-layer,
  .industry-item--other {
    display: none;
  }

  .brand-item {
    opacity: 1;
  }
}
</style>
