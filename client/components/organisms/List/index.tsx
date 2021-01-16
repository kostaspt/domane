import Result from '@components/molecules/Result';
import { availabilitiesSelector } from '@store/availabilities/slice';
import { domainsSelector } from '@store/domains/slice';
import { SimpleGrid } from '@chakra-ui/react';
import React from 'react';
import { useSelector } from 'react-redux';

export default function List() {
  const ids = useSelector(domainsSelector.selectIds);
  const domains = useSelector(domainsSelector.selectEntities);
  const availabilities = useSelector(availabilitiesSelector.selectEntities);

  return (
    <SimpleGrid width="100%" paddingY={6} columns={1} spacing={3}>
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
