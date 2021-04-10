import { HStack, Link } from '@chakra-ui/react';

type DomainViewProps = {
  domain: string;
};

export default function DomainView({ domain }: DomainViewProps) {
  return (
    <HStack spacing={3} data-testid="domain-view-text">
      <span>😞</span>
      <HStack spacing={1}>
        <Link href={`http://${domain}`} color="brand" isExternal>
          {domain}
        </Link>
        <span>is taken.</span>
      </HStack>
    </HStack>
  );
}
