import { Box } from '@chakra-ui/react';
import React, { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

export default function ListColumn({ children }: Props) {
  return <Box display="flex" flexDirection="column" alignItems="center" children={children} />;
}
