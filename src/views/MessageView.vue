<script setup>
import { computed, shallowRef } from 'vue'
import MessageFilterBar from '../components/message/MessageFilterBar.vue'
import MessageGrid from '../components/message/MessageGrid.vue'
import MessageHero from '../components/message/MessageHero.vue'
import { messageArticles, messageCategories, messageHero } from '../data/messageContent'

const allCategory = messageCategories[0]
const activeCategory = shallowRef(allCategory)
const keyword = shallowRef('')

const visibleArticles = computed(() => {
  const query = keyword.value.trim().toLowerCase()

  return messageArticles.filter((article) => {
    const matchesCategory = activeCategory.value === allCategory || article.category === activeCategory.value
    const matchesKeyword = !query || `${article.title} ${article.desc}`.toLowerCase().includes(query)

    return matchesCategory && matchesKeyword
  })
})

function updateCategory(category) {
  activeCategory.value = category
}
</script>

<template>
  <MessageHero :content="messageHero" />
  <section class="message-section">
    <MessageFilterBar
      :categories="messageCategories"
      :active-category="activeCategory"
      :keyword="keyword"
      @change-category="updateCategory"
      @update-keyword="keyword = $event"
    />
    <MessageGrid :articles="visibleArticles" />
  </section>
</template>

<style scoped>
.message-section {
  background: #fff;
}
</style>
