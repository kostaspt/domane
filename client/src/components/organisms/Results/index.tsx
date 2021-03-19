import { SimpleGrid } from '@chakra-ui/react';
import ListColumn from '@components/molecules/ListColumn';
import ListHeader from '@components/molecules/ListHeader';
import List from '@components/organisms/List';
import { selectIds as domainsExtensionsSelectIds } from '@store/domainsExtensions/slice';
import { selectIds as domainsSimilarsSelectIds } from '@store/domainsSimilars/slice';
import { useSelector } from 'react-redux';

export default function Results() {
  const extensionsIds = useSelector(domainsExtensionsSelectIds);
  const similarsIds = useSelector(domainsSimilarsSelectIds);

  return (
    <SimpleGrid width="100%" columns={{ sm: 1, md: 2 }} spacing={6} paddingTop={6}>
      <ListColumn>
        <ListHeader>Extension</ListHeader>
        <List ids={extensionsIds} />
      </ListColumn>
      <ListColumn>
        <ListHeader>Similar</ListHeader>
        <List ids={similarsIds} />
      </ListColumn>
    </SimpleGrid>
  );
}
