<script setup>
import { computed, shallowRef } from 'vue'
import { useRoute } from 'vue-router'
import SiteHeader from './components/SiteHeader.vue'
import FloatingContact from './components/home/FloatingContact.vue'
import FooterSection from './components/home/FooterSection.vue'
import MarketingConsultDialog from './components/consulting/MarketingConsultDialog.vue'

const route = useRoute()
const consultingOpen = shallowRef(false)
const isAdminRoute = computed(() => route.path.startsWith('/admin'))

function openConsulting() {
  consultingOpen.value = true
}

function closeConsulting() {
  consultingOpen.value = false
}
</script>

<template>
  <RouterView v-if="isAdminRoute" />
  <template v-else>
    <SiteHeader @open-consult="openConsulting" />
    <main>
      <RouterView @open-consult="openConsulting" />
    </main>
    <FloatingContact @open-consult="openConsulting" />
    <FooterSection />
    <MarketingConsultDialog v-if="consultingOpen" @close="closeConsulting" />
  </template>
</template>
