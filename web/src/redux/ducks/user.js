import axios from "axios";
//Actions
const LOGIN_USER_REQUEST = "restaurant/login/LOGIN_USER_REQUEST";
const LOGIN_USER_SUCCESS = "restaurant/login/LOGIN_USER_SUCCESS";
const LOGIN_USER_FAILURE = "restaurant/login/LOGIN_USER_FAILURE";
const initialState = {
  id: "",
  error: "Usuario no se encuentra",
};
//Reducer
export function userReducer(state = initialState, action = {}) {
  switch (action.type) {
    case LOGIN_USER_SUCCESS:
      return {
        ...state,
        id: action.payload,
      };
    case LOGIN_USER_FAILURE:
      return {
        ...state,
        error: action.payload,
      };
    default:
      return state;
  }
}
//Actions
const fetchUserId = () => {
  return {
    type: LOGIN_USER_REQUEST,
  };
};
const fetchUserIdSuccess = (id) => {
  return {
    type: LOGIN_USER_SUCCESS,
    payload: id,
  };
};
const fetchUserIdFailure = (error) => {
  return {
    type: LOGIN_USER_FAILURE,
    payload: error,
  };
};

export const fetchUser = () => {
  return function (dispatch) {
    axios
      .get("http://localhost:8080/")
      .then((response) => {
        dispatch(fetchUserIdSuccess(response.data));
      })
      .catch((error) => {
        dispatch(fetchUserIdFailure(error));
      });
  };
};
