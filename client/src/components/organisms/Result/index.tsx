import { useDisclosure } from '@chakra-ui/hooks';
import DomainCard from '@components/molecules/DomainCard';
import DomainModal from '@components/molecules/DomainModal';
import { RootState } from '@store';
import { selectors as availabilitiesSelectors } from '@store/availabilities/slice';
import { selectors as domainsSelectors } from '@store/domains/slice';
import { useSelector } from 'react-redux';

const dummyDomain = { domain: '' };

const dummyAvailability = {
  isAvailable: undefined,
};

type ResultProps = {
  id: string | number;
};

export default function Result({ id }: ResultProps) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { domain } = useSelector((s: RootState) => domainsSelectors.selectById(s, id)) ?? dummyDomain;
  const { isAvailable } =
    useSelector((s: RootState) => availabilitiesSelectors.selectById(s, domain ?? '')) ?? dummyAvailability;

  if (domain.length === 0) {
    return null;
  }

  return (
    <>
      <DomainCard domain={domain} isAvailable={isAvailable} handleClick={onOpen} />
      <DomainModal domain={domain} isAvailable={isAvailable} isOpen={isOpen} handleClose={onClose} />
    </>
  );
}
