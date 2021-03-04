import {
  Box,
  Table,
  Tbody,
  Thead,
  Tr,
  Th,
  Td,
  Button,
  Modal,
  ModalHeader,
  ModalOverlay,
  ModalContent,
  ModalCloseButton,
  ModalBody,
  useDisclosure,
} from "@chakra-ui/react";
const CRUDTable = () => {
  return (
    <Box>
      <CRUDButton
        handleModal={handleModal}
        handleDelete={handleDelete}
        item={item}
      />
      <Table>
        <Thead>
          <Tr>
            <Th>Nombre</Th>
            <Th>Direccion</Th>
            <Th>Numero</Th>
            <Th>CP</Th>
            <Th>RFC</Th>
          </Tr>
        </Thead>
        <Tbody>
          {restaurants.map((restaurant) => (
            <Tr
              key={restaurant.ID}
              onClick={() => handleItemSelect(restaurant.ID)}
            >
              <Td>{restaurant.Name}</Td>
              <Td>{restaurant.Address}</Td>
              <Td>{restaurant.Number}</Td>
              <Td>{restaurant.CP}</Td>
              <Td>{restaurant.RFC}</Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
      <ModalCreate
        onClose={onClose}
        isOpen={isOpen}
        type={modalType}
        data={
          item.selected && modalType === types.modify
            ? restaurants.find((r) => r.ID === item.id)
            : null
        }
        load={loadRestaurants}
      />
    </Box>
  );
};
export default CRUDTable;
