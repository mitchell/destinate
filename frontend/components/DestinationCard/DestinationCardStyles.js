import { StyleSheet, Dimensions } from 'react-native';

const { width, height } = Dimensions.get('window');

export default StyleSheet.create({
  cardContainer: {
    width: '100%',
    height: height * 0.5,
    paddingHorizontal: 30,
    paddingVertical: 20,
    borderColor: '#DDDDDD',
    borderWidth: 0.5,
    borderRadius: 10,
    shadowColor: '#aaa',
    shadowOffset: {
      width: 0,
      height: 1,
    },
    shadowRadius: 3,
    shadowOpacity: 0.5,
    backgroundColor: '#fff',
    alignItems: 'center',
    marginTop: 10,
  },

  name: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 10,
  },

  open: {
    fontSize: 12,
    marginBottom: 10,
  },

  vicinity: {
    marginTop: 20,
    fontSize: 12,
  },

  image: {
    width: '100%',
    height: 0.6 * width,
  }
});
