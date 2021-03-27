import { SimpleGrid, Text } from '@chakra-ui/react';
import Result from '@components/organisms/Result';
import { EntityId } from '@reduxjs/toolkit';

type Props = {
  ids: EntityId[];
};

export default function List({ ids }: Props) {
  if (ids.length === 0) {
    return <Text textAlign="center">Could not find any...</Text>;
  }

  return (
    <SimpleGrid width="100%" columns={1} spacing={3}>
      {ids.map((id) => (
        <Result id={id} key={id} />
      ))}
    </SimpleGrid>
  );
}
