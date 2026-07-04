<script setup>
defineProps({
  content: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <section class="message-banner">
    <img v-if="content.image" :src="content.image" :alt="content.title" />
    <h1>{{ content.title }}</h1>
  </section>
</template>

<style scoped>
.message-banner {
  position: relative;
  height: 340px;
  margin-top: var(--header-height);
  overflow: hidden;
  background: #dae3f0;
}

.message-banner img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  object-position: center;
}

.message-banner h1 {
  position: absolute;
  left: max(24px, calc((100vw - 788px) / 2));
  top: 142px;
  width: 846px;
  margin: 0;
  color: #333;
  font-size: 40px;
  line-height: 56px;
  font-weight: 600;
}

@media (max-width: 900px) {
  .message-banner {
    height: 260px;
  }

  .message-banner h1 {
    left: 24px;
    right: 24px;
    top: 108px;
    width: auto;
    font-size: 30px;
    line-height: 1.35;
  }
}

@media (max-width: 480px) {
  .message-banner {
    height: 210px;
  }

  .message-banner h1 {
    top: 82px;
    font-size: 24px;
  }
}
</style>
