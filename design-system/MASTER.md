# WSD Social Clone Design System

Source: https://www.wsdsocial.com/zh-cn
Generated with UI-UX-Pro-Max rules, then overridden to match the referenced official site.

## Product And Tone
- Product type: Chinese corporate landing site for a social media digital marketing agency.
- Audience: brand owners, marketing teams, prospective employees, enterprise visitors.
- Visual tone: clean B2B confidence, white corporate shell, high-contrast black/gold campaign hero, red conversion accents.
- Core memory point: white fixed navigation and full-width promotional banner followed by spacious service storytelling.

## Color Tokens
```css
--color-brand: #ff3d45;
--color-brand-dark: #e93039;
--color-brand-soft: #fff0f1;
--color-ink: #111111;
--color-ink-muted: #62656d;
--color-ink-light: #8b8e96;
--color-surface: #ffffff;
--color-page: #f7f8fb;
--color-section: #f2f5f9;
--color-border: #edf0f5;
--color-gold: #ffd06b;
--color-black: #020202;
--shadow-card: 0 14px 34px rgba(26, 38, 66, 0.09);
--shadow-soft: 0 8px 26px rgba(255, 61, 69, 0.18);
```

## Typography
- Font family: `PingFang SC`, `Microsoft YaHei`, `Noto Sans SC`, system sans-serif.
- Body: 16px, line-height 1.6, color `--color-ink`.
- Navigation: 16px regular, active red with 4px underline.
- Section title desktop: 38-44px, font-weight 800, centered.
- Hero campaign text is image-driven; fallback text must stay high contrast.
- Letter spacing: 0.

## Layout
- Header: fixed, 88px desktop, white, centered 1040px-1120px content.
- Main content starts below header. Hero full width with 16:6.6 desktop ratio and no card wrapper.
- Desktop container: max-width 1040px, side gutter 24px.
- Section vertical rhythm: 88-120px desktop, 56-72px mobile.
- Cards: 8px radius max unless original image already has softer corners.
- Floating contact rail: fixed right center, 44px square controls, red icons on white.

## Components
- Header: logo left, nav center, red filled consultation CTA, red outline login, language pill.
- HeroCarousel: full-width image slides, dot pagination over bottom center, 5s interval, pause not required.
- MarketingModel: large centered H1 plus three pale module cards connected by visual arrows on desktop.
- ServiceSection: light gray-blue background, tabs on top, large service image panel, red CTA.
- BrandCases: two-row responsive case cards with image, number, title, summary, detail button.
- IntroStats: background image or pale band, three numeric metrics: 12 years, 3000+ clients, 120+ honors.
- BrandWall: category tabs and dense logo grid.
- ContactCta: red gradient band, concise lead capture CTA.
- Footer: dark corporate footer with nav columns and copyright.

## Interaction
- Hover and pressed states: 180-260ms, opacity/transform only.
- Buttons never change layout size on hover.
- Keyboard focus: 2px red outline with 3px offset.
- Reduced motion: disable carousel transition and card entrance animations.

## Responsive
- 1440 desktop: nav visible, hero image crops center, service cards in wide layout.
- 1024 tablet: nav remains if space allows; grids reduce to 2-3 columns.
- 768 and below: nav collapses to compact menu button, hero ratio becomes taller, case cards 1 column.
- 375 mobile: all touch targets >=44px, no horizontal scroll, floating rail moves lower and smaller.

## Implementation Rules
- Use Vue 3 Composition API with `<script setup>` and JavaScript because the project request is Vue3 + JS.
- Keep global tokens in `src/style.css` and section-specific styles scoped in SFCs.
- Use original public image URLs where possible for visual fidelity.
- All meaningful images need alt text.
- Avoid emoji icons and decorative gradient blobs.
