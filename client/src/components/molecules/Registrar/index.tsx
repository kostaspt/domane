import { Link, useColorModeValue } from '@chakra-ui/react';
import RegistrarLogo from '@components/atoms/RegistrarLogo';

type Registrar = {
  name: string;
  url: string;
};

const registrarUrls = new Map<string, Registrar>([
  [
    'namecheap',
    {
      name: 'Namecheap',
      url: `https://www.namecheap.com/cart/addtocart.aspx?producttype=domains&action=register&domainlist=%s`,
    },
  ],
  [
    'godaddy',
    {
      name: 'GoDaddy',
      url: `https://www.godaddy.com/domainsearch/find?domainToCheck=%s`,
    },
  ],
  [
    'dynadot',
    {
      name: 'Dynadot',
      url: `https://www.dynadot.com/domain/search.html?domain=%s`,
    },
  ],
]);

type RegistrarProps = {
  name: string;
  domain: string;
};

export default function Registrar({ name, domain }: RegistrarProps) {
  const backgroundColor = useColorModeValue('white', 'gray.700');
  const backgroundColorHover = useColorModeValue('gray.100', 'gray.600');
  const borderColor = useColorModeValue('gray.200', 'gray.900');

  const registrar = registrarUrls.get(name);

  if (!registrar) {
    return null;
  }

  const url = registrar.url.replace('%s', domain);

  if (!url) {
    return null;
  }

  return (
    <Link
      href={url}
      alignItems="center"
      background={backgroundColor}
      border="1px"
      borderColor={borderColor}
      display="flex"
      isExternal
      paddingX={8}
      paddingY={4}
      _hover={{ bg: backgroundColorHover }}
    >
      <RegistrarLogo name={name} width={100} />
    </Link>
  );
}
