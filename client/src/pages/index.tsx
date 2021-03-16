import { Box, Container } from '@chakra-ui/react';
import Wrapper from '@/components/molecules/Wrapper';
import Header from '@/components/organisms/Header';
import Results from '@/components/organisms/Results';
import Search from '@/components/organisms/Search';
import { selectors } from '@/store/domains/slice';
import Head from 'next/head';
import { useSelector } from 'react-redux';

export default function HomePage({ cookies }: any) {
  const hasResults = useSelector(selectors.selectTotal) > 0;

  return (
    <Wrapper cookies={cookies}>
      <Head>
        <title>Domane</title>
      </Head>

      <Container maxWidth="72rem" centerContent>
        <Container maxWidth="xl" centerContent>
          <Box width="100%" marginY={6}>
            <Header />
          </Box>
          <Search />
        </Container>
        {hasResults && <Results />}
      </Container>
    </Wrapper>
  );
}

export { getServerSideProps } from '@/components/molecules/Wrapper';
