# WSD Social About Page

Source target: official `关于我们` route on https://www.wsdsocial.com/zh-cn/about.

## Structure

- Route: `/about`, official alias `/zh-cn/about`.
- Header `营销咨询` remains a modal trigger.
- Sections:
  - Hero: 520px desktop, local `/assets/wsd/about/banner.webp`.
  - Anchor nav: 1040px wide, 64px tall, sticky under global header.
  - Group intro: 1040px content, background visual `/assets/wsd/about/intro-bg.png`.
  - W/S/D meanings: 3 columns.
  - Values: 2x2 centered grid.
  - History: official desktop video animation with local poster and mobile static fallback.
  - Honors: 3-column industry-awards block.
  - Sub-brands: 3-column logo cards.
  - Team: compact closing section.

## Visual Rules

- Use official red `#ff4848` for active anchors and honor tabs.
- Main section title: 40px / 48px, weight 600.
- Body copy: 16px / 30px, #666.
- Keep desktop containers aligned to 1040px unless official card grid needs 3 columns.

## Responsive

- <= 1100px: use `calc(100% - 40px)` containers.
- <= 760px: stack values, honors, and brand cards.
- No horizontal scroll at 375px.
