import { Box, Container } from '@chakra-ui/react';
import Wrapper from '@components/molecules/Wrapper';
import Header from '@components/organisms/Header';
import List from '@components/organisms/List';
import Search from '@components/organisms/Search';
import { domainsSelector } from '@store/domains/slice';
import Head from 'next/head';
import { useSelector } from 'react-redux';

export default function HomePage({ cookies }: any) {
  const hasResults = useSelector(domainsSelector.selectTotal) > 0;

  return (
    <Wrapper cookies={cookies}>
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
    </Wrapper>
  );
}

export { getServerSideProps } from '@components/molecules/Wrapper';
