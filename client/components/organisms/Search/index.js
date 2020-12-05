import { name as searchName, updateQuery } from '@store/search/slice';
import { Box, TextInput } from 'grommet';
import { Search as SearchIcon } from 'grommet-icons';
import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';

export default function Search() {
  const query = useSelector((s) => s[searchName].query);

  const dispatch = useDispatch();
  const handleQueryChanged = useCallback((event) => dispatch(updateQuery(event.target.value)), []);

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
