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
  Checkbox,
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
const Empleados = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [modalType, setModalType] = useState("");
  const [empleados, setEmpleados] = useState([]);
  const [item, setItem] = useState({ id: null, selected: false });
  useEffect(() => {
    loadEmpleados();
  }, []);
  const handleModal = (type) => {
    setModalType(type);
    onOpen();
  };
  const handleItemSelect = (id) => {
    if (item.id === id) {
      setItem({ ...item, selected: !item.selected });
    } else {
      setItem({ id: id, selected: true });
    }
  };
  const loadEmpleados = async () => {
    axios.get(`${HOSTNAME}/empleados`).then(({ data }) => {
      setEmpleados(data);
    });
  };
  const handleDelete = async () => {
    axios.delete(`${HOSTNAME}/empleados/${item.id}`).then(() => {
      loadRoles();
      setItem({ ...item, selected: false });
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
            <Th>Nombre</Th>
            <Th>RestaurantID</Th>
            <Th>RoleID</Th>
            <Th>Username</Th>
            <Th>Password</Th>
          </Tr>
        </Thead>
        <Tbody>
          {empleados.map((emp) => (
            <Tr key={emp.ID} onClick={() => handleItemSelect(emp.ID)}>
              <Td>{emp.Name}</Td>
              <Td>{emp.RestaurantID}</Td>
              <Td>{emp.RoleID}</Td>
              <Td>{emp.Username}</Td>
              <Td>******</Td>
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
            ? empleados.find((r) => r.ID === item.id)
            : null
        }
        load={loadEmpleados}
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
            ? "Crear Empleados"
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
        ID: data === null ? 0 : data.ID,
        name: data === null ? "" : data.Name,
        RestaurantID: data === null ? "" : data.RestaurantID,
        RoleID: data === null ? "" : data.RoleID,
        username: data === null ? "" : data.Username,
        password: "",
        change: false,
      }}
      onSubmit={async (values) => {
        if (type === types.create) {
          axios
            .post(`${HOSTNAME}/empleados`, values)
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
            .put(`${HOSTNAME}/empleados/${values.change}`, values)
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
          <InputField label="Nombre" type="text" id="name" name="name" />
          <InputField
            label="Restaurant"
            type="number"
            id="RestaurantID"
            name="RestaurantID"
          />
          <InputField label="RoleID" type="number" id="RoleID" name="RoleID" />
          <InputField
            label="Username"
            type="text"
            id="username"
            name="username"
          />
          <InputField
            label="Contraseña"
            type="password"
            id="password"
            name="password"
          />
          <Box>
            <Checkbox name="change" id="change">
              Cambiar Contraseña
            </Checkbox>
          </Box>

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
export default Empleados;
