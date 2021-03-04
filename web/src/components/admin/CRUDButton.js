import { Flex, Button } from "@chakra-ui/react";
const types = {
  create: "create",
  modify: "modify",
};
const CRUDButton = ({ handleModal, handleDelete, item }) => {
  return (
    <Flex>
      <Button onClick={() => handleModal(types.create)}>Crear</Button>
      <Button
        onClick={() => handleModal(types.modify)}
        disabled={!item.selected}
      >
        Modificar
      </Button>
      <Button onClick={() => handleDelete()} disabled={!item.selected}>
        Borrar
      </Button>
    </Flex>
  );
};
export default CRUDButton;
