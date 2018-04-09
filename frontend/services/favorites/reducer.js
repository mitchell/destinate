import { ADD_FAVORITE } from './actions.js';

const favorites = (state = [], action) => {
  switch(action.type) {

  case ADD_FAVORITE:
    return [action.favorite, ...state];

  default:
    return state;
  }
};

export default favorites;
