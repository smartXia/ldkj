<script setup>
import { computed, reactive, shallowRef } from 'vue'
import { submitPublicForm } from '../../services/publicApi'

const demandOptions = ['整合营销', '新品推广', '电商大促', '社交种草', '账号运营', '口碑管理', '其他']
const form = reactive({
  name: '',
  phone: '',
  company: '',
  demand: '',
})
const dropdownOpen = shallowRef(false)
const status = shallowRef('idle')
const errorMessage = shallowRef('')

const demandLabel = computed(() => form.demand || '请选择营销诉求')
const isSubmitting = computed(() => status.value === 'submitting')

function selectDemand(option) {
  form.demand = option
  dropdownOpen.value = false
}

function validateForm() {
  if (!form.name.trim()) return '请填写您的姓名'
  if (!form.phone.trim()) return '请填写手机号码'
  if (!form.company.trim()) return '请填写公司名称'
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
      source: 'home_footer',
      name: form.name.trim(),
      phone: form.phone.trim(),
      company: form.company.trim(),
      demand: form.demand,
      pageUrl: window.location.href,
    })
    status.value = 'success'
    resetForm()
  } catch (error) {
    status.value = 'error'
    errorMessage.value = '提交失败，请稍后再试或直接联系我们。'
  }
}
</script>

<template>
  <section id="consult" class="consultation-container">
    <h2 class="consultation-title">预约营销咨询</h2>
    <p class="consultation-desc">留下您的联系方式，获取一对一营销诊断</p>

    <div class="consultation-main">
      <div class="consultation-left">
      </div>

      <form class="consultation-form" novalidate @submit.prevent="submitForm">
        <div class="form-row">
          <label class="field field-half">
            <input v-model="form.name" type="text" placeholder="您的姓名 *" autocomplete="name" />
          </label>
          <label class="field field-half">
            <input v-model="form.phone" type="tel" placeholder="请填写手机号码 *" autocomplete="tel" />
          </label>
        </div>

        <label class="field">
          <input v-model="form.company" type="text" placeholder="请填写公司名称 *" autocomplete="organization" />
        </label>

        <div class="field select-field" @mouseleave="dropdownOpen = false">
          <button
            class="select-trigger"
            type="button"
            :class="{ muted: !form.demand }"
            @click="dropdownOpen = !dropdownOpen"
          >
            <span>{{ demandLabel }}</span>
            <i aria-hidden="true"></i>
          </button>
          <div v-if="dropdownOpen" class="select-menu">
            <button v-for="option in demandOptions" :key="option" type="button" @click="selectDemand(option)">
              {{ option }}
            </button>
          </div>
        </div>

        <button class="submit-button" type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? '提交中...' : '立即预约' }}
        </button>
        <p v-if="status === 'success'" class="submit-tip success" role="status">预约信息已提交，我们会尽快与您联系。</p>
        <p v-else-if="status === 'error'" class="submit-tip error" role="alert">{{ errorMessage }}</p>
      </form>
    </div>
  </section>
</template>

<style scoped>
.consultation-container {
  height: 584px;
  padding: 80px 0 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #fff;
}

.consultation-title {
  margin: 0 0 20px;
  color: #000;
  font-size: 40px;
  font-weight: 600;
  line-height: 48px;
}

.consultation-desc {
  margin: 0 0 48px;
  color: #666;
  font-size: 20px;
  line-height: 28px;
  text-align: center;
}

.consultation-main {
  width: 1200px;
  display: flex;
  justify-content: center;
}

.consultation-left {
  width: 280px;
  height: 280px;
  margin-right: 120px;
  flex: 0 0 auto;
}

.consultation-left img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  display: block;
}

.consultation-form {
  width: 400px;
  padding: 22px 0;
  display: grid;
  gap: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.field {
  position: relative;
  height: 40px;
  display: block;
}

.field input {
  width: 100%;
  height: 40px;
  border: 1px solid #eee;
  border-radius: 20px;
  padding: 0 16px;
  outline: none;
  color: #333;
  font-size: 14px;
  background: #fff;
  transition: border-color .2s;
}

.field input:focus {
  border-color: #ff4848;
}

.field input::placeholder {
  color: #bbb;
}

.select-trigger {
  width: 100%;
  height: 40px;
  border: 1px solid #eee;
  border-radius: 20px;
  padding: 0 42px 0 16px;
  display: flex;
  align-items: center;
  background: #fff;
  color: #333;
  font-size: 14px;
  text-align: left;
  cursor: pointer;
}

.select-trigger.muted {
  color: #bbb;
}

.select-trigger span {
  position: static;
  height: auto;
  color: inherit;
  pointer-events: auto;
}

.select-trigger i {
  position: absolute;
  right: 16px;
  top: 50%;
  width: 8px;
  height: 8px;
  border-right: 1px solid #d8d8d8;
  border-bottom: 1px solid #d8d8d8;
  transform: translateY(-70%) rotate(45deg);
}

.select-menu {
  position: absolute;
  left: 0;
  right: 0;
  top: 46px;
  z-index: 5;
  padding: 8px 0;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 10px 30px rgba(0, 0, 0, .12);
}

.select-menu button {
  width: 100%;
  height: 34px;
  border: 0;
  padding: 0 16px;
  background: transparent;
  color: #555;
  font-size: 14px;
  text-align: left;
  cursor: pointer;
}

.select-menu button:hover {
  color: #ff4848;
  background: rgba(255, 72, 72, .06);
}

.submit-button {
  height: 44px;
  border: 0;
  border-radius: 22px;
  background: #ff4848;
  color: #fff;
  font-size: 18px;
  font-weight: 600;
  cursor: pointer;
  transition: background .2s, opacity .2s;
}

.submit-button:hover {
  background: #ff3434;
}

.submit-button:disabled {
  opacity: .72;
  cursor: wait;
}

.submit-tip {
  margin: -8px 0 0;
  font-size: 13px;
  text-align: center;
}

.submit-tip.success {
  color: #14955f;
}

.submit-tip.error {
  color: #ff4848;
}

@media (max-width: 767px) {
  .consultation-container {
    height: auto;
    padding: 32px 0;
  }

  .consultation-title {
    margin-bottom: 4px;
    font-size: 20px;
    line-height: 28px;
  }

  .consultation-desc {
    margin-bottom: 20px;
    font-size: 12px;
    line-height: 16px;
  }

  .consultation-main {
    width: 100%;
  }

  .consultation-left {
    display: none;
  }

  .consultation-form {
    width: 100%;
    padding: 0 24px;
    gap: 14px;
  }

  .form-row {
    grid-template-columns: 1fr;
    gap: 14px;
  }
}
</style>
