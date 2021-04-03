import { SimpleGrid } from '@chakra-ui/react';
import ListColumn from '@components/molecules/ListColumn';
import ListHeader from '@components/molecules/ListHeader';
import List from '@components/organisms/List';
import {
  selectIds as domainsExtensionsSelectIds,
  selectStatus as domainsExtensionsSelectStatus,
} from '@store/domainsExtensions/slice';
import {
  selectIds as domainsSimilarsSelectIds,
  selectStatus as domainsSimilarsSelectStatus,
} from '@store/domainsSimilars/slice';
import { STATUS_LOADING } from '@store/status';
import { useSelector } from 'react-redux';

export default function Results() {
  const extensionsIds = useSelector(domainsExtensionsSelectIds);
  const extensionsStatus = useSelector(domainsExtensionsSelectStatus);
  const similarsIds = useSelector(domainsSimilarsSelectIds);
  const similarsStatus = useSelector(domainsSimilarsSelectStatus);

  return (
    <SimpleGrid width="100%" columns={{ sm: 1, md: 2 }} spacing={6} marginY={16}>
      <ListColumn>
        <ListHeader isLoading={extensionsStatus === STATUS_LOADING} count={extensionsIds.length}>
          Extension
        </ListHeader>
        <List ids={extensionsIds} />
      </ListColumn>
      <ListColumn>
        <ListHeader isLoading={similarsStatus === STATUS_LOADING} count={similarsIds.length}>
          Similar
        </ListHeader>
        <List ids={similarsIds} />
      </ListColumn>
    </SimpleGrid>
  );
}
