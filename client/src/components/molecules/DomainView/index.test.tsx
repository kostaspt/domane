import { render, screen } from '@testing-library/react';
import DomainView from './index';

it('can present a taken domain', () => {
  render(<DomainView domain="test.com" />);

  expect(screen.getByTestId('domain-view-text')).toHaveTextContent(['😞', 'test.com', 'is taken.'].join(''));
});
