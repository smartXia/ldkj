<script setup>
import HeroCarousel from '../components/home/HeroCarousel.vue'
import MarketingModel from '../components/home/MarketingModel.vue'
import ServiceSection from '../components/home/ServiceSection.vue'
import BrandCases from '../components/home/BrandCases.vue'
import IntroStats from '../components/home/IntroStats.vue'
import BrandWall from '../components/home/BrandWall.vue'
import ContactCta from '../components/home/ContactCta.vue'

const emit = defineEmits(['open-consult'])
</script>

<template>
  <HeroCarousel />
  <MarketingModel />
  <ServiceSection />
  <BrandCases />
  <IntroStats />
  <BrandWall />
  <ContactCta @open-consult="emit('open-consult')" />
</template>
