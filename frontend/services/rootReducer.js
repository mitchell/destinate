import { combineReducers } from 'redux';

import destinations from './destinations/reducer.js';
import favorites from './favorites/reducer.js';

const root = combineReducers({
  destinations,
  favorites,
});

export default root;
