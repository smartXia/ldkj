<script setup>
import { nextTick, onBeforeUnmount, onMounted, useTemplateRef, watch } from 'vue'

const props = defineProps({
  cases: {
    type: Array,
    required: true,
  },
})

const emit = defineEmits(['select'])
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
      threshold: 0.18,
      rootMargin: '0px 0px -8% 0px',
    }
  )

  gridRef.value.querySelectorAll('.case-card').forEach((card) => {
    observer.observe(card)
  })
}

onMounted(setupRevealObserver)

watch(
  () => props.cases,
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
  <div class="content">
    <div v-if="cases.length" ref="grid" class="case-grid">
      <article
        v-for="(item, index) in cases"
        :key="item.id"
        class="case-card"
        :style="{ transitionDelay: `${Math.min((index % 3) * 90, 180)}ms` }"
      >
        <button class="case-link" type="button" :aria-label="`查看案例详情：${item.title}`" @click="emit('select', item)">
          <img v-if="item.image" :src="item.image" :alt="item.title" loading="lazy" />
          <div class="case-info">
            <h2>{{ item.title }}</h2>
            <div class="case-meta">
              <span>{{ item.industry }}</span>
              <span>{{ item.method }}</span>
              <span>{{ item.platform }}</span>
            </div>
          </div>
        </button>
      </article>
    </div>

    <div v-else class="empty-state">
      <strong>暂无匹配案例</strong>
      <p>请切换筛选条件查看更多客户案例</p>
    </div>
  </div>
</template>

<style scoped>
.content {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 0 auto 140px;
}

.case-grid {
  display: grid;
  grid-template-columns: repeat(3, 376px);
  gap: 60px 36px;
}

.case-card,
.case-link {
  display: block;
}

.case-card {
  opacity: 0;
  transform: translateY(34px);
  transition:
    opacity 620ms cubic-bezier(0.22, 1, 0.36, 1),
    transform 620ms cubic-bezier(0.22, 1, 0.36, 1);
}

.case-card.is-visible {
  opacity: 1;
  transform: translateY(0);
}

.case-link {
  width: 100%;
  border: 0;
  outline: 0;
  padding: 0;
  background: transparent;
  color: inherit;
  text-align: left;
  cursor: pointer;
  appearance: none;
  -webkit-appearance: none;
  box-shadow: none;
  font: inherit;
}

.case-link:focus {
  outline: 0;
}

.case-link:focus-visible {
  outline: 2px solid var(--color-brand);
  outline-offset: 4px;
}

.case-link img {
  width: 376px;
  height: 250px;
  display: block;
  object-fit: cover;
  background: #f2f2f2;
  transition: transform var(--motion);
}

.case-link:hover img {
  transform: translateY(-2px);
}

.case-info {
  padding-top: 18px;
}

.case-info h2 {
  min-height: 52px;
  margin: 0;
  color: #111;
  font-size: 18px;
  line-height: 1.45;
  font-weight: 700;
}

.case-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0;
  margin-top: 14px;
  color: #777;
  font-size: 14px;
  line-height: 1.5;
}

.case-meta span + span::before {
  content: "";
  display: inline-block;
  width: 1px;
  height: 12px;
  margin: 0 12px;
  vertical-align: -1px;
  background: #d8d8d8;
}

.empty-state {
  min-height: 260px;
  display: grid;
  place-content: center;
  text-align: center;
  color: #111;
}

.empty-state strong {
  font-size: 22px;
}

.empty-state p {
  margin: 8px 0 0;
  color: #777;
}

@media (max-width: 1240px) {
  .content {
    width: 788px;
  }

  .case-grid {
    grid-template-columns: repeat(2, 376px);
  }
}

@media (max-width: 900px) {
  .content {
    max-width: calc(100% - 32px);
    margin-bottom: 88px;
  }

  .case-grid {
    grid-template-columns: 1fr;
    gap: 42px;
  }

  .case-link img {
    width: 100%;
    height: auto;
    aspect-ratio: 376 / 250;
  }

  .case-info h2 {
    min-height: 0;
  }
}

@media (prefers-reduced-motion: reduce) {
  .case-card {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
</style>
