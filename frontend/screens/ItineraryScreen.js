import React from 'react';
import { connect } from 'react-redux';
import { ScrollView, StyleSheet } from 'react-native';
import { ExpoLikesView } from '@expo/samples';

import DestinationCard from '../components/DestinationCard/DestinationCard';
import SettingsButton from '../components/SettingsButton/SettingsButton';

class ItineraryScreen extends React.Component {
  static navigationOptions = {
    title: 'Itinerary',
  };

  render() {
    return (
      <ScrollView style={styles.container}>
        {
          this.props.favorites[0] &&
          this.props.favorites.map((favorite) => { return <SettingsButton value={favorite.name}/>; })
        }
      </ScrollView>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    padding: 30,
    flex: 1,
    paddingTop: 15,
    backgroundColor: '#fff',
  },
});

const mapStateToProps = state => ({
  favorites: state.favorites,
});

export default connect(
  mapStateToProps,
  undefined,
)(ItineraryScreen);
