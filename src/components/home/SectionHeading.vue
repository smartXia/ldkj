<script setup>
defineProps({ title: String, subtitle: String })
</script>

<template>
  <div class="section-heading">
    <h2>{{ title }}</h2>
    <p v-if="subtitle">{{ subtitle }}</p>
  </div>
</template>

<style scoped>
.section-heading {
  text-align: center;
  margin-bottom: 54px;
}

.section-heading h2 {
  margin: 0;
  font-size: clamp(32px, 3.4vw, 44px);
  line-height: 1.2;
  font-weight: 800;
  color: var(--color-ink);
}

.section-heading p {
  margin: 18px 0 0;
  color: var(--color-ink-muted);
  font-size: 18px;
}
</style>
