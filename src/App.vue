<script setup>
import { shallowRef } from 'vue'
import SiteHeader from './components/SiteHeader.vue'
import FloatingContact from './components/home/FloatingContact.vue'
import FooterSection from './components/home/FooterSection.vue'
import MarketingConsultDialog from './components/consulting/MarketingConsultDialog.vue'

const consultingOpen = shallowRef(false)

function openConsulting() {
  consultingOpen.value = true
}

function closeConsulting() {
  consultingOpen.value = false
}
</script>

<template>
  <SiteHeader @open-consult="openConsulting" />
  <main>
    <RouterView @open-consult="openConsulting" />
  </main>
  <FloatingContact @open-consult="openConsulting" />
  <FooterSection />
  <MarketingConsultDialog v-if="consultingOpen" @close="closeConsulting" />
</template>
