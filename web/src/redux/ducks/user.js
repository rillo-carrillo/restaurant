import axios from "axios";
import Cookies from "js-cookie";
import { HOSTNAME } from "../../consts";
//Actions
const LOGIN_USER_REQUEST = "restaurant/login/LOGIN_USER_REQUEST";
const LOGIN_USER_SUCCESS = "restaurant/login/LOGIN_USER_SUCCESS";
const LOGIN_USER_FAILURE = "restaurant/login/LOGIN_USER_FAILURE";
const GET_USER_INFO = "restaurant/login/GET_USER_INFO";
const LOGOUT_USER = "restaurant/login/LOGOUT_USER";
const initialState = {
  id: "",
  role: 0,
  error: "",
};
const config = {
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
};
//Reducer
export function userReducer(state = initialState, action = {}) {
  switch (action.type) {
    case LOGIN_USER_SUCCESS:
      return {
        ...state,
        id: action.payload.ID,
        role: action.payload.RoleID,
        error: "",
      };
    case LOGIN_USER_FAILURE:
      return {
        ...state,
        error: action.payload,
      };
    case GET_USER_INFO:
      return {
        ...state,
        id: action.payload.ID,
        role: action.payload.RoleID,
      };
    case LOGOUT_USER:
      return {
        ...state,
        id: "",
        role: 0,
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
const fetchUserIdSuccess = (data) => {
  return {
    type: LOGIN_USER_SUCCESS,
    payload: data,
  };
};
const fetchUserIdFailure = (error) => {
  return {
    type: LOGIN_USER_FAILURE,
    payload: error,
  };
};
const fetchMeAction = (data) => {
  return {
    type: GET_USER_INFO,
    payload: data,
  };
};
const logoutUserAction = () => {
  return {
    type: LOGOUT_USER,
  };
};
export const fetchUser = (values) => {
  return function (dispatch) {
    axios
      .post(
        `${HOSTNAME}/login`,
        {
          ...values,
        },
        config
      )
      .then(({ data, headers }) => {
        if (data.Message !== undefined) {
          dispatch(fetchUserIdFailure(data.Message));
          Cookies.set("access_token", headers["x-access-token"]);
        } else {
          dispatch(fetchUserIdSuccess(data));
        }
      })
      .catch((error) => {
        console.log(error);
      });
  };
};
export const fetchMe = () => {
  return function (dispatch) {
    axios
      .get(`${HOSTNAME}/me`, config)
      .then(({ data }) => {
        dispatch(fetchMeAction(data));
      })
      .catch((error) => {
        console.log(error);
      });
  };
};
export const logoutUser = () => {
  return function (dispatch) {
    axios.delete(`${HOSTNAME}/login`, config).then(() => {
      Cookies.remove("cnr");
      dispatch(logoutUserAction());
    });
  };
};
