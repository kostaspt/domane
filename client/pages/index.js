import Header from '@components/organisms/Header';
import List from '@components/organisms/List';
import { Main } from 'grommet';
import Head from 'next/head';

export default function HomePage() {
  return (
    <div className="container">
      <Head>
        <title>Domane</title>
      </Head>

      <Main>
        <Header />
        <List />
      </Main>
    </div>
  );
}
