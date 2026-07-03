<script setup>
import AboutAnchorNav from '../components/about/AboutAnchorNav.vue'
import AboutBrands from '../components/about/AboutBrands.vue'
import AboutHero from '../components/about/AboutHero.vue'
import AboutHistory from '../components/about/AboutHistory.vue'
import AboutHonors from '../components/about/AboutHonors.vue'
import AboutIntro from '../components/about/AboutIntro.vue'
import AboutTeam from '../components/about/AboutTeam.vue'
import AboutValues from '../components/about/AboutValues.vue'
import { useI18n } from '../composables/useI18n'

const { messages } = useI18n()
</script>

<template>
  <AboutHero :content="messages.about.hero" />
  <AboutAnchorNav :items="messages.about.anchors" />
  <AboutIntro :intro="messages.about.intro" :meanings="messages.about.meanings" />
  <AboutValues :title="messages.about.valuesTitle" :values="messages.about.values" />
  <AboutHistory :content="messages.about.history" />
  <AboutHonors />
  <AboutBrands :brands="messages.about.brands" />
  <AboutTeam :title="messages.about.teamTitle" />
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
