import { renderWithStore } from '@test/utils';
import { screen } from '@testing-library/react';
import Results from './index';

it('should display a list of results', () => {
  renderWithStore(<Results />, {
    initialState: {
      availabilities: {
        ids: ['test.com', 'test.net', 'try.com'],
        entities: {
          'test.com': {
            domain: 'test.com',
            isAvailable: true,
          },
          'test.net': {
            domain: 'test.net',
            isAvailable: null,
          },
          'try.com': {
            domain: 'try.com',
            isAvailable: false,
          },
        },
      },
      domainsExtensions: {
        ids: ['0-test.com', '1-test.net'],
      },
      domainsSimilars: {
        ids: ['0-try.com'],
      },
      domains: {
        ids: ['0-test.com', '1-test.net', '0-try.com'],
        entities: {
          '0-test.com': {
            domain: 'test.com',
            kind: 'extensions',
            position: 0,
          },
          '1-test.net': {
            domain: 'test.net',
            kind: 'extensions',
            position: 1,
          },
          '0-try.com': {
            domain: 'try.com',
            kind: 'similars',
            position: 0,
          },
        },
      },
    },
  });

  expect(screen.getAllByTestId('list-header')[0]).toHaveTextContent('Extension');
  expect(screen.getAllByTestId('list-header')[1]).toHaveTextContent('Similar');

  expect(screen.getAllByTestId('domain-name')[0]).toHaveTextContent('test.com');
  expect(screen.getAllByTestId('domain-availability')[0]).toHaveTextContent('Available');
  expect(screen.getAllByTestId('domain-name')[1]).toHaveTextContent('test.net');
  expect(screen.getAllByTestId('domain-availability')[1]).toHaveTextContent('Unknown');
  expect(screen.getAllByTestId('domain-name')[2]).toHaveTextContent('try.com');
  expect(screen.getAllByTestId('domain-availability')[2]).toHaveTextContent('Taken');
});
