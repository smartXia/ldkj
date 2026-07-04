<script setup>
import { computed, onBeforeUnmount, onMounted, reactive, shallowRef } from 'vue'
import { getConsultContent, submitPublicForm } from '../../services/publicApi'

const emit = defineEmits(['close'])

const defaultDemandOptions = ['整合营销', '新品推广', '电商大促', '社交种草', '账号运营', '舆情优化', '其他']
const form = reactive({
  name: '',
  phone: '',
  company: '',
  demand: '',
})
const dropdownOpen = shallowRef(false)
const status = shallowRef('idle')
const errorMessage = shallowRef('')
const consultContent = shallowRef(null)

const demandOptions = computed(() => consultContent.value?.demandOptions?.length ? consultContent.value.demandOptions : defaultDemandOptions)
const demandLabel = computed(() => form.demand || consultContent.value?.demandPlaceholder || '请选择营销诉求')
const isSubmitting = computed(() => status.value === 'submitting')

function close() {
  emit('close')
}

function selectDemand(option) {
  form.demand = option
  dropdownOpen.value = false
}

function validateForm() {
  if (!form.name.trim()) return '请填写您的姓名'
  if (!form.phone.trim()) return '请填写手机号码'
  if (!form.company.trim()) return '请填写公司名称'
  if (!form.demand) return '请选择营销诉求'
  return ''
}

function resetForm() {
  form.name = ''
  form.phone = ''
  form.company = ''
  form.demand = ''
}

async function submitForm() {
  const validationError = validateForm()
  if (validationError) {
    errorMessage.value = validationError
    status.value = 'error'
    return
  }

  status.value = 'submitting'
  errorMessage.value = ''

  try {
    await submitPublicForm({
      source: 'marketing_consult_dialog',
      name: form.name.trim(),
      phone: form.phone.trim(),
      company: form.company.trim(),
      requirement: form.demand,
    })
    status.value = 'success'
    resetForm()
  } catch (error) {
    status.value = 'error'
    errorMessage.value = consultContent.value?.errorText || '提交失败，请稍后再试或直接联系我们。'
  }
}

function closeDropdown() {
  dropdownOpen.value = false
}

function handleKeydown(event) {
  if (event.key !== 'Escape') return

  if (dropdownOpen.value) {
    dropdownOpen.value = false
    return
  }

  close()
}

onMounted(async () => {
  document.body.classList.add('modal-open')
  window.addEventListener('keydown', handleKeydown)
  window.addEventListener('click', closeDropdown)
  consultContent.value = await getConsultContent()
})

onBeforeUnmount(() => {
  document.body.classList.remove('modal-open')
  window.removeEventListener('keydown', handleKeydown)
  window.removeEventListener('click', closeDropdown)
})
</script>

<template>
  <Teleport to="body">
    <div class="consult-modal" role="dialog" aria-modal="true" aria-labelledby="consult-title">
      <button class="consult-backdrop" type="button" aria-label="鍏抽棴钀ラ攢鍜ㄨ" @click="close"></button>

      <section class="consult-dialog">
        <button class="consult-close" type="button" aria-label="鍏抽棴" @click="close"></button>

        <div class="consult-right">
          <h2 id="consult-title">{{ consultContent?.title || '预约咨询' }}</h2>
          <form class="consult-form" novalidate @submit.prevent="submitForm">
            <label>
              <span class="field-label">姓名</span>
              <input v-model="form.name" type="text" placeholder="您的姓名" autocomplete="name" />
            </label>

            <label>
              <span class="field-label">手机号码</span>
              <input v-model="form.phone" type="tel" placeholder="请填写手机号码" autocomplete="tel" />
            </label>

            <label>
              <span class="field-label">公司名称</span>
              <input v-model="form.company" type="text" placeholder="请填写公司名称" autocomplete="organization" />
            </label>

            <label class="demand-field" @click.stop>
              <span class="field-label">营销诉求</span>
              <button class="select-trigger" type="button" :class="{ muted: !form.demand }" :disabled="isSubmitting" @click="dropdownOpen = !dropdownOpen">
                {{ demandLabel }}
                <i aria-hidden="true"></i>
              </button>
              <div v-if="dropdownOpen" class="select-menu">
                <button v-for="option in demandOptions" :key="option" type="button" @click="selectDemand(option)">
                  {{ option }}
                </button>
              </div>
            </label>

            <button class="submit-button" type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? (consultContent?.submittingText || '提交中...') : (consultContent?.submitText || '立即预约') }}
            </button>
            <p v-if="status === 'success'" class="submit-tip success" role="status">{{ consultContent?.successText || '预约信息已提交，我们会尽快与您联系。' }}</p>
            <p v-else-if="status === 'error'" class="submit-tip error" role="alert">{{ errorMessage }}</p>
          </form>
        </div>
      </section>
    </div>
  </Teleport>
</template>

<style scoped>
.consult-modal {
  position: fixed;
  inset: 0;
  z-index: 180;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 58px;
}

.consult-backdrop {
  position: fixed;
  inset: 0;
  border: 0;
  background: rgba(0, 0, 0, 0.5);
}

.consult-dialog {
  position: relative;
  z-index: 1;
  width: 460px;
  min-height: 500px;
  background: #fff;
}

.consult-close {
  position: absolute;
  right: 18px;
  top: 18px;
  z-index: 3;
  width: 36px;
  height: 36px;
  border: 0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.92);
}

.consult-close::before,
.consult-close::after {
  content: "";
  position: absolute;
  left: 10px;
  top: 17px;
  width: 16px;
  height: 2px;
  background: #222;
}

.consult-close::before { transform: rotate(45deg); }
.consult-close::after { transform: rotate(-45deg); }

.consult-left {
  position: relative;
  width: 420px;
  height: 600px;
  overflow: hidden;
  flex: 0 0 420px;
  background: #eaf0f6;
}

.left-bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.left-content {
  position: relative;
  z-index: 1;
  padding: 50px 48px 0 56px;
}

.consult-logo {
  width: 124px;
  height: 40px;
  display: block;
  object-fit: contain;
}

.left-content h2 {
  margin: 20px 0 0;
  color: #333;
  font-size: 26px;
  line-height: 30px;
  font-weight: 600;
}

.left-tip {
  margin: 40px 0 0;
  color: #333;
  font-size: 12px;
  line-height: 18px;
}

.benefit-list {
  list-style: none;
  margin: 30px 0 0;
  padding: 0 20px;
}

.benefit-list li {
  height: 20px;
  margin: 0 0 12px;
  display: flex;
  align-items: center;
  gap: 12px;
  color: #333;
  font-size: 14px;
  line-height: 20px;
}

.benefit-list li:last-child {
  margin-bottom: 0;
}

.benefit-list img {
  width: 14px;
  height: 14px;
  flex: 0 0 auto;
}

.qr-box {
  width: 316px;
  margin-top: 56px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.qr-box img {
  width: 108px;
  height: 108px;
  display: block;
}

.qr-box span {
  margin-top: 8px;
  color: #333;
  font-size: 12px;
  line-height: 16px;
}

.consult-right {
  width: 100%;
  padding: 70px 40px 0;
}

.consult-right h2 {
  margin: 0 0 36px;
  color: #333;
  font-size: 22px;
  line-height: 22px;
  font-weight: 600;
}

.consult-form {
  display: grid;
  width: 360px;
  gap: 16px;
}

.consult-form label {
  display: grid;
  gap: 2px;
  color: #333;
  font-size: 14px;
  line-height: 22px;
}

.field-label::before {
  content: "*";
  color: var(--color-brand);
  margin-right: 1px;
}

.consult-form input,
.select-trigger {
  width: 100%;
  height: 40px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background: #fff;
  color: #333;
  padding: 0 15px;
  font-size: 14px;
  line-height: 38px;
  outline: none;
}

.consult-form input::placeholder,
.select-trigger.muted {
  color: #b8bcc4;
}

.consult-form input:focus,
.select-trigger:focus-visible {
  border-color: var(--color-brand);
}

.demand-field {
  position: relative;
}

.select-trigger {
  display: flex;
  align-items: center;
  justify-content: space-between;
  text-align: left;
}

.select-trigger i {
  width: 8px;
  height: 8px;
  border-right: 1px solid #999;
  border-bottom: 1px solid #999;
  transform: translateY(-2px) rotate(45deg);
}

.select-menu {
  position: absolute;
  left: 0;
  right: 0;
  bottom: calc(100% - 22px);
  z-index: 4;
  padding: 12px 0;
  border: 1px solid #ececec;
  border-radius: 4px;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.select-menu button {
  width: 100%;
  height: 34px;
  border: 0;
  background: transparent;
  color: #333;
  padding: 0 20px;
  text-align: left;
  font-size: 14px;
  line-height: 34px;
}

.select-menu button:hover {
  color: var(--color-brand);
  background: var(--color-brand-soft);
}

.submit-button {
  height: 40px;
  border: 1px solid #ff4848;
  border-radius: 4px;
  background: #ff4848;
  color: #fff;
  font-size: 16px;
  line-height: 16px;
  font-weight: 600;
}

.submit-button:disabled,
.select-trigger:disabled {
  cursor: not-allowed;
  opacity: 0.72;
}

.submit-tip {
  margin: 0;
  font-size: 13px;
}

.submit-tip.success {
  color: var(--color-brand);
}

.submit-tip.error {
  color: #d93025;
}

@media (max-width: 900px) {
  .consult-modal {
    align-items: flex-start;
    padding: 24px 16px;
    overflow: auto;
  }

  .consult-dialog {
    width: min(100%, 520px);
    height: auto;
    display: block;
  }

  .consult-left,
  .consult-right {
    width: 100%;
    height: auto;
    flex-basis: auto;
    min-height: 0;
  }

  .left-content {
    padding: 34px 28px 30px;
  }

  .qr-box {
    width: 100%;
    margin-top: 28px;
  }

  .consult-right {
    padding: 34px 28px 34px;
  }

  .consult-form {
    width: 100%;
  }

  .select-menu {
    bottom: auto;
    top: calc(100% + 4px);
  }
}
</style>

