<script setup>
import { computed, nextTick, onMounted, onUnmounted, shallowRef, watch } from 'vue'
import SectionHeading from './SectionHeading.vue'

const props = defineProps({
  services: {
    type: Array,
    default: () => [],
  },
})

const tabs = ['社媒广告', '达人内容', '营销培训', '社交电商', '口碑管理', '营销技术']
const activeIndex = shallowRef(0)
const sectionRef = shallowRef(null)
const videoRef = shallowRef(null)
const visible = shallowRef(false)
let observer

const platformCards = [
  { name: '微信', desc: '全域经营', className: 'wechat' },
  { name: '微博', desc: '流量曝光', className: 'weibo' },
  { name: '小红书', desc: '场景种草', className: 'redbook' },
  { name: '哔哩哔哩', desc: '年轻潮流', className: 'bilibili' },
  { name: '知乎', desc: '信任背书', className: 'zhihu' }
]

const platformLogos = [
  { name: 'f', className: 'facebook' },
  { name: '◎', className: 'instagram' },
  { name: '𝕏', className: 'x' },
  { name: 'in', className: 'linkedin' },
  { name: 'r', className: 'reddit' },
  { name: 'S', className: 'snapchat' },
  { name: 'p', className: 'pinterest' },
  { name: '...', className: 'more' }
]

const trainingItems = [
  {
    title: '品牌社媒运营',
    desc: '品牌号/品牌矩阵账号运营培训\n助力声量与销量双丰收',
    icon: 'doc'
  },
  {
    title: '社交电商运营',
    desc: '品牌小红书电商直播培训\n轻松实现种销一体',
    icon: 'shop'
  },
  {
    title: '品牌KOS营销',
    desc: 'KOS号/门店账号运营培训\n帮助门店获得生意结果',
    icon: 'grid'
  },
  {
    title: '社交广告运营',
    desc: '营销投放技能培训\n精准覆盖提高转化效果',
    icon: 'gear'
  }
]

const trainingPlatforms = [
  { label: '微信', className: 'wechat' },
  { label: '微博', className: 'weibo' },
  { label: '小红书', className: 'redbook' },
  { label: '哔哩', className: 'bilibili' },
  { label: '知乎', className: 'zhihu' }
]

const serviceItems = computed(() => props.services)
const activeService = computed(() => serviceItems.value[activeIndex.value] || {})
const videoSrc = computed(() => '')
const isAdPanel = computed(() => activeIndex.value === 0)
const isTrainingPanel = computed(() => activeIndex.value === 2)

function setActive(index) {
  activeIndex.value = index
}

async function playCurrentVideo() {
  await nextTick()
  const video = videoRef.value
  if (!video) return
  video.load()
  video.play().catch(() => {})
}

function toggleVideo() {
  const video = videoRef.value
  if (!video) return
  if (video.paused) video.play().catch(() => {})
  else video.pause()
}

watch(activeIndex, playCurrentVideo)

onMounted(() => {
  observer = new IntersectionObserver(([entry]) => {
    if (!entry.isIntersecting) return
    visible.value = true
    playCurrentVideo()
  }, { threshold: 0.22 })

  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => observer?.disconnect())
</script>

<template>
  <section ref="sectionRef" class="service-section" :class="{ visible }">
    <SectionHeading title="社交营销服务" subtitle="连接品牌与世界 激发社交营销力" />

    <div class="service-tabs-wrap">
      <div class="service-tabs" role="tablist" aria-label="社交营销服务">
        <button
          v-for="(tab, index) in tabs"
          :key="tab"
          type="button"
          role="tab"
          :aria-selected="activeIndex === index"
          :class="{ active: activeIndex === index }"
          @click="setActive(index)"
        >
          {{ tab }}
        </button>
      </div>
    </div>

    <div class="service-body">
      <div class="service-inner">
        <article class="service-copy">
          <Transition name="service-fade" mode="out-in">
            <div :key="activeIndex" class="panel-content">
              <template v-if="isAdPanel">
                <h3>社交媒体深度商业合作伙伴</h3>
                <div class="ad-platforms">
                  <div v-for="item in platformCards" :key="item.name" class="platform-card">
                    <span class="platform-app" :class="item.className">{{ item.name }}</span>
                    <span class="platform-desc">{{ item.desc }}</span>
                  </div>
                </div>
                <div class="oversea-platforms" aria-label="海外社交平台">
                  <span
                    v-for="item in platformLogos"
                    :key="item.className"
                    class="oversea-icon"
                    :class="item.className"
                  >
                    {{ item.name }}
                  </span>
                </div>
              </template>

              <template v-else-if="isTrainingPanel">
                <h3>营销培训服务提供的服务项目及平台</h3>
                <div class="training-grid">
                  <div v-for="item in trainingItems" :key="item.title" class="training-card">
                    <span class="training-icon" :class="item.icon"></span>
                    <span class="training-text">
                      <strong>{{ item.title }}</strong>
                      <small>{{ item.desc }}</small>
                    </span>
                  </div>
                </div>
                <div class="training-platform-bar">
                  <strong>培训课程涵盖主流营销平台</strong>
                  <span class="training-platform-list">
                    <span
                      v-for="item in trainingPlatforms"
                      :key="item.label"
                      class="training-platform"
                      :class="item.className"
                    >
                      {{ item.label }}
                    </span>
                  </span>
                </div>
              </template>

              <template v-else>
                <h3>{{ activeService.title }}</h3>
                <img v-if="activeService.image" class="service-image" :src="activeService.image" :alt="activeService.title" loading="lazy" />
              </template>

              <a class="service-more" href="#consult">了解更多详情 »</a>
            </div>
          </Transition>
        </article>

        <button v-if="videoSrc" class="phone-demo" type="button" aria-label="播放或暂停手机视频" @click="toggleVideo">
          <span class="phone-notch"></span>
          <video
            ref="videoRef"
            class="phone-video"
            :src="videoSrc"
            muted
            loop
            playsinline
            preload="auto"
          ></video>
        </button>
      </div>
    </div>
  </section>
</template>

<style scoped>
.service-section {
  padding: 70px 0 0;
  background: #fff;
  overflow: visible;
}

.service-tabs-wrap {
  margin-top: 50px;
  background: #f8f9fb;
}

.service-tabs {
  width: min(1188px, calc(100% - 48px));
  height: 72px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(6, 1fr);
}

.service-tabs button {
  position: relative;
  min-height: 72px;
  border: 0;
  background: transparent;
  color: #11151d;
  font-size: 18px;
  font-weight: 800;
  transition: color var(--motion-fast), background var(--motion-fast);
}

.service-tabs button.active {
  color: var(--color-brand);
  background: #ffe8e9;
}

.service-tabs button.active::after {
  content: "";
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 3px;
  background: var(--color-brand);
}

.service-body {
  min-height: 594px;
  background: #eef2f7;
  overflow: visible;
}

.service-inner {
  position: relative;
  width: min(1188px, calc(100% - 48px));
  height: 594px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 760px 300px;
  justify-content: space-between;
  align-items: start;
}

.service-copy {
  padding-top: 60px;
  opacity: 0;
  transform: translateY(38px);
  transition: opacity 880ms ease, transform 880ms cubic-bezier(.16, 1, .3, 1);
}

.panel-content {
  min-height: 420px;
}

.panel-content h3 {
  margin: 0 0 28px;
  font-size: 28px;
  line-height: 40px;
  font-weight: 800;
  color: var(--color-ink);
}

.ad-platforms {
  width: 748px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 18px;
}

.platform-card {
  height: 180px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 18px;
  background: #fff;
  box-shadow: 0 12px 28px rgba(24, 37, 62, 0.1);
}

.platform-app {
  width: 62px;
  height: 62px;
  border-radius: 14px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 17px;
  font-weight: 800;
}

.platform-desc {
  color: #8a8f99;
  font-size: 16px;
  font-weight: 500;
}

.wechat { background: #18c76b; }
.weibo { background: linear-gradient(135deg, #ffd54d 0%, #ff363a 100%); }
.redbook { background: #ff3046; }
.bilibili { background: #fb5a9c; }
.zhihu { background: #1478ff; }

.oversea-platforms {
  width: 748px;
  height: 70px;
  margin-top: 18px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 14px;
  background: #fff;
  box-shadow: 0 12px 26px rgba(24, 37, 62, 0.1);
}

.oversea-icon {
  width: 42px;
  height: 42px;
  border-radius: 8px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 21px;
  line-height: 1;
  font-weight: 900;
}

.oversea-icon.facebook { background: #1877f2; }
.oversea-icon.instagram { background: radial-gradient(circle at 30% 105%, #feda75 0%, #fa7e1e 28%, #d62976 58%, #962fbf 78%, #4f5bd5 100%); }
.oversea-icon.x { background: #000; }
.oversea-icon.linkedin { background: #0a66c2; font-size: 18px; }
.oversea-icon.reddit { background: #ff4500; text-transform: uppercase; }
.oversea-icon.snapchat { background: #fffc00; color: #111; }
.oversea-icon.pinterest { background: #e60023; }
.oversea-icon.more { background: #ffeaed; color: #ff5660; font-size: 19px; }

.training-grid {
  width: 650px;
  margin-left: 34px;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 14px;
}

.training-card {
  height: 112px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  gap: 22px;
  padding: 0 26px;
  background: #fff;
  box-shadow: 0 14px 30px rgba(26, 38, 66, 0.11);
}

.training-icon {
  position: relative;
  flex: 0 0 auto;
  width: 48px;
  height: 48px;
  border-radius: 13px;
  background: linear-gradient(135deg, #ff868a 0%, #ff3d45 100%);
}

.training-icon::after {
  content: "";
  position: absolute;
  right: -5px;
  bottom: -5px;
  width: 22px;
  height: 22px;
  border: 3px solid #fff;
  border-radius: 50%;
  background: rgba(255, 61, 69, 0.82);
}

.training-icon.doc::before {
  content: "";
  position: absolute;
  left: 11px;
  top: 11px;
  width: 25px;
  height: 4px;
  border-radius: 4px;
  background: #fff;
  box-shadow: 0 10px 0 #fff, 0 20px 0 #fff;
}

.training-icon.shop::before,
.training-icon.gear::before,
.training-icon.grid::before {
  content: "";
  position: absolute;
  inset: 12px;
  border: 5px solid #fff;
  border-radius: 50%;
}

.training-icon.grid::before {
  inset: 10px;
  border: 0;
  border-radius: 4px;
  background: #fff;
  box-shadow: 18px 0 0 #fff, 0 18px 0 #fff, 18px 18px 0 #fff;
  width: 12px;
  height: 12px;
}

.training-text {
  display: grid;
  gap: 4px;
  color: #3c4350;
}

.training-text strong {
  font-size: 17px;
  line-height: 24px;
  color: #30343b;
}

.training-text small {
  white-space: pre-line;
  color: #5c6470;
  font-size: 13px;
  line-height: 20px;
}

.training-platform-bar {
  width: 650px;
  height: 46px;
  margin: 14px 0 0 34px;
  border-radius: 999px;
  display: grid;
  grid-template-columns: 300px 1fr;
  align-items: center;
  overflow: hidden;
  background: #ff5158;
  box-shadow: 0 12px 26px rgba(255, 61, 69, 0.18);
}

.training-platform-bar strong {
  color: #fff;
  font-size: 17px;
  text-align: center;
}

.training-platform-list {
  height: 38px;
  margin-right: 8px;
  border: 3px solid rgba(255, 93, 100, 0.82);
  border-radius: 999px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  background: #fff;
}

.training-platform {
  min-width: 36px;
  height: 28px;
  border-radius: 7px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 7px;
  color: #fff;
  font-size: 12px;
  font-weight: 800;
}

.service-image {
  width: 748px;
  height: auto;
  display: block;
  object-fit: contain;
  filter: drop-shadow(0 18px 28px rgba(30, 48, 78, 0.08));
}

.service-more {
  min-width: 168px;
  min-height: 44px;
  margin-top: 36px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: var(--color-brand);
  color: #fff;
  font-size: 17px;
  font-weight: 800;
  box-shadow: var(--shadow-soft);
}

.phone-demo {
  position: relative;
  justify-self: end;
  width: 270px;
  height: 548px;
  margin-top: 28px;
  border: 8px solid #30333a;
  border-radius: 34px;
  padding: 0;
  overflow: hidden;
  background: #111;
  box-shadow: 0 24px 42px rgba(13, 22, 36, 0.18);
  opacity: 0;
  transform: translateX(42px);
  transition: opacity 880ms ease 130ms, transform 880ms cubic-bezier(.16, 1, .3, 1) 130ms;
}

.phone-demo::before,
.phone-demo::after {
  content: "";
  position: absolute;
  left: -11px;
  z-index: 4;
  width: 4px;
  border-radius: 4px;
  background: #30333a;
}

.phone-demo::before {
  top: 96px;
  height: 40px;
}

.phone-demo::after {
  top: 144px;
  height: 52px;
}

.phone-notch {
  position: absolute;
  left: 50%;
  top: 11px;
  z-index: 3;
  width: 58px;
  height: 18px;
  border-radius: 999px;
  background: #30333a;
  transform: translateX(-50%);
}

.phone-video {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
}

.service-section.visible .service-copy,
.service-section.visible .phone-demo {
  opacity: 1;
  transform: translate(0, 0);
}

.service-fade-enter-active,
.service-fade-leave-active {
  transition: opacity 240ms ease, transform 240ms ease;
}

.service-fade-enter-from,
.service-fade-leave-to {
  opacity: 0;
  transform: translateY(8px);
}

@media (max-width: 1180px) {
  .service-tabs,
  .service-inner {
    width: min(1040px, calc(100% - 48px));
  }

  .service-inner {
    grid-template-columns: minmax(0, 1fr) 270px;
    gap: 28px;
  }

  .ad-platforms,
  .oversea-platforms,
  .service-image {
    width: 100%;
  }

  .training-grid,
  .training-platform-bar {
    width: calc(100% - 34px);
  }
}

@media (max-width: 900px) {
  .service-body {
    min-height: auto;
  }

  .service-inner {
    height: auto;
    grid-template-columns: 1fr;
    padding: 44px 0 64px;
  }

  .service-copy {
    padding-top: 0;
  }

  .phone-demo {
    justify-self: center;
    margin-top: 10px;
  }
}

@media (max-width: 760px) {
  .service-section {
    padding: 23px 0 0;
    background: #fff;
  }

  :deep(.section-heading) {
    margin-bottom: 12px;
  }

  :deep(.section-heading h2) {
    font-size: 20px;
    line-height: 28px;
  }

  :deep(.section-heading p) {
    margin-top: 2px;
    font-size: 11px;
    line-height: 16px;
  }

  .service-tabs-wrap {
    margin-top: 0;
    background: #fff;
  }

  .service-tabs {
    width: 100%;
    height: 41px;
    padding: 0 10px;
    display: flex;
    overflow-x: auto;
    scrollbar-width: none;
    gap: 0;
  }

  .service-tabs::-webkit-scrollbar {
    display: none;
  }

  .service-tabs button {
    flex: 0 0 auto;
    min-width: 56px;
    min-height: 41px;
    padding: 0 8px;
    background: transparent;
    color: #111;
    font-size: 11px;
    font-weight: 700;
  }

  .service-tabs button.active {
    color: var(--color-brand);
    background: transparent;
  }

  .service-tabs button.active::after {
    height: 2px;
  }

  .service-body {
    min-height: 0;
    background: #f3f6fb;
  }

  .service-inner {
    width: 100%;
    padding: 24px 12px 22px;
    display: block;
  }

  .service-copy {
    opacity: 1;
    transform: none;
  }

  .panel-content {
    min-height: 0;
  }

  .panel-content h3 {
    margin: 0 0 14px;
    text-align: left;
    font-size: 14px;
    line-height: 20px;
  }

  .ad-platforms {
    width: 100%;
    grid-template-columns: repeat(5, minmax(0, 1fr));
    gap: 7px;
  }

  .platform-card {
    height: 53px;
    border-radius: 5px;
    gap: 5px;
    box-shadow: 0 4px 14px rgba(24, 37, 62, 0.08);
  }

  .platform-app {
    width: 26px;
    height: 26px;
    border-radius: 5px;
    font-size: 9px;
  }

  .platform-desc {
    font-size: 9px;
    line-height: 12px;
    white-space: nowrap;
  }

  .oversea-platforms {
    width: 100%;
    height: 34px;
    margin-top: 8px;
    border-radius: 5px;
    gap: 5px;
    box-shadow: 0 4px 14px rgba(24, 37, 62, 0.08);
  }

  .oversea-icon {
    width: 20px;
    height: 20px;
    border-radius: 4px;
    font-size: 11px;
  }

  .training-grid {
    width: 100%;
    margin-left: 0;
    grid-template-columns: 1fr;
    gap: 8px;
  }

  .training-card {
    height: 64px;
    padding: 0 14px;
    gap: 12px;
    border-radius: 7px;
  }

  .training-icon {
    width: 32px;
    height: 32px;
    border-radius: 8px;
  }

  .training-text strong {
    font-size: 13px;
    line-height: 18px;
  }

  .training-text small {
    font-size: 11px;
    line-height: 15px;
  }

  .training-grid,
  .training-platform-bar {
    width: 100%;
    margin-left: 0;
  }

  .training-platform-bar {
    height: auto;
    grid-template-columns: 1fr;
    gap: 8px;
    padding: 8px;
  }

  .service-image {
    width: 100%;
  }

  .service-more {
    min-width: 104px;
    min-height: 26px;
    margin: 14px auto 0;
    font-size: 11px;
  }

  .phone-demo {
    display: none;
  }
}

@media (prefers-reduced-motion: reduce) {
  .service-copy,
  .phone-demo {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
</style>
