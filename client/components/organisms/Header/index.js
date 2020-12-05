import { Anchor, Header as GrommetHeader, Heading } from 'grommet';
import Link from 'next/link';

export default function Header() {
  return (
    <GrommetHeader>
      <Link href="/">
        <Heading as="h1" size="small" color="brand">
          <Anchor label="Domane">Domane</Anchor>
        </Heading>
      </Link>
      <Anchor color="dark-1" weight="normal" label="GitHub" href="https://github.com/kostaspt/domane" />
    </GrommetHeader>
  );
}
