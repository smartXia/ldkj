<script setup>
import { computed, shallowRef } from 'vue'

const tabs = ['行业奖项', '媒体奖项', '企业认证']
const activeTab = shallowRef(tabs[0])
const activeHonorIndex = shallowRef(-1)

const honorGroups = {
  行业奖项: [
    {
      title: '长城奖',
      image: '/assets/wsd/about/honor-01.webp',
      awards: ['内容营销类案例 铜奖', '跨界营销类案例 优秀奖', '整合营销类案例 优秀奖'],
    },
    {
      title: '金投赏',
      image: '/assets/wsd/about/honor-02.webp',
      awards: ['经典案例奖', '产品力营销 银奖', '产品力营销 提名奖', '长期品牌媒介管理 铜奖'],
    },
    {
      title: '虎啸奖',
      image: '/assets/wsd/about/honor-03.webp',
      awards: ['3C家电行业年度营销服务商', '家居日用类案例 金奖', '节日营销类案例 金奖', '家电类案例 银奖', '电子产品与设备类案例 银奖'],
    },
    {
      title: 'IAI传鉴国际广告奖',
      image: '/assets/wsd/about/honor-04.webp',
      awards: ['年度整合营销公司', '国际旅游案例奖', '新锐品牌营销 金奖'],
    },
    {
      title: 'ADMEN',
      image: '/assets/wsd/about/honor-05.webp',
      awards: ['年度最具商业价值数字营销公司', '实战金案奖'],
    },
    {
      title: '非凡奖',
      image: '/assets/wsd/about/honor-06.webp',
      awards: ['最佳快消行业 小红书营销奖', '最佳家电行业 微博营销奖', '优选案例奖'],
    },
  ],
  媒体奖项: [
    {
      title: '微博',
      image: '/assets/wsd/about/honor-01.webp',
      awards: ['年度商业传播案例', '年度内容营销案例', '年度创新营销案例'],
    },
    {
      title: '小红书',
      image: '/assets/wsd/about/honor-02.webp',
      awards: ['优质服务商', '年度影响力机构', '数字营销案例奖'],
    },
    {
      title: '知乎',
      image: '/assets/wsd/about/honor-03.webp',
      awards: ['品牌营销传播奖', '社交营销案例奖', '内容种草案例奖'],
    },
  ],
  企业认证: [
    {
      title: '先进基层党组织',
      image: '/assets/wsd/about/honor-04.webp',
      awards: ['小红书头部服务商', '腾讯广告核心服务商', '企业微信合作服务商'],
    },
    {
      title: '国家级高新技术企业',
      image: '/assets/wsd/about/honor-05.webp',
      awards: ['行业年度服务机构', '年度数字营销公司', '商业价值认证'],
    },
    {
      title: '中国广告协会',
      image: '/assets/wsd/about/honor-06.webp',
      awards: ['平台合作认证', '社交生态合作伙伴', '整合营销服务认证'],
    },
  ],
}

const visibleHonors = computed(() => honorGroups[activeTab.value])

function changeTab(tab) {
  activeTab.value = tab
  activeHonorIndex.value = -1
}
</script>

<template>
  <section id="about-honors" class="about-honors">
    <h2>企业荣誉</h2>
    <div class="honor-tabs" aria-label="荣誉分类">
      <button
        v-for="tab in tabs"
        :key="tab"
        type="button"
        :class="{ active: tab === activeTab }"
        @click="changeTab(tab)"
      >
        {{ tab }}
      </button>
    </div>

    <ul class="honor-list" @mouseleave="activeHonorIndex = -1">
      <li
        v-for="(item, index) in visibleHonors"
        :key="item.title"
        class="honor-item"
        :class="{ active: index === activeHonorIndex }"
        tabindex="0"
        @mouseenter="activeHonorIndex = index"
        @focus="activeHonorIndex = index"
        @click="activeHonorIndex = index"
      >
        <img :src="item.image" :alt="item.title" loading="lazy" />
        <div class="honor-copy">
          <h3>{{ item.title }}</h3>
          <p v-for="award in item.awards" :key="award">· {{ award }}</p>
        </div>
      </li>
    </ul>
  </section>
</template>

<style scoped>
.about-honors {
  min-height: 872px;
  padding: 0;
  text-align: center;
  background: #fff;
}

.about-honors h2 {
  margin: 0 0 50px;
  color: #000;
  font-size: 40px;
  line-height: 56px;
  font-weight: 700;
}

.honor-tabs {
  width: 800px;
  height: 52px;
  margin: 0 auto 42px;
  padding: 4px 8px;
  display: flex;
  overflow: hidden;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 0 10px rgba(31, 38, 56, 0.08);
}

.honor-tabs button {
  flex: 1;
  border: 0;
  border-radius: 6px;
  color: #1f2530;
  background: transparent;
  font-size: 18px;
  line-height: 44px;
  font-weight: 500;
  cursor: pointer;
  transition: color var(--motion-fast), background var(--motion-fast);
}

.honor-tabs button.active {
  color: #fff;
  background: #ff4848;
}

.honor-list {
  width: min(1200px, calc(100% - 96px));
  margin: 0 auto;
  padding: 0;
  list-style: none;
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 30px;
  text-align: left;
}

.honor-item {
  position: relative;
  height: 240px;
  border-radius: 14px;
  background: #fff;
  box-shadow: 0 12px 34px rgba(31, 38, 56, 0.06);
  display: grid;
  place-items: center;
  cursor: pointer;
  overflow: hidden;
  transition: background var(--motion), box-shadow var(--motion), transform var(--motion);
}

.honor-item img {
  position: relative;
  z-index: 1;
  width: 245px;
  max-height: 132px;
  display: block;
  object-fit: contain;
  transition: opacity var(--motion), transform var(--motion);
}

.honor-copy {
  position: absolute;
  inset: 0;
  z-index: 2;
  padding: 32px 30px;
  color: #fff;
  opacity: 0;
  transform: translateY(12px);
  transition: opacity var(--motion), transform var(--motion);
}

.honor-copy h3 {
  margin: 0 0 18px;
  color: #fff;
  font-size: 32px;
  line-height: 42px;
  font-weight: 700;
}

.honor-copy p {
  margin: 0;
  color: #fff;
  font-size: 18px;
  line-height: 34px;
  font-weight: 500;
}

.honor-item:hover,
.honor-item.active {
  background: linear-gradient(149deg, #ff9c7a 0%, #fb5151 78%, #fb4848 100%);
  box-shadow: 0 16px 38px rgba(251, 72, 72, 0.18);
  transform: translateY(-1px);
}

.honor-item:hover img,
.honor-item.active img {
  opacity: 0;
  transform: scale(0.96);
}

.honor-item:hover .honor-copy,
.honor-item.active .honor-copy {
  opacity: 1;
  transform: translateY(0);
}

.honor-item:focus-visible {
  outline: 2px solid #ff4848;
  outline-offset: 4px;
}

@media (max-width: 1120px) {
  .honor-list {
    width: calc(100% - 40px);
  }

  .honor-copy {
    padding: 28px 24px;
  }

  .honor-copy h3 {
    font-size: 28px;
  }

  .honor-copy p {
    font-size: 16px;
    line-height: 30px;
  }
}

@media (max-width: 820px) {
  .about-honors {
    min-height: 0;
    padding: 33px 20px 34px;
  }

  .about-honors h2 {
    margin-bottom: 20px;
    font-size: 24px;
    line-height: 29px;
  }

  .honor-tabs {
    width: 100%;
    height: 40px;
    margin-bottom: 20px;
    padding: 4px;
  }

  .honor-tabs button {
    border-radius: 5px;
    font-size: 14px;
    line-height: 32px;
  }

  .honor-list {
    width: 100%;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 12px;
  }

  .honor-item {
    height: 107px;
    border-radius: 0;
    box-shadow: none;
  }

  .honor-item img {
    width: 132px;
    max-height: 64px;
  }

  .honor-copy {
    padding: 12px;
  }

  .honor-copy h3 {
    margin-bottom: 6px;
    font-size: 16px;
    line-height: 20px;
  }

  .honor-copy p {
    font-size: 12px;
    line-height: 16px;
  }
}
</style>
