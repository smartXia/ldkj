const local = (name) => `/assets/wsd/${name}`

export const navItems = [
  { label: '首页', path: '/' },
  { label: '客户案例', path: '/case' },
  { label: '营销资讯', path: '/message' },
  { label: '关于我们', path: '/about' },
  { label: '加入我们', href: '#jobs' },
]

export const heroSlides = [
  { title: '南京灵动信息技术有限公司斩获第十七届虎啸奖16项大奖！', image: local('hero-01.jpg') },
  { title: '南京灵动信息技术有限公司企业开放日', image: local('hero-03.png') },
  { title: '第三届金敦奖', image: local('hero-02.jpg') },
  { title: '金投赏', image: local('hero-04.jpg') },
  { title: '小红书种草大赏', image: local('hero-05.png') }
]

export const marketingModules = [
  { title: '社交生态占位', image: local('marketing-01.png') },
  { title: '用户信任构建', image: local('marketing-02.png') },
  { title: '品牌资产沉淀', image: local('marketing-03.png') }
]

export const services = [
  { title: '社交媒体深度商业合作伙伴', image: local('service-01.png') },
  { title: '达人内容服务的策略及流程', image: local('service-02.png') },
  { title: '营销培训服务提供的服务项目及平台', image: local('service-03.png') },
  { title: '社交电商服务提供的服务项目', image: local('service-04.png') },
  { title: '口碑管理服务提供的服务项目', image: local('service-05.png') },
  { title: '营销技术服务提供的服务项目', image: local('service-06.png') }
]

export const brandCases = [
  ['01', '新品/爆品打造', '通过新品营销将产品打造成电商热卖爆款', local('case-01.png')],
  ['02', '种草全域转化', '通过社交内容种草，对全渠道生意产生积极影响', local('case-02.png')],
  ['03', 'TOP1赛道抢占', '通过精准的营销策略，成功登顶赛道的TOP1地位', local('case-03.png')],
  ['04', '线索获取', '通过精准识别并打动意向用户，获取生意线索', local('case-04.png')],
  ['05', '直播电商', '通过私域直播经营，实现站内电商转化', local('case-05.png')],
  ['06', '品牌社媒运营', '通过社交平台账号运营，提升品牌知名度和用户好感度', local('case-06.png')],
  ['07', '社交营销培训', '通过系统化的营销培训，解锁社交营销新玩法', local('case-07.png')],
  ['08', '事件营销', '通过打造或参与有影响力事件进行的营销活动', local('case-08.png')]
].map(([number, title, summary, image]) => ({ number, title, summary, image }))

export const logoList = [
  ['追觅', 'logo-01.png'],
  ['学而思', 'logo-02.jpeg'],
  ['武汉城投', 'logo-03.jpg'],
  ['潘婷', 'logo-04.jpg'],
  ['林氏家居', 'logo-05.png'],
  ['可优比', 'logo-06.png'],
  ['华帝', 'logo-07.jpg'],
  ['海信', 'logo-08.png'],
  ['高露洁', 'logo-09.png'],
  ['新东方', 'logo-10.png'],
  ['方太', 'logo-11.png'],
  ['vivo', 'logo-12.png']
].map(([name, file]) => ({ name, image: local(file) }))
