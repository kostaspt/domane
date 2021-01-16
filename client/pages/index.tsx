import { Box, Container } from '@chakra-ui/react';
import Header from '@components/organisms/Header';
import List from '@components/organisms/List';
import Search from '@components/organisms/Search';
import { domainsSelector } from '@store/domains/slice';
import Head from 'next/head';
import { useSelector } from 'react-redux';

export default function HomePage() {
  const hasResults = useSelector(domainsSelector.selectTotal) > 0;

  return (
    <div className="container">
      <Head>
        <title>Domane</title>
      </Head>

      <Container maxW="xl" centerContent>
        <Box width="100%" marginY={6}>
          <Header />
        </Box>
        <Search />
        {hasResults && <List />}
      </Container>
    </div>
  );
}
