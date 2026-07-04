import { writeFileSync } from 'node:fs'
import { caseCards } from '../src/data/caseContent.js'

const outputPath = new URL('../server/seed_cases.sql', import.meta.url)

function sql(value) {
  return `'${String(value ?? '').replace(/\\/g, '\\\\').replace(/'/g, "''")}'`
}

function caseSummary(item) {
  return `${item.title}，所属行业：${item.industry}，营销方式：${item.method}，投放平台：${item.platform}。`
}

function caseContent(item) {
  return [
    `<p>${caseSummary(item)}</p>`,
    `<p>本案例围绕品牌阶段目标，结合平台内容语境、达人资源与转化链路，形成从策略到执行的完整社交营销方案。</p>`,
    `<p><img src="${item.detailImage}" alt="${item.title}" /></p>`,
  ].join('\n')
}

function metrics(item) {
  return JSON.stringify({
    detailImage: item.detailImage,
    source: 'src/data/caseContent.js',
  })
}

const values = caseCards.map((item, index) => {
  const sortOrder = (index + 1) * 10
  return `  (${[
    sql(item.title),
    sql(item.id),
    sql(item.industry),
    sql(item.platform),
    sql(item.method),
    sql(item.image),
    sql(caseSummary(item)),
    sql(caseContent(item)),
    sql(metrics(item)),
    sortOrder,
    sql('published'),
  ].join(', ')})`
})

const sqlText = `-- Generated from src/data/caseContent.js.
-- Re-run with: node scripts/generate-case-seed.mjs

SET NAMES utf8mb4;

INSERT INTO cases (
  title,
  slug,
  industry,
  platform,
  strategy,
  cover_url,
  summary,
  content,
  core_metrics,
  sort_order,
  status
) VALUES
${values.join(',\n')}
ON DUPLICATE KEY UPDATE
  title = VALUES(title),
  industry = VALUES(industry),
  platform = VALUES(platform),
  strategy = VALUES(strategy),
  cover_url = VALUES(cover_url),
  summary = VALUES(summary),
  content = VALUES(content),
  core_metrics = VALUES(core_metrics),
  sort_order = VALUES(sort_order),
  status = VALUES(status);
`

writeFileSync(outputPath, sqlText, 'utf8')
console.log(`Wrote ${caseCards.length} cases to ${outputPath.pathname}`)
