# WSD Social Case Page Design

Source targets:
- Official reference: https://www.wsdsocial.com/zh-cn/case
- UI-UX-Pro-Max generated baseline: Minimal Single Column, strong focus, high contrast, mobile-first, visible interaction states.

## Design Decision

Use UI-UX-Pro-Max as the UX framework, but override the visual style to match the official site exactly:
- Not brutalist/pink; use the WSD red-white corporate system from `design-system/MASTER.md`.
- Keep one clear content journey: banner -> filters -> case list -> official footer.
- Keep page width and spacing faithful to official desktop layout.

## Official Layout Targets

- Header height: 64px desktop.
- Banner starts below header and is 340px high at 1280px viewport.
- Banner image: `assets/cn/subPage/img_case_banner_pc.jpg`; headline text is embedded in the image on the official site.
- Case list section starts immediately after banner.
- Desktop case content width: 1200px centered for large desktop, falling back to 788px at narrower desktop widths.
- Filter block: 3 rows, labels `行业`, `营销方式`, `平台`.
- Filter row height cluster: about 136px total, no card shadow wrapper.
- Case grid: 3 columns on large desktop, 2 columns on medium desktop, 1 column on mobile.
- Case image size desktop: 376px x 250px.
- Column gap: 36px.
- Vertical rhythm between rows: about 60px.
- Card text sits below the image; no heavy card box, no shadow card wrapper.
- Tag text order: title, industry, marketing method, platform.

## Content Model

Filters:
- 行业: 全部, 3C家电, 美妆日化, 母婴宠物, 食品饮料, 房产家居, 网服, 教育, 其他
- 营销方式: 全部, 新品/爆品打造, 种草全域转化, TOP1赛道抢占, 线索获取, 直播电商, 品牌社媒运营, 社交营销培训, 事件营销
- 平台: 全部, 腾讯, 微博, 小红书, B站, 知乎, 多媒体投放

Case list uses official image/title/tag data where available.

## Visual Rules

- Background: white for banner/list transition, footer dark.
- Filters: plain text chips, active state red text with pale red pill background.
- Radius: 0 on images, 20px max on filter pills.
- Typography: Chinese corporate sans stack from global CSS.
- Case title: 18px, bold, line-height around 1.45.
- Case metadata: 14px muted text, inline with spacing.
- Hover: image lift/scale should be very subtle; no large shadow card style.

## Responsive Rules

- Under 900px: content width becomes page gutter width, cards become 1 column.
- Mobile banner uses a taller crop and preserves official image mood.
- Filter rows stack label above chips below 640px.
- No horizontal scroll at 375px.

## Verification

- Desktop 1280px: banner 340px, list content centered at 788px, 2 cards per row.
- Mobile 375px: single column, filter chips wrap, menu button visible.
- Build must pass.
