import Result from '@components/molecules/Result';
import { searchSelector } from '@store/search/slice';
import { Box } from 'grommet';
import React from 'react';
import { useSelector } from 'react-redux';

export default function List() {
  const positions = useSelector((s) => searchSelector.selectIds(s));
  const results = useSelector((s) => searchSelector.selectEntities(s));

  return (
    <Box pad="medium">
      {positions.map((position) => (
        <Result data={results[position]} key={position} />
      ))}
    </Box>
  );
}
