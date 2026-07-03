<script setup>
import { onBeforeUnmount, onMounted, shallowRef } from 'vue'

const props = defineProps({
  items: {
    type: Array,
    required: true,
  },
})

const activeItem = shallowRef(getItemId(props.items[0]))
let observer
let pinRaf = 0

function getItemId(item) {
  return typeof item === 'string' ? item : item.id
}

function getItemLabel(item) {
  return typeof item === 'string' ? item : item.label
}

function setPinnedState() {
  if (pinRaf) return

  pinRaf = window.requestAnimationFrame(() => {
    pinRaf = 0
    const shell = document.querySelector('.about-anchor-shell')
    if (!shell) return

    const headerHeight = Number.parseFloat(getComputedStyle(document.documentElement).getPropertyValue('--header-height')) || 64
    const pinned = shell.getBoundingClientRect().top <= headerHeight + 1
    document.body.classList.toggle('about-subnav-pinned', pinned)
  })
}

function handleClick(item, event) {
  event.preventDefault()
  const id = getItemId(item)
  activeItem.value = id

  const target = document.getElementById(id)
  if (!target) return

  const headerHeight = Number.parseFloat(getComputedStyle(document.documentElement).getPropertyValue('--header-height')) || 64
  const navHeight = 66
  const top = target.getBoundingClientRect().top + window.scrollY - headerHeight - navHeight

  window.scrollTo({
    top,
    behavior: 'smooth',
  })
}

onMounted(() => {
  const sections = props.items
    .map((item) => document.getElementById(getItemId(item)))
    .filter(Boolean)

  observer = new IntersectionObserver(
    (entries) => {
      const visible = entries
        .filter((entry) => entry.isIntersecting)
        .sort((a, b) => b.intersectionRatio - a.intersectionRatio)[0]

      if (visible?.target?.id) activeItem.value = visible.target.id
    },
    {
      rootMargin: '-34% 0px -54% 0px',
      threshold: [0.08, 0.18, 0.32, 0.5],
    }
  )

  sections.forEach((section) => observer.observe(section))
  setPinnedState()
  window.addEventListener('scroll', setPinnedState, { passive: true })
  window.addEventListener('resize', setPinnedState)
})

onBeforeUnmount(() => {
  observer?.disconnect()
  window.removeEventListener('scroll', setPinnedState)
  window.removeEventListener('resize', setPinnedState)
  if (pinRaf) window.cancelAnimationFrame(pinRaf)
  document.body.classList.remove('about-subnav-pinned')
})
</script>

<template>
  <div class="about-anchor-shell">
    <nav class="about-anchor-nav" aria-label="关于我们页面导航">
      <a
        v-for="item in items"
        :key="getItemId(item)"
        :class="{ active: getItemId(item) === activeItem }"
        :href="`#${getItemId(item)}`"
        @click="handleClick(item, $event)"
      >
        <span>{{ getItemLabel(item) }}</span>
      </a>
    </nav>
  </div>
</template>

<style scoped>
.about-anchor-shell {
  position: sticky;
  top: var(--header-height);
  z-index: 45;
  height: 66px;
  background: #fff;
  border-bottom: 1px solid #f0f1f5;
}

.about-anchor-nav {
  width: 1040px;
  height: 64px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.about-anchor-nav a {
  position: relative;
  min-width: 80px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #333;
  font-size: 20px;
  line-height: 64px;
  transition: color var(--motion-fast);
}

.about-anchor-nav a + a::before {
  content: "";
  position: absolute;
  left: -84px;
  top: 50%;
  width: 74px;
  height: 1px;
  background: #edf0f5;
}

.about-anchor-nav a.active,
.about-anchor-nav a:hover {
  color: #ff4848;
}

@media (max-width: 1100px) {
  .about-anchor-shell {
    height: 56px;
  }

  .about-anchor-nav {
    width: 100%;
    height: 56px;
    padding: 0 18px;
    overflow-x: auto;
    overflow-y: hidden;
    gap: 30px;
    justify-content: flex-start;
    scrollbar-width: none;
    -webkit-overflow-scrolling: touch;
  }

  .about-anchor-nav::-webkit-scrollbar {
    display: none;
  }

  .about-anchor-nav a {
    flex: 0 0 auto;
    min-width: auto;
    font-size: 15px;
    line-height: 56px;
    white-space: nowrap;
  }

  .about-anchor-nav a + a::before {
    left: -22px;
    width: 14px;
  }
}

@media (max-width: 760px) {
  .about-anchor-shell {
    display: none;
  }
}
</style>
