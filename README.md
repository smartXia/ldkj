# WSD Social Vue Clone

这是一个基于 Vue 3、Vue Router 和 Vite 的微思敦官网复刻项目，覆盖首页、客户案例、营销资讯、关于我们和营销咨询弹窗等主要体验。

## 功能

- 首页：品牌轮播、营销模型、服务能力、案例入口、客户墙和咨询转化区。
- 客户案例：支持按行业、营销方式、平台筛选，并可打开案例详情弹窗。
- 营销资讯：支持分类切换和关键词搜索。
- 关于我们：包含集团介绍、价值观、发展历程、荣誉、子品牌和团队板块。
- 全局能力：固定导航、移动端菜单、语言切换、右侧联系入口和营销咨询表单。

## 环境要求

- Node.js 24.9.0 已验证可用。
- npm 随 Node.js 安装。

## 本地启动

```bash
npm install
npm run dev
```

默认开发地址为 `http://localhost:5173/`。如果端口被占用，Vite 会自动切换到下一个可用端口。

## 常用命令

```bash
npm test
npm run build
npm run preview
```

- `npm test`：检查关键中文内容，防止案例页文案出现错码或 README 回退为模板说明。
- `npm run build`：生成生产构建，输出到 `dist/`。
- `npm run preview`：本地预览生产构建。

## 页面路由

| 页面 | 路由 |
| --- | --- |
| 首页 | `/`、`/zh-cn` |
| 客户案例 | `/case`、`/zh-cn/case` |
| 营销资讯 | `/message`、`/zh-cn/message` |
| 关于我们 | `/about`、`/zh-cn/about` |

## 项目结构

```text
src/
  components/      页面组件和通用组件
  composables/     国际化等组合式逻辑
  data/            页面内容数据
  router/          Vue Router 配置
  views/           路由页面
public/assets/wsd/ 官方视觉素材
design-system/     页面设计与实现约束
tests/             内容回归测试
```

## 验证记录

当前版本已通过：

- `npm test`
- `npm run build`
