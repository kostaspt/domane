import {
  Heading,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  SimpleGrid,
} from '@chakra-ui/react';
import Registrar from '@components/molecules/Registrar';

type DomainModalProps = {
  domain: string;
  isOpen: boolean;
  handleClose: () => void;
};

export default function DomainModal({ domain, isOpen, handleClose }: DomainModalProps) {
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
          <Heading as="h6" size="xs" marginBottom={4} textAlign="center">
            Buy now
          </Heading>
          <SimpleGrid columns={3} spacing={4}>
            {['namecheap', 'godaddy', 'dynadot'].map((name) => (
              <Registrar name={name} domain={domain} key={name} />
            ))}
          </SimpleGrid>
        </ModalBody>
        <ModalFooter />
      </ModalContent>
    </Modal>
  );
}
