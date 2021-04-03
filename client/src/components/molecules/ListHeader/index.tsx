import { Badge, Flex, Heading, Spinner, Box } from '@chakra-ui/react';
import { ReactNode } from 'react';

type Props = {
  count?: number;
  isLoading?: boolean;
  children: ReactNode;
};

export default function ListHeader({ isLoading = false, count = 0, children }: Props) {
  return (
    <Flex marginBottom={6}>
      <Heading as="h3" size="md" textAlign="center" isTruncated data-testid="list-header" marginRight={2}>
        {children}
      </Heading>

      <Flex alignItems="center">
        <Badge>
          <Box paddingY={0.5} paddingX={1} lineHeight={1}>
            {isLoading ? <Spinner size="xs" /> : count}
          </Box>
        </Badge>
      </Flex>
    </Flex>
  );
}
