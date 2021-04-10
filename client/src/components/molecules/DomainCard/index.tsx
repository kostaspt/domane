import { Box, Grid, GridItem, useColorMode, useColorModeValue } from '@chakra-ui/react';
import { useMemo } from 'react';

type DomainCardProps = {
  domain: string;
  isAvailable?: boolean;
  handleClick: () => void;
};

export default function DomainCard({ domain, isAvailable, handleClick }: DomainCardProps) {
  const { colorMode } = useColorMode();

  const backgroundColor = useColorModeValue('white', 'gray.700');
  const backgroundColorHover = useColorModeValue('gray.100', 'gray.600');
  const borderColor = useColorModeValue('gray.200', 'gray.900');
  const textColor = useColorModeValue('gray.900', 'gray.200');

  const [labelColor, labelText] = useMemo(() => {
    switch (isAvailable) {
      case true:
        return ['green', 'Available'];
      case false:
        return ['red', 'Taken'];
      default:
        return ['gray', 'Unknown'];
    }
  }, [isAvailable]);

  return (
    <Box
      background={backgroundColor}
      borderColor={borderColor}
      borderWidth={1}
      borderRadius="base"
      color={textColor}
      display="flex"
      flexDirection="row"
      justifyContent="center"
      paddingX={4}
      paddingY={2}
      shadow={colorMode ? 'md' : 'none'}
      _hover={{ bg: backgroundColorHover, cursor: 'pointer' }}
      onClick={handleClick}
    >
      <Grid templateColumns="3fr 1fr" gap={6} width="100%">
        <GridItem display="flex" alignItems="center" data-testid="domain-name">
          {domain}
        </GridItem>
        <GridItem display="flex" justifyContent="flex-end">
          <Box
            background={`${labelColor}.200`}
            borderRadius="base"
            display="inline-flex"
            color={`${labelColor}.800`}
            paddingX={2}
            paddingY={1}
            data-testid="domain-availability"
          >
            {labelText}
          </Box>
        </GridItem>
      </Grid>
    </Box>
  );
}
