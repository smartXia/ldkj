# WSD Social Marketing Consult Module

Source target: official `营销咨询` dialog on https://www.wsdsocial.com/zh-cn/case.

## Structure

- Triggered by desktop header `营销咨询` button and consulting-related floating entries.
- Modal overlay covers viewport with a dark scrim.
- Dialog size on desktop: 860px x 600px.
- Dialog position: centered, slightly below top at desktop viewport.
- Dialog has two columns:
  - Left: 420px visual panel.
  - Right: 440px form panel.

## Left Panel

- Background image: `/assets/wsd/consulting/side-bg.png`.
- Logo: existing `/assets/wsd/logo.png`, 124px x 40px.
- Title: `好品牌 会社交`.
- Tip text: `预约咨询，获取以下信息`.
- Benefit list with checked icon:
  - 定制化营销解决方案
  - 1对1营销顾问快速反馈
  - 行业报告、神秘礼物，更多惊喜等着你
- QR code: `/assets/wsd/consulting/marketing-qrcode.png`, 108px x 108px.
- QR caption: `直接联系营销顾问`.

## Right Panel

- Title: `预约咨询`.
- Form labels:
  - 姓名
  - 手机号码
  - 公司名称
  - 营销述求
- Placeholders:
  - 您的姓名
  - 请填写手机号码
  - 请填写公司名称
  - 请选择营销诉求
- Dropdown options:
  - 整合营销
  - 新品推广
  - 电商大促
  - 社交种草
  - 账号运营
  - 舆情优化
  - 其他
- Primary CTA: `立即预约`.

## Visual Details

- Keep white right panel, red primary CTA, light gray form borders.
- Left panel text is dark over pale pink official background.
- Close button is a small icon button in top-right of dialog area.
- Use no emoji icons.
- All controls >=44px high.

## Behavior

- Click trigger opens modal.
- Close by close button, overlay click, or Esc.
- Body scroll locks while modal is open.
- Form is front-end only for now; submit shows a lightweight success state.
- Dropdown opens inline under the field and closes after option selection.

## Responsive

- <= 860px: modal width follows viewport, columns stack.
- Mobile hides no key content; QR code remains visible below benefit list.
- No horizontal scroll at 375px.
