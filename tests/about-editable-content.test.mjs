import assert from 'node:assert/strict'
import { existsSync, readFileSync } from 'node:fs'

const aboutView = readFileSync(new URL('../src/views/AboutView.vue', import.meta.url), 'utf8')
const publicApi = readFileSync(new URL('../src/services/publicApi.js', import.meta.url), 'utf8')
const serverApp = readFileSync(new URL('../server/app.go', import.meta.url), 'utf8')

assert.match(serverApp, /HandleFunc\("\/api\/public\/about",\s*a\.publicAbout\)/)
assert.match(publicApi, /export function getAboutContent/)
assert.match(publicApi, /request\('\/api\/public\/about'\)/)

assert.match(aboutView, /getAboutContent/)
assert.match(aboutView, /const aboutContent = ref/)
assert.match(aboutView, /<AboutHero :content="aboutContent\.hero"/)
assert.match(aboutView, /<AboutHonors[\s\S]*:honors="aboutContent\.honors"/)

const adminPagePath = new URL('../src/admin/pages/AboutPageView.vue', import.meta.url)
assert.equal(existsSync(adminPagePath), true)

