import {
  Box,
  Table,
  Tbody,
  Thead,
  Tr,
  Th,
  Td,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  useDisclosure,
  Button,
} from "@chakra-ui/react";
import { Formik, Form } from "formik";
import InputField from "../../InputField";
import axios from "axios";
import { useEffect, useState } from "react";
import { HOSTNAME } from "../../../consts";
import CRUDButton from "../CRUDButton";
const types = {
  create: "create",
  modify: "modify",
};
const Roles = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [modalType, setModalType] = useState("");
  const [roles, setRoles] = useState([]);
  const [item, setItem] = useState({ id: null, selected: false });
  useEffect(async () => {
    loadRoles();
  }, []);
  const handleModal = (type) => {
    setModalType(type);
    onOpen();
  };
  const handleDelete = async () => {
    axios.delete(`${HOSTNAME}/roles/${item.id}`).then(() => {
      loadRoles();
      setItem({ ...item, selected: false });
    });
  };
  const handleItemSelect = (id) => {
    if (item.id === id) {
      setItem({ ...item, selected: !item.selected });
    } else {
      setItem({ id: id, selected: true });
    }
  };
  const loadRoles = async () => {
    axios.get(`${HOSTNAME}/roles`).then(({ data }) => {
      setRoles(data);
    });
  };
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
            <Th w="10">Codigo</Th>
            <Th>Nombre</Th>
          </Tr>
        </Thead>
        <Tbody>
          {roles.map((role) => (
            <Tr key={role.Code} onClick={() => handleItemSelect(role.Code)}>
              <Td>{role.Code}</Td>
              <Td>{role.Name}</Td>
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
            ? roles.find((r) => r.Code === item.id)
            : null
        }
        load={loadRoles}
      />
    </Box>
  );
};
const ModalCreate = ({ onClose, type, isOpen, data, load }) => {
  return (
    <Modal size="3xl" onClose={onClose} isOpen={isOpen}>
      <ModalOverlay />
      <ModalContent background="white">
        <ModalHeader>
          {type === types.create
            ? "Crear Rol"
            : data === null
            ? null
            : `Modificar ${data.Name}`}
        </ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <CreateForm data={data} type={type} load={load} onClose={onClose} />
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
const CreateForm = ({ data, type, load, onClose }) => {
  return (
    <Formik
      initialValues={{
        code: data === null ? 0 : data.Code,
        name: data === null ? "" : data.Name,
      }}
      onSubmit={async (values) => {
        if (type === types.create) {
          axios
            .post(`${HOSTNAME}/roles`, values)
            .then(() => {
              load();
              onClose();
            })
            .catch((error) => {
              console.log(error);
            });
        }
        if (type === types.modify) {
          axios
            .put(`${HOSTNAME}/roles`, values)
            .then(() => {
              load();
              onClose();
            })
            .catch((error) => {
              console.log(error);
            });
        }
      }}
    >
      {({ isSubmiting }) => (
        <Form>
          <InputField label="Codigo" type="number" id="code" name="code" />
          <InputField label="Nombre" type="text" id="name" name="name" />
          <Button
            colorScheme="teal"
            type="submit"
            isLoading={isSubmiting}
            mt="10px"
          >
            Guardar
          </Button>
        </Form>
      )}
    </Formik>
  );
};
export default Roles;
