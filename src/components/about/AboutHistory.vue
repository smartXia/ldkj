<script setup>
import { onBeforeUnmount, onMounted, useTemplateRef } from 'vue'

defineProps({
  content: {
    type: Object,
    required: true,
  },
})

const historyRef = useTemplateRef('history')
const videoRef = useTemplateRef('video')
let observer

function playVideo() {
  const video = videoRef.value
  if (!video) return

  video.currentTime = 0
  video.play().catch(() => {})
}

onMounted(() => {
  if (!historyRef.value) return

  observer = new IntersectionObserver(
    (entries) => {
      const entry = entries[0]
      if (!entry?.isIntersecting) return
      historyRef.value.classList.add('is-visible')
      playVideo()
    },
    { threshold: 0.45 }
  )

  observer.observe(historyRef.value)
})

onBeforeUnmount(() => {
  observer?.disconnect()
})
</script>

<template>
  <section id="about-history" ref="history" class="about-history">
    <h2>{{ content.title }}</h2>
    <video
      ref="video"
      class="history-video"
      :src="content.video"
      :poster="content.poster"
      muted
      playsinline
      preload="metadata"
      :aria-label="content.title"
    ></video>
    <img class="history-fallback" :src="content.image" :alt="content.title" />
  </section>
</template>

<style scoped>
.about-history {
  position: relative;
  height: auto;
  min-height: 573px;
  aspect-ratio: 1920 / 860;
  overflow: hidden;
  text-align: center;
  background: #111;
}

.about-history h2 {
  position: absolute;
  width: 1px;
  height: 1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
}

.history-video {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  opacity: 0;
  transform: scale(1.015);
  transition: opacity 520ms ease, transform 900ms ease;
}

.about-history.is-visible .history-video {
  opacity: 1;
  transform: scale(1);
}

.history-fallback {
  display: none;
}

@media (max-width: 768px) {
  .about-history {
    width: 100%;
    min-height: 0;
    aspect-ratio: auto;
    height: auto;
    padding: 33px 0 30px;
    overflow: hidden;
    background: #fff;
  }

  .about-history h2 {
    position: static;
    width: auto;
    height: auto;
    clip: auto;
    margin: 0 0 28px;
    color: #000;
    font-size: 24px;
    line-height: 29px;
    font-weight: 600;
    white-space: normal;
  }

  .history-video {
    display: none;
  }

  .history-fallback {
    width: min(100% - 66px, 316px);
    aspect-ratio: 1149 / 1117;
    display: block;
    margin: 0 auto;
    object-fit: contain;
  }
}

@media (max-width: 480px) {
  .about-history h2 {
    font-size: 24px;
    line-height: 29px;
  }
}
</style>
