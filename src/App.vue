<script setup>
import { onMounted, shallowRef } from 'vue'
import SiteHeader from './components/SiteHeader.vue'
import FooterSection from './components/home/FooterSection.vue'
import MarketingConsultDialog from './components/consulting/MarketingConsultDialog.vue'
import { getHomeContent } from './services/publicApi'

const consultingOpen = shallowRef(false)
const siteConfig = shallowRef({})

function openConsulting() {
  consultingOpen.value = true
}

function closeConsulting() {
  consultingOpen.value = false
}

onMounted(async () => {
  siteConfig.value = await getHomeContent()
})
</script>

<template>
  <SiteHeader :site="siteConfig" @open-consult="openConsulting" />
  <main>
    <RouterView @open-consult="openConsulting" />
  </main>
  <FooterSection :site="siteConfig.site" />
  <MarketingConsultDialog v-if="consultingOpen" @close="closeConsulting" />
</template>
