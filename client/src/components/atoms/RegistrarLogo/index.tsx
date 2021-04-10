import theme from '@/theme';
import DynadotLogo from '@assets/logos/dynadot.svg';
import GoDaddyLogo from '@assets/logos/godaddy.svg';
import NamecheapLogo from '@assets/logos/namecheap.svg';
import { useColorModeValue } from '@chakra-ui/react';
import { getColor } from '@chakra-ui/theme-tools';
import isPropValid from '@emotion/is-prop-valid';
import styled, { StyledOptions } from '@emotion/styled';

type RegistrarLogoProps = {
  name: string;
  width?: number;
};

const styledOptions: StyledOptions<any> = {
  shouldForwardProp: (prop: PropertyKey) => isPropValid(prop.toString()),
};

const StyledGoDaddyLogo = styled(GoDaddyLogo, styledOptions)`
  .godaddy_svg__text {
    fill: ${({ textColor }) => textColor};
  }
`;

const StyledNamecheapLogo = styled(NamecheapLogo, styledOptions)`
  .namecheap_svg__text {
    fill: ${({ textColor }) => textColor};
  }
`;

const StyledDynadotLogo = styled(DynadotLogo, styledOptions)`
  .dynadot_svg__text {
    fill: ${({ textColor }) => textColor};
  }
`;

export default function RegistrarLogo({ name, width }: RegistrarLogoProps) {
  const textColor = getColor(theme, useColorModeValue('gray.900', 'gray.200'));

  switch (name) {
    case 'godaddy':
      return <StyledGoDaddyLogo textColor={textColor} width={width} />;
    case 'namecheap':
      return <StyledNamecheapLogo textColor={textColor} width={width} />;
    case 'dynadot':
      return <StyledDynadotLogo textColor={textColor} width={width} />;
    default:
      return null;
  }
}
