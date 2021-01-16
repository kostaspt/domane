import { Box, Link, Grid, GridItem } from '@chakra-ui/react';
import { useMemo } from 'react';

type ResultProps = {
  domain: string;
  isAvailable: boolean | undefined;
};

export default function Result({ domain, isAvailable }: ResultProps) {
  const labelColor = useMemo(() => {
    switch (isAvailable) {
      case true:
        return 'green';
      case false:
        return 'red';
      default:
        return 'gray';
    }
  }, [isAvailable]);

  const labelText = useMemo(() => {
    switch (isAvailable) {
      case true:
        return 'Available';
      case false:
        return 'Unavailable';
      default:
        return 'Unknown';
    }
  }, [isAvailable]);

  return (
    <Link
      href={`https://www.namecheap.com/cart/addtocart.aspx?producttype=domains&action=register&domainlist=${domain}`}
      isExternal
      _hover={{ textDecoration: 'none' }}
    >
      <Box
        borderColor="gray.200"
        borderWidth={1}
        borderRadius="base"
        color="black"
        display="flex"
        flexDirection="row"
        justifyContent="center"
        paddingX={4}
        paddingY={2}
        shadow="md"
        _hover={{ bg: 'gray.100' }}
      >
        <Grid templateColumns="1fr 3fr" gap={6} width="100%">
          <GridItem display="flex" justifyContent="flex-end">
            <Box
              background={`${labelColor}.200`}
              borderRadius="base"
              display="inline-flex"
              color={`${labelColor}.800`}
              paddingX={2}
              paddingY={1}
            >
              {labelText}
            </Box>
          </GridItem>
          <GridItem>{domain}</GridItem>
        </Grid>
      </Box>
    </Link>
  );
}
