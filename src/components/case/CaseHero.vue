<script setup>
defineProps({
  content: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <section class="case-banner" :aria-label="content.title">
    <h1 class="sr-only">{{ content.title }}</h1>
    <img :src="content.image" :alt="content.title" />
  </section>
</template>

<style scoped>
.case-banner {
  height: 340px;
  margin-top: var(--header-height);
  overflow: hidden;
  background: #f3f3f3;
}

.case-banner img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  object-position: center;
}

.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
}

@media (max-width: 768px) {
  .case-banner {
    height: 260px;
  }
}

@media (max-width: 480px) {
  .case-banner {
    height: 210px;
  }
}
</style>
