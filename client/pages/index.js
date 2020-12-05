import Header from '@components/organisms/Header';
import List from '@components/organisms/List';
import Search from '@components/organisms/Search';
import { Main } from 'grommet';
import Head from 'next/head';

export default function HomePage() {
  return (
    <div className="container">
      <Head>
        <title>Domane</title>
      </Head>

      <Main alignContent="center" width="600px" margin={{ vertical: 0, horizontal: 'auto' }}>
        <Header />
        <Search />
        <List />
      </Main>
    </div>
  );
}
