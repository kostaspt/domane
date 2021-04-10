import { render, screen } from '@testing-library/react';
import DomainBuy from './index';

it('can present an available domain', () => {
  render(<DomainBuy domain="test.com" isAvailable={true} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent(
    ['🎉', 'test.com is available! You can buy it via:'].join('')
  );
});

it('can present a domain with undefined availability', () => {
  render(<DomainBuy domain="test.com" isAvailable={undefined} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent(
    ['🤔', 'test.com could be available. Check it using:'].join('')
  );
});

it('can present a domain with null availability', () => {
  render(<DomainBuy domain="test.com" isAvailable={null} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent(
    ['🤔', 'test.com could be available. Check it using:'].join('')
  );
});
