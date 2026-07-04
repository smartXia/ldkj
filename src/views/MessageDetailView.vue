<script setup>
import { computed, onMounted, shallowRef } from 'vue'
import { useRoute } from 'vue-router'
import { getArticleById } from '../services/publicApi'

const props = defineProps({
  id: {
    type: String,
    default: '',
  },
})

const route = useRoute()
const article = shallowRef(null)
const articleId = computed(() => props.id || route.params.id || route.params.articleId || '')

onMounted(async () => {
  article.value = await getArticleById(articleId.value)
})
</script>

<template>
  <article v-if="article" class="article-page">
    <header class="article-hero">
      <a href="/message">{{ article.category }}</a>
      <h1>{{ article.title }}</h1>
      <p>
        <span>{{ article.author }}</span>
        <time :datetime="article.date">{{ article.date }}</time>
      </p>
    </header>

    <img class="article-cover" :src="article.image" :alt="article.title" />

    <section class="article-body">
      <p v-for="paragraph in article.content || [article.desc]" :key="paragraph">{{ paragraph }}</p>
    </section>
  </article>
</template>

<style scoped>
.article-page {
  background: #fff;
  color: #111;
}

.article-hero,
.article-body {
  width: 820px;
  max-width: calc(100% - 48px);
  margin: 0 auto;
}

.article-hero {
  padding: 140px 0 40px;
}

.article-hero a {
  color: #ff4848;
  font-weight: 700;
}

.article-hero h1 {
  margin: 18px 0;
  font-size: 42px;
  line-height: 1.25;
}

.article-hero p {
  display: flex;
  gap: 18px;
  margin: 0;
  color: #777;
}

.article-cover {
  width: 820px;
  max-width: calc(100% - 48px);
  margin: 0 auto 48px;
  display: block;
  border-radius: 8px;
}

.article-body {
  padding-bottom: 96px;
}

.article-body p {
  margin: 0 0 22px;
  color: #444;
  font-size: 18px;
  line-height: 2;
}

@media (max-width: 640px) {
  .article-hero h1 {
    font-size: 30px;
  }

  .article-hero p {
    flex-direction: column;
    gap: 4px;
  }
}
</style>
