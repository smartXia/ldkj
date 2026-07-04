# Website Platform Implementation Plan

> 面向 AI 代理的工作者：使用子 agent 分块实现。每个任务必须只修改自己负责的目录，不能回退其他人的改动。

**目标：** 将当前静态 Vue 官网补齐为可运行的前台 + Go API + MySQL + 管理端方案，覆盖需求 README 中尚未落地的核心功能。

**架构：** 前台继续使用 Vue 3/Vite。新增 `server/` Go 服务提供公开接口、管理接口、MySQL 持久化、本地 `/oss` 上传、表单邮件通知。新增 `src/admin/` 管理端页面，通过管理 API 完成内容、表单、站点设置和 SEO 管理。

**技术栈：** Vue 3、Vite、Go、MySQL、本地文件存储。

---

## 任务 1：Go 后端与数据库

**负责目录：**
- `server/**`

**要求：**
- 创建 Go module 和 HTTP API。
- 使用 MySQL，默认读取环境变量，提供 README 中临时测试库的示例配置。
- 实现公开接口：站点配置、Banner、案例列表/详情、资讯列表/详情、表单提交。
- 实现管理接口：登录、案例 CRUD、资讯 CRUD、Banner 更新、站点设置、SEO、表单列表、状态更新、CSV 导出、文件上传到当前运行目录 `/oss`。
- 实现基础鉴权、输入校验、SQL 注入防护、XSS 基础处理、CSRF 适配说明。
- 提供 `server/schema.sql`、`server/README.md`、后端测试。

## 任务 2：管理端

**负责目录：**
- `src/admin/**`
- `src/services/adminApi.js`

**要求：**
- 实现管理端登录页和后台布局。
- 实现 Dashboard、案例管理、资讯管理、Banner 管理、表单管理、站点设置、SEO 设置、上传控件。
- 使用 `src/services/adminApi.js` 封装管理接口。
- 管理端界面以高密度、安静、实用为主，不做营销页风格。
- 不修改 `src/router/index.js`，路由由主线程整合。

## 任务 3：前台接口化与缺失页面

**负责目录：**
- `src/services/publicApi.js`
- `src/views/**`
- `src/components/home/**`
- `src/components/case/**`
- `src/components/message/**`
- `src/data/**`

**要求：**
- 封装公开 API，失败时使用现有静态数据兜底。
- 首页和咨询弹窗表单提交到后端接口，增加基础必填校验和提交状态。
- 新增或调整服务页、服务详情页、合作页、案例详情页、资讯详情页的前台实现。
- 列表页支持分页，每页 9 至 12 条。
- 资讯分类调整为灵动动态、营销洞察、行业资讯。
- 不修改 `src/router/index.js`，路由由主线程整合。

## 主线程整合

- 整合前台和管理端路由。
- 更新 `package.json` 脚本。
- 补充根 README 的启动说明。
- 运行 `npm.cmd run build`、后端 `go test ./...`。
