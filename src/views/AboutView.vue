<script setup>
import { onMounted, ref } from 'vue'
import AboutAnchorNav from '../components/about/AboutAnchorNav.vue'
import AboutBrands from '../components/about/AboutBrands.vue'
import AboutHero from '../components/about/AboutHero.vue'
import AboutHistory from '../components/about/AboutHistory.vue'
import AboutHonors from '../components/about/AboutHonors.vue'
import AboutIntro from '../components/about/AboutIntro.vue'
import AboutTeam from '../components/about/AboutTeam.vue'
import AboutValues from '../components/about/AboutValues.vue'
import { useI18n } from '../composables/useI18n'
import { getAboutContent } from '../services/publicApi'

const { messages } = useI18n()
const aboutContent = ref(messages.value.about)

onMounted(async () => {
  aboutContent.value = await getAboutContent(messages.value.about)
})
</script>

<template>
  <AboutHero :content="aboutContent.hero" />
  <AboutAnchorNav :items="aboutContent.anchors" />
  <AboutIntro :intro="aboutContent.intro" :meanings="aboutContent.meanings" />
  <AboutValues :title="aboutContent.valuesTitle" :values="aboutContent.values" />
  <AboutHistory :content="aboutContent.history" />
  <AboutHonors :honors="aboutContent.honors" />
  <AboutBrands :brands="aboutContent.brands" />
  <AboutTeam :title="aboutContent.teamTitle" />
</template>

<style scoped>
:global(main) {
  overflow: visible;
}

:global(body.about-subnav-pinned .site-header) {
  opacity: 0;
  pointer-events: none;
  transform: translateY(-100%);
}

:global(body.about-subnav-pinned .about-anchor-shell) {
  top: 0 !important;
}
</style>
