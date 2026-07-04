<script setup>
import { computed, shallowRef } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from '../composables/useI18n'

const emit = defineEmits(['open-consult'])
defineProps({
  site: {
    type: Object,
    default: () => ({}),
  },
})
const route = useRoute()
const menuOpen = shallowRef(false)
const languageOpen = shallowRef(false)
const { languages, currentLanguage, messages, locale, setLanguage } = useI18n()

const routeNavItems = computed(() => messages.value.navItems.filter((item) => item.path))
const hashNavItems = computed(() => messages.value.navItems.filter((item) => item.href))

function isActive(item) {
  return item.path === route.path
}

function closeMenu() {
  menuOpen.value = false
  languageOpen.value = false
}

function openConsult() {
  closeMenu()
  emit('open-consult')
}

function toggleLanguage() {
  languageOpen.value = !languageOpen.value
}

function selectLanguage(code) {
  setLanguage(code)
  languageOpen.value = false
}

function closeLanguageOnFocusout(event) {
  if (!event.currentTarget.contains(event.relatedTarget)) {
    languageOpen.value = false
  }
}
</script>

<template>
  <header class="site-header">
    <div class="header-inner">
      <RouterLink class="brand" to="/" aria-label="微思敦首页" @click="closeMenu">
        <img v-if="site.logo" :src="site.logo" :alt="site.site_title || '网站 Logo'" />
        <span v-else>{{ site.site_title || '南京灵动' }}</span>
      </RouterLink>

      <button
        class="menu-toggle"
        type="button"
        aria-label="打开导航"
        :aria-expanded="menuOpen"
        @click="menuOpen = !menuOpen"
      >
        <span></span><span></span><span></span>
      </button>

      <nav class="nav" :class="{ open: menuOpen }" aria-label="主导航">
        <RouterLink
          v-for="item in routeNavItems"
          :key="item.label"
          :to="item.path"
          :class="{ active: isActive(item) }"
          @click="closeMenu"
        >
          {{ item.label }}
        </RouterLink>
        <a
          v-for="item in hashNavItems"
          :key="item.label"
          :href="item.href"
          @click="closeMenu"
        >
          {{ item.label }}
        </a>
      </nav>

      <div class="header-actions">
        <button class="btn primary" type="button" @click="openConsult">{{ messages.header.consult }}</button>
        <a class="btn ghost" href="#login">{{ messages.header.login }}</a>
        <div class="language-switch" @focusout="closeLanguageOnFocusout" @keydown.esc="languageOpen = false">
          <button
            class="language"
            :class="{ open: languageOpen }"
            type="button"
            aria-haspopup="listbox"
            :aria-expanded="languageOpen"
            @click="toggleLanguage"
          >
            {{ currentLanguage.label }}
            <span class="language-arrow" aria-hidden="true"></span>
          </button>
          <div v-show="languageOpen" class="language-menu" role="listbox">
            <button
              v-for="item in languages"
              :key="item.code"
              type="button"
              role="option"
              :aria-selected="item.code === locale"
              :class="{ active: item.code === locale }"
              @click="selectLanguage(item.code)"
            >
              {{ item.label }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<style scoped>
.site-header {
  position: fixed;
  inset: 0 0 auto;
  z-index: 50;
  height: var(--header-height);
  background: var(--color-surface);
  box-shadow: none;
}

.header-inner {
  max-width: none;
  height: 100%;
  margin: 0 auto;
  padding: 0 40px;
  display: flex;
  align-items: center;
  gap: 44px;
}

.brand img {
  width: 124px;
  height: 40px;
  object-fit: contain;
  display: block;
}

.brand span {
  color: var(--color-ink);
  font-size: 18px;
  font-weight: 800;
  white-space: nowrap;
}

.nav {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 42px;
  height: 100%;
}

.nav a {
  position: relative;
  height: 100%;
  display: inline-flex;
  align-items: center;
  color: var(--color-ink);
  font-size: 16px;
  font-weight: 400;
  transition: color var(--motion-fast);
}

.nav a.active,
.nav a:hover {
  color: var(--color-brand);
}

.nav a.active::after {
  content: "";
  position: absolute;
  left: 50%;
  bottom: 0;
  width: 58px;
  height: 2px;
  border-radius: 0;
  background: var(--color-brand);
  transform: translateX(-50%);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.btn,
.language {
  min-height: 44px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  font-size: 16px;
  white-space: nowrap;
  transition: transform var(--motion-fast), box-shadow var(--motion-fast), border-color var(--motion-fast);
}

.btn {
  padding: 0 20px;
  border: 1px solid transparent;
}

.btn.primary {
  color: #fff;
  background: var(--color-brand);
  box-shadow: var(--shadow-soft);
  font-weight: 700;
}

.btn.ghost {
  color: var(--color-brand);
  border-color: rgba(255, 61, 69, 0.28);
  background: #fff;
}

.language-switch {
  position: relative;
}

.language {
  gap: 18px;
  border: 1px solid var(--color-border);
  background: #fff;
  color: var(--color-ink);
  min-width: 118px;
  padding: 0 20px;
}

.language-arrow {
  width: 9px;
  height: 9px;
  border-right: 2px solid #222;
  border-bottom: 2px solid #222;
  transform: rotate(45deg) translateY(-2px);
  transition: transform var(--motion-fast);
}

.language.open .language-arrow {
  transform: rotate(225deg) translate(-2px, -1px);
}

.language-menu {
  position: absolute;
  top: calc(100% + 14px);
  right: 0;
  z-index: 60;
  width: 118px;
  padding: 8px 0;
  border-radius: 3px;
  background: #fff;
  box-shadow: 0 12px 28px rgba(31, 38, 56, 0.1);
  overflow: hidden;
}

.language-menu button {
  width: 100%;
  height: 44px;
  padding: 0 24px;
  border: 0;
  background: #fff;
  color: var(--color-ink);
  text-align: left;
  font-size: 16px;
  cursor: pointer;
  transition: color var(--motion-fast), background var(--motion-fast);
}

.language-menu button:hover,
.language-menu button.active {
  color: var(--color-brand);
  background: #fff6f6;
}

.btn:hover,
.language:hover {
  transform: translateY(-1px);
}

.menu-toggle {
  display: none;
  width: 44px;
  height: 44px;
  border: 0;
  background: transparent;
}

.menu-toggle span {
  display: block;
  width: 22px;
  height: 2px;
  margin: 5px auto;
  background: var(--color-ink);
}

@media (max-width: 920px) {
  .site-header {
    height: var(--header-height);
    left: 0;
    right: 0;
    width: 100%;
  }
  .header-inner {
    justify-content: flex-start;
    padding: 0 13px;
    gap: 0;
  }
  .brand {
    order: 1;
  }
  .menu-toggle {
    order: 3;
    display: block;
    width: 24px;
    height: 24px;
    margin-left: 10px;
  }
  .menu-toggle span {
    width: 14px;
    height: 1px;
    margin: 4px auto;
  }
  .nav {
    position: absolute;
    left: 0;
    right: 0;
    top: var(--header-height);
    height: auto;
    padding: 18px 24px 24px;
    flex-direction: column;
    align-items: flex-start;
    gap: 18px;
    background: #fff;
    border-top: 1px solid var(--color-border);
    transform: translateY(-140%);
    opacity: 0;
    pointer-events: none;
    transition: transform var(--motion), opacity var(--motion);
  }
  .nav.open { transform: translateY(0); opacity: 1; pointer-events: auto; }
  .nav a { height: auto; }
  .nav a.active::after { bottom: -8px; left: 0; transform: none; }
  .brand img {
    width: 72px;
    height: 23px;
  }
  .header-actions {
    order: 2;
    margin-left: auto;
    gap: 0;
  }
  .btn.primary {
    min-height: 24px;
    padding: 0;
    color: var(--color-brand);
    background: transparent;
    box-shadow: none;
    font-size: 12px;
    font-weight: 700;
  }
  .header-actions .ghost,
  .language-switch { display: none; }
}
</style>
