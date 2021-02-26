import React from "react";
import { connect } from "react-redux";
import {
  Box,
  Button,
  Container,
  Flex,
  FormControl,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import InputField from "../components/InputField";
import { Formik, Form } from "formik";
const Login = ({ user: { error } }) => {
  return (
    <Container maxW="md" mt="15%" borderRadius="md" boxShadow="dark-lg">
      <Box padding="4">
        <h1>Sign Up</h1>
        <Formik
          initialValues={{
            username: "",
            password: "",
          }}
          onSubmit={async (values, { setFieldError }) => {
            console.log(values);
            if (error !== "") {
              setFieldError("password", error);
            }
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
                err
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
