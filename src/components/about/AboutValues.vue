<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
  values: {
    type: Array,
    required: true,
  },
})
</script>

<template>
  <section id="about-values" class="about-values">
    <h2>{{ title }}</h2>
    <ul>
      <li v-for="item in values" :key="item.title">
        <img :src="item.icon" :alt="item.title" />
        <h3>{{ item.title }}</h3>
        <p v-for="line in item.lines" :key="line">{{ line }}</p>
      </li>
    </ul>
  </section>
</template>

<style scoped>
.about-values {
  min-height: 430px;
  padding: 8px 0 78px;
  text-align: center;
  background: #fff;
}

.about-values h2 {
  margin: 0 0 60px;
  color: #000;
  font-size: 40px;
  line-height: 56px;
  font-weight: 700;
}

.about-values ul {
  width: 1160px;
  margin: 0 auto;
  padding: 0;
  list-style: none;
  display: grid;
  grid-template-columns: repeat(4, 260px);
  justify-content: space-between;
  gap: 0;
}

.about-values li {
  min-height: 220px;
}

.about-values img {
  width: 58px;
  height: 58px;
  margin-bottom: 42px;
  object-fit: contain;
}

.about-values h3 {
  margin: 0 0 22px;
  color: #000;
  font-size: 24px;
  line-height: 34px;
  font-weight: 700;
}

.about-values p {
  margin: 0;
  color: rgba(102, 102, 102, 0.78);
  font-size: 13px;
  line-height: 22px;
}

@media (max-width: 760px) {
  .about-values {
    min-height: 0;
    padding: 33px 20px 32px;
  }

  .about-values h2 {
    margin-bottom: 20px;
    font-size: 24px;
    line-height: 29px;
  }

  .about-values ul {
    width: 100%;
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 10px 10px;
  }

  .about-values li {
    min-height: 175px;
    padding-top: 20px;
  }

  .about-values img {
    width: 47px;
    height: 47px;
    margin-bottom: 6px;
  }

  .about-values h3 {
    margin-bottom: 7px;
    font-size: 14px;
    line-height: 20px;
  }

  .about-values p {
    width: min(137px, 100%);
    margin: 0 auto;
    font-size: 12px;
    line-height: 16px;
  }
}

@media (max-width: 1180px) and (min-width: 761px) {
  .about-values ul {
    width: calc(100% - 40px);
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 24px;
  }

  .about-values h3 {
    font-size: 22px;
  }
}
</style>
