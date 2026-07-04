const caseImage = (name) => `/assets/wsd/case/${name}`
const detailImage = (name) => `/assets/wsd/case/${name}`

export const caseHero = {
  title: '客户案例',
  image: caseImage('banner.jpg'),
}

export const allOption = '全部'

export const caseFilterGroups = [
  {
    key: 'industry',
    label: '行业',
    options: [allOption, '3C家电', '美妆日化', '母婴宠物', '食品饮料', '房产家居', '网服', '教育', '其他'],
  },
  {
    key: 'method',
    label: '营销方式',
    options: [
      allOption,
      '新品/爆品打造',
      '种草全域转化',
      'TOP1赛道抢占',
      '线索获取',
      '直播电商',
      '品牌社媒运营',
      '社交营销培训',
      '事件营销',
    ],
  },
  {
    key: 'platform',
    label: '平台',
    options: [allOption, '腾讯', '微博', '小红书', 'B站', '知乎', '多媒体投放'],
  },
]

export const caseCards = [
  {
    id: 'mageline',
    title: '麦吉丽借势腾讯大剧 IP《庆余年 2》，引爆新品声量',
    image: caseImage('case-01.jpg'),
    detailImage: detailImage('detail-01.jpg'),
    industry: '美妆日化',
    method: '新品/爆品打造',
    platform: '腾讯',
  },
  {
    id: 'miniso',
    title: '名创优品洞察圈层情绪，打造社交热议爆款',
    image: caseImage('case-02.jpg'),
    detailImage: detailImage('detail-02.jpg'),
    industry: '其他',
    method: '事件营销',
    platform: '小红书',
  },
  {
    id: 'hurricane',
    title: '追觅小红书种草破圈，撬动高潜消费场景',
    image: caseImage('case-03.jpg'),
    detailImage: detailImage('detail-03.jpg'),
    industry: '3C家电',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'anker',
    title: '安克创新进军中国市场，打造闪充品类心智',
    image: caseImage('case-04.jpg'),
    detailImage: detailImage('detail-04.jpg'),
    industry: '3C家电',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'honeywell',
    title: '霍尼韦尔以内容制胜，建立宠物空气净化标杆',
    image: caseImage('case-05.jpg'),
    detailImage: detailImage('detail-05.jpg'),
    industry: '3C家电',
    method: '种草全域转化',
    platform: 'B站',
  },
  {
    id: 'oct',
    title: '武汉华侨城创新地产行业社交营销打法',
    image: caseImage('case-06.jpg'),
    detailImage: detailImage('detail-06.jpg'),
    industry: '房产家居',
    method: '线索获取',
    platform: '多媒体投放',
  },
  {
    id: 'yeehoo',
    title: '英氏 6 个月拉新人群，资产规模实现高速增长',
    image: caseImage('case-07.jpg'),
    detailImage: detailImage('detail-07.jpg'),
    industry: '母婴宠物',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'jmgo',
    title: '坚果投影用真实场景和小创意撬动大生意',
    image: caseImage('case-08.jpg'),
    detailImage: detailImage('detail-08.jpg'),
    industry: '3C家电',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'joyoung',
    title: '九阳新品上市内容种草，带动电商转化提升',
    image: caseImage('case-09.jpg'),
    detailImage: detailImage('detail-09.jpg'),
    industry: '3C家电',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'zuoyebang',
    title: '作业帮学习机抢占红海赛道，实现生意突破',
    image: caseImage('case-10.jpg'),
    detailImage: detailImage('detail-10.jpg'),
    industry: '教育',
    method: 'TOP1赛道抢占',
    platform: '小红书',
  },
  {
    id: 'peacebird',
    title: '太平鸟春季新品上线，以情绪共鸣促进转化',
    image: caseImage('case-11.jpg'),
    detailImage: detailImage('detail-11.jpg'),
    industry: '其他',
    method: '新品/爆品打造',
    platform: '小红书',
  },
  {
    id: 'wuyinliangpin',
    title: '吾饮良品以情绪化种草带动区域门店热销',
    image: caseImage('case-12.jpg'),
    detailImage: detailImage('detail-12.jpg'),
    industry: '食品饮料',
    method: '事件营销',
    platform: '小红书',
  },
]
