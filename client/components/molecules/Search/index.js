import { Box, TextInput } from 'grommet';
import { Search as SearchIcon } from 'grommet-icons';
import { useCallback, useEffect, useState } from 'react';

export default function Search({ handleQueryChanged }) {
  const [query, setQuery] = useState('');

  const handleChange = useCallback((event) => setQuery(event.target.value), []);

  useEffect(() => handleQueryChanged(query), [query]);

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
      <TextInput plain placeholder="" value={query} onChange={handleChange} />
    </Box>
  );
}
