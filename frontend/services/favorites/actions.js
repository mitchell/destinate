export const ADD_FAVORITE = 'ADD_FAVORITE';

export const addFavorite = (favorite) => {
  return {
    type: ADD_FAVORITE,
    favorite,
  };
};
