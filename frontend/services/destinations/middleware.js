import request from 'superagent';
import { GET_DESTINATIONS, getDestinationsSuccess } from './actions.js';

const path = 'https://k8tnbgmjyg.execute-api.us-west-1.amazonaws.com/dev/destinations';

const destinationsMiddleware = store => next => action => {
  next(action);

  switch(action.type) {

  case GET_DESTINATIONS:
    request.get(path)
      .then((res) => {
        next(getDestinationsSuccess(res.body.Results));
      });
    break;

  default:
    break;
  }
};

export default destinationsMiddleware;
