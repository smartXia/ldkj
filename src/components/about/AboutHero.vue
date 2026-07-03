<script setup>
defineProps({
  content: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <section class="about-hero">
    <img :src="content.image" alt="" aria-hidden="true" />
    <div class="about-hero-copy">
      <h1>{{ content.title }}</h1>
      <p>
        <span v-for="line in content.desc" :key="line">{{ line }}</span>
      </p>
    </div>
  </section>
</template>

<style scoped>
.about-hero {
  position: relative;
  height: 520px;
  margin-top: var(--header-height);
  overflow: hidden;
  background: #eef2f7;
}

.about-hero img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  object-position: center;
}

.about-hero-copy {
  position: absolute;
  left: 50%;
  top: 186px;
  width: 1040px;
  transform: translateX(-50%);
  text-align: center;
  color: #333;
}

.about-hero h1 {
  margin: 0 0 16px;
  font-size: 40px;
  line-height: 56px;
  font-weight: 600;
}

.about-hero p {
  width: 580px;
  margin: 0 auto;
  font-size: 20px;
  line-height: 32px;
}

.about-hero span {
  display: block;
}

@media (max-width: 900px) {
  .about-hero {
    height: 265px;
  }

  .about-hero-copy {
    top: 104px;
    width: calc(100% - 40px);
  }

  .about-hero h1 {
    margin-bottom: 0;
    font-size: 24px;
    line-height: 34px;
  }

  .about-hero p {
    display: none;
  }
}
</style>
