import { Anchor, Box } from 'grommet';
import { useMemo } from 'react';

export default function Result({ domain, isAvailable }) {
  const labelColor = useMemo(() => {
    switch (isAvailable) {
      case true:
        return 'status-ok';
      case false:
        return 'status-error';
      default:
        return 'status-unknown';
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
    <Box
      as={Anchor}
      color="black"
      direction="row"
      elevation="small"
      flex={true}
      gap="small"
      hoverIndicator={true}
      href={`https://www.namecheap.com/cart/addtocart.aspx?producttype=domains&action=register&domainlist=${domain}`}
      justify="center"
      pad={{ horizontal: 'medium', vertical: 'small' }}
      rel="nofollow"
      target="_blank"
    >
      {domain}
      <Box background={labelColor} pad={{ horizontal: 'xsmall', vertical: 'xxsmall' }} round="xsmall">
        {labelText}
      </Box>
    </Box>
  );
}
