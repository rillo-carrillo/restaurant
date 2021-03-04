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
import { Formik, Form } from "formik";
import { useEffect, useState } from "react";
import InputField from "./InputField";
import CRUDButton from "./CRUDButton";
import { HOSTNAME } from "../consts";
const types = {
  create: "create",
  modify: "modify",
};

const CRUDTable = ({ tableHead, apiPath, dataKeys, modalTitle }) => {
  const [data, setData] = useState([]);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [item, setItem] = useState({ id: null, selected: false });
  const [modalType, setModalType] = useState("");
  const [fieldTypes, setFieldTypes] = useState({});

  const ConfigFields = () => {
    let fields = {};
    tableHead.forEach((field) => {
      let f = field.split("|");
      if (f.length > 1) {
        fields[f[0]] = f[1];
      } else {
        fields[f[0]] = "text";
      }
    });
    setFieldTypes({ ...fields });
  };
  useEffect(() => {
    loadData();
    ConfigFields();
  }, []);
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
  const loadData = async () => {
    axios.get(`${HOSTNAME}/${apiPath}`).then((res) => {
      setData(res.data);
    });
  };
  const handleDelete = () => {
    axios.delete(`${HOSTNAME}/${apiPath}/${item.id}`).then(() => {
      loadData();
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
            {tableHead.map((head, idx) => (
              <Th key={idx}>{head.split("|")[0]}</Th>
            ))}
          </Tr>
        </Thead>
        <Tbody>
          {data.map((d) => (
            <Tr
              key={d[dataKeys[0]]}
              onClick={() => handleItemSelect(d[dataKeys[0]])}
            >
              {dataKeys.map((key, idx) =>
                key !== "ID" ? <Td key={idx}>{d[key]}</Td> : null
              )}
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
            ? data.find((r) => r[dataKeys[0]] === item.id)
            : null
        }
        load={loadData}
        modalTitle={modalTitle}
      >
        <CreateForm
          data={
            item.selected && modalType === types.modify
              ? data.find((r) => r[dataKeys[0]] === item.id)
              : null
          }
          dataTypes={fieldTypes}
          dataKeys={dataKeys}
          type={modalType}
          load={loadData}
          onClose={onClose}
          apiPath={apiPath}
          tableHead={tableHead}
        />
      </ModalCreate>
    </Box>
  );
};
const ModalCreate = ({ onClose, type, isOpen, data, modalTitle, children }) => {
  return (
    <Modal size="3xl" onClose={onClose} isOpen={isOpen}>
      <ModalOverlay />
      <ModalContent background="white">
        <ModalHeader>
          {type === types.create
            ? modalTitle
            : data === null
            ? null
            : `Modificar ${data.Name}`}
        </ModalHeader>
        <ModalCloseButton />
        <ModalBody>{children}</ModalBody>
      </ModalContent>
    </Modal>
  );
};

const CreateForm = ({
  data,
  dataKeys,
  type,
  load,
  onClose,
  apiPath,
  tableHead,
  dataTypes,
}) => {
  const loadInitValues = () => {
    let initValues = {};
    if (data === null) {
      dataKeys.forEach((element) => {
        initValues[element] = null;
      });
    } else {
      dataKeys.forEach((element) => {
        initValues[element] = data[element];
      });
    }
    return initValues;
  };
  console.log(tableHead);
  console.log(dataTypes);
  return (
    <Formik
      initialValues={loadInitValues()}
      onSubmit={async (values) => {
        if (type === types.create) {
          if (values.Password === "*****") {
            values.Password = "";
          }
          axios
            .post(`${HOSTNAME}/${apiPath}`, values)
            .then(() => {
              load();
              onClose();
            })
            .catch((error) => {
              console.log(error);
            });
        }
        if (type === types.modify) {
          if (values.Password === "*****") {
            values.Password = "";
          }
          axios
            .put(`${HOSTNAME}/${apiPath}`, values)
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
      {({ isSubmitting }) => (
        <Form>
          {tableHead.map((head, id) =>
            tableHead.length === dataKeys.length ? (
              <InputField
                label={head.split("|")[0]}
                type={dataTypes[head.split("|")[0]]}
                id={dataKeys[id]}
                name={dataKeys[id]}
                key={id}
              />
            ) : (
              <InputField
                label={head.split("|")[0]}
                type={dataTypes[head.split("|")[0]]}
                id={dataKeys[id + 1]}
                name={dataKeys[id + 1]}
                key={id}
              />
            )
          )}

          <Button
            colorScheme="teal"
            type="submit"
            isLoading={isSubmitting}
            mt="10px"
          >
            Guardar
          </Button>
        </Form>
      )}
    </Formik>
  );
};
export default CRUDTable;
