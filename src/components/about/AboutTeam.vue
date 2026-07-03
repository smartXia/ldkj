<script setup>
import { onBeforeUnmount, onMounted, shallowRef } from 'vue'

defineProps({
  title: {
    type: String,
    required: true,
  },
})

const teamImages = [
  '/assets/wsd/about/team-01.png',
  '/assets/wsd/about/team-02.png',
  '/assets/wsd/about/team-03.png',
  '/assets/wsd/about/team-04.jpg',
  '/assets/wsd/about/team-05.png',
  '/assets/wsd/about/team-06.png',
  '/assets/wsd/about/team-07.jpg',
  '/assets/wsd/about/team-08.png',
]

const activeIndex = shallowRef(2)
let timer

function getOffset(index) {
  const total = teamImages.length
  let offset = index - activeIndex.value
  if (offset > total / 2) offset -= total
  if (offset < -total / 2) offset += total
  return offset
}

function slideClass(index) {
  const offset = getOffset(index)
  return {
    active: offset === 0,
    prev: offset === -1,
    next: offset === 1,
    hidden: Math.abs(offset) > 1,
  }
}

function goTo(index) {
  activeIndex.value = (index + teamImages.length) % teamImages.length
}

function next() {
  goTo(activeIndex.value + 1)
}

function prev() {
  goTo(activeIndex.value - 1)
}

function start() {
  stop()
  timer = window.setInterval(next, 3600)
}

function stop() {
  if (timer) window.clearInterval(timer)
}

onMounted(start)
onBeforeUnmount(stop)
</script>

<template>
  <section id="about-team" class="about-team">
    <h2>{{ title }}</h2>

    <div class="team-carousel" @mouseenter="stop" @mouseleave="start">
      <button class="team-arrow team-arrow--prev" type="button" aria-label="上一张" @click="prev">
        ‹
      </button>
      <button class="team-arrow team-arrow--next" type="button" aria-label="下一张" @click="next">
        ›
      </button>

      <div class="team-stage">
        <img
          v-for="(image, index) in teamImages"
          :key="image"
          :src="image"
          :alt="`${title} ${index + 1}`"
          :class="slideClass(index)"
          loading="lazy"
          @click="goTo(index)"
        />
      </div>

      <div class="team-dots" :aria-label="`${title}轮播分页`">
        <button
          v-for="(image, index) in teamImages"
          :key="image"
          type="button"
          :aria-label="`查看第 ${index + 1} 张`"
          :class="{ active: index === activeIndex }"
          @click="goTo(index)"
        ></button>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about-team {
  padding: 18px 0 96px;
  text-align: center;
  background: #fff;
}

.about-team h2 {
  margin: 0 0 46px;
  color: #000;
  font-size: 40px;
  line-height: 56px;
  font-weight: 700;
}

.team-carousel {
  position: relative;
  width: 1400px;
  height: 528px;
  margin: 0 auto;
}

.team-stage {
  position: relative;
  height: 450px;
  overflow: hidden;
}

.team-stage img {
  position: absolute;
  top: 38px;
  left: 50%;
  width: 900px;
  height: 450px;
  object-fit: cover;
  border-radius: 6px;
  opacity: 0;
  transform: translateX(-50%) scale(0.82);
  transition: transform var(--motion), opacity var(--motion), filter var(--motion);
  cursor: pointer;
  filter: brightness(0.92);
}

.team-stage img.active {
  z-index: 3;
  opacity: 1;
  transform: translateX(-50%) scale(1);
  filter: none;
}

.team-stage img.prev {
  z-index: 2;
  opacity: 1;
  transform: translateX(-106%) scale(0.86);
}

.team-stage img.next {
  z-index: 2;
  opacity: 1;
  transform: translateX(6%) scale(0.86);
}

.team-stage img.hidden {
  pointer-events: none;
}

.team-arrow {
  position: absolute;
  top: 258px;
  z-index: 5;
  width: 42px;
  height: 42px;
  border: 0;
  border-radius: 50%;
  display: grid;
  place-items: center;
  color: rgba(255, 255, 255, 0.9);
  background: rgba(0, 0, 0, 0.22);
  font-size: 48px;
  line-height: 40px;
  font-family: Arial, sans-serif;
  cursor: pointer;
  transition: background var(--motion-fast), transform var(--motion-fast);
}

.team-arrow:hover {
  background: rgba(255, 72, 72, 0.9);
  transform: translateY(-1px);
}

.team-arrow--prev {
  left: 312px;
}

.team-arrow--next {
  right: 312px;
}

.team-dots {
  margin-top: 28px;
  display: flex;
  justify-content: center;
  gap: 8px;
}

.team-dots button {
  width: 8px;
  height: 8px;
  padding: 0;
  border: 0;
  border-radius: 50%;
  background: #dedede;
  cursor: pointer;
  transition: background var(--motion-fast), transform var(--motion-fast);
}

.team-dots button.active {
  background: #ff4848;
  transform: scale(1.05);
}

@media (max-width: 1440px) {
  .team-carousel {
    width: 1200px;
  }

  .team-stage img {
    width: 900px;
  }

  .team-arrow--prev {
    left: 232px;
  }

  .team-arrow--next {
    right: 232px;
  }
}

@media (max-width: 1240px) {
  .team-carousel {
    width: calc(100% - 40px);
    height: 460px;
  }

  .team-stage {
    height: 388px;
  }

  .team-stage img {
    width: 78vw;
    height: 388px;
  }

  .team-arrow--prev {
    left: 12%;
  }

  .team-arrow--next {
    right: 12%;
  }
}

@media (max-width: 720px) {
  .about-team {
    padding: 33px 0 5px;
  }

  .about-team h2 {
    margin-bottom: 20px;
    font-size: 24px;
    line-height: 29px;
  }

  .team-carousel {
    width: 100%;
    height: 206px;
  }

  .team-stage {
    height: 155px;
  }

  .team-stage img {
    top: 0;
    width: calc(100% - 32px);
    height: 155px;
  }

  .team-stage img.prev,
  .team-stage img.next {
    opacity: 0;
  }

  .team-arrow {
    top: 56px;
    width: 34px;
    height: 34px;
    font-size: 34px;
  }

  .team-arrow--prev {
    left: 28px;
  }

  .team-arrow--next {
    right: 28px;
  }
}
</style>
