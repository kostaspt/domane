import { updateQuery } from '@store/search/slice';
import { Box, TextInput } from 'grommet';
import { Search as SearchIcon } from 'grommet-icons';
import { useCallback, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

export default function Search() {
  const [query, setQuery] = useState('');

  useEffect(() => dispatch(updateQuery(query)), [query]);

  const dispatch = useDispatch();
  const handleQueryChanged = useCallback((event) => setQuery(event.target.value), []);

  return (
    <Box
      align="center"
      border={{ side: 'all' }}
      direction="row"
      pad={{ horizontal: 'small', vertical: 'xsmall' }}
      round="9999px"
      width="large"
    >
      <SearchIcon color="brand" />
      <TextInput plain placeholder="" value={query} onChange={handleQueryChanged} />
    </Box>
  );
}
