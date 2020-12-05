import { Anchor, Box } from 'grommet';

export default function Result({ data }) {
  return (
    <Box
      as={Anchor}
      href={`https://www.namecheap.com/cart/addtocart.aspx?producttype=domains&action=register&domainlist=${data.domain}`}
      align="center"
      color="black"
      elevation="small"
      flex={true}
      hoverIndicator={true}
      pad={{ horizontal: 'medium', vertical: 'small' }}
      rel="nofollow"
      target="_blank"
    >
      {data.domain}
    </Box>
  );
}
