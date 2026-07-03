<script setup>
const emit = defineEmits(['open-consult'])

const sideActions = [
  {
    label: '电话咨询',
    href: 'tel:4000035180',
    icon: 'phone',
  },
  {
    label: '微信咨询',
    action: 'consult',
    icon: 'wechat',
  },
  {
    label: '邮件咨询',
    href: 'mailto:service@wsdsocial.com',
    icon: 'mail',
  },
  {
    label: '在线客服',
    action: 'consult',
    icon: 'service',
  },
]

function handleAction(action) {
  if (action.action === 'consult') {
    emit('open-consult')
  }
}
</script>

<template>
  <aside class="float-rail" aria-label="联系微思敦">
    <a
      v-for="action in sideActions"
      :key="action.label"
      :href="action.href || 'javascript:void(0)'"
      :aria-label="action.label"
      @click="handleAction(action)"
    >
      <span :class="['icon', action.icon]" aria-hidden="true"></span>
    </a>
    <a class="back-top" href="#" aria-label="返回顶部"></a>
  </aside>
</template>

<style scoped>
.float-rail {
  position: fixed;
  right: 28px;
  top: 50%;
  z-index: 60;
  display: flex;
  flex-direction: column;
  gap: 8px;
  transform: translateY(-50%);
}

.float-rail a {
  width: 44px;
  height: 44px;
  border-radius: 6px;
  display: grid;
  place-items: center;
  background: #fff;
  color: var(--color-brand);
  box-shadow: 0 8px 20px rgba(20, 30, 48, 0.1);
  transition: transform var(--motion-fast), box-shadow var(--motion-fast);
}

.float-rail a:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 28px rgba(20, 30, 48, 0.14);
}

.icon {
  position: relative;
  width: 22px;
  height: 22px;
  display: block;
}

.icon.phone::before {
  content: "";
  position: absolute;
  inset: 3px 6px;
  border: 2px solid currentColor;
  border-radius: 10px 10px 6px 6px;
  transform: rotate(-28deg);
}

.icon.wechat::before,
.icon.service::before {
  content: "";
  position: absolute;
  left: 2px;
  top: 4px;
  width: 17px;
  height: 12px;
  border: 2px solid currentColor;
  border-radius: 999px;
}

.icon.wechat::after,
.icon.service::after {
  content: "";
  position: absolute;
  left: 7px;
  bottom: 3px;
  width: 5px;
  height: 5px;
  border-left: 2px solid currentColor;
  border-bottom: 2px solid currentColor;
  transform: rotate(-25deg);
}

.icon.mail::before {
  content: "";
  position: absolute;
  left: 2px;
  top: 5px;
  width: 18px;
  height: 13px;
  border: 2px solid currentColor;
  border-radius: 2px;
}

.icon.mail::after {
  content: "";
  position: absolute;
  left: 5px;
  top: 8px;
  width: 12px;
  height: 8px;
  border-left: 2px solid currentColor;
  border-bottom: 2px solid currentColor;
  transform: rotate(-45deg);
}

.back-top::before {
  content: "";
  width: 10px;
  height: 10px;
  border-top: 2px solid var(--color-brand);
  border-left: 2px solid var(--color-brand);
  transform: translateY(3px) rotate(45deg);
}

@media (max-width: 720px) {
  .float-rail {
    display: none;
  }
}
</style>
