import Result from '@components/molecules/Result';
import { domainsSelector } from '@store/domains/slice';
import { Box } from 'grommet';
import React from 'react';
import { useSelector } from 'react-redux';

export default function List() {
  const positions = useSelector((s) => domainsSelector.selectIds(s));
  const results = useSelector((s) => domainsSelector.selectEntities(s));

  return (
    <Box pad="medium">
      {positions.map((position) => (
        <Result data={results[position]} key={position} />
      ))}
    </Box>
  );
}
