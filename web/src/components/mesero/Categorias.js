import { useEffect, useState } from "react";
import { HOSTNAME } from "../../consts";
import { Flex, Center } from "@chakra-ui/react";
import axios from "axios";
const Categorias = () => {
  const [categories, setCategories] = useState([]);
  useEffect(async () => {
    await axios.get(`${HOSTNAME}` + "/categorias").then(({ data }) => {
      setCategories(data);
    });
  }, []);
  return (
    <Flex>
      {categories.map((category, id) => (
        <Center
          w="md"
          m={4}
          key={id}
          boxShadow="dark-lg"
          h="100px"
          borderRadius="md"
        >
          {category.Name}
        </Center>
      ))}
    </Flex>
  );
};

export default Categorias;
