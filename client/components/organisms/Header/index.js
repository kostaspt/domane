import LogoSvg from '@assets/logo.svg';
import { Anchor, Header as GrommetHeader } from 'grommet';
import Link from 'next/link';
import styled from 'styled-components';

const Logo = styled(LogoSvg)`
  fill: ${({ theme }) => theme.global.colors['brand']};
  width: 8rem;
`;

export default function Header() {
  return (
    <GrommetHeader>
      <Link href="/">
        <Logo />
      </Link>
      <Anchor color="dark-1" weight="normal" label="GitHub" href="https://github.com/kostaspt/domane" />
    </GrommetHeader>
  );
}
