const serviceImage = (name) => `/assets/wsd/${name}`

export const publicServices = [
  {
    id: 'social-media-partner',
    title: '社交媒体深度商业合作伙伴',
    summary: '连接主流社交平台资源，为品牌提供策略、投放、内容与转化一体化服务。',
    image: serviceImage('service-01.png'),
    highlights: ['平台商业合作', '账号与内容策略', '整合投放管理'],
    process: ['需求诊断', '平台组合', '内容排期', '效果复盘'],
  },
  {
    id: 'creator-content',
    title: '达人内容服务的策略及流程',
    summary: '围绕品牌目标筛选达人、共创内容，并通过精细化运营提升种草效率。',
    image: serviceImage('service-02.png'),
    highlights: ['达人矩阵搭建', '内容脚本共创', '投后数据复盘'],
    process: ['人群洞察', '达人筛选', '脚本执行', '内容扩散'],
  },
  {
    id: 'marketing-training',
    title: '营销培训服务提供的服务项目及平台',
    summary: '为企业团队提供社交营销方法论、平台规则和实战案例培训。',
    image: serviceImage('service-03.png'),
    highlights: ['平台规则解读', '案例拆解', '团队实战工作坊'],
    process: ['课题确认', '课程定制', '现场授课', '作业点评'],
  },
  {
    id: 'social-commerce',
    title: '社交电商服务提供的服务项目',
    summary: '打通内容种草、直播承接和私域转化，让社交声量沉淀为生意增长。',
    image: serviceImage('service-04.png'),
    highlights: ['直播电商', '私域承接', '转化链路优化'],
    process: ['货盘诊断', '内容预热', '直播执行', '成交复盘'],
  },
  {
    id: 'reputation-management',
    title: '口碑管理服务提供的服务项目',
    summary: '监测品牌舆情与用户反馈，构建长期可信的社交口碑资产。',
    image: serviceImage('service-05.png'),
    highlights: ['舆情监测', '口碑内容建设', '风险响应'],
    process: ['声量扫描', '风险分级', '内容优化', '长期维护'],
  },
  {
    id: 'marketing-technology',
    title: '营销技术服务提供的服务项目',
    summary: '以数据工具和自动化能力辅助投放、内容管理和营销效果评估。',
    image: serviceImage('service-06.png'),
    highlights: ['数据看板', '线索管理', '自动化分析'],
    process: ['指标定义', '数据接入', '看板搭建', '持续迭代'],
  },
]

export const partnerFaqs = [
  {
    question: '提交合作需求后多久会联系？',
    answer: '工作日通常会在 24 小时内完成初步沟通，并根据需求安排策略顾问对接。',
  },
  {
    question: '是否支持指定平台或行业专项合作？',
    answer: '支持。表单中可填写目标平台、行业赛道、预算范围和希望达成的业务目标。',
  },
  {
    question: '合作前需要准备哪些资料？',
    answer: '建议准备品牌介绍、产品资料、阶段目标、已有账号或投放数据，便于更快完成诊断。',
  },
]
