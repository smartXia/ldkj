<script setup>
import { computed, reactive, shallowRef } from 'vue'
import CaseHero from '../components/case/CaseHero.vue'
import CaseFilterBar from '../components/case/CaseFilterBar.vue'
import CaseGrid from '../components/case/CaseGrid.vue'
import CaseDetailModal from '../components/case/CaseDetailModal.vue'
import { allOption, caseCards, caseFilterGroups, caseHero } from '../data/caseContent'

const selectedFilters = reactive(
  Object.fromEntries(caseFilterGroups.map((group) => [group.key, allOption]))
)
const activeCase = shallowRef(null)

const filteredCases = computed(() =>
  caseCards.filter((item) =>
    caseFilterGroups.every((group) => {
      const value = selectedFilters[group.key]
      return value === allOption || item[group.key] === value
    })
  )
)

function updateFilter({ key, value }) {
  selectedFilters[key] = value
}

function openCaseDetail(item) {
  activeCase.value = item
}

function closeCaseDetail() {
  activeCase.value = null
}
</script>

<template>
  <CaseHero :content="caseHero" />
  <section class="case-list-container">
    <CaseFilterBar
      :groups="caseFilterGroups"
      :selected="selectedFilters"
      @change="updateFilter"
    />
    <CaseGrid :cases="filteredCases" @select="openCaseDetail" />
  </section>
  <CaseDetailModal v-if="activeCase" :case-item="activeCase" @close="closeCaseDetail" />
</template>

<style scoped>
.case-list-container {
  background: #fff;
}
</style>
