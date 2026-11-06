(function ($) {
  'use strict';

  $(document).ready(function () {
    const $searchInputs = $('.td-search--offline input');
    let worker;
    let requestID = 0;
    let activeSearch;

    function dispose($input) {
      bootstrap.Popover.getInstance($input[0])?.dispose();
      $input.off('shown.bs.popover.asoSearch');
    }

    function clear() {
      activeSearch = undefined;
      $searchInputs.each(function () {
        dispose($(this));
      });
    }

    function show($input, query, results, message) {
      dispose($input);
      const $html = $('<div>');
      $html.append(
        $('<div>')
          .css({ display: 'flex', justifyContent: 'space-between', marginBottom: '1em' })
          .append($('<span>').text('Search results').css({ fontWeight: 'bold' }))
          .append($('<span>').addClass('td-offline-search-results__close-button'))
      );
      const $body = $('<div>').css({
        maxHeight: `calc(100vh - ${$input.offset().top - $(window).scrollTop() + 180}px)`,
        overflowY: 'auto',
      });
      $html.append($body);

      if (message || results.length === 0) {
        $body.append(
          $('<p>').attr('role', 'status').text(message || `No results found for query "${query}"`)
        );
      } else {
        results.forEach((result) => {
          const href = $input.data('offline-search-base-href') + result.ref.replace(/^\//, '');
          $body.append(
            $('<div>').addClass('mt-4')
              .append($('<small>').addClass('d-block text-body-secondary').text(result.ref))
              .append($('<a>').addClass('d-block').css({ fontSize: '1.2rem' })
                .attr('href', href).text(result.title))
              .append($('<p>').text(result.excerpt))
          );
        });
      }

      $input.one('shown.bs.popover.asoSearch', () => {
        $('.td-offline-search-results__close-button').on('click', () => {
          $input.val('');
          clear();
        });
      });
      new bootstrap.Popover($input[0], {
        content: $html[0],
        html: true,
        customClass: 'td-offline-search-results',
        placement: 'bottom',
      }).show();
    }

    function fail(error) {
      console.error('Offline search failed:', error);
      if (activeSearch) {
        show(activeSearch.$input, activeSearch.query, [],
          'Search failed to load. Please try again.');
      }
    }

    function getWorker() {
      if (!worker) {
        const config = JSON.parse(document.getElementById('aso-search-config').textContent);
        worker = new Worker(new URL(config.workerURL, document.baseURI));
        worker.onmessage = ({ data }) => {
          if (!activeSearch || data.id !== activeSearch.id) {
            return;
          }
          if (data.error) {
            fail(data.error);
          } else {
            show(activeSearch.$input, activeSearch.query, data.results);
          }
        };
        worker.onerror = (event) => {
          worker.terminate();
          worker = undefined;
          fail(event.message);
        };
      }
      return worker;
    }

    // Discard pending results as soon as the user edits or clears a query.
    $searchInputs.on('input', clear);
    $searchInputs.on('change', (event) => {
      clear();
      const $input = $(event.target);
      const query = $input.val().trim();
      if (!query) {
        return;
      }
      activeSearch = { id: ++requestID, $input, query };
      show($input, query, [], 'Preparing search...');
      try {
        getWorker().postMessage({
          id: activeSearch.id,
          query,
          indexURL: new URL($input.data('offline-search-index-json-src'), document.baseURI).href,
          maxResults: $input.data('offline-search-max-results'),
        });
      } catch (error) {
        fail(error);
      }
      $input.blur();
    });
    $searchInputs.closest('form').on('submit', () => false);
  });
})(jQuery);
