export const GET_DESTINATIONS = 'GET_DESTINATIONS';
export const GET_DESTINATIONS_SUCCESS = 'GET_DESTINATIONS_SUCCESS';

export const getDestinations = () => {
  return {
    type: GET_DESTINATIONS,
  };
};

export const getDestinationsSuccess = (destinations) => {
  return {
    type: GET_DESTINATIONS_SUCCESS,
    destinations,
  };
};
