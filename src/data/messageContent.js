const local = (name) => `/assets/wsd/message/${name}`

export const messageHero = {
  title: '关注行业动态，及时获取前沿资讯',
  image: local('banner.webp'),
}

export const messageCategories = ['全部资讯', '企业动态', '营销观察', '案例故事']

export const messageArticles = [
  {
    id: 'skhd-iai',
    category: '企业动态',
    title: '荣誉加冕！拾刻互动携手合生元、安琪酵母斩获第26届IAI双奖！',
    desc: '拾刻互动斩获26届IAI品牌营销金奖、种草营销优秀奖！',
    date: '2026-06-23 17:28:36',
    image: local('news-01.webp'),
  },
  {
    id: 'wsdcyxy',
    category: '企业动态',
    title: '携手共建产业学院，微思敦集团与武汉东湖学院正式签约！',
    desc: '倾力打造校企协同育人标杆示范项目！',
    date: '2026-06-16 14:20:25',
    image: local('news-02.webp'),
  },
  {
    id: 'reddit-marketing-for-chinese-overseas-brands',
    category: '营销观察',
    title: '中国出海品牌如何做好Reddit营销【2026最新版指南】',
    desc: '面向中国出海品牌的 Reddit 营销实战指南。从品牌为何翻车讲起，拆解广告价值（27% 购买意愿提升、46% 信任度提升）、本地化与素人种草两大策略，以及 Reddit 在 SEO 与 AI 搜索中',
    date: '2026-06-23 13:58:35',
    image: local('news-03.webp'),
  },
  {
    id: 'LinkedIn',
    category: '营销观察',
    title: 'B2B出海企业领英营销第一步：Linkedin与微信有什么不同？',
    desc: '把国内做微信那套思维，硬套在LinkedIn上。搞明白这两个平台的结构性差异，任何一个跨境B2B内容策略才算真正迈出了第一步。',
    date: '2026-06-23 09:44:38',
    image: local('news-04.webp'),
  },
  {
    id: 'hys-tvgg-nzzr',
    category: '营销观察',
    title: '腾讯广告联动海月昇发布「2026房企年中助燃计划」',
    desc: '年中助燃窗口期转瞬即逝，抢占先机，刻不容缓',
    date: '2026-06-09 15:51:59',
    image: local('news-05.webp'),
  },
  {
    id: 'rednote',
    category: '企业动态',
    title: '微思敦再度认证「小红书优质合作伙伴」，多赛道实力领跑！',
    desc: '微思敦再次斩获产品种草类目「美妆个护」「3C及电器」「日用百货」赛道、客资收集类目「生活服务」赛道，以及「种草直达」类目优质合作伙伴认证。',
    date: '2026-06-18 17:10:28',
    image: local('news-06.webp'),
  },
  {
    id: 'jxhd-sjb',
    category: '案例故事',
    title: '世界杯来了，看品牌率先上场大显神通',
    desc: '从绿茵场上的激烈对抗到屏幕前的情绪共振，品牌营销的战场早已从单纯的曝光转向更深层的情感链接。',
    date: '2026-06-12 08:44:11',
    image: local('news-07.webp'),
  },
  {
    id: 'Revive',
    category: '案例故事',
    title: '闯美成功！WSD让小众护肤品牌出现在北美华人的化妆台',
    desc: '一个美国小众护肤品牌，以2500元的高客单价进入美国市场，要想在短短两个月内实现声量与销量的双重突围，该如何做？',
    date: '2026-06-17 11:46:07',
    image: local('news-08.webp'),
  },
  {
    id: 'qp-smqxbps',
    category: '营销观察',
    title: '三层情绪源头、六大进化方向：一文读懂数码品牌增长新逻辑',
    desc: '点击领取《2026潮数码情绪白皮书》',
    date: '2026-06-10 12:34:32',
    image: local('news-09.webp'),
  },
  {
    id: 'Spokesperson',
    category: '营销观察',
    title: '不止于流量，代言人与品牌如何互相成就？',
    desc: '越来越多的品牌意识到，与其盲目追求“流量”，不如与一位“对的人”互相成就、合作共赢。',
    date: '2026-06-09 15:08:25',
    image: local('news-10.webp'),
  },
  {
    id: 'wsd-sjyxxly-3',
    category: '企业动态',
    title: '公益续力，逐梦前行｜微思敦社交营销训练营三期圆满结业',
    desc: '从观赛到实战，以热爱赴新程',
    date: '2026-06-08 16:30:57',
    image: local('news-11.webp'),
  },
  {
    id: 'cjrb-hxj',
    category: '企业动态',
    title: '《长江日报》报道：武昌滨江这家企业一举斩获16项行业大奖',
    desc: '此次获奖彰显了武汉本土数字经济企业在全国市场的强劲竞争力',
    date: '2026-06-05 18:02:13',
    image: local('news-12.webp'),
  },
].map((article) => ({
  ...article,
  author: '好品牌 会社交',
  href: `https://www.wsdsocial.com/zh-cn/article/${article.id}`,
}))
