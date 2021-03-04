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
import axios from "axios";
import { useEffect, useState } from "react";
import { HOSTNAME } from "../../../consts";
import { Formik, Form } from "formik";
import InputField from "../../InputField";
import CRUDButton from "../CRUDButton";
const types = {
  create: "create",
  modify: "modify",
};
const Restaurantes = () => {
  const [restaurants, setRestaurants] = useState([]);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [item, setItem] = useState({ id: null, selected: false });
  const [modalType, setModalType] = useState("");
  useEffect(async () => {
    loadRestaurants();
  }, []);
  const loadRestaurants = async () => {
    axios.get(`${HOSTNAME}/restaurants`).then(({ data }) => {
      setRestaurants(data);
    });
  };
  const handleItemSelect = (id) => {
    if (item.id === id) {
      setItem({ ...item, selected: !item.selected });
    } else {
      setItem({ id: id, selected: true });
    }
  };
  const handleModal = (type) => {
    setModalType(type);
    onOpen();
  };
  const handleDelete = async () => {
    axios.delete(`${HOSTNAME}/restaurants/${item.id}`).then(() => {
      loadRestaurants();
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

const ModalCreate = ({ onClose, type, isOpen, data, load }) => {
  return (
    <Modal size="3xl" onClose={onClose} isOpen={isOpen}>
      <ModalOverlay />
      <ModalContent background="white">
        <ModalHeader>
          {type === types.create
            ? "Crear Restaurante"
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
        id: data === null ? 0 : data.ID,
        name: data === null ? "" : data.Name,
        address: data === null ? "" : data.Address,
        number: data === null ? 0 : data.Number,
        CP: data === null ? 0 : data.CP,
        rfc: data === null ? "" : data.RFC,
      }}
      onSubmit={async (values) => {
        if (type === types.create) {
          axios
            .post(`${HOSTNAME}/restaurants`, values)
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
            .put(`${HOSTNAME}/restaurants`, values)
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
            label="Direccion"
            type="text"
            id="address"
            name="address"
          />
          <InputField label="Numero" type="number" id="number" name="number" />
          <InputField label="Codigo Postal" type="text" id="CP" name="CP" />
          <InputField label="RFC" type="text" id="rfc" name="rfc" />
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
export default Restaurantes;
