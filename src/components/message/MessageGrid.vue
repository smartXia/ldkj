<script setup>
import { nextTick, onBeforeUnmount, onMounted, useTemplateRef, watch } from 'vue'

const props = defineProps({
  articles: {
    type: Array,
    required: true,
  },
})

const gridRef = useTemplateRef('grid')
let observer

function setupRevealObserver() {
  if (!gridRef.value) return

  observer?.disconnect()
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (!entry.isIntersecting) return
        entry.target.classList.add('is-visible')
        observer.unobserve(entry.target)
      })
    },
    {
      threshold: 0.16,
      rootMargin: '0px 0px -8% 0px',
    }
  )

  gridRef.value.querySelectorAll('.message-card').forEach((card) => {
    observer.observe(card)
  })
}

onMounted(setupRevealObserver)

watch(
  () => props.articles,
  async () => {
    await nextTick()
    setupRevealObserver()
  },
  { deep: false }
)

onBeforeUnmount(() => {
  observer?.disconnect()
})
</script>

<template>
  <div class="message-list">
    <div v-if="articles.length" ref="grid" class="message-grid">
      <a
        v-for="(article, index) in articles"
        :key="article.id"
        class="message-card"
        :href="article.href"
        target="_blank"
        rel="noreferrer"
        :style="{ transitionDelay: `${Math.min((index % 3) * 90, 180)}ms` }"
      >
        <span class="message-image">
          <img :src="article.image" :alt="article.title" loading="lazy" />
        </span>
        <span class="message-content">
          <strong>{{ article.title }}</strong>
          <em>{{ article.desc }}</em>
        </span>
        <span class="message-footer">
          <span>{{ article.author }}</span>
          <time :datetime="article.date">{{ article.date }}</time>
        </span>
      </a>
    </div>

    <div v-else class="message-empty">
      <strong>暂无匹配资讯</strong>
      <p>请切换分类或关键词查看更多内容</p>
    </div>
  </div>
</template>

<style scoped>
.message-list {
  width: 1200px;
  margin: 0 auto 60px;
}

.message-grid {
  display: grid;
  grid-template-columns: repeat(3, 376px);
  column-gap: 36px;
}

.message-card {
  width: 376px;
  height: 398px;
  margin-bottom: 40px;
  display: block;
  overflow: hidden;
  border-radius: 8px;
  background: #fff;
  color: #333;
  opacity: 0;
  transform: translateY(34px);
  transition:
    opacity 620ms cubic-bezier(0.22, 1, 0.36, 1),
    transform 620ms cubic-bezier(0.22, 1, 0.36, 1),
    box-shadow var(--motion-fast);
}

.message-card.is-visible {
  opacity: 1;
  transform: translateY(0);
}

.message-card:hover {
  box-shadow: 0 14px 34px rgba(26, 38, 66, 0.1);
}

.message-image {
  width: 376px;
  height: 188px;
  display: inline-flex;
  overflow: hidden;
  background: #f8f8f8;
}

.message-image img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  transition: transform var(--motion);
}

.message-card:hover .message-image img {
  transform: scale(1.03);
}

.message-content {
  height: 140px;
  margin-bottom: 30px;
  display: block;
  padding: 20px 20px 0;
}

.message-content strong {
  height: 56px;
  margin-bottom: 20px;
  display: -webkit-box;
  overflow: hidden;
  color: #333;
  font-size: 20px;
  line-height: 28px;
  font-weight: 600;
  font-style: normal;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.message-content em {
  height: 44px;
  display: -webkit-box;
  overflow: hidden;
  color: #777;
  font-size: 16px;
  line-height: 22px;
  font-style: normal;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.message-footer {
  height: 40px;
  display: flex;
  padding: 0 20px 20px;
  color: #777;
  font-size: 14px;
  line-height: 20px;
}

.message-footer span {
  width: 198px;
  margin-right: 8px;
}

.message-footer time {
  white-space: nowrap;
}

.message-empty {
  min-height: 260px;
  display: grid;
  place-content: center;
  text-align: center;
  color: #333;
}

.message-empty strong {
  font-size: 22px;
}

.message-empty p {
  margin: 8px 0 0;
  color: #777;
}

@media (max-width: 1240px) {
  .message-list {
    width: 788px;
  }

  .message-grid {
    grid-template-columns: repeat(2, 376px);
  }
}

@media (max-width: 900px) {
  .message-list {
    width: calc(100% - 32px);
    margin-bottom: 80px;
  }

  .message-grid {
    grid-template-columns: 1fr;
  }

  .message-card {
    width: 100%;
    height: auto;
  }

  .message-image {
    width: 100%;
    height: auto;
    aspect-ratio: 376 / 188;
  }
}

@media (max-width: 480px) {
  .message-content strong {
    font-size: 18px;
  }

  .message-footer {
    height: auto;
    flex-direction: column;
    gap: 2px;
  }

  .message-footer span {
    width: auto;
  }
}

@media (prefers-reduced-motion: reduce) {
  .message-card {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
</style>
