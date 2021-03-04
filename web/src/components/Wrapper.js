import { connect, useDispatch } from "react-redux";
import { useRole } from "../hooks/useRole";
import { useMe } from "../hooks/useMe";
import { Button } from "@chakra-ui/react";
import { logoutUser } from "../redux/ducks/user";
const Wrapper = ({
  children,

  user: { role },
}) => {
  useMe();
  useRole(role);
  const dispatch = useDispatch();
  return (
    <>
      {role > 0 ? (
        <Button
          onClick={() => {
            dispatch(logoutUser());
          }}
        >
          LogOut
        </Button>
      ) : null}
      {children}
    </>
  );
};
function mapStateToProps(state) {
  return {
    user: state.user,
  };
}
export default connect(mapStateToProps, null)(Wrapper);
