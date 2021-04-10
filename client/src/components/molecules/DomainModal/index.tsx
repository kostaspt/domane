import {
  Heading,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from '@chakra-ui/react';
import DomainBuy from '@components/molecules/DomainBuy';
import DomainView from '@components/molecules/DomainView';

type DomainModalProps = {
  domain: string;
  isAvailable?: boolean;
  isOpen: boolean;
  handleClose: () => void;
};

export default function DomainModal({ domain, isAvailable, isOpen, handleClose }: DomainModalProps) {
  return (
    <Modal isOpen={isOpen} onClose={handleClose} size="xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          <Heading as="h4" size="lg">
            {domain}
          </Heading>
        </ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          {isAvailable === false ? (
            <DomainView domain={domain} />
          ) : (
            <DomainBuy domain={domain} isAvailable={isAvailable} />
          )}
        </ModalBody>
        <ModalFooter />
      </ModalContent>
    </Modal>
  );
}
