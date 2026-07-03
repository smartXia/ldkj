<script setup>
import { computed, onMounted, onUnmounted, shallowRef } from 'vue'

const title = '\u5fae\u601d\u6566'
const stats = [
  { value: 12, suffix: '', label: '\u5e74', icon: 'years' },
  { value: 3000, suffix: '+', label: '\u670d\u52a1\n\u5ba2\u6237', icon: 'clients' },
  { value: 120, suffix: '+', label: '\u884c\u4e1a\n\u8363\u8a89', icon: 'honors' }
]

const sectionRef = shallowRef(null)
const visible = shallowRef(false)
const counts = shallowRef(stats.map(() => 0))
let observer
let rafId

const displayStats = computed(() =>
  stats.map((item, index) => ({
    ...item,
    displayValue: counts.value[index]
  }))
)

function easeOutCubic(value) {
  return 1 - Math.pow(1 - value, 3)
}

function animateCounts() {
  const start = performance.now()
  const duration = 1100

  function step(now) {
    const progress = Math.min((now - start) / duration, 1)
    const eased = easeOutCubic(progress)
    counts.value = stats.map((item) => Math.round(item.value * eased))
    if (progress < 1) {
      rafId = requestAnimationFrame(step)
    } else {
      counts.value = stats.map((item) => item.value)
    }
  }

  cancelAnimationFrame(rafId)
  rafId = requestAnimationFrame(step)
}

onMounted(() => {
  observer = new IntersectionObserver(([entry]) => {
    if (!entry.isIntersecting || visible.value) return
    visible.value = true
    animateCounts()
  }, { threshold: 0.35 })

  if (sectionRef.value) observer.observe(sectionRef.value)
})

onUnmounted(() => {
  observer?.disconnect()
  cancelAnimationFrame(rafId)
})
</script>

<template>
  <section ref="sectionRef" class="intro-section" :class="{ visible }">
    <div class="intro-shell">
      <h2>{{ title }}</h2>

      <ul class="stats" :aria-label="title">
        <li v-for="(item, index) in displayStats" :key="item.icon" :style="{ '--delay': `${index * 140}ms` }">
          <span class="stat-icon" :class="item.icon" aria-hidden="true">
            <i></i>
          </span>
          <strong>
            {{ item.displayValue }}<em v-if="item.suffix">{{ item.suffix }}</em>
          </strong>
          <span class="stat-label">{{ item.label }}</span>
        </li>
      </ul>
    </div>
  </section>
</template>

<style scoped>
.intro-section {
  position: relative;
  min-height: 536px;
  color: #fff;
  background: #383e44 url('/assets/wsd/intro-bg.png') center / cover no-repeat;
  overflow: hidden;
}

.intro-shell {
  position: relative;
  width: min(960px, calc(100% - 48px));
  min-height: 536px;
  margin: 0 auto;
  display: grid;
  grid-template-rows: 184px 1fr;
  align-items: start;
}

.intro-shell h2 {
  margin: 84px 0 0;
  text-align: center;
  font-size: 42px;
  line-height: 1.2;
  font-weight: 900;
  letter-spacing: 0;
  text-shadow: 0 3px 8px rgba(0, 0, 0, 0.24);
  opacity: 0;
  transform: translateY(26px);
  transition: opacity 760ms ease, transform 760ms cubic-bezier(.16, 1, .3, 1);
}

.stats {
  position: relative;
  display: grid;
  grid-template-columns: 286px 430px 244px;
  align-items: center;
  justify-content: center;
  gap: 0;
  margin: 96px 0 0;
  padding: 0;
  list-style: none;
}

.stats::before,
.stats::after {
  content: "";
  width: 2px;
  height: 64px;
  background: rgba(255, 255, 255, 0.72);
  position: absolute;
  top: 10px;
}

.stats::before {
  left: 286px;
}

.stats::after {
  left: 716px;
}

.stats li {
  min-width: 0;
  display: flex;
  align-items: end;
  gap: 0;
  opacity: 0;
  transform: translateX(-34px);
  transition:
    opacity 760ms ease var(--delay),
    transform 760ms cubic-bezier(.16, 1, .3, 1) var(--delay);
}

.stat-icon {
  position: relative;
  flex: 0 0 auto;
  width: 58px;
  height: 58px;
  margin-right: 24px;
  align-self: center;
  filter: drop-shadow(0 8px 10px rgba(0, 0, 0, 0.16));
  transform: translateY(-2px) scale(.86) rotate(-4deg);
  transition: transform 780ms cubic-bezier(.16, 1, .3, 1) var(--delay);
}

.stat-icon::before,
.stat-icon::after,
.stat-icon i {
  content: "";
  position: absolute;
  border-radius: 12px;
}

.stat-icon::before {
  inset: 7px 8px 8px 8px;
  background: linear-gradient(145deg, #ff7779 0%, #f43b45 58%, #c82733 100%);
  transform: skewY(-14deg) rotate(10deg);
}

.stat-icon::after {
  right: 5px;
  top: 13px;
  width: 22px;
  height: 38px;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.32), rgba(255, 255, 255, 0));
  transform: skewY(24deg);
}

.stat-icon i {
  left: 16px;
  top: 24px;
  z-index: 2;
  width: 28px;
  height: 20px;
  border: 5px solid rgba(255, 255, 255, 0.86);
  border-top: 0;
  border-radius: 0 0 12px 12px;
}

.stat-icon.clients i {
  left: 13px;
  top: 21px;
  width: 33px;
  height: 25px;
  border: 0;
  border-radius: 50%;
  background:
    radial-gradient(circle at 28% 38%, #fff 0 5px, transparent 6px),
    radial-gradient(circle at 68% 38%, #fff 0 5px, transparent 6px),
    radial-gradient(ellipse at 50% 82%, rgba(255,255,255,.92) 0 14px, transparent 15px);
}

.stat-icon.honors i {
  left: 16px;
  top: 17px;
  width: 27px;
  height: 31px;
  border: 4px solid rgba(255,255,255,.9);
  border-radius: 8px;
}

.stats strong {
  display: inline-flex;
  align-items: baseline;
  color: #fff;
  font-size: 76px;
  line-height: .88;
  font-weight: 900;
  letter-spacing: 0;
  text-shadow: 0 8px 18px rgba(0, 0, 0, 0.18);
}

.stats strong em {
  margin-left: 8px;
  font-style: normal;
  font-size: 58px;
  line-height: 1;
}

.stat-label {
  flex: 0 0 auto;
  align-self: end;
  margin: 0 0 5px 18px;
  white-space: pre-line;
  color: #fff;
  font-size: 16px;
  line-height: 24px;
  font-weight: 500;
}

.intro-section.visible .intro-shell h2,
.intro-section.visible .stats li {
  opacity: 1;
  transform: translateX(0) translateY(0);
}

.intro-section.visible .stat-icon {
  transform: translateY(-2px) scale(1) rotate(-4deg);
}

@media (max-width: 980px) {
  .intro-shell {
    width: min(100% - 32px, 720px);
    min-height: 640px;
    grid-template-rows: 150px 1fr;
  }

  .stats {
    grid-template-columns: 1fr;
    gap: 34px;
    margin-top: 60px;
  }

  .stats::before,
  .stats::after {
    display: none;
  }

  .stats li {
    justify-self: center;
    width: 320px;
  }
}

@media (max-width: 520px) {
  .intro-section {
    min-height: 154px;
  }

  .intro-shell {
    width: 100%;
    min-height: 154px;
    grid-template-rows: 58px 1fr;
  }

  .intro-shell h2 {
    margin-top: 27px;
    font-size: 17px;
    line-height: 22px;
  }

  .stats {
    width: 100%;
    margin: 26px auto 0;
    padding: 0 12px;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 0;
  }

  .stats li,
  .stats li:nth-child(2),
  .stats li:nth-child(3) {
    width: auto;
    justify-self: center;
    align-items: center;
  }

  .stat-icon {
    width: 21px;
    height: 21px;
    margin-right: 5px;
  }

  .stat-icon::before,
  .stat-icon::after,
  .stat-icon i {
    border-radius: 5px;
  }

  .stat-icon i {
    display: none;
  }

  .stats strong {
    font-size: 28px;
    line-height: 1;
  }

  .stats strong em {
    margin-left: 1px;
    font-size: 16px;
  }

  .stat-label {
    margin: 0 0 1px 3px;
    font-size: 9px;
    line-height: 12px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .intro-shell h2,
  .stats li,
  .stat-icon {
    opacity: 1;
    transform: none;
    transition: none;
  }
}
</style>
