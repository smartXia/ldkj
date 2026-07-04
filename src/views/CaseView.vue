<script setup>
import { computed, onMounted, reactive, shallowRef } from 'vue'
import CaseHero from '../components/case/CaseHero.vue'
import CaseFilterBar from '../components/case/CaseFilterBar.vue'
import CaseGrid from '../components/case/CaseGrid.vue'
import { allOption } from '../data/caseContent'
import { getCases } from '../services/publicApi'

const pageSize = 9
const caseHero = shallowRef({ title: '', image: '' })
const caseFilterGroups = shallowRef([])
const caseCards = shallowRef([])
const selectedFilters = reactive({})
const currentPage = shallowRef(1)

const filteredCases = computed(() =>
  caseCards.value.filter((item) =>
    caseFilterGroups.value.every((group) => {
      const value = selectedFilters[group.key]
      return value === allOption || item[group.key] === value
    })
  )
)

const pageCount = computed(() => Math.max(1, Math.ceil(filteredCases.value.length / pageSize)))
const visibleCases = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredCases.value.slice(start, start + pageSize)
})

function updateFilter({ key, value }) {
  selectedFilters[key] = value
  currentPage.value = 1
}

function updatePage(page) {
  currentPage.value = page
}

onMounted(async () => {
  const data = await getCases()
  caseHero.value = data.hero
  caseFilterGroups.value = data.filters
  caseCards.value = data.items
  data.filters.forEach((group) => {
    selectedFilters[group.key] = allOption
  })
})
</script>

<template>
  <CaseHero :content="caseHero" />
  <section class="case-list-container">
    <CaseFilterBar
      :groups="caseFilterGroups"
      :selected="selectedFilters"
      @change="updateFilter"
    />
    <CaseGrid :cases="visibleCases" />
    <nav v-if="pageCount > 1" class="pager" aria-label="案例分页">
      <button
        v-for="page in pageCount"
        :key="page"
        type="button"
        :class="{ active: page === currentPage }"
        @click="updatePage(page)"
      >
        {{ page }}
      </button>
    </nav>
  </section>
</template>

<style scoped>
.case-list-container {
  background: #fff;
}

.pager {
  display: flex;
  justify-content: center;
  gap: 10px;
  padding: 0 0 96px;
}

.pager button {
  min-width: 38px;
  height: 38px;
  border: 1px solid #e9e9e9;
  border-radius: 19px;
  background: #fff;
  color: #555;
  cursor: pointer;
}

.pager button.active {
  border-color: #ff4848;
  background: #ff4848;
  color: #fff;
}
</style>
