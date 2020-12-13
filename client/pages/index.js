import Header from '@components/organisms/Header';
import List from '@components/organisms/List';
import Search from '@components/organisms/Search';
import { domainsSelector } from '@store/domains/slice';
import { Box, Main } from 'grommet';
import Head from 'next/head';
import { useSelector } from 'react-redux';

export default function HomePage() {
  const hasResults = useSelector((s) => domainsSelector.selectTotal(s)) > 0;

  return (
    <div className="container">
      <Head>
        <title>Domane</title>
      </Head>

      <Main alignContent="center" width="600px" margin={{ vertical: 0, horizontal: 'auto' }}>
        <Box margin={{ vertical: 'medium' }}>
          <Header />
        </Box>
        <Search />
        {hasResults && <List />}
      </Main>
    </div>
  );
}
