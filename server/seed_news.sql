-- Mock news seed data for the marketing news section.
-- Import after schema.sql. Re-run safely because slug has a unique key.

SET NAMES utf8mb4;

INSERT INTO news (
  title,
  slug,
  category,
  author,
  cover_url,
  summary,
  content,
  status,
  published_at
) VALUES
  (
    '2026 社交内容营销趋势：从流量采买走向长期经营',
    'mock-social-content-marketing-trends-2026',
    'insight',
    '灵动研究院',
    '/assets/wsd/message/news-01.webp',
    '品牌正在从单次曝光转向内容资产沉淀，达人、社区和私域链路需要被放在同一张增长地图里管理。',
    '<h2>从短期声量到长期资产</h2><p>2026 年，社交内容营销的核心变化不再只是追逐平台流量，而是把内容、达人、人群和转化数据纳入持续经营。品牌需要用更稳定的内容主题、更清晰的人群分层和更完整的复盘机制，沉淀可复用的增长资产。</p><p>对营销团队来说，预算配置也会从单点投放转为组合打法：品牌内容负责建立认知，达人内容负责建立信任，搜索和私域承接负责提升转化效率。</p>',
    'published',
    '2026-07-01 10:00:00'
  ),
  (
    '小红书种草进入精细化阶段，品牌如何提升内容转化率',
    'mock-rednote-seeding-conversion-playbook',
    'insight',
    '内容策略组',
    '/assets/wsd/message/news-02.webp',
    '小红书种草不再只是铺量，品牌需要围绕关键词、场景、人群和评论区反馈建立更细的内容管理机制。',
    '<h2>种草需要更强的结构化运营</h2><p>当用户决策路径变长，单篇爆文很难独立承担转化目标。品牌需要围绕核心关键词设计内容矩阵，并持续监测搜索结果页、收藏率、评论关键词和私信咨询变化。</p><p>更有效的做法是把达人笔记、品牌号内容和店铺承接统一规划，减少用户从兴趣到下单之间的流失。</p>',
    'published',
    '2026-06-29 14:30:00'
  ),
  (
    'B2B 企业做 LinkedIn 内容，第一步不是发产品介绍',
    'mock-b2b-linkedin-content-first-step',
    'insight',
    '出海营销组',
    '/assets/wsd/message/news-03.webp',
    'B2B 出海品牌在 LinkedIn 上应先建立专业议题和可信观点，再逐步导入案例、解决方案和线索表单。',
    '<h2>先建立观点，再承接线索</h2><p>LinkedIn 用户更关注专业判断、行业经验和可信背书。对于 B2B 企业来说，第一阶段不适合大量发布产品硬广，而应围绕客户痛点、行业变化和方法论输出稳定内容。</p><p>当账号形成清晰定位后，再通过案例拆解、白皮书下载和活动报名承接线索，转化效率会更稳。</p>',
    'published',
    '2026-06-27 09:20:00'
  ),
  (
    '短视频投放复盘：如何判断一条素材是否值得继续放量',
    'mock-short-video-ad-creative-review',
    'industry',
    '投放运营组',
    '/assets/wsd/message/news-04.webp',
    '判断短视频素材价值不能只看点击率，还要结合完播、互动、转化成本和评论情绪做综合评估。',
    '<h2>素材复盘要看完整链路</h2><p>一条短视频素材是否值得继续投放，需要同时观察前端吸引力和后端转化效率。点击率高但评论负面、转化成本高的素材，并不一定适合扩大预算。</p><p>建议按照前三秒留存、完播率、互动率、落地页转化率和成交成本建立素材评分模型，让投放决策更稳定。</p>',
    'published',
    '2026-06-25 16:45:00'
  ),
  (
    '品牌直播间增长放缓后，内容团队还能做什么',
    'mock-brand-livestream-growth-after-slowdown',
    'industry',
    '电商内容组',
    '/assets/wsd/message/news-05.webp',
    '直播间增长放缓时，内容团队应从货品表达、短视频预热、达人背书和复购内容四个方向重新拆解效率。',
    '<h2>直播增长需要内容前置</h2><p>直播间不是孤立的交易场景。增长放缓往往意味着用户对货品价值、品牌信任或购买理由的理解不足。此时内容团队需要提前承担解释和种草任务。</p><p>通过短视频预热、达人体验、直播切片和用户反馈二次传播，可以为直播间持续补充高意向流量。</p>',
    'published',
    '2026-06-23 11:15:00'
  ),
  (
    '灵动信息完成多平台内容服务能力升级',
    'mock-lingdong-multiplatform-content-service-upgrade',
    'news',
    '灵动信息',
    '/assets/wsd/message/news-06.webp',
    '灵动信息围绕小红书、抖音、B站、微信生态和海外社媒，升级内容策略、达人管理和数据复盘能力。',
    '<h2>多平台协同服务升级</h2><p>近日，灵动信息完成多平台内容服务能力升级，进一步打通策略、创意、达人、投放和复盘环节。新服务体系将围绕品牌阶段目标，提供更精细的平台组合方案。</p><p>未来，团队将继续加强行业案例沉淀和数据分析能力，帮助客户在复杂内容环境中提升增长确定性。</p>',
    'published',
    '2026-06-21 15:00:00'
  ),
  (
    'AI 生成内容进入营销工作流，品牌需要先建立审核标准',
    'mock-ai-content-marketing-review-standard',
    'industry',
    '营销技术组',
    '/assets/wsd/message/news-07.webp',
    'AI 可以提升内容生产效率，但品牌必须明确事实校验、调性一致性、版权风险和平台规范边界。',
    '<h2>效率提升不能替代品牌判断</h2><p>AI 工具正在进入选题、脚本、图片和复盘等营销工作流。它能提升效率，但不能替代品牌对事实、调性和合规边界的判断。</p><p>建议企业先建立内容审核标准，再逐步把 AI 用于资料整理、初稿生成和版本测试，降低内容风险。</p>',
    'published',
    '2026-06-19 13:40:00'
  ),
  (
    '新品上市期如何配置达人矩阵：头部声量与腰尾部信任并重',
    'mock-new-product-kol-matrix-launch',
    'insight',
    '达人运营组',
    '/assets/wsd/message/news-08.webp',
    '新品上市期的达人矩阵应兼顾破圈声量、垂类解释和真实体验，避免预算过度集中在单一达人层级。',
    '<h2>达人矩阵要匹配上市节奏</h2><p>新品上市通常同时需要认知扩散和使用理由解释。头部达人适合制造声量，垂类达人适合解释卖点，腰尾部达人适合提供真实体验和评论区互动。</p><p>更稳妥的做法是按预热、首发、扩散和长尾四个阶段配置不同达人资源，并持续优化内容角度。</p>',
    'published',
    '2026-06-17 10:10:00'
  ),
  (
    '社区运营的关键不是高频发帖，而是形成可讨论的主题',
    'mock-community-operation-topic-design',
    'insight',
    '社群策略组',
    '/assets/wsd/message/news-09.webp',
    '好的社区运营需要持续设计用户愿意回应的话题，而不是单向发布品牌信息。',
    '<h2>让用户有话可说</h2><p>社区运营的难点不在于发帖频率，而在于能否持续提出用户愿意参与的话题。品牌需要把产品卖点翻译成用户场景、情绪和经验问题。</p><p>当社区形成稳定讨论主题后，内容反馈会反过来帮助品牌优化产品表达和投放素材。</p>',
    'published',
    '2026-06-15 18:25:00'
  ),
  (
    '618 后复盘：品牌如何把大促流量转化为长期用户资产',
    'mock-618-review-long-term-user-asset',
    'industry',
    '增长复盘组',
    '/assets/wsd/message/news-10.webp',
    '大促结束后，品牌应及时复盘人群质量、内容触点和复购路径，把短期成交沉淀为长期经营资产。',
    '<h2>大促复盘要回到人群经营</h2><p>618 结束并不意味着营销周期结束。品牌需要复盘哪些内容带来了高质量人群，哪些渠道贡献了复购潜力，哪些用户问题暴露了承接短板。</p><p>把大促期间积累的搜索词、评论、私信和成交数据整理成经营资产，才能支持下一阶段的持续增长。</p>',
    'published',
    '2026-06-13 09:50:00'
  )
ON DUPLICATE KEY UPDATE
  title = VALUES(title),
  category = VALUES(category),
  author = VALUES(author),
  cover_url = VALUES(cover_url),
  summary = VALUES(summary),
  content = VALUES(content),
  status = VALUES(status),
  published_at = VALUES(published_at);
