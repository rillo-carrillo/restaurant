import { useEffect } from "react";
import { useDispatch } from "react-redux";
import { fetchMe } from "../redux/ducks/user";
export const useMe = () => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchMe());
  }, []);
};
