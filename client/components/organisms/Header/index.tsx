import { Box, Link } from '@chakra-ui/react';
import Logo from '@components/atoms/Logo';
import styled from '@emotion/styled';
import { updateQuery } from '@store/query/slice';
import { useCallback } from 'react';
import { useDispatch } from 'react-redux';

const StyledLogo = styled(Logo)`
  color: ${({ theme }: any) => theme.colors.brand[500]};
`;

export default function Header() {
  const dispatch = useDispatch();
  const handleClick = useCallback(() => dispatch(updateQuery('')), []);

  return (
    <Box display="flex" justifyContent="space-between">
      <Link onClick={handleClick}>
        <StyledLogo height="auto" width={32} />
      </Link>
      <Link color="black" href="https://github.com/kostaspt/domane">
        GitHub
      </Link>
    </Box>
  );
}
