# WSD Social Message Page

Source target: official `营销资讯` route on https://www.wsdsocial.com/zh-cn/message.

## Structure

- Route: `/message`, official alias `/zh-cn/message`.
- Header `营销咨询` remains a modal trigger, separate from this route.
- Page sections:
  - Banner: 340px desktop height, local image `/assets/wsd/message/banner.webp`.
  - Title: `关注行业动态，及时获取前沿资讯`, 40px / 56px, weight 600.
  - Filter/search row: centered 788px container.
- Article grid: three columns on desktop, each card 376px x 398px, column gap 36px.

## Controls

- Categories:
  - 全部资讯
  - 企业动态
  - 营销观察
  - 案例故事
- Active tab uses red `#ff4848` and 600 weight.
- Search placeholder: `请输入关键词搜索资讯`.

## Card

- Image: 376px x 188px.
- Content padding: 20px.
- Title: 20px / 28px, two-line clamp.
- Description: 16px / 22px, two-line clamp.
- Footer: author `好品牌 会社交` + timestamp, 14px.
- Cards reveal on viewport entry.

## Responsive

- <= 1240px: two columns.
- <= 900px: stack controls and cards to one column.
- Maintain no horizontal scroll at 375px.
