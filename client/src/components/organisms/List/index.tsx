import { SimpleGrid } from '@chakra-ui/react';
import Result from '@components/molecules/Result';
import { EntityId } from '@reduxjs/toolkit';
import { selectors as availabilitiesSelectors } from '@store/availabilities/slice';
import { selectors as domainsSelectors } from '@store/domains/slice';
import React from 'react';
import { useSelector } from 'react-redux';

type Props = {
  ids: EntityId[];
};

export default function List({ ids }: Props) {
  const domains = useSelector(domainsSelectors.selectEntities);
  const availabilities = useSelector(availabilitiesSelectors.selectEntities);

  return (
    <SimpleGrid width="100%" maxWidth="md" paddingY={6} columns={1} spacing={3}>
      {ids.map((id) => {
        if (!domains[id]) {
          return null;
        }

        const { domain } = domains[id] ?? {};

        if (!domain) {
          return null;
        }

        return <Result domain={domain} isAvailable={availabilities[domain]?.isAvailable} key={id} />;
      })}
    </SimpleGrid>
  );
}
