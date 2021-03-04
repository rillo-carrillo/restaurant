import { Box, Flex, Select, Spacer } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import axios from "axios";
import { HOSTNAME } from "../../consts";
import { setRestaurantId } from "../../redux/ducks/restaurant";
const Navbar = () => {
  const [restaurants, setRestaurants] = useState([]);
  const dispatch = useDispatch();
  useEffect(async () => {
    await axios.get(`${HOSTNAME}/restaurants`).then(({ data }) => {
      setRestaurants(data);
      dispatch(setRestaurantId(data[0].ID));
    });
  }, []);
  const handleSelect = ({ target: { value } }) => {
    dispatch(setRestaurantId(value));
  };
  return (
    <Flex bg="AppWorkspace">
      <Spacer />
      <Box mr={4}>
        <Select onChange={handleSelect}>
          {restaurants.map((rest, id) => (
            <option key={id} defaultValue={id === 1} value={rest.id}>
              {rest.Name}
            </option>
          ))}
        </Select>
      </Box>
    </Flex>
  );
};
export default Navbar;
