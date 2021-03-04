import React from "react";

import { connect, useDispatch } from "react-redux";
import { fetchUser } from "../../redux/ducks/user";
import { Box, Button, Container, Flex } from "@chakra-ui/react";
import { Formik, Form } from "formik";
import InputField from "../../components/InputField";
const Login = ({ user: { error } }) => {
  const dispatch = useDispatch();
  return (
    <Container maxW="md" mt="15%" borderRadius="md" boxShadow="dark-lg">
      <Box padding="4">
        <h1>Sign Up</h1>
        <Formik
          initialValues={{
            username: "",
            password: "",
          }}
          onSubmit={async (values) => {
            dispatch(fetchUser(values));
          }}
        >
          {({ isSubmiting }) => (
            <Form>
              <InputField
                id="username"
                label="username"
                name="username"
                type="text"
              />
              <InputField
                id="password"
                label="password"
                name="password"
                type="password"
                error={error}
              />
              <Flex>
                <Button
                  mt={4}
                  colorScheme="teal"
                  type="submit"
                  isLoading={isSubmiting}
                  flex="1"
                >
                  Enviar
                </Button>
              </Flex>
            </Form>
          )}
        </Formik>
      </Box>
    </Container>
  );
};
function mapStateToProps(state) {
  return {
    user: state.user,
  };
}
export default connect(mapStateToProps, null)(Login);
