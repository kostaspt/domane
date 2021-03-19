import { Heading } from '@chakra-ui/react';
import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

export default function ListHeader({ children }: Props) {
  return <Heading as="h3" size="lg" textAlign="center" isTruncated children={children} data-testid="list-header" />;
}
