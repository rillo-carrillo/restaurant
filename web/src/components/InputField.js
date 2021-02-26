import React from "react";
import PropTypes from "prop-types";
import {
  FormControl,
  FormLabel,
  Input,
  Box,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useField } from "formik";
const InputField = ({ id, label, err, ...props }) => {
  const [field, { error }] = useField(props);
  let C = Input;
  console.log(error !== undefined);
  console.log(error);
  return (
    <Box mt={4}>
      <FormControl id={id} isInvalid={!!error}>
        <FormLabel htmlFor={field.name}>{label}</FormLabel>
        <C {...field} {...props} id={field.name} />
        {error !== undefined ? (
          <FormErrorMessage>{error}</FormErrorMessage>
        ) : null}
      </FormControl>
    </Box>
  );
};

InputField.propTypes = {
  id: PropTypes.string,
  label: PropTypes.string,
  err: PropTypes.bool,
};

export default InputField;
