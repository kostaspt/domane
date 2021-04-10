import { render, screen } from '@testing-library/react';
import DomainBuy from './index';

it('can present available domain', () => {
  render(<DomainBuy domain="test.com" isAvailable={true} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent('🎉test.com is available! You can buy it via:');
});

it('can present undefined domain', () => {
  render(<DomainBuy domain="test.com" isAvailable={undefined} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent('🤔test.com could be available. Check it using:');
});

it('can present null domain', () => {
  render(<DomainBuy domain="test.com" isAvailable={null} />);

  expect(screen.getByTestId('domain-buy-text')).toHaveTextContent('🤔test.com could be available. Check it using:');
});
