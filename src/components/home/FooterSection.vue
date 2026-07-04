<script setup>
import { computed } from 'vue'

const props = defineProps({
  site: {
    type: Object,
    default: () => ({}),
  },
})

const footerLinks = computed(() => {
  const raw = props.site.footer_links
  if (Array.isArray(raw)) return raw
  if (!raw) return []
  try {
    const parsed = JSON.parse(raw)
    return Array.isArray(parsed) ? parsed : []
  } catch {
    return []
  }
})

const hasFooter = computed(() =>
  props.site.contact || props.site.copyright || props.site.icp_no || footerLinks.value.length
)
</script>

<template>
  <footer v-if="hasFooter" class="site-footer">
    <div class="footer-inner">
      <section v-if="site.contact" class="footer-block">
        <h2>联系我们</h2>
        <p>{{ site.contact }}</p>
      </section>

      <nav v-if="footerLinks.length" class="footer-links" aria-label="页脚链接">
        <a
          v-for="link in footerLinks"
          :key="link.label || link.title || link.href"
          :href="link.href || link.url || 'javascript:void(0)'"
        >
          {{ link.label || link.title || link.name }}
        </a>
      </nav>

      <p v-if="site.icp_no" class="footer-icp">{{ site.icp_no }}</p>
      <p v-if="site.copyright" class="footer-copy">{{ site.copyright }}</p>
    </div>
  </footer>
</template>

<style scoped>
.site-footer {
  padding: 44px 0 28px;
  background: #242832;
  color: rgba(255, 255, 255, 0.72);
}

.footer-inner {
  width: min(960px, calc(100% - 48px));
  margin: 0 auto;
  display: grid;
  gap: 18px;
  text-align: center;
}

.footer-block h2 {
  margin: 0 0 10px;
  color: #fff;
  font-size: 18px;
}

.footer-block p,
.footer-icp,
.footer-copy {
  margin: 0;
  font-size: 14px;
  line-height: 1.7;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 14px;
}

.footer-links a {
  color: rgba(255, 255, 255, 0.72);
  font-size: 14px;
}

@media (max-width: 820px) {
  .site-footer {
    padding: 28px 0 20px;
  }

  .footer-inner {
    width: calc(100% - 24px);
  }

  .footer-block p,
  .footer-icp,
  .footer-copy,
  .footer-links a {
    font-size: 12px;
  }
}
</style>
