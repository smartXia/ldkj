<script setup>
defineProps({
  groups: {
    type: Array,
    required: true,
  },
  selected: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['change'])
</script>

<template>
  <div class="filter-section" aria-label="案例筛选">
    <div v-for="group in groups" :key="group.key" class="filter-row">
      <div class="filter-label">{{ group.label }}</div>
      <div class="filter-options">
        <button
          v-for="option in group.options"
          :key="option"
          type="button"
          :class="{ active: selected[group.key] === option }"
          @click="emit('change', { key: group.key, value: option })"
        >
          {{ option }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.filter-section {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 36px auto 24px;
}

.filter-row {
  display: grid;
  grid-template-columns: 68px minmax(0, 1fr);
  align-items: start;
  gap: 16px;
}

.filter-row + .filter-row {
  margin-top: 14px;
}

.filter-label {
  min-height: 28px;
  display: inline-flex;
  align-items: center;
  color: #111;
  font-size: 15px;
  font-weight: 700;
}

.filter-options {
  display: flex;
  flex-wrap: wrap;
  gap: 8px 12px;
}

.filter-options button {
  min-height: 28px;
  border: 0;
  border-radius: 16px;
  background: transparent;
  color: #555;
  padding: 0 12px;
  font-size: 14px;
  line-height: 28px;
  transition: color var(--motion-fast), background var(--motion-fast);
}

.filter-options button.active,
.filter-options button:hover {
  color: var(--color-brand);
  background: var(--color-brand-soft);
}

@media (max-width: 640px) {
  .filter-section {
    max-width: calc(100% - 32px);
    margin-top: 28px;
  }

  .filter-row {
    grid-template-columns: 1fr;
    gap: 8px;
  }
}
</style>
