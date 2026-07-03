<script setup>
import { computed, shallowRef } from 'vue'

defineProps({
  intro: {
    type: Object,
    required: true,
  },
  meanings: {
    type: Array,
    required: true,
  },
})

const activeMeaningIndex = shallowRef(null)
const visualStateClass = computed(() =>
  activeMeaningIndex.value === null ? 'intro-visual--idle' : `intro-visual--active-${activeMeaningIndex.value}`
)

function setActiveMeaning(index) {
  activeMeaningIndex.value = index
}

function clearActiveMeaning() {
  activeMeaningIndex.value = null
}

function handleVisualFocusout(event) {
  if (!event.currentTarget.contains(event.relatedTarget)) {
    clearActiveMeaning()
  }
}
</script>

<template>
  <section id="about-intro" class="about-intro">
    <div class="about-container">
      <div class="intro-copy">
        <h2>{{ intro.title }}</h2>
        <p v-for="paragraph in intro.paragraphs" :key="paragraph">{{ paragraph }}</p>
      </div>

      <div
        class="intro-visual"
        :class="visualStateClass"
        aria-label="WSD集团理念"
        @mouseleave="clearActiveMeaning"
        @focusout="handleVisualFocusout"
      >
        <img class="intro-bg" :src="intro.image" alt="" />
        <article
          v-for="(item, index) in meanings"
          :key="item.cn"
          class="meaning-card"
          :class="[
            `meaning-card--${item.title.charAt(0).toLowerCase()}`,
            { active: index === activeMeaningIndex },
          ]"
          tabindex="0"
          @mouseenter="setActiveMeaning(index)"
          @focus="setActiveMeaning(index)"
        >
          <div class="meaning-heading">
            <span>{{ item.title.replace(' ·', '') }}</span>
            <strong>{{ item.cn }}</strong>
          </div>
          <img :src="item.icon" :alt="item.cn" />
          <p>{{ item.desc }}</p>
        </article>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about-intro {
  position: relative;
  min-height: 514px;
  overflow: hidden;
  background: #fff;
}

.about-container {
  position: relative;
  width: 1414px;
  min-height: 514px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 650px 1fr;
}

.intro-copy {
  padding-top: 48px;
}

.intro-copy h2 {
  margin: 0 0 28px;
  color: #333;
  font-size: 36px;
  line-height: 50px;
  font-weight: 700;
}

.intro-copy p {
  margin: 0 0 12px;
  color: #4f5560;
  font-size: 16px;
  line-height: 30px;
}

.intro-visual {
  position: relative;
  min-height: 514px;
}

.intro-bg {
  position: absolute;
  right: 8px;
  top: -62px;
  width: 642px;
  height: 642px;
  object-fit: contain;
  opacity: 0.86;
  pointer-events: none;
}

.meaning-card {
  position: absolute;
  z-index: 1;
  width: 238px;
  height: 88px;
  padding: 0 72px 0 20px;
  border: 0;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 10px 28px rgba(31, 38, 56, 0.08);
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  outline: none;
  overflow: hidden;
  transition:
    left 420ms cubic-bezier(0.22, 1, 0.36, 1),
    top 420ms cubic-bezier(0.22, 1, 0.36, 1),
    width 420ms cubic-bezier(0.22, 1, 0.36, 1),
    height 420ms cubic-bezier(0.22, 1, 0.36, 1),
    padding 420ms cubic-bezier(0.22, 1, 0.36, 1),
    background 260ms ease,
    box-shadow 260ms ease;
}

.meaning-card:focus-visible {
  box-shadow: 0 12px 34px rgba(255, 92, 82, 0.18);
}

.meaning-card.active {
  z-index: 5;
  width: 328px;
  height: 176px;
  padding: 30px 22px 24px;
  align-items: flex-start;
  flex-direction: column;
  gap: 0;
  background: linear-gradient(135deg, #fff 0%, #fff2f2 100%);
  box-shadow: 0 18px 42px rgba(31, 38, 56, 0.09);
}

.meaning-heading {
  display: flex;
  align-items: center;
  gap: 10px;
  min-width: 0;
  max-width: 100%;
}

.meaning-card span {
  color: #ff745f;
  font-size: 20px;
  line-height: 28px;
  font-weight: 700;
  white-space: nowrap;
}

.meaning-card span::after {
  content: "•";
  margin-left: 9px;
  color: #ff3d45;
}

.meaning-card strong {
  color: #1f2530;
  font-size: 13px;
  line-height: 20px;
  white-space: nowrap;
}

.meaning-card img {
  position: absolute;
  right: 20px;
  top: 25px;
  width: 38px;
  height: 38px;
  margin-left: 0;
  flex: 0 0 auto;
  transition:
    width 420ms cubic-bezier(0.22, 1, 0.36, 1),
    height 420ms cubic-bezier(0.22, 1, 0.36, 1),
    transform 420ms cubic-bezier(0.22, 1, 0.36, 1);
}

.meaning-card.active img {
  right: 24px;
  top: 26px;
  width: 44px;
  height: 44px;
  margin-left: 0;
}

.meaning-card p {
  margin: 24px 0 0;
  color: #3e4652;
  font-size: 13px;
  line-height: 26px;
  opacity: 0;
  transform: translateY(8px);
  transition: opacity 180ms ease, transform 260ms ease;
}

.meaning-card.active p {
  opacity: 1;
  transform: translateY(0);
  transition-delay: 130ms;
}

.meaning-card--w {
  left: 32px;
  top: 54px;
}

.meaning-card--s {
  left: 370px;
  top: 168px;
}

.meaning-card--d {
  left: 32px;
  top: 300px;
}

.intro-visual--active-0 .meaning-card--w {
  left: 32px;
  top: 54px;
}

.intro-visual--active-0 .meaning-card--s {
  left: 370px;
  top: 168px;
}

.intro-visual--active-0 .meaning-card--d {
  left: 32px;
  top: 300px;
}

.intro-visual--active-1 .meaning-card--w {
  left: 32px;
  top: 54px;
}

.intro-visual--active-1 .meaning-card--s {
  left: 286px;
  top: 152px;
  height: 218px;
}

.intro-visual--active-1 .meaning-card--d {
  left: 22px;
  top: 286px;
}

.intro-visual--active-2 .meaning-card--w {
  left: 32px;
  top: 54px;
}

.intro-visual--active-2 .meaning-card--s {
  left: 324px;
  top: 196px;
}

.intro-visual--active-2 .meaning-card--d {
  left: 32px;
  top: 196px;
  height: 218px;
}

@media (max-width: 1480px) {
  .about-container {
    width: 1200px;
    grid-template-columns: 610px 1fr;
  }

  .meaning-card--s,
  .intro-visual--active-0 .meaning-card--s {
    left: 318px;
  }

  .intro-visual--active-1 .meaning-card--s {
    left: 214px;
  }

  .intro-visual--active-2 .meaning-card--s {
    left: 278px;
  }
}

@media (max-width: 1240px) {
  .about-container {
    width: calc(100% - 40px);
    grid-template-columns: 1fr;
    padding-bottom: 36px;
  }

  .intro-visual {
    min-height: 420px;
  }

  .intro-bg {
    left: 50%;
    right: auto;
    top: -40px;
    width: 560px;
    height: 560px;
    transform: translateX(-50%);
  }

  .meaning-card--w,
  .intro-visual--active-0 .meaning-card--w,
  .intro-visual--active-1 .meaning-card--w,
  .intro-visual--active-2 .meaning-card--w {
    left: calc(50% - 290px);
  }

  .meaning-card--d,
  .intro-visual--active-0 .meaning-card--d,
  .intro-visual--active-1 .meaning-card--d,
  .intro-visual--active-2 .meaning-card--d {
    left: calc(50% - 290px);
  }

  .meaning-card--s,
  .intro-visual--active-0 .meaning-card--s,
  .intro-visual--active-2 .meaning-card--s {
    left: calc(50% + 60px);
  }

  .intro-visual--active-1 .meaning-card--s {
    left: calc(50% - 70px);
  }
}

@media (max-width: 640px) {
  .about-intro {
    min-height: 1104px;
  }

  .intro-copy {
    padding-top: 34px;
  }

  .intro-copy h2 {
    font-size: 30px;
    line-height: 40px;
  }

  .intro-copy p {
    font-size: 14px;
    line-height: 26px;
  }

  .intro-visual {
    min-height: 430px;
  }

  .intro-bg {
    width: 420px;
    height: 420px;
  }

  .meaning-card,
  .meaning-card.active {
    left: 0 !important;
    width: min(100%, 328px);
  }

  .meaning-card.active {
    height: auto;
    min-height: 176px;
  }

  .meaning-card--w,
  .intro-visual--active-0 .meaning-card--w,
  .intro-visual--active-1 .meaning-card--w,
  .intro-visual--active-2 .meaning-card--w {
    top: 40px;
  }

  .meaning-card--s,
  .intro-visual--active-0 .meaning-card--s,
  .intro-visual--active-1 .meaning-card--s,
  .intro-visual--active-2 .meaning-card--s {
    top: 168px;
  }

  .meaning-card--d,
  .intro-visual--active-0 .meaning-card--d,
  .intro-visual--active-1 .meaning-card--d,
  .intro-visual--active-2 .meaning-card--d {
    top: 296px;
  }

  .intro-visual--active-0 .meaning-card--w,
  .intro-visual--active-1 .meaning-card--s,
  .intro-visual--active-2 .meaning-card--d {
    top: 40px;
  }
}
</style>
