import { GET_DESTINATIONS_SUCCESS } from './actions.js';

const destinations = (state = [], action) => {
  switch(action.type) {

  case GET_DESTINATIONS_SUCCESS:
    return action.destinations;

  default:
    return state;
  }
};

export default destinations;
