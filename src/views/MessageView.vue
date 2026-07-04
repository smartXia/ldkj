<script setup>
import { computed, onMounted, shallowRef, watch } from 'vue'
import MessageFilterBar from '../components/message/MessageFilterBar.vue'
import MessageGrid from '../components/message/MessageGrid.vue'
import MessageHero from '../components/message/MessageHero.vue'
import { getArticles } from '../services/publicApi'

const pageSize = 9
const articles = shallowRef([])
const categories = shallowRef(['全部资讯'])
const hero = shallowRef({ title: '', image: '' })
const allCategory = computed(() => categories.value[0] || '全部资讯')
const activeCategory = shallowRef('全部资讯')
const keyword = shallowRef('')
const currentPage = shallowRef(1)

const filteredArticles = computed(() => {
  const query = keyword.value.trim().toLowerCase()

  return articles.value.filter((article) => {
    const matchesCategory = activeCategory.value === allCategory.value || article.category === activeCategory.value
    const matchesKeyword = !query || `${article.title} ${article.desc}`.toLowerCase().includes(query)

    return matchesCategory && matchesKeyword
  })
})

const pageCount = computed(() => Math.max(1, Math.ceil(filteredArticles.value.length / pageSize)))
const visibleArticles = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredArticles.value.slice(start, start + pageSize)
})

function updateCategory(category) {
  activeCategory.value = category
  currentPage.value = 1
}

function updatePage(page) {
  currentPage.value = page
}

watch(keyword, () => {
  currentPage.value = 1
})

onMounted(async () => {
  const data = await getArticles()
  articles.value = data.items
  categories.value = data.categories
  hero.value = data.hero
  activeCategory.value = categories.value[0]
})
</script>

<template>
  <MessageHero :content="hero" />
  <section class="message-section">
    <MessageFilterBar
      :categories="categories"
      :active-category="activeCategory"
      :keyword="keyword"
      @change-category="updateCategory"
      @update-keyword="keyword = $event"
    />
    <MessageGrid :articles="visibleArticles" />
    <nav v-if="pageCount > 1" class="pager" aria-label="资讯分页">
      <button
        v-for="page in pageCount"
        :key="page"
        type="button"
        :class="{ active: page === currentPage }"
        @click="updatePage(page)"
      >
        {{ page }}
      </button>
    </nav>
  </section>
</template>

<style scoped>
.message-section {
  background: #fff;
}

.pager {
  display: flex;
  justify-content: center;
  gap: 10px;
  padding: 0 0 96px;
}

.pager button {
  min-width: 38px;
  height: 38px;
  border: 1px solid #e9e9e9;
  border-radius: 19px;
  background: #fff;
  color: #555;
  cursor: pointer;
}

.pager button.active {
  border-color: #ff4848;
  background: #ff4848;
  color: #fff;
}
</style>
