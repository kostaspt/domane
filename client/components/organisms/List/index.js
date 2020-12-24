import Result from '@components/molecules/Result';
import { availabilitiesSelector } from '@store/availabilities/slice';
import { domainsSelector } from '@store/domains/slice';
import { Box } from 'grommet';
import React from 'react';
import { useSelector } from 'react-redux';

export default function List() {
  const ids = useSelector(domainsSelector.selectIds);
  const domains = useSelector(domainsSelector.selectEntities);
  const availabilities = useSelector(availabilitiesSelector.selectEntities);

  return (
    <Box pad="medium">
      {ids.map((id) => (
        <Result domain={domains[id].domain} isAvailable={availabilities[domains[id].domain]?.isAvailable} key={id} />
      ))}
    </Box>
  );
}
