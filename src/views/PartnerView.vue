<script setup>
import { reactive, shallowRef } from 'vue'
import { submitPublicForm } from '../services/publicApi'

const form = reactive({
  name: '',
  company: '',
  title: '',
  phone: '',
  email: '',
  demand: '',
})
const status = shallowRef('idle')
const errorMessage = shallowRef('')

function validateForm() {
  if (!form.name.trim()) return '请填写姓名'
  if (!form.company.trim()) return '请填写公司名称'
  if (!form.phone.trim()) return '请填写联系电话'
  return ''
}

function resetForm() {
  form.name = ''
  form.company = ''
  form.title = ''
  form.phone = ''
  form.email = ''
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
      ...form,
      source: 'partner',
      pageUrl: window.location.href,
    })
    status.value = 'success'
    resetForm()
  } catch (error) {
    status.value = 'error'
    errorMessage.value = '提交失败，请稍后再试。'
  }
}

</script>

<template>
  <section class="partner-page">
    <header class="partner-hero">
      <p>Partnership</p>
      <h1>开启合作</h1>
      <span>告诉我们你的品牌目标、平台方向和当前挑战，灵动信息会尽快安排顾问对接。</span>
    </header>

    <div class="partner-main">
      <form class="partner-form" novalidate @submit.prevent="submitForm">
        <label>
          <span>姓名 *</span>
          <input v-model="form.name" autocomplete="name" />
        </label>
        <label>
          <span>公司 *</span>
          <input v-model="form.company" autocomplete="organization" />
        </label>
        <label>
          <span>职位</span>
          <input v-model="form.title" autocomplete="organization-title" />
        </label>
        <label>
          <span>电话 *</span>
          <input v-model="form.phone" type="tel" autocomplete="tel" />
        </label>
        <label>
          <span>邮箱</span>
          <input v-model="form.email" type="email" autocomplete="email" />
        </label>
        <label class="wide">
          <span>需求描述</span>
          <textarea v-model="form.demand" rows="5"></textarea>
        </label>
        <div class="partner-submit">
          <button type="submit" :disabled="status === 'submitting'">
            {{ status === 'submitting' ? '提交中...' : '提交合作需求' }}
          </button>
          <p v-if="status === 'success'" class="form-tip success">提交成功，我们会尽快联系您。</p>
          <p v-else-if="status === 'error'" class="form-tip error">{{ errorMessage }}</p>
        </div>
      </form>

      <aside class="partner-aside">
        <h2>合作流程</h2>
        <ol>
          <li>需求提交</li>
          <li>顾问沟通</li>
          <li>策略方案</li>
          <li>项目启动</li>
        </ol>
        <RouterLink class="faq-entry" to="/faq">查看常见问题</RouterLink>
      </aside>
    </div>
  </section>
</template>

<style scoped>
.partner-page {
  background: #fff;
  color: #111;
}

.partner-hero {
  padding: 140px max(24px, calc((100vw - 1200px) / 2)) 72px;
  background: #fff8f8;
}

.partner-hero p {
  margin: 0 0 12px;
  color: #ff4848;
  font-weight: 700;
}

.partner-hero h1 {
  margin: 0;
  font-size: 48px;
}

.partner-hero span {
  display: block;
  max-width: 680px;
  margin-top: 18px;
  color: #666;
  font-size: 18px;
  line-height: 1.8;
}

.partner-main {
  width: 1200px;
  max-width: calc(100% - 48px);
  margin: 72px auto 120px;
  display: grid;
  grid-template-columns: 1fr 360px;
  gap: 56px;
}

.partner-form {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 18px;
}

.partner-form label {
  display: grid;
  gap: 8px;
  color: #555;
}

.partner-form .wide {
  grid-column: 1 / -1;
}

.partner-form input,
.partner-form textarea {
  width: 100%;
  min-width: 0;
  border: 1px solid #e6e6e6;
  border-radius: 8px;
  padding: 13px 14px;
  color: #222;
  font: inherit;
}

.partner-form input:focus,
.partner-form textarea:focus {
  border-color: #ff4848;
  outline: none;
}

.partner-submit {
  grid-column: 1 / -1;
  display: flex;
  align-items: center;
  gap: 16px;
  flex-wrap: wrap;
}

.partner-form button {
  width: max-content;
  border: 0;
  border-radius: 22px;
  padding: 13px 28px;
  background: #ff4848;
  color: #fff;
  font-weight: 700;
  white-space: nowrap;
}

.partner-form button:disabled {
  opacity: .72;
}

.form-tip {
  align-self: center;
  margin: 0;
}

.form-tip.success {
  color: #14955f;
}

.form-tip.error {
  color: #ff4848;
}

.partner-aside h2 {
  margin: 0 0 18px;
  font-size: 24px;
}

.partner-aside ol {
  margin: 0 0 38px;
  padding-left: 20px;
  color: #555;
  line-height: 2;
}

.faq-entry {
  width: max-content;
  min-height: 44px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 24px;
  background: #ff4848;
  color: #fff;
  font-weight: 700;
}

@media (max-width: 860px) {
  .partner-main,
  .partner-form {
    grid-template-columns: 1fr;
  }

  .partner-hero {
    padding: 112px 24px 56px;
  }

  .partner-hero h1 {
    font-size: 40px;
  }

  .partner-main {
    margin: 48px auto 88px;
  }
}

@media (max-width: 520px) {
  .partner-main {
    max-width: calc(100% - 32px);
  }

  .partner-hero {
    padding: 96px 16px 44px;
  }

  .partner-hero h1 {
    font-size: 34px;
  }

  .partner-submit {
    align-items: stretch;
    gap: 12px;
  }

  .partner-form button {
    width: 100%;
  }
}
</style>
