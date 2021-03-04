import CRUDTable from "../../CRUDTable";

const Restaurantes = () => {
  const headers = ["Nombre", "Direccion", "Numero|number", "CP", "RFC"];
  const keys = ["ID", "Name", "Address", "Number", "CP", "RFC"];
  return (
    <CRUDTable
      tableHead={headers}
      apiPath="restaurants"
      dataKeys={keys}
      modalTitle="Crear Restaurante"
    />
  );
};

export default Restaurantes;
