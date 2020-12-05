import { searchSelector } from '@store/search/slice';
import Result from '@components/molecules/Result';
import { Grid } from 'grommet';
import React from 'react';
import { useSelector } from 'react-redux';

export default function List() {
  const positions = useSelector((s) => searchSelector.selectIds(s));
  const results = useSelector((s) => searchSelector.selectEntities(s));

  if (searchSelector.selectTotal === 0) {
    return null;
  }

  return (
    <Grid columns={{ count: 'fill', size: 'medium' }} pad="medium" gap="medium">
      {positions.map((position) => (
        <Result data={results[position]} key={position} />
      ))}
    </Grid>
  );
}
