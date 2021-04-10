import { HStack, SimpleGrid } from '@chakra-ui/react';
import Registrar from '@components/molecules/Registrar';

type DomainBuyProps = {
  domain: string;
  isAvailable?: boolean | null;
};

export default function DomainBuy({ domain, isAvailable }: DomainBuyProps) {
  const icon = isAvailable === true ? '🎉' : '🤔';
  const text = isAvailable === true ? 'is available! You can buy it via' : 'could be available. Check it using';

  return (
    <>
      <HStack spacing={3} marginBottom={4} data-testid="domain-buy-text">
        <span>{icon}</span>
        <span>
          {domain} {text}:
        </span>
      </HStack>
      <SimpleGrid columns={3} spacing={4}>
        {['namecheap', 'godaddy', 'dynadot'].map((name) => (
          <Registrar name={name} domain={domain} key={name} />
        ))}
      </SimpleGrid>
    </>
  );
}
