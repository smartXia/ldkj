# 后台资讯富文本编辑器实现计划

> **面向 AI 代理的工作者：** 必需子技能：使用 superpowers:subagent-driven-development（推荐）或 superpowers:executing-plans 逐任务实现此计划。步骤使用复选框（`- [ ]`）语法来跟踪进度。

**目标：** 让后台资讯管理的正文满足富文本编辑器录入能力。

**架构：** 在后台通用内容管理表单中新增 `richtext` 字段类型，由独立的 `RichTextEditor.vue` 组件负责工具栏、可编辑区域和 HTML 值同步。资讯页面只声明 `content` 字段为 `richtext`，不改服务端接口。

**技术栈：** Vue 3、Vite、浏览器原生 `contenteditable` 和 `document.execCommand`。

---

### 任务 1：结构测试

**文件：**
- 创建：`tests/admin-richtext-check.mjs`

- [ ] **步骤 1：编写失败的测试**

```js
import { readFile } from 'node:fs/promises'
import assert from 'node:assert/strict'

const newsView = await readFile(new URL('../src/admin/pages/NewsView.vue', import.meta.url), 'utf8')
const manager = await readFile(new URL('../src/admin/components/ContentManager.vue', import.meta.url), 'utf8')
const editor = await readFile(new URL('../src/admin/components/RichTextEditor.vue', import.meta.url), 'utf8')

assert.match(newsView, /key:\s*'content'[\s\S]*type:\s*'richtext'/)
assert.match(manager, /RichTextEditor/)
assert.match(manager, /field\.type === 'richtext'/)
assert.match(editor, /contenteditable="true"/)
assert.match(editor, /defineModel/)
assert.match(editor, /execCommand/)
assert.match(editor, /源码/)
```

- [ ] **步骤 2：运行测试验证失败**

运行：`node tests/admin-richtext-check.mjs`
预期：FAIL，原因是 `RichTextEditor.vue` 不存在，且资讯正文仍为 `textarea`。

### 任务 2：富文本编辑器组件

**文件：**
- 创建：`src/admin/components/RichTextEditor.vue`

- [ ] **步骤 1：实现最少组件**

实现内容：
- 工具栏支持段落、二级标题、三级标题、加粗、斜体、下划线、引用、有序列表、无序列表、链接、图片、清除格式、源码视图。
- 使用 `defineModel({ type: String, default: '' })` 与父表单同步 HTML。
- `contenteditable` 区域输入时写回 HTML，外部值变化时刷新编辑区。

- [ ] **步骤 2：运行测试验证通过**

运行：`node tests/admin-richtext-check.mjs`
预期：PASS。

### 任务 3：接入资讯表单

**文件：**
- 修改：`src/admin/components/ContentManager.vue`
- 修改：`src/admin/pages/NewsView.vue`
- 修改：`src/admin/styles.css`

- [ ] **步骤 1：接入字段类型**

在 `ContentManager.vue` 导入 `RichTextEditor`，并在字段渲染中对 `field.type === 'richtext'` 使用组件。

- [ ] **步骤 2：切换资讯正文**

将 `NewsView.vue` 的 `content` 字段从 `textarea` 改为 `richtext`。

- [ ] **步骤 3：补充后台编辑器样式**

在 `src/admin/styles.css` 增加编辑器容器、工具栏、按钮、正文区域和源码文本域样式，保持后台管理的克制工作台风格。

- [ ] **步骤 4：验证构建**

运行：`npm run build`
预期：Vite 构建成功，退出码为 0。
