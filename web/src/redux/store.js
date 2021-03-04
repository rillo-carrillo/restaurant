import { createStore, applyMiddleware } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";
import rootReducer from "./ducks/index";
// const devTools =
//   process.env.NODE_ENV === "development"
//     ? window.__REDUX_DEVTOOLS_EXTENSION__ &&
//       window.__REDUX_DEVTOOLS_EXTENSION__()
//     : null;
const initialState = {};
const store = createStore(
  rootReducer,
  initialState,
  //applyMiddleware(thunk)
  composeWithDevTools(applyMiddleware(thunk))
);
export function useStore() {
  return store;
}
