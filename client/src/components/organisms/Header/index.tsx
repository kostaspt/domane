import { MoonIcon, SunIcon } from '@chakra-ui/icons';
import { Box, Link, SimpleGrid, useColorMode, useColorModeValue } from '@chakra-ui/react';
import GitHubIcon from '@components/atoms/GitHubIcon';
import Logo from '@components/atoms/Logo';
import styled from '@emotion/styled';
import { updateQuery } from '@store/query/slice';
import { useCallback } from 'react';
import { useDispatch } from 'react-redux';

const StyledLogo = styled(Logo)`
  color: ${({ theme }: any) => theme.colors.brand};
`;

export default function Header() {
  const { colorMode, toggleColorMode } = useColorMode();

  const linkColor = useColorModeValue('gray.900', 'gray.200');
  const linkColorHover = useColorModeValue('gray.700', 'gray.400');

  const dispatch = useDispatch();
  const handleClick = useCallback(() => dispatch(updateQuery('')), []);

  return (
    <Box as="header" display="flex" justifyContent="space-between">
      <Link onClick={handleClick}>
        <StyledLogo height="auto" width={32} />
      </Link>
      <SimpleGrid columns={2} spacing={3}>
        <Link color={linkColor} href="https://github.com/kostaspt/domane" target="_blank">
          <GitHubIcon _hover={{ color: linkColorHover }} />
        </Link>
        <Link as="div" onClick={toggleColorMode}>
          {colorMode === 'light' ? (
            <SunIcon _hover={{ color: linkColorHover }} />
          ) : (
            <MoonIcon _hover={{ color: linkColorHover }} />
          )}
        </Link>
      </SimpleGrid>
    </Box>
  );
}
