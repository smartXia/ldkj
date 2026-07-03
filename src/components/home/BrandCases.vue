<script setup>
import { onMounted, onUnmounted, shallowRef } from 'vue'
import { brandCases } from '../../data/siteContent'
import SectionHeading from './SectionHeading.vue'

const sectionTitle = '\u54c1\u724c\u5b9a\u5236\u670d\u52a1'
const sectionSubtitle = '\u4ee5\u54c1\u724c\u76ee\u6807\u4e3a\u9a71\u52a8\u7684\u5b9a\u5236\u5316\u8425\u9500\u89e3\u51b3\u65b9\u6848'
const detailLabel = '\u67e5\u770b\u6848\u4f8b\u8be6\u60c5 \u300b'
const moreLabel = '\u67e5\u770b\u66f4\u591a\u6848\u4f8b \u300b'
const dotsLabel = '\u54c1\u724c\u5b9a\u5236\u670d\u52a1\u6848\u4f8b\u5207\u6362'

const activeIndex = shallowRef(0)
const moveDirection = shallowRef('next')
let timer

const brandBadges = [
  { logo: '\u7c73', name: '\u7c73\u5bb6', className: 'mijia' },
  { logo: '\u73c0', name: '\u73c0\u83b1\u96c5', className: 'proya' },
  { logo: 'XGIMI', name: '\u6781\u7c73', className: 'xgimi' },
  { logo: '\u7f8e\u56e2', name: '\u7f8e\u56e2', className: 'meituan' },
  { logo: '\u6296', name: '\u6296\u97f3', className: 'douyin' },
  { logo: '\u6d77', name: '\u6d77\u4fe1', className: 'hisense' },
  { logo: '\u65b0', name: '\u65b0\u4e1c\u65b9', className: 'xdf' },
  { logo: '\u65b9', name: '\u65b9\u592a', className: 'fotile' }
]

function setActive(index) {
  if (index === activeIndex.value) return
  moveDirection.value = index > activeIndex.value || (activeIndex.value === brandCases.length - 1 && index === 0)
    ? 'next'
    : 'prev'
  activeIndex.value = index
}

function nextCase() {
  setActive((activeIndex.value + 1) % brandCases.length)
}

function startAutoplay() {
  stopAutoplay()
  timer = window.setInterval(nextCase, 3600)
}

function stopAutoplay() {
  if (!timer) return
  window.clearInterval(timer)
  timer = undefined
}

function previewCase(index) {
  stopAutoplay()
  setActive(index)
}

function dotLabel(title) {
  return `\u9884\u89c8${title}`
}

onMounted(startAutoplay)
onUnmounted(stopAutoplay)
</script>

<template>
  <section class="case-section">
    <SectionHeading :title="sectionTitle" :subtitle="sectionSubtitle" />

    <div class="case-viewport" @mouseenter="stopAutoplay" @mouseleave="startAutoplay">
      <div class="case-rail" :class="`move-${moveDirection}`">
        <article
          v-for="(item, index) in brandCases"
          :key="item.number"
          class="case-item"
          :class="{ active: activeIndex === index }"
          @mouseenter="previewCase(index)"
          @focusin="previewCase(index)"
        >
          <div class="case-expanded">
            <div class="case-photo">
              <img :src="item.image" :alt="item.summary" loading="lazy" />
              <div class="brand-badge">
                <span :class="brandBadges[index].className">{{ brandBadges[index].logo }}</span>
                <strong>{{ brandBadges[index].name }}</strong>
              </div>
            </div>
            <div class="case-active-body">
              <span class="case-number">{{ item.number }}</span>
              <h3>{{ item.title }}</h3>
              <p>{{ item.summary }}</p>
              <RouterLink class="case-detail" to="/case">{{ detailLabel }}</RouterLink>
            </div>
          </div>

          <div class="case-static">
            <span class="case-number muted">{{ item.number }}</span>
            <h3>{{ item.title }}</h3>
            <p>{{ item.summary }}</p>
          </div>
        </article>
      </div>
    </div>

    <div class="case-dots" :aria-label="dotsLabel">
      <button
        v-for="(item, index) in brandCases"
        :key="item.number"
        type="button"
        :aria-label="dotLabel(item.title)"
        :aria-current="activeIndex === index"
        :class="{ active: activeIndex === index }"
        @mouseenter="previewCase(index)"
        @focus="previewCase(index)"
      ></button>
    </div>

    <RouterLink class="more-cases" to="/case">{{ moreLabel }}</RouterLink>
  </section>
</template>

<style scoped>
.case-section {
  padding: 96px 0 86px;
  background: #fff;
  overflow: visible;
}

.case-viewport {
  width: 100%;
  margin-top: 62px;
  overflow: hidden;
}

.case-rail {
  display: flex;
  align-items: flex-end;
  gap: 16px;
  width: max-content;
  height: 320px;
}

.case-item {
  position: relative;
  flex: 0 0 208px;
  height: 320px;
  border-radius: 28px;
  background: transparent;
  color: #050506;
  cursor: pointer;
  overflow: visible;
  transform: translateZ(0);
  transition:
    flex-basis 560ms cubic-bezier(.22, 1, .36, 1);
  will-change: flex-basis;
}

.case-item.active {
  flex-basis: 464px;
  cursor: default;
}

.case-static {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 280px;
  padding: 50px 16px 24px;
  border-radius: 28px;
  background: #f0f2f5;
  opacity: 1;
  transform: translateX(0);
  transition:
    opacity 260ms ease,
    transform 520ms cubic-bezier(.2, .8, .2, 1);
  will-change: opacity, transform;
}

.case-item.active .case-static {
  opacity: 0;
  pointer-events: none;
}

.case-rail.move-next .case-item.active .case-static {
  transform: translateX(-26px);
}

.case-rail.move-prev .case-item.active .case-static {
  transform: translateX(26px);
}

.case-expanded {
  position: absolute;
  inset: 0;
  display: grid;
  grid-template-columns: 240px 224px;
  width: 464px;
  height: 320px;
  opacity: 0;
  pointer-events: none;
  transform: translateX(44px);
  transition:
    opacity 240ms ease 120ms,
    transform 560ms cubic-bezier(.22, 1, .36, 1);
  will-change: opacity, transform;
}

.case-rail.move-prev .case-expanded {
  transform: translateX(-44px);
}

.case-item.active .case-expanded {
  opacity: 1;
  pointer-events: auto;
  transform: translateX(0);
}

.case-number {
  position: absolute;
  top: 28px;
  left: 16px;
  z-index: 1;
  color: rgba(255, 255, 255, 0.76);
  font-size: 40px;
  line-height: 1;
  font-weight: 900;
}

.case-static .case-number {
  color: rgba(255, 255, 255, 0.96);
}

.case-static h3,
.case-active-body h3 {
  position: relative;
  z-index: 2;
  margin: 0 0 18px;
  font-size: 25px;
  line-height: 32px;
  font-weight: 900;
  color: #020202;
}

.case-static p,
.case-active-body p {
  position: relative;
  z-index: 2;
  margin: 0;
  color: #646a73;
  font-size: 17px;
  line-height: 28px;
}

.case-photo {
  position: relative;
  height: 320px;
  border-radius: 26px;
  overflow: hidden;
  background: #ddd;
  box-shadow: 0 18px 34px rgba(22, 32, 52, 0.12);
}

.case-photo::after {
  content: "";
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.2);
}

.case-photo img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
}

.brand-badge {
  position: absolute;
  left: 32px;
  bottom: 32px;
  z-index: 2;
  display: flex;
  align-items: center;
  gap: 14px;
  color: #fff;
}

.brand-badge span {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 18px;
  font-weight: 900;
  white-space: nowrap;
}

.brand-badge span.xgimi {
  background: #fff;
  color: #101010;
  font-size: 13px;
}

.brand-badge span.mijia,
.brand-badge span.proya,
.brand-badge span.hisense,
.brand-badge span.xdf,
.brand-badge span.fotile {
  background: #18c76b;
}

.brand-badge span.meituan {
  background: #ffd21f;
  color: #111;
}

.brand-badge span.douyin {
  background: #111;
}

.brand-badge strong {
  font-size: 18px;
  font-weight: 800;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.case-active-body {
  position: relative;
  height: 280px;
  margin-top: 40px;
  border-radius: 0 28px 28px 0;
  padding: 64px 28px 28px 32px;
  background: #ff444b;
  color: #fff;
}

.case-active-body .case-number {
  top: 40px;
  left: 32px;
}

.case-active-body h3 {
  color: #fff;
  white-space: nowrap;
  font-size: 24px;
  line-height: 32px;
}

.case-active-body p {
  color: #fff;
  font-size: 16px;
  line-height: 26px;
  font-weight: 700;
}

.case-detail {
  position: absolute;
  left: 32px;
  bottom: 30px;
  min-width: 160px;
  height: 40px;
  border: 1px solid #fff;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 16px;
  font-weight: 800;
}

.case-dots {
  margin: 39px 0 32px;
  display: flex;
  justify-content: center;
  gap: 16px;
}

.case-dots button {
  width: 10px;
  height: 10px;
  border: 0;
  border-radius: 50%;
  padding: 0;
  background: #ececec;
}

.case-dots button.active {
  background: var(--color-brand);
}

.more-cases {
  width: fit-content;
  min-height: 30px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  font-size: 19px;
  font-weight: 500;
}

@media (max-width: 980px) {
  .case-viewport {
    overflow-x: auto;
    padding: 0 20px 12px;
  }
}

@media (max-width: 640px) {
  .case-section {
    padding: 30px 0 28px;
  }

  :deep(.section-heading) {
    margin-bottom: 18px;
  }

  :deep(.section-heading h2) {
    font-size: 20px;
    line-height: 28px;
  }

  :deep(.section-heading p) {
    width: calc(100% - 44px);
    margin: 3px auto 0;
    font-size: 11px;
    line-height: 16px;
  }

  .case-viewport {
    width: calc(100% - 24px);
    margin: 0 auto;
    padding: 0;
    overflow: hidden;
  }

  .case-rail {
    width: 100%;
    height: 160px;
    display: block;
  }

  .case-item {
    display: none;
  }

  .case-item.active {
    display: block;
    width: 100%;
    height: 160px;
    border-radius: 10px;
  }

  .case-expanded {
    width: 100%;
    height: 160px;
    grid-template-columns: 50% 50%;
    opacity: 1;
    transform: none;
    pointer-events: auto;
  }

  .case-photo {
    height: 160px;
    border-radius: 10px 0 0 10px;
  }

  .case-active-body {
    height: 137px;
    margin-top: 0;
    align-self: end;
    border-radius: 0 10px 10px 0;
    padding: 31px 11px 14px 13px;
  }

  .case-active-body .case-number {
    top: 12px;
    left: 13px;
    font-size: 28px;
  }

  .case-active-body h3 {
    margin-bottom: 7px;
    font-size: 15px;
    line-height: 20px;
    white-space: normal;
  }

  .case-active-body p {
    font-size: 10px;
    line-height: 15px;
  }

  .case-detail {
    left: 13px;
    bottom: 12px;
    min-width: 86px;
    height: 24px;
    font-size: 10px;
  }

  .brand-badge {
    left: 18px;
    bottom: 18px;
    gap: 7px;
  }

  .brand-badge span {
    width: 28px;
    height: 28px;
    border-radius: 6px;
    font-size: 10px;
  }

  .brand-badge strong {
    font-size: 11px;
  }

  .case-dots {
    margin: 14px 0 14px;
    gap: 10px;
  }

  .case-dots button {
    width: 5px;
    height: 5px;
  }

  .more-cases {
    min-height: 18px;
    font-size: 11px;
  }
}
</style>
