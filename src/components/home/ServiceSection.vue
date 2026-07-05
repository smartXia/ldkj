<script setup>
import SectionHeading from './SectionHeading.vue'

defineProps({
  services: {
    type: Array,
    default: () => [],
  },
})

const caseImage = '/assets/wsd/case-01.png'

const serviceCards = [
  {
    title: '新品/爆品打造',
    desc: '通过新品营销将产品打造成为电商热卖爆款',
    featured: true,
  },
  {
    title: '种草全域转化',
    desc: '通过社交内容种草，对全渠道生意产生积极影响',
  },
  {
    title: 'TOP1赛道抢占',
    desc: '通过精准的营销策略，成功登顶赛道的TOP1地位',
  },
  {
    title: '线索获取',
    desc: '通过精准识别并打动意向用户，获取生意线索',
  },
  {
    title: '直播电商',
    desc: '通过私域直播经营，实现站内电商转化',
  },
  {
    title: '品牌社媒运营',
    desc: '通过社交平台账号运营，提升品牌知名度和用户好感度',
  },
  {
    title: '社交营销培训',
    desc: '通过系统化的营销培训，解锁社交营销新玩法',
  },
  {
    title: '事件营销',
    desc: '通过打造或参与有影响力事件进行的营销活动',
  },
]
</script>

<template>
  <section class="service-section" aria-labelledby="service-title">
    <div class="phone-peek" aria-hidden="true">
      <span class="phone-speaker"></span>
      <span class="phone-line"></span>
    </div>

    <SectionHeading
      id="service-title"
      title="品牌定制服务"
      subtitle="以品牌目标为驱动的定制化营销解决方案"
    />

    <div class="service-viewport">
      <div class="service-rail" aria-label="品牌定制服务列表">
        <article class="case-card">
          <img :src="caseImage" alt="米家案例" loading="lazy" />
          <div class="case-logo">
            <span>米</span>
          </div>
          <strong>米家</strong>
        </article>

        <article
          v-for="(item, index) in serviceCards"
          :key="item.title"
          class="service-card"
          :class="{ featured: item.featured }"
        >
          <span class="service-index">{{ String(index + 1).padStart(2, '0') }}</span>
          <h3>{{ item.title }}</h3>
          <p>{{ item.desc }}</p>
          <a v-if="item.featured" href="#case" class="case-link">查看案例详情</a>
        </article>
      </div>
    </div>

    <div class="service-dots" aria-hidden="true">
      <span class="active"></span>
      <span v-for="dot in 7" :key="dot"></span>
    </div>

    <a class="service-more" href="#case">查看更多案例</a>
  </section>
</template>

<style scoped>
.service-section {
  position: relative;
  padding: 126px 0 88px;
  background: #fff;
  overflow: hidden;
}

.phone-peek {
  position: absolute;
  top: -96px;
  left: 50%;
  width: 256px;
  height: 256px;
  border: 9px solid #33363b;
  border-radius: 32px;
  background:
    linear-gradient(#fff, #fff) padding-box,
    linear-gradient(135deg, #f7f7f7, #e7e9ee) border-box;
  box-shadow: 0 32px 58px rgba(20, 23, 31, 0.16);
  transform: translateX(220px);
}

.phone-speaker {
  position: absolute;
  top: 16px;
  left: 50%;
  width: 58px;
  height: 16px;
  border-radius: 999px;
  background: #33363b;
  transform: translateX(-50%);
}

.phone-line {
  position: absolute;
  left: 50%;
  bottom: 13px;
  width: 104px;
  height: 4px;
  border-radius: 999px;
  background: #111;
  transform: translateX(-50%);
}

:deep(.section-heading) {
  margin-bottom: 76px;
}

:deep(.section-heading h2) {
  font-size: 42px;
  letter-spacing: 0;
}

:deep(.section-heading p) {
  margin-top: 18px;
  color: #767b86;
  font-size: 20px;
}

.service-viewport {
  width: 100%;
  overflow-x: auto;
  scrollbar-width: none;
}

.service-viewport::-webkit-scrollbar {
  display: none;
}

.service-rail {
  width: max-content;
  min-width: 100%;
  display: flex;
  align-items: stretch;
  gap: 16px;
  padding: 0 max(78px, calc((100vw - 1920px) / 2 + 78px));
}

.case-card,
.service-card {
  flex: 0 0 216px;
  height: 252px;
  border-radius: 24px;
  overflow: hidden;
}

.case-card {
  position: relative;
  color: #fff;
  background: #1f242b;
  box-shadow: 0 18px 34px rgba(17, 24, 39, 0.15);
}

.case-card::after {
  content: "";
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.28);
}

.case-card img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
}

.case-logo {
  position: absolute;
  left: 30px;
  bottom: 28px;
  z-index: 1;
  width: 52px;
  height: 52px;
  border-radius: 14px;
  display: grid;
  place-items: center;
  background: #55d579;
  box-shadow: 0 10px 20px rgba(12, 88, 39, 0.22);
}

.case-logo span {
  color: #fff;
  font-size: 27px;
  line-height: 1;
  font-weight: 900;
}

.case-card strong {
  position: absolute;
  left: 93px;
  bottom: 43px;
  z-index: 1;
  font-size: 18px;
  line-height: 1;
  font-weight: 800;
}

.service-card {
  position: relative;
  padding: 64px 14px 24px;
  background: #f2f3f6;
}

.service-card.featured {
  flex-basis: 202px;
  border-radius: 0 24px 24px 0;
  background: #ef4f4f;
  color: #fff;
}

.service-index {
  position: absolute;
  left: 16px;
  top: 34px;
  color: rgba(255, 255, 255, 0.48);
  font-size: 33px;
  line-height: 1;
  font-weight: 900;
}

.service-card:not(.featured) .service-index {
  color: #fff;
}

.service-card h3 {
  position: relative;
  z-index: 1;
  margin: 0;
  color: #05070b;
  font-size: 24px;
  line-height: 1.22;
  font-weight: 900;
  letter-spacing: 0;
}

.service-card.featured h3 {
  color: #fff;
}

.service-card p {
  position: relative;
  z-index: 1;
  margin: 20px 0 0;
  color: #737881;
  font-size: 16px;
  line-height: 1.55;
  font-weight: 500;
}

.service-card.featured p {
  color: #fff;
  font-weight: 700;
}

.case-link {
  position: absolute;
  left: 22px;
  bottom: 27px;
  min-width: 146px;
  height: 36px;
  border: 2px solid rgba(255, 255, 255, 0.9);
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 14px;
  font-weight: 800;
}

.case-link::after,
.service-more::after {
  content: ">>";
  margin-left: 4px;
  letter-spacing: -2px;
}

.service-dots {
  margin-top: 34px;
  display: flex;
  justify-content: center;
  gap: 14px;
}

.service-dots span {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: #f0f0f0;
}

.service-dots .active {
  background: #ef4f4f;
}

.service-more {
  width: max-content;
  margin: 46px auto 0;
  display: flex;
  align-items: center;
  color: #8a8d94;
  font-size: 18px;
  line-height: 1;
  font-weight: 600;
}

@media (max-width: 1180px) {
  .service-section {
    padding-top: 110px;
  }

  .service-rail {
    padding-inline: 48px;
  }
}

@media (max-width: 760px) {
  .service-section {
    padding: 76px 0 56px;
  }

  .phone-peek {
    top: -176px;
    width: 210px;
    height: 230px;
    transform: translateX(54px);
  }

  :deep(.section-heading) {
    margin-bottom: 34px;
  }

  :deep(.section-heading h2) {
    font-size: 28px;
  }

  :deep(.section-heading p) {
    margin-top: 10px;
    font-size: 14px;
  }

  .service-rail {
    gap: 10px;
    padding-inline: 18px;
  }

  .case-card,
  .service-card {
    flex-basis: 156px;
    height: 198px;
    border-radius: 18px;
  }

  .service-card.featured {
    flex-basis: 166px;
    border-radius: 0 18px 18px 0;
  }

  .service-card {
    padding: 52px 12px 18px;
  }

  .service-index {
    top: 24px;
    font-size: 26px;
  }

  .service-card h3 {
    font-size: 18px;
  }

  .service-card p {
    margin-top: 12px;
    font-size: 12px;
  }

  .case-link {
    left: 14px;
    bottom: 18px;
    min-width: 128px;
    height: 32px;
    font-size: 12px;
  }

  .case-logo {
    left: 18px;
    bottom: 22px;
    width: 42px;
    height: 42px;
  }

  .case-card strong {
    left: 68px;
    bottom: 35px;
    font-size: 15px;
  }

  .service-dots {
    margin-top: 28px;
    gap: 10px;
  }

  .service-more {
    margin-top: 34px;
    font-size: 15px;
  }
}
</style>
