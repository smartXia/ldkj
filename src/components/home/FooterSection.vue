<script setup>
const partners = ['腾讯广告', '微博广告', '小红书商业化', 'B站营销', '知乎营销', '活动行', '广告门', '数英网']
const addresses = [
  '深圳：深圳市南山区科兴科学园B1栋703室',
  '武汉：武汉市武昌区徐东大街8号匠心城17层',
  '上海：上海市徐汇区龙启路158号灿星大厦713室',
  '北京：北京市朝阳区百子湾路33号院6号楼万利中心A101'
]
const socials = ['微信', '微博', '小红书', 'bilibili', '知']
const qrCodes = []
</script>

<template>
  <footer class="site-footer">
    <div class="footer-inner">
      <div class="footer-left">
        <section class="footer-block">
          <h2>合作伙伴</h2>
          <div class="partner-links">
            <a v-for="partner in partners" :key="partner" href="javascript:void(0)">{{ partner }}</a>
          </div>
        </section>

        <section class="footer-block">
          <h2>办公地址</h2>
          <p v-for="address in addresses" :key="address">{{ address }}</p>
        </section>
      </div>

      <section class="footer-block footer-contact">
        <h2>联系我们</h2>
        <p>@微思敦</p>
        <div class="social-icons" aria-label="社交媒体">
          <span v-for="item in socials" :key="item">{{ item }}</span>
        </div>
        <p>400-003-5180</p>
        <p>service@wsdsocial.com</p>
        <div class="qr-row">
          <figure v-for="[label, src] in qrCodes" :key="label">
            <img :src="src" :alt="label" loading="lazy" />
            <figcaption>{{ label }}</figcaption>
          </figure>
        </div>
      </section>
    </div>

    <div class="footer-bottom">
      <nav aria-label="页脚导航">
        <a href="#about">关于我们</a>
        <a href="#jobs">加入我们</a>
        <a href="#client">客户中心</a>
      </nav>
      <p>鄂ICP备15001665号 Copyright © 2026 WSD. All Rights Reserved. 武汉微思敦网络技术有限公司 版权所有</p>
    </div>
  </footer>
</template>

<style scoped>
.site-footer {
  padding: 72px 0 34px;
  background: #3d424a;
  color: rgba(255, 255, 255, 0.72);
}

.footer-inner {
  width: min(960px, calc(100% - 48px));
  margin: 0 auto;
  display: grid;
  grid-template-columns: 560px 300px;
  justify-content: space-between;
  gap: 70px;
}

.footer-left {
  display: grid;
  gap: 58px;
}

.footer-block h2 {
  position: relative;
  margin: 0 0 22px;
  padding-left: 16px;
  color: #fff;
  font-size: 20px;
  line-height: 1.4;
  font-weight: 900;
}

.footer-block h2::before {
  content: "";
  position: absolute;
  left: 0;
  top: 4px;
  width: 4px;
  height: 22px;
  border-radius: 999px;
  background: var(--color-brand);
}

.partner-links {
  display: flex;
  flex-wrap: wrap;
}

.partner-links a {
  color: rgba(255, 255, 255, 0.72);
  font-size: 14px;
}

.partner-links a + a::before {
  content: "|";
  margin: 0 8px;
  color: rgba(255, 255, 255, 0.38);
}

.footer-block p {
  margin: 0 0 13px;
  color: rgba(255, 255, 255, 0.72);
  font-size: 14px;
  line-height: 1.7;
}

.social-icons {
  display: flex;
  gap: 10px;
  margin: 10px 0 18px;
}

.social-icons span {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: grid;
  place-items: center;
  background: rgba(255, 255, 255, 0.78);
  color: #3d424a;
  font-size: 10px;
  font-weight: 900;
}

.qr-row {
  display: flex;
  gap: 18px;
  margin-top: 18px;
}

.qr-row figure {
  margin: 0;
}

.qr-row img {
  width: 96px;
  height: 96px;
  display: block;
  background: #fff;
}

.qr-row figcaption {
  margin-top: 8px;
  color: rgba(255, 255, 255, 0.72);
  text-align: center;
  font-size: 13px;
}

.footer-bottom {
  width: min(960px, calc(100% - 48px));
  margin: 64px auto 0;
  text-align: center;
}

.footer-bottom nav {
  display: flex;
  justify-content: center;
  gap: 16px;
  margin-bottom: 18px;
}

.footer-bottom a,
.footer-bottom p {
  color: rgba(255, 255, 255, 0.66);
  font-size: 13px;
}

@media (max-width: 820px) {
  .site-footer {
    padding: 22px 0 18px;
  }

  .footer-inner {
    width: calc(100% - 24px);
    grid-template-columns: 1fr;
    gap: 22px;
  }

  .footer-left {
    gap: 22px;
  }

  .footer-contact {
    order: -1;
  }

  .footer-block h2 {
    margin-bottom: 12px;
    font-size: 14px;
  }

  .footer-block h2::before {
    top: 3px;
    width: 3px;
    height: 16px;
  }

  .footer-block p,
  .partner-links a {
    margin-bottom: 5px;
    font-size: 11px;
    line-height: 1.5;
  }

  .social-icons span {
    width: 22px;
    height: 22px;
    font-size: 8px;
  }

  .qr-row {
    gap: 12px;
  }

  .qr-row img {
    width: 70px;
    height: 70px;
  }

  .qr-row figcaption,
  .footer-bottom a,
  .footer-bottom p {
    font-size: 10px;
  }

  .footer-bottom {
    width: calc(100% - 24px);
    margin-top: 22px;
  }
}
</style>
