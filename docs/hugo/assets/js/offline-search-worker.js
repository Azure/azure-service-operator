'use strict';

let indexPromise;

async function loadIndex(indexURL) {
  if (!indexPromise) {
    indexPromise = (async () => {
      // Match the integrity pin in Docsy's head.html before executing CDN code.
      const scriptResponse = await fetch('https://unpkg.com/lunr@2.3.9/lunr.min.js', {
        integrity: 'sha384-203J0SNzyqHby3iU6hzvzltrWi/M41wOP5Gu+BiJMz5nwKykbkUx8Kp7iti0Lpli',
      });
      if (!scriptResponse.ok) {
        throw new Error(`Lunr request failed: HTTP ${scriptResponse.status}`);
      }
      const scriptURL = URL.createObjectURL(new Blob(
        [await scriptResponse.text()], { type: 'text/javascript' }
      ));
      try {
        importScripts(scriptURL);
      } finally {
        URL.revokeObjectURL(scriptURL);
      }
      const response = await fetch(indexURL);
      if (!response.ok) {
        throw new Error(`Search index request failed: HTTP ${response.status}`);
      }
      const documents = await response.json();
      const details = new Map();
      const index = lunr(function () {
        this.ref('ref');
        this.field('title', { boost: 5 });
        this.field('categories', { boost: 3 });
        this.field('tags', { boost: 3 });
        this.field('description', { boost: 2 });
        this.field('body');

        documents.forEach((doc) => {
          this.add(doc);
          details.set(doc.ref, { title: doc.title, excerpt: doc.excerpt });
        });
      });
      return { index, details };
    })().catch((error) => {
      indexPromise = undefined;
      throw error;
    });
  }
  return indexPromise;
}

self.onmessage = async ({ data: { id, query, indexURL, maxResults } }) => {
  try {
    const { index, details } = await loadIndex(indexURL);
    const matches = index.query((q) => {
      lunr.tokenizer(query.toLowerCase()).forEach((token) => {
        const term = token.toString();
        q.term(term, { boost: 100 });
        q.term(term, {
          wildcard: lunr.Query.wildcard.LEADING | lunr.Query.wildcard.TRAILING,
          boost: 10,
        });
        q.term(term, { editDistance: 2 });
      });
    });
    const results = matches.slice(0, maxResults).map(({ ref }) => ({
      ref,
      ...details.get(ref),
    }));
    self.postMessage({ id, results });
  } catch (error) {
    self.postMessage({ id, error: error.message });
  }
};
