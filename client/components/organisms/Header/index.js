import Search from '@components/molecules/Search';
import { updateQuery } from '@store/search/slice';
import { Anchor, Box, Heading } from 'grommet';
import Link from 'next/link';
import { useCallback } from 'react';
import { useDispatch } from 'react-redux';

export default function Header() {
  const dispatch = useDispatch();

  const handleQueryChanged = useCallback((query) => dispatch(updateQuery(query), 500), []);

  return (
    <Box align="center" pad={{ bottom: 'large' }}>
      <Link href="/">
        <Heading level={1} color="brand">
          <Anchor>Domane</Anchor>
        </Heading>
      </Link>
      <Search handleQueryChanged={handleQueryChanged} />
    </Box>
  );
}
