import { Grid, GridItem } from "@chakra-ui/react";
import { useState } from "react";
import Restaurantes from "../../components/admin/menu/Restaurantes";
import Roles from "../../components/admin/menu/Roles";
import Navbar from "../../components/admin/Navbar";
import Sidebar from "../../components/admin/Sidebar";
import Empleados from "../../components/admin/menu/Empleados";
import Ordenes from "../../components/admin/menu/Ordenes";
import Menu from "../../components/admin/menu/Menu";
import Proveedores from "../../components/admin/menu/Proveedores";
const Admin = () => {
  const [menu, setMenu] = useState(0);
  const handleMenuChange = (id) => {
    setMenu(id);
  };
  const showMenu = () => {
    if (menu === 1) {
      return <Restaurantes />;
    } else if (menu === 2) {
      return <Roles />;
    } else if (menu === 3) {
      return <Empleados />;
    } else if (menu === 4) {
      return <Ordenes />;
    } else if (menu === 5) {
      return <Menu />;
    } else if (menu === 6) {
      return <Proveedores />;
    }
    return null;
  };
  return (
    <Grid h="100%" templateRows="repeat(6,1fr)" templateColumns="repeat(6,1fr)">
      <GridItem rowSpan={6} marginRight="2">
        <Sidebar setMenu={handleMenuChange} />
      </GridItem>
      <GridItem colStart={2} colEnd={7} rowSpan={6}>
        <Navbar />
        {showMenu()}
      </GridItem>
    </Grid>
  );
};

export default Admin;
