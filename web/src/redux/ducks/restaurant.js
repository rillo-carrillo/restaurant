const RESTAURANT_SELECT = "restaurant/restaurant/RESTAURANT_SELECT";

const initialState = {
  id: 0,
};
export function restaurantReducer(state = initialState, action = {}) {
  switch (action.type) {
    case RESTAURANT_SELECT:
      return {
        ...state,
        id: action.payload,
      };
    default:
      return state;
  }
}

export const setRestaurantId = (id) => {
  return {
    type: RESTAURANT_SELECT,
    payload: id,
  };
};
