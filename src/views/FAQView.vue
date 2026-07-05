<script setup>
import { computed, onMounted, shallowRef } from 'vue'
import { getPartnerContent } from '../services/publicApi'

const faqs = shallowRef([])
const openedFaqs = shallowRef(new Set([0]))
const loading = shallowRef(true)

const categories = [
  { label: '合作流程', keywords: ['联系', '周期', '项目', '启动', '合作'] },
  { label: '服务内容', keywords: ['服务', '内容', '方案', '策略', '账号'] },
  { label: '交付与支持', keywords: ['交付', '报告', '复盘', '团队', '培训'] },
]

const groupedFaqs = computed(() => {
  const groups = categories.map((category) => ({ ...category, items: [] }))
  const fallback = { label: '其他问题', items: [] }

  faqs.value.forEach((item, index) => {
    const text = `${item.question || ''}${item.answer || ''}`
    const group = groups.find((category) => category.keywords.some((keyword) => text.includes(keyword))) || fallback
    group.items.push({ ...item, index })
  })

  return [...groups, fallback].filter((group) => group.items.length)
})

function toggleFaq(index) {
  const next = new Set(openedFaqs.value)
  if (next.has(index)) next.delete(index)
  else next.add(index)
  openedFaqs.value = next
}

onMounted(async () => {
  const data = await getPartnerContent()
  faqs.value = Array.isArray(data.faqs) ? data.faqs : []
  loading.value = false
})
</script>

<template>
  <main class="faq-page">
    <section class="faq-hero">
      <p>FAQ</p>
      <h1>常见问题</h1>
      <span>围绕合作流程、服务交付与项目支持，快速了解南京灵动如何与品牌团队协作。</span>
    </section>

    <section class="faq-main" aria-label="常见问题列表">
      <aside class="faq-summary">
        <strong>问题索引</strong>
        <a v-for="group in groupedFaqs" :key="group.label" :href="`#${group.label}`">{{ group.label }}</a>
        <RouterLink class="consult-link" to="/cooperate">提交合作咨询</RouterLink>
      </aside>

      <div class="faq-content">
        <div v-if="loading" class="faq-empty">加载中...</div>
        <div v-else-if="!faqs.length" class="faq-empty">暂无常见问题</div>

        <section v-for="group in groupedFaqs" v-else :id="group.label" :key="group.label" class="faq-group">
          <h2>{{ group.label }}</h2>
          <article v-for="item in group.items" :key="item.question" class="faq-item">
            <button type="button" :aria-expanded="openedFaqs.has(item.index)" @click="toggleFaq(item.index)">
              <span>{{ item.question }}</span>
              <i aria-hidden="true">{{ openedFaqs.has(item.index) ? '-' : '+' }}</i>
            </button>
            <p v-if="openedFaqs.has(item.index)">{{ item.answer }}</p>
          </article>
        </section>
      </div>
    </section>
  </main>
</template>

<style scoped>
.faq-page {
  min-height: 100vh;
  background: #fff;
  color: #11151d;
}

.faq-hero {
  padding: 146px max(24px, calc((100vw - 1180px) / 2)) 76px;
  background: linear-gradient(180deg, #fff6f6 0%, #fff 100%);
}

.faq-hero p {
  margin: 0 0 14px;
  color: var(--color-brand);
  font-size: 15px;
  font-weight: 800;
  letter-spacing: .08em;
}

.faq-hero h1 {
  margin: 0;
  font-size: clamp(38px, 5vw, 64px);
  line-height: 1.08;
  font-weight: 900;
  letter-spacing: 0;
}

.faq-hero span {
  display: block;
  max-width: 680px;
  margin-top: 20px;
  color: #626977;
  font-size: 18px;
  line-height: 1.8;
}

.faq-main {
  width: min(1180px, calc(100% - 48px));
  margin: 70px auto 118px;
  display: grid;
  grid-template-columns: 260px minmax(0, 1fr);
  gap: 64px;
  align-items: start;
}

.faq-summary {
  position: sticky;
  top: calc(var(--header-height) + 28px);
  padding: 24px;
  border: 1px solid #f0f1f4;
  border-radius: 8px;
  display: grid;
  gap: 16px;
  background: #fff;
  box-shadow: 0 18px 42px rgba(28, 37, 55, 0.07);
}

.faq-summary strong {
  font-size: 18px;
  font-weight: 900;
}

.faq-summary a {
  color: #6d7480;
  font-size: 15px;
}

.faq-summary a:hover {
  color: var(--color-brand);
}

.consult-link {
  height: 42px;
  margin-top: 8px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: var(--color-brand);
  color: #fff !important;
  font-weight: 800;
}

.faq-content {
  display: grid;
  gap: 46px;
}

.faq-group h2 {
  margin: 0 0 18px;
  font-size: 26px;
  line-height: 1.2;
  font-weight: 900;
}

.faq-item {
  border-top: 1px solid #eceef2;
}

.faq-item:last-child {
  border-bottom: 1px solid #eceef2;
}

.faq-item button {
  width: 100%;
  min-height: 76px;
  border: 0;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
  background: transparent;
  color: #11151d;
  text-align: left;
  font: inherit;
  cursor: pointer;
}

.faq-item button span {
  font-size: 19px;
  line-height: 1.5;
  font-weight: 800;
}

.faq-item button i {
  flex: 0 0 auto;
  width: 34px;
  height: 34px;
  border-radius: 50%;
  display: grid;
  place-items: center;
  background: #f6f7f9;
  color: var(--color-brand);
  font-style: normal;
  font-size: 22px;
  font-weight: 700;
}

.faq-item p {
  max-width: 760px;
  margin: -4px 0 28px;
  color: #626977;
  font-size: 16px;
  line-height: 1.9;
}

.faq-empty {
  min-height: 220px;
  border-radius: 8px;
  display: grid;
  place-items: center;
  background: #f8f9fb;
  color: #7a808b;
  font-weight: 700;
}

@media (max-width: 900px) {
  .faq-main {
    grid-template-columns: 1fr;
    gap: 34px;
    margin-top: 46px;
  }

  .faq-summary {
    position: static;
    display: flex;
    flex-wrap: wrap;
    align-items: center;
  }

  .consult-link {
    padding: 0 20px;
  }
}

@media (max-width: 560px) {
  .faq-hero {
    padding: 104px 18px 50px;
  }

  .faq-hero span {
    font-size: 15px;
  }

  .faq-main {
    width: calc(100% - 32px);
    margin-bottom: 76px;
  }

  .faq-summary {
    padding: 18px;
  }

  .faq-item button {
    min-height: 66px;
  }

  .faq-item button span {
    font-size: 16px;
  }
}
</style>
