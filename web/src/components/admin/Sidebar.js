import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  List,
  ListIcon,
  ListItem,
  Icon,
} from "@chakra-ui/react";
import { HiOutlineOfficeBuilding } from "react-icons/hi";
import { GoPerson } from "react-icons/go";
import { MdBorderColor, MdRestaurantMenu } from "react-icons/md";
import { ImUsers, ImTruck } from "react-icons/im";
const Sidebar = ({ setMenu }) => {
  return (
    <Box w="100%" height="100%" marginRight="2em">
      <Accordion allowToggle>
        <AccordionItem>
          <AccordionButton>
            Administracion
            <AccordionIcon />
          </AccordionButton>
          <AccordionPanel>
            <List>
              <ListItem cursor="pointer" onClick={() => setMenu(1)}>
                <Icon as={HiOutlineOfficeBuilding} marginRight={2} />
                Restaurantes
              </ListItem>
              <ListItem cursor="pointer" onClick={() => setMenu(2)}>
                <Icon as={GoPerson} marginRight={2} />
                Roles
              </ListItem>
              <ListItem cursor="pointer" onClick={() => setMenu(3)}>
                <Icon as={ImUsers} marginRight={2} />
                Empleados
              </ListItem>
              <ListItem cursor="pointer" onClick={() => setMenu(4)}>
                <Icon as={MdBorderColor} marginRight={2} />
                Ordenes
              </ListItem>
            </List>
          </AccordionPanel>
        </AccordionItem>
        <AccordionItem>
          <AccordionButton>
            Inventario
            <AccordionIcon />
          </AccordionButton>
          <AccordionPanel>
            <List>
              <ListItem cursor="pointer" onClick={() => setMenu(5)}>
                <Icon as={MdRestaurantMenu} marginRight={2} />
                Menu
              </ListItem>
              <ListItem cursor="pointer" onClick={() => setMenu(6)}>
                <Icon as={ImTruck} marginRight={2} />
                Proveedores
              </ListItem>
            </List>
          </AccordionPanel>
        </AccordionItem>
      </Accordion>
    </Box>
  );
};
export default Sidebar;
