import LogoSvg from '@assets/logo.svg';
import { updateQuery } from '@store/query/slice';
import { Anchor, Header as GrommetHeader } from 'grommet';
import { useCallback } from 'react';
import { useDispatch } from 'react-redux';
import styled from 'styled-components';

const Logo = styled(LogoSvg)`
  fill: ${({ theme }) => theme.global.colors['brand']};
  width: 8rem;
`;

export default function Header() {
  const dispatch = useDispatch();
  const handleClick = useCallback(() => dispatch(updateQuery('')), []);

  return (
    <GrommetHeader>
      <Anchor onClick={handleClick}>
        <Logo />
      </Anchor>
      <Anchor color="dark-1" weight="normal" label="GitHub" href="https://github.com/kostaspt/domane" />
    </GrommetHeader>
  );
}
