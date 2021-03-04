import CRUDTable from "../../CRUDTable";
const Empleados = () => {
  const headers = [
    "Nombre",
    "RestaurantID|number",
    "RoleID|number",
    "Username",
    "Password|password",
  ];
  const keys = ["ID", "Name", "RestaurantID", "RoleID", "Username", "Password"];
  return (
    <CRUDTable
      tableHead={headers}
      apiPath="empleados"
      dataKeys={keys}
      modalTitle="Crear Empleado"
    />
  );
};
export default Empleados;
