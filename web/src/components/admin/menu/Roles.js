import CRUDTable from "../../CRUDTable";
const Roles = () => {
  const headers = ["Codigo|number", "Nombre"];
  const keys = ["Code", "Name"];
  return (
    <CRUDTable
      tableHead={headers}
      apiPath="roles"
      dataKeys={keys}
      modalTitle="Crear Rol"
    />
  );
};
export default Roles;
