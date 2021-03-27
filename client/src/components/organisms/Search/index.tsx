import { Search2Icon } from '@chakra-ui/icons';
import { Input, InputGroup, InputLeftElement, Stack } from '@chakra-ui/react';
import { selectQueryText, updateQuery } from '@store/query/slice';
import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';

export default function Search() {
  const query: string = useSelector(selectQueryText);

  const dispatch = useDispatch();
  const handleQueryChanged = useCallback((e) => dispatch(updateQuery(e.target.value)), []);

  return (
    <Stack width="100%" spacing={4}>
      <InputGroup>
        <InputLeftElement pointerEvents="none" children={<Search2Icon color="brand" />} />
        <Input
          borderColor="gray.500"
          borderRadius="full"
          placeholder="Search for domain..."
          onChange={handleQueryChanged}
          value={query}
          type="text"
          aria-label="Search for domain:"
        />
      </InputGroup>
    </Stack>
  );
}
