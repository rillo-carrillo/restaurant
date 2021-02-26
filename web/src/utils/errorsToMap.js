export function errorToMap(errors) {
  const err = new Map();
  errors.forEach((error) => {
    err[error.field] = error.message;
  });
  return err;
}
