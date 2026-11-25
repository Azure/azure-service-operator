const assert = require('node:assert/strict');
const { readFileSync, readdirSync } = require('node:fs');
const { basename, join } = require('node:path');
const { test } = require('node:test');

const publicDir = join(__dirname, '../public');

function searchIndex(pagePath) {
  const html = readFileSync(join(publicDir, pagePath), 'utf8');
  const urls = [...html.matchAll(/data-offline-search-index-json-src="([^"]+)"/g)]
    .map((match) => match[1]);
  assert.ok(urls.length > 0, `No search inputs on ${pagePath}`);
  assert.equal(new Set(urls).size, 1, `Search inputs disagree on ${pagePath}`);
  const path = join(publicDir, basename(urls[0]));
  return { html, url: urls[0], documents: JSON.parse(readFileSync(path, 'utf8')) };
}

test('non-reference pages use the documentation index and label', () => {
  const guide = searchIndex('guide/authentication/index.html');
  const home = searchIndex('index.html');
  assert.equal(guide.url, home.url);
  assert.match(guide.url, /offline-search-index-documentation/);
  assert.match(guide.html, /placeholder="Search documentation\.\.\."/);
  assert.match(guide.html, /aria-label="Search documentation\.\.\."/);
  assert.ok(guide.documents.length > 0);
  assert.ok(guide.documents.every((doc) => !/\/reference(?:\/|$)/.test(doc.ref)));
  assert.ok(guide.documents.some((doc) => doc.ref.endsWith('/guide/authentication/')));
});

test('reference landing and nested pages use the reference index and label', () => {
  const reference = searchIndex('reference/index.html');
  const nestedPath = readdirSync(join(publicDir, 'reference'), { recursive: true })
    .find((path) => path !== 'index.html' && path.endsWith('/index.html'));
  assert.ok(nestedPath, 'No nested reference pages');
  const nested = searchIndex(join('reference', nestedPath));
  assert.equal(reference.url, nested.url);
  assert.match(reference.url, /offline-search-index-reference/);
  assert.match(nested.html, /placeholder="Search API reference\.\.\."/);
  assert.match(nested.html, /aria-label="Search API reference\.\.\."/);
  assert.ok(reference.documents.length > 0);
  assert.ok(reference.documents.every((doc) => /\/reference(?:\/|$)/.test(doc.ref)));
  assert.ok(reference.documents.some((doc) => doc.ref.endsWith('/reference/')));
});

test('documentation and reference indexes do not overlap or contain duplicates', () => {
  const guide = searchIndex('guide/authentication/index.html');
  const reference = searchIndex('reference/index.html');
  assert.notEqual(guide.url, reference.url);
  const documents = [...guide.documents, ...reference.documents];
  assert.equal(new Set(documents.map((doc) => doc.ref)).size, documents.length);
});
