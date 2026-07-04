<script setup>
import { onMounted, shallowRef } from 'vue'
import HeroCarousel from '../components/home/HeroCarousel.vue'
import BrandCases from '../components/home/BrandCases.vue'
import BrandWall from '../components/home/BrandWall.vue'
import ContactCta from '../components/home/ContactCta.vue'
import { getCases, getHomeContent } from '../services/publicApi'

const emit = defineEmits(['open-consult'])
const homeContent = shallowRef({ heroSlides: [], logoList: [] })
const cases = shallowRef([])

onMounted(async () => {
  const [home, caseContent] = await Promise.all([
    getHomeContent(),
    getCases(),
  ])
  homeContent.value = home
  cases.value = caseContent.items || []
})
</script>

<template>
  <HeroCarousel :slides="homeContent.heroSlides" />
  <BrandCases :cases="cases" />
  <BrandWall :logos="homeContent.logoList" />
  <ContactCta @open-consult="emit('open-consult')" />
</template>
