<script setup>
import { onBeforeUnmount, onMounted } from 'vue'

defineProps({
  caseItem: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['close'])

function handleKeydown(event) {
  if (event.key === 'Escape') {
    emit('close')
  }
}

onMounted(() => {
  document.body.classList.add('modal-open')
  window.addEventListener('keydown', handleKeydown)
})

onBeforeUnmount(() => {
  document.body.classList.remove('modal-open')
  window.removeEventListener('keydown', handleKeydown)
})
</script>

<template>
  <Teleport to="body">
    <div class="case-modal" role="dialog" aria-modal="true" :aria-label="caseItem.title">
      <button class="modal-backdrop" type="button" aria-label="关闭案例详情" @click="emit('close')"></button>
      <button class="modal-close" type="button" aria-label="关闭" @click="emit('close')"></button>
      <div class="modal-panel">
        <img :src="caseItem.detailImage" :alt="caseItem.title" />
      </div>
    </div>
  </Teleport>
</template>

<style scoped>
.case-modal {
  position: fixed;
  inset: 0;
  z-index: 200;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding: 56px 100px 72px;
}

.modal-backdrop {
  position: fixed;
  inset: 0;
  border: 0;
  background: rgba(0, 0, 0, 0.58);
  cursor: default;
}

.modal-panel {
  position: relative;
  z-index: 1;
  width: min(1080px, calc(100vw - 200px));
  max-height: calc(100vh - 112px);
  overflow: auto;
  background: #fff;
}

.modal-panel img {
  width: 100%;
  height: auto;
  display: block;
}

.modal-close {
  position: fixed;
  top: 38px;
  right: 40px;
  z-index: 2;
  width: 48px;
  height: 48px;
  border: 0;
  border-radius: 50%;
  background: rgba(39, 42, 49, 0.92);
}

.modal-close::before,
.modal-close::after {
  content: "";
  position: absolute;
  left: 15px;
  top: 23px;
  width: 18px;
  height: 2px;
  background: #fff;
}

.modal-close::before {
  transform: rotate(45deg);
}

.modal-close::after {
  transform: rotate(-45deg);
}

@media (max-width: 768px) {
  .case-modal {
    padding: 52px 16px 36px;
  }

  .modal-panel {
    width: 100%;
    max-height: calc(100vh - 88px);
  }

  .modal-close {
    top: 12px;
    right: 12px;
  }
}
</style>
