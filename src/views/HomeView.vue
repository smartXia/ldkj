<script setup>
import { onMounted, shallowRef } from 'vue'
import HeroCarousel from '../components/home/HeroCarousel.vue'
import MarketingModel from '../components/home/MarketingModel.vue'
import ServiceSection from '../components/home/ServiceSection.vue'
import BrandCases from '../components/home/BrandCases.vue'
import IntroStats from '../components/home/IntroStats.vue'
import BrandWall from '../components/home/BrandWall.vue'
import ContactCta from '../components/home/ContactCta.vue'
import HomeAbout from '../components/home/HomeAbout.vue'
import { getCases, getHomeContent, getServices } from '../services/publicApi'

const emit = defineEmits(['open-consult'])
const homeContent = shallowRef({ heroSlides: [], marketingModules: [], logoList: [] })
const services = shallowRef([])
const cases = shallowRef([])

onMounted(async () => {
  const [home, serviceContent, caseContent] = await Promise.all([
    getHomeContent(),
    getServices(),
    getCases(),
  ])
  homeContent.value = home
  services.value = serviceContent
  cases.value = caseContent.items || []
})
</script>

<template>
  <HeroCarousel :slides="homeContent.heroSlides" />
  <MarketingModel :modules="homeContent.marketingModules" />
  <ServiceSection :services="services" />
  <BrandCases :cases="cases" />
  <IntroStats />
  <BrandWall :logos="homeContent.logoList" />
  <ContactCta @open-consult="emit('open-consult')" />
  <HomeAbout />
</template>
