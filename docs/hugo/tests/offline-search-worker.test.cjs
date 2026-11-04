const assert = require('node:assert/strict');
const { readFileSync } = require('node:fs');
const { test } = require('node:test');
const { runInNewContext } = require('node:vm');

const source = readFileSync(
  `${__dirname}/../assets/js/offline-search-worker.js`, 'utf8'
);

function createWorker(responses = [], scriptResponses = [], importError) {
  const documents = [
    { ref: '/guide/authentication/', title: 'Authentication', excerpt: 'Credentials' },
    { ref: '/reference/managedidentity/', title: 'Managed Identity', excerpt: 'API reference' },
  ];
  const state = {
    fetches: 0, imports: 0, builds: 0, messages: [], terms: [],
    scriptFetches: [], scriptBlobs: [], revokedURLs: [],
  };
  const lunr = function (build) {
    state.builds++;
    build.call({ ref() {}, field() {}, add() {} });
    return {
      query(callback) {
        callback({ term: (term, options) => state.terms.push({ term, ...options }) });
        return documents.map(({ ref }) => ({ ref }));
      },
    };
  };
  lunr.tokenizer = (query) => query.split(/\s+/);
  lunr.Query = { wildcard: { LEADING: 1, TRAILING: 2 } };
  const self = { postMessage: (message) => state.messages.push(message) };
  runInNewContext(source, {
    self,
    lunr,
    Blob,
    URL: {
      createObjectURL(blob) {
        state.scriptBlobs.push(blob);
        return 'blob:verified-lunr';
      },
      revokeObjectURL: (url) => state.revokedURLs.push(url),
    },
    importScripts: (url) => {
      assert.equal(url, 'blob:verified-lunr');
      state.imports++;
      if (importError) {
        throw importError;
      }
    },
    fetch: async (url, options) => {
      if (url === 'https://unpkg.com/lunr@2.3.9/lunr.min.js') {
        state.scriptFetches.push({ url, options });
        const response = scriptResponses.shift();
        if (response instanceof Error) {
          throw response;
        }
        return response || { ok: true, text: async () => 'verified Lunr source' };
      }
      state.fetches++;
      const response = responses.shift();
      if (response instanceof Error) {
        throw response;
      }
      return response || { ok: true, json: async () => documents };
    },
  });
  return {
    state,
    search: (id, maxResults = 10) => self.onmessage({
      data: { id, query: 'IDENTITY', indexURL: 'https://example.test/index.json', maxResults },
    }),
  };
}

test('does not fetch or build an index until a query arrives', () => {
  const { state } = createWorker();
  assert.equal(state.imports, 0);
  assert.equal(state.scriptFetches.length, 0);
  assert.equal(state.fetches, 0);
  assert.equal(state.builds, 0);
});

test('shares initialization across queries and returns bounded result details', async () => {
  const { state, search } = createWorker();
  await Promise.all([search(1, 1), search(2)]);
  assert.equal(state.fetches, 1);
  assert.equal(state.scriptFetches.length, 1);
  assert.equal(state.imports, 1);
  assert.equal(state.builds, 1);
  assert.equal(state.messages[0].id, 1);
  assert.equal(state.messages[0].results.length, 1);
  assert.equal(state.messages[0].results[0].title, 'Authentication');
  assert.equal(state.messages[1].results[1].excerpt, 'API reference');
  assert.equal(state.terms[0].term, 'identity');
  assert.equal(state.terms[0].boost, 100);
  assert.equal(state.terms[1].wildcard, 3);
  assert.equal(state.terms[2].editDistance, 2);
  await search(3);
  assert.equal(state.fetches, 1);
  assert.equal(state.builds, 1);
});

test('requires the pinned integrity digest before importing Lunr from a blob', async () => {
  const { state, search } = createWorker();
  await search(1);
  assert.equal(state.scriptFetches[0].options.integrity,
    'sha384-203J0SNzyqHby3iU6hzvzltrWi/M41wOP5Gu+BiJMz5nwKykbkUx8Kp7iti0Lpli');
  assert.equal(state.scriptBlobs[0].type, 'text/javascript');
  assert.equal(await state.scriptBlobs[0].text(), 'verified Lunr source');
  assert.deepEqual(state.revokedURLs, ['blob:verified-lunr']);
});

test('does not execute Lunr or fetch the index when integrity verification fails, and retries', async () => {
  const { state, search } = createWorker([], [new TypeError('Integrity mismatch')]);
  await search(1);
  assert.equal(state.messages[0].error, 'Integrity mismatch');
  assert.equal(state.messages[0].results, undefined);
  assert.equal(state.imports, 0);
  assert.equal(state.scriptBlobs.length, 0);
  assert.equal(state.fetches, 0);
  assert.equal(state.builds, 0);
  await search(2);
  assert.equal(state.scriptFetches.length, 2);
  assert.equal(state.imports, 1);
  assert.equal(state.messages[1].results.length, 2);
});

test('reports Lunr HTTP failures without executing the response', async () => {
  const { state, search } = createWorker([], [{ ok: false, status: 503 }]);
  await search(1);
  assert.equal(state.messages[0].error, 'Lunr request failed: HTTP 503');
  assert.equal(state.imports, 0);
  assert.equal(state.fetches, 0);
});

test('revokes the script blob URL even if importing Lunr fails', async () => {
  const { state, search } = createWorker([], [], new Error('Import failed'));
  await search(1);
  assert.equal(state.messages[0].error, 'Import failed');
  assert.deepEqual(state.revokedURLs, ['blob:verified-lunr']);
  assert.equal(state.fetches, 0);
});

test('reports HTTP failures and retries on a later query', async () => {
  const { state, search } = createWorker([{ ok: false, status: 503 }]);
  await search(1);
  assert.equal(state.messages[0].id, 1);
  assert.match(state.messages[0].error, /HTTP 503/);
  await search(2);
  assert.equal(state.fetches, 2);
  assert.equal(state.messages[1].results.length, 2);
});

test('reports network failures instead of returning empty results', async () => {
  const { state, search } = createWorker([new Error('Network unavailable')]);
  await search(1);
  assert.equal(state.messages[0].error, 'Network unavailable');
  assert.equal(state.messages[0].results, undefined);
});
