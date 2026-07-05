<script setup>
const highlights = [
  { value: '2014', label: '成立于南京' },
  { value: '500+', label: '营销与技术人才' },
  { value: 'ACT', label: '自研营销模型' },
]

const strengths = [
  '社交广告、达人内容、口碑管理、社交电商与营销技术一体化服务',
  '深度连接微信、小红书、B站、Instagram、LinkedIn、Reddit 等社交平台',
  '长期服务宝洁、小米、名创优品、安克创新、小熊电器等知名品牌',
]
</script>

<template>
  <section class="home-about" aria-labelledby="home-about-title">
    <div class="home-about-inner">
      <div class="about-copy">
        <span class="about-kicker">ABOUT WSD</span>
        <h2 id="home-about-title">关于我们</h2>
        <p>
          灵动集团创立于 2014 年，以社交为桥梁、以科技为引擎，持续为品牌提供系统化、可衡量的社交营销解决方案。
        </p>
        <p>
          我们聚合广告、内容与交易三维能力，帮助品牌在复杂的社交生态里建立用户关系、沉淀品牌资产，并实现可持续增长。
        </p>

        <RouterLink class="about-link" to="/about">了解关于我们</RouterLink>
      </div>

      <div class="about-panel">
        <img class="about-image" src="/assets/wsd/about/company-bg-pc.png" alt="" loading="lazy" />
        <ul class="about-stats" aria-label="灵动集团概览">
          <li v-for="item in highlights" :key="item.label">
            <strong>{{ item.value }}</strong>
            <span>{{ item.label }}</span>
          </li>
        </ul>
      </div>

      <ul class="about-strengths" aria-label="关于我们亮点">
        <li v-for="item in strengths" :key="item">
          <span aria-hidden="true"></span>
          <p>{{ item }}</p>
        </li>
      </ul>
    </div>
  </section>
</template>

<style scoped>
.home-about {
  position: relative;
  padding: 92px 0 96px;
  background: #f5f6f8;
  overflow: hidden;
}

.home-about::before {
  content: "";
  position: absolute;
  inset: 0 0 auto;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(17, 17, 17, 0.12), transparent);
}

.home-about-inner {
  position: relative;
  width: min(1188px, calc(100% - 48px));
  margin: 0 auto;
  display: grid;
  grid-template-columns: minmax(0, 500px) minmax(420px, 1fr);
  grid-template-areas:
    "copy panel"
    "strengths panel";
  gap: 36px 72px;
  align-items: center;
}

.about-copy {
  grid-area: copy;
}

.about-kicker {
  display: inline-flex;
  align-items: center;
  min-height: 22px;
  color: var(--color-brand);
  font-size: 13px;
  line-height: 22px;
  font-weight: 900;
  letter-spacing: 0;
}

.about-copy h2 {
  margin: 12px 0 24px;
  color: #08090b;
  font-size: 42px;
  line-height: 1.18;
  font-weight: 900;
  letter-spacing: 0;
}

.about-copy p {
  margin: 0 0 14px;
  color: #4a505a;
  font-size: 17px;
  line-height: 30px;
}

.about-link {
  width: fit-content;
  min-width: 158px;
  min-height: 44px;
  margin-top: 22px;
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 24px;
  background: var(--color-brand);
  color: #fff;
  font-size: 16px;
  font-weight: 800;
  box-shadow: var(--shadow-soft);
  transition: transform var(--motion), box-shadow var(--motion);
}

.about-link:hover {
  transform: translateY(-2px);
  box-shadow: 0 14px 32px rgba(255, 61, 69, 0.24);
}

.about-panel {
  position: relative;
  grid-area: panel;
  min-height: 430px;
  border-radius: 8px;
  overflow: hidden;
  background: #151922;
  box-shadow: 0 22px 46px rgba(20, 27, 41, 0.16);
}

.about-panel::after {
  content: "";
  position: absolute;
  inset: 0;
  background:
    linear-gradient(90deg, rgba(13, 16, 23, 0.56), rgba(13, 16, 23, 0.08) 55%, rgba(13, 16, 23, 0.3)),
    linear-gradient(0deg, rgba(13, 16, 23, 0.54), transparent 46%);
}

.about-image {
  width: 100%;
  height: 430px;
  display: block;
  object-fit: cover;
}

.about-stats {
  position: absolute;
  left: 28px;
  right: 28px;
  bottom: 28px;
  z-index: 1;
  margin: 0;
  padding: 0;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  list-style: none;
}

.about-stats li {
  min-height: 92px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-radius: 8px;
  padding: 16px 18px;
  background: rgba(255, 255, 255, 0.11);
  color: #fff;
  backdrop-filter: blur(12px);
}

.about-stats strong {
  display: block;
  font-size: 32px;
  line-height: 36px;
  font-weight: 900;
}

.about-stats span {
  display: block;
  margin-top: 6px;
  color: rgba(255, 255, 255, 0.78);
  font-size: 13px;
  line-height: 18px;
}

.about-strengths {
  grid-area: strengths;
  margin: 0;
  padding: 0;
  display: grid;
  gap: 12px;
  list-style: none;
}

.about-strengths li {
  min-height: 58px;
  border-radius: 8px;
  display: grid;
  grid-template-columns: 26px 1fr;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background: #fff;
  box-shadow: 0 8px 22px rgba(23, 32, 50, 0.06);
}

.about-strengths span {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  background:
    radial-gradient(circle at 50% 50%, #fff 0 3px, transparent 4px),
    var(--color-brand);
}

.about-strengths p {
  margin: 0;
  color: #343a44;
  font-size: 15px;
  line-height: 24px;
}

@media (max-width: 980px) {
  .home-about-inner {
    grid-template-columns: 1fr;
    grid-template-areas:
      "copy"
      "panel"
      "strengths";
    gap: 28px;
  }
}

@media (max-width: 640px) {
  .home-about {
    padding: 38px 0 42px;
  }

  .home-about-inner {
    width: calc(100% - 28px);
    gap: 18px;
  }

  .about-kicker {
    font-size: 11px;
    line-height: 18px;
  }

  .about-copy h2 {
    margin: 8px 0 12px;
    font-size: 24px;
    line-height: 32px;
  }

  .about-copy p {
    font-size: 13px;
    line-height: 24px;
  }

  .about-link {
    min-width: 120px;
    min-height: 34px;
    margin-top: 10px;
    padding: 0 18px;
    font-size: 13px;
  }

  .about-panel {
    min-height: 286px;
  }

  .about-image {
    height: 286px;
  }

  .about-stats {
    left: 12px;
    right: 12px;
    bottom: 12px;
    gap: 7px;
  }

  .about-stats li {
    min-height: 66px;
    padding: 10px 8px;
  }

  .about-stats strong {
    font-size: 20px;
    line-height: 24px;
  }

  .about-stats span {
    margin-top: 4px;
    font-size: 10px;
    line-height: 14px;
  }

  .about-strengths li {
    min-height: 52px;
    grid-template-columns: 20px 1fr;
    gap: 10px;
    padding: 10px 12px;
  }

  .about-strengths span {
    width: 20px;
    height: 20px;
  }

  .about-strengths p {
    font-size: 12px;
    line-height: 20px;
  }
}
</style>
