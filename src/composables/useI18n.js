import { computed, shallowRef } from 'vue'

const currentLanguageCode = shallowRef('zh')

export const languages = [
  { code: 'zh', label: '中文' },
  // { code: 'en', label: 'English' },
  // { code: 'ja', label: '日本語' },
]

const localAbout = (name) => `/assets/wsd/about/${name}`

const sharedAboutAssets = {
  heroImage: localAbout('banner.webp'),
  introImage: localAbout('intro-bg.png'),
  historyImage: localAbout('history.webp'),
  historyVideo: localAbout('history.mp4'),
  historyPoster: localAbout('history-poster.png'),
}

const aboutAnchors = {
  zh: [
    ['about-intro', '集团介绍'],
    ['about-values', '价值观'],
    ['about-history', '发展历程'],
    ['about-honors', '企业荣誉'],
    ['about-brands', '子品牌'],
    ['about-team', '团队风采'],
  ],
  en: [
    ['about-intro', 'Group Intro'],
    ['about-values', 'Values'],
    ['about-history', 'Milestones'],
    ['about-honors', 'Honors'],
    ['about-brands', 'Sub-brands'],
    ['about-team', 'Team Life'],
  ],
  ja: [
    ['about-intro', 'グループ紹介'],
    ['about-values', '価値観'],
    ['about-history', '沿革'],
    ['about-honors', '企業栄誉'],
    ['about-brands', 'サブブランド'],
    ['about-team', 'チーム風采'],
  ],
}

const content = {
  zh: {
    navItems: [
      { label: '首页', path: '/' },
      { label: '客户案例', path: '/case' },
      { label: '营销资讯', path: '/message' },
      { label: '关于我们', path: '/about' },
      { label: '加入我们', href: '#jobs' },
    ],
    header: {
      consult: '营销咨询',
      login: '登录/注册',
    },
    about: {
      anchors: aboutAnchors.zh.map(([id, label]) => ({ id, label })),
      hero: {
        title: '灵动丨好品牌 会社交',
        desc: [
          '使命：以社交为桥梁，以科技为引擎，助力品牌实现可持续增长。',
          '愿景：成为全球品牌更值得信赖的社交营销伙伴。',
        ],
        image: sharedAboutAssets.heroImage,
      },
      intro: {
        title: '集团介绍',
        image: sharedAboutAssets.introImage,
        paragraphs: [
          '灵动集团创立于2014年，汇聚500多位营销人才。',
          '集团下设灵动、WSD、皓量科技、思多睿、海月昇、取泊、游友广告、拾刻互动、吉讯互动等子公司，核心业务涵盖社交广告、达人营销、口碑营销、社交电商、营销培训、营销技术等。',
          '灵动集团作为微信、小红书、B站、Instagram、LinkedIn、Reddit等社交媒体的深度商业合作伙伴，长期服务宝洁、名创优品、小米、安克创新、小熊电器等知名品牌，多次获得行业大奖认可。',
          '灵动独创ACT营销模型，依托自研灵动客户中心营销平台，聚合广告Advertising、内容Content、交易Transaction三维能力，帮助品牌更科学地实现跨越式增长。',
          '我们将持续深耕社交营销，以系统标准化服务流程，为品牌高质量发展提供解决方案。',
        ],
      },
      meanings: [
        {
          title: 'Wisdom ·',
          cn: '智慧',
          icon: localAbout('meaning-w.webp'),
          desc: '智慧（Wisdom） 是我们对知识与创新的追求。我们致力于为客户提供智慧化的营销解决方案和服务，帮助品牌营销提质、提速、提效，实现跨越式增长。',
        },
        {
          title: 'Social ·',
          cn: '社交',
          icon: localAbout('meaning-s.webp'),
          desc: '社交（Social） 代表着我们与品牌、消费者之间的联系与互动。社交媒体作为我们的主要平台，我们以品牌目标为驱动，以用户需求为核心，利用社交媒体的力量来传递品牌价值，与消费者建立良好的、长期的用户关系。',
        },
        {
          title: 'Digital ·',
          cn: '数字化',
          icon: localAbout('meaning-d.webp'),
          desc: '数字化（Digital） 是我们在数字时代的专业营销方式和技术能力。我们独创ACT营销模型，以广告（Advertising）、内容（Content）和交易（Transaction）为核心的三种营销手段，构建可持续、可循环、可衡量的系统性营销方式。',
        },
      ],
      valuesTitle: '价值观',
      values: [
        { title: '成就客户', icon: localAbout('value-customer.webp'), lines: ['专注专业、求真务实', '最好的关系是彼此成就'] },
        { title: '开放合作', icon: localAbout('value-open.webp'), lines: ['保持开放的心态', '创造各种达成合作的可能'] },
        { title: '平等包容', icon: localAbout('value-equality.webp'), lines: ['对外合作上，平等互利相互促进彼此成就', '对内协作上，自由平等包容互助共同成长'] },
        { title: '保持好奇心', icon: localAbout('value-curiosity.webp'), lines: ['保持谦逊，不断学习，打破原有的边界', '不断探索，拥抱变化，挑战各种可能性'] },
      ],
      history: {
        title: '发展历程',
        image: sharedAboutAssets.historyImage,
        video: sharedAboutAssets.historyVideo,
        poster: sharedAboutAssets.historyPoster,
      },
      brands: [
        ['灵动官方Logo', '好品牌 会社交', '以品牌目标为驱动、以用户需求为核心、以社交媒体为平台，为品牌提供可持续、可循环、可衡量的系统性社交营销服务。', localAbout('brand-wsd.webp')],
        ['WSD官方Logo', '跨境社交营销专家', '专注海外社交营销，以创意与文化洞察，助力品牌在全球建立真实连接，实现有效传播与可持续增长。', localAbout('brand-global-wsd.webp')],
        ['皓量科技官方Logo', '社交广告服务商', '深耕社交广告营销领域，整合广告产品、AI和数据技术，助力品牌广告提效，实现营收增长。', localAbout('brand-hlkj.webp')],
        ['思多睿官方Logo', '亲子生态营销专家', '通过不同场景与内容，深入触达家庭消费群体，打造亲子生态营销闭环。', localAbout('brand-sdr.webp')],
        ['海月昇官方Logo', '房产泛生活社交营销机构', '提供多平台一站式社交营销服务，助力房产及泛生活行业品牌增长。', localAbout('brand-hys.webp')],
        ['取泊文化官方Logo', '种草营销厂牌', '深耕3C领域，用内容创造品牌与消费者之间的情感链接，为品牌做确定性增长。', localAbout('brand-qp.webp')],
        ['游友广告官方Logo', '网服社交营销专家', '聚焦网服、游戏、电商、社交营销技术革新，为品牌提供全链路营销解决方案。', localAbout('brand-yy.webp')],
        ['拾刻互动官方Logo', '社交整合营销品牌', '主要服务国内外头部集团客户，用精准策略、精细服务、精质成效，提供专业社交生态整合营销服务。', localAbout('brand-skhd.webp')],
        ['吉讯互动官方Logo', '腾讯生态营销专家', '专注腾讯全域生态，提供洞察、策略、执行、运营一站式闭环服务，助力品牌长期经营增长。', localAbout('brand-jx.webp')],
      ].map(([alt, title, desc, logo]) => ({ alt, title, desc, logo })),
      teamTitle: '团队风采',
    },
  },
  en: {
    navItems: [
      { label: 'Home', path: '/' },
      { label: 'Cases', path: '/case' },
      { label: 'Insights', path: '/message' },
      { label: 'About Us', path: '/about' },
      { label: 'Careers', href: '#jobs' },
    ],
    header: {
      consult: 'Consulting',
      login: 'Login / Register',
    },
    about: {
      anchors: aboutAnchors.en.map(([id, label]) => ({ id, label })),
      hero: {
        title: 'WSD | Good Brands Socialize Better',
        desc: [
          'Mission: connecting brands and people through social media, powered by technology.',
          'Vision: to become a trusted global partner for social marketing growth.',
        ],
        image: sharedAboutAssets.heroImage,
      },
      intro: {
        title: 'Group Introduction',
        image: sharedAboutAssets.introImage,
        paragraphs: [
          'Founded in 2014, WSD Group brings together more than 500 marketing professionals.',
          'The group operates multiple subsidiaries including WSD, Global WSD, Haoliang Technology, Story, Haiyuesheng, Qubo, Youyou Ads, Think Interaction and For Jixun, covering social ads, influencer marketing, word-of-mouth, social commerce, training and marketing technology.',
          'As a deep commercial partner of WeChat, Xiaohongshu, Bilibili, Instagram, LinkedIn, Reddit and other social platforms, WSD has served many well-known brands and has been recognized by leading industry awards.',
          'With its original ACT marketing model and self-developed customer-center platform, WSD combines Advertising, Content and Transaction capabilities to help brands achieve sustainable growth.',
          'We will continue to focus on social marketing and provide systematic, standardized solutions for high-quality brand development.',
        ],
      },
      meanings: [
        { title: 'Wisdom ·', cn: 'Wisdom', icon: localAbout('meaning-w.webp'), desc: 'Wisdom reflects our pursuit of knowledge and innovation. We provide intelligent marketing solutions and services to help brands improve quality, speed and efficiency.' },
        { title: 'Social ·', cn: 'Social', icon: localAbout('meaning-s.webp'), desc: 'Social represents the connection and interaction between brands and consumers. We use social media to deliver brand value and build long-term user relationships.' },
        { title: 'Digital ·', cn: 'Digital', icon: localAbout('meaning-d.webp'), desc: 'Digital represents our professional methods and technical capabilities in the digital era. With ACT, we build sustainable, measurable marketing systems.' },
      ],
      valuesTitle: 'Values',
      values: [
        { title: 'Customer Success', icon: localAbout('value-customer.webp'), lines: ['Professional focus and pragmatic execution', 'The best relationship is mutual achievement'] },
        { title: 'Open Collaboration', icon: localAbout('value-open.webp'), lines: ['Keep an open mindset', 'Create more possibilities for cooperation'] },
        { title: 'Equality & Inclusion', icon: localAbout('value-equality.webp'), lines: ['Mutual benefit in external partnerships', 'Free, equal and inclusive collaboration internally'] },
        { title: 'Stay Curious', icon: localAbout('value-curiosity.webp'), lines: ['Keep learning and break boundaries', 'Explore change and challenge possibilities'] },
      ],
      history: { title: 'Milestones', image: sharedAboutAssets.historyImage, video: sharedAboutAssets.historyVideo, poster: sharedAboutAssets.historyPoster },
      brands: [
        ['WSD official logo', 'Good Brands Socialize Better', 'A systematic social marketing service driven by brand goals, user needs and social media platforms.', localAbout('brand-wsd.webp')],
        ['Global WSD official logo', 'Cross-border Social Marketing Expert', 'Helping brands build real global connections through creativity and cultural insight.', localAbout('brand-global-wsd.webp')],
        ['Haoliang Technology logo', 'Social Advertising Service Provider', 'Integrating ad products, AI and data technology to improve advertising efficiency.', localAbout('brand-hlkj.webp')],
        ['Story official logo', 'Parent-child Ecosystem Marketing Expert', 'Reaching family consumers through scenarios and content.', localAbout('brand-sdr.webp')],
        ['Haiyuesheng official logo', 'Real Estate Lifestyle Social Marketing Agency', 'One-stop social marketing services for real estate and lifestyle brands.', localAbout('brand-hys.webp')],
        ['Qubo Culture official logo', 'Seeding Marketing Label', 'Building emotional links between brands and consumers through content.', localAbout('brand-qp.webp')],
        ['Youyou Ads official logo', 'Online Service Social Marketing Expert', 'Full-funnel marketing solutions for online services, games and commerce.', localAbout('brand-yy.webp')],
        ['Think Interaction official logo', 'Integrated Social Marketing Brand', 'Professional social ecosystem marketing services for leading enterprise clients.', localAbout('brand-skhd.webp')],
        ['For Jixun official logo', 'Tencent Ecosystem Marketing Expert', 'Closed-loop services across Tencent ecosystem insight, strategy, execution and operation.', localAbout('brand-jx.webp')],
      ].map(([alt, title, desc, logo]) => ({ alt, title, desc, logo })),
      teamTitle: 'Team Life',
    },
  },
  ja: {
    navItems: [
      { label: 'ホーム', path: '/' },
      { label: '事例', path: '/case' },
      { label: 'マーケ情報', path: '/message' },
      { label: '会社情報', path: '/about' },
      { label: '採用情報', href: '#jobs' },
    ],
    header: {
      consult: '相談する',
      login: 'ログイン/登録',
    },
    about: {
      anchors: aboutAnchors.ja.map(([id, label]) => ({ id, label })),
      hero: {
        title: 'WSD | 良いブランドは、よくつながる',
        desc: [
          '使命：ソーシャルを橋渡しに、テクノロジーを推進力に、ブランドの持続的成長を支援します。',
          'ビジョン：世界中のブランドに信頼されるソーシャルマーケティングパートナーへ。',
        ],
        image: sharedAboutAssets.heroImage,
      },
      intro: {
        title: 'グループ紹介',
        image: sharedAboutAssets.introImage,
        paragraphs: [
          'WSDグループは2014年に設立され、500名以上のマーケティング人材が集まっています。',
          'グループは複数の子会社を展開し、ソーシャル広告、インフルエンサー施策、口コミ、ソーシャルコマース、研修、マーケティング技術などをカバーしています。',
          'WeChat、小紅書、Bilibili、Instagram、LinkedIn、Redditなどのプラットフォームと深く連携し、多くの有名ブランドを支援してきました。',
          '独自のACTモデルにより、Advertising、Content、Transactionを統合し、ブランドの持続的成長を支援します。',
          '今後もソーシャルマーケティングを深耕し、高品質なブランド成長ソリューションを提供します。',
        ],
      },
      meanings: [
        { title: 'Wisdom ·', cn: '知恵', icon: localAbout('meaning-w.webp'), desc: 'Wisdomは知識と革新への追求を表します。スマートなマーケティングソリューションでブランド成長を支援します。' },
        { title: 'Social ·', cn: 'ソーシャル', icon: localAbout('meaning-s.webp'), desc: 'Socialはブランドと消費者のつながりを表します。ソーシャルメディアを通じてブランド価値を届けます。' },
        { title: 'Digital ·', cn: 'デジタル', icon: localAbout('meaning-d.webp'), desc: 'Digitalはデジタル時代の専門性と技術力を表します。ACTモデルで持続可能なマーケティングを構築します。' },
      ],
      valuesTitle: '価値観',
      values: [
        { title: '顧客の成功', icon: localAbout('value-customer.webp'), lines: ['専門性と実行力', '互いに成果を生む関係'] },
        { title: 'オープンな協力', icon: localAbout('value-open.webp'), lines: ['開かれた姿勢を保つ', '協力の可能性を広げる'] },
        { title: '平等と包容', icon: localAbout('value-equality.webp'), lines: ['外部とは互恵的に協力', '内部では自由で包容的に成長'] },
        { title: '好奇心を保つ', icon: localAbout('value-curiosity.webp'), lines: ['学び続け、境界を越える', '変化を受け入れ、可能性に挑む'] },
      ],
      history: { title: '沿革', image: sharedAboutAssets.historyImage, video: sharedAboutAssets.historyVideo, poster: sharedAboutAssets.historyPoster },
      brands: [
        ['WSD公式ロゴ', '良いブランドは、よくつながる', 'ブランド目標とユーザーニーズを中心にした体系的なソーシャルマーケティングサービス。', localAbout('brand-wsd.webp')],
        ['Global WSD公式ロゴ', '越境ソーシャルマーケティング専門家', '創意と文化理解で、グローバルなつながりを支援します。', localAbout('brand-global-wsd.webp')],
        ['皓量科技公式ロゴ', 'ソーシャル広告サービス', '広告商品、AI、データ技術で広告効果を高めます。', localAbout('brand-hlkj.webp')],
        ['Story公式ロゴ', '親子エコシステムマーケティング専門家', 'シーンとコンテンツで家庭消費者に深く届けます。', localAbout('brand-sdr.webp')],
        ['海月昇公式ロゴ', '不動産・ライフスタイル領域の代理店', '複数プラットフォームの一括マーケティング支援。', localAbout('brand-hys.webp')],
        ['取泊文化公式ロゴ', 'シーディングマーケティングブランド', 'コンテンツでブランドと消費者の感情的なつながりを作ります。', localAbout('brand-qp.webp')],
        ['游友广告公式ロゴ', 'ネットサービス領域の専門家', 'ゲーム、EC、ネットサービス向けの全体施策。', localAbout('brand-yy.webp')],
        ['拾刻互动公式ロゴ', '統合ソーシャルマーケティングブランド', '主要企業向けに専門的な統合サービスを提供します。', localAbout('brand-skhd.webp')],
        ['吉讯互动公式ロゴ', 'Tencentエコシステム専門家', '洞察、戦略、実行、運用まで一貫して支援します。', localAbout('brand-jx.webp')],
      ].map(([alt, title, desc, logo]) => ({ alt, title, desc, logo })),
      teamTitle: 'チーム風采',
    },
  },
}

export function useI18n() {
  const locale = computed(() => currentLanguageCode.value)
  const messages = computed(() => content[currentLanguageCode.value] ?? content.zh)
  const currentLanguage = computed(() =>
    languages.find((item) => item.code === currentLanguageCode.value) ?? languages[0]
  )

  function setLanguage(code) {
    if (content[code]) currentLanguageCode.value = code
  }

  return {
    locale,
    languages,
    currentLanguage,
    messages,
    setLanguage,
  }
}
