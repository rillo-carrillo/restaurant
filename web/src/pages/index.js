import Login from "./login";
import { connect } from "react-redux";
const Index = ({ user }) => {
  return <>{user.id === "" ? <Login /> : <div>LOGEADO</div>}</>;
};

function mapStateToProps(state) {
  return {
    user: state.user,
  };
}

export default connect(mapStateToProps, null)(Index);
