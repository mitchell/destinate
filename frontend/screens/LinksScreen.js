import React from 'react';
import { connect } from 'react-redux';
import { ScrollView, StyleSheet } from 'react-native';
import { ExpoLinksView } from '@expo/samples';

import DestinationCard from '../components/DestinationCard/DestinationCard';

class LinksScreen extends React.Component {
  static navigationOptions = {
    title: 'Links',
  };

  render() {
    return (
      <ScrollView style={styles.container}>
        {
          this.props.favorites[0] &&
          this.props.favorites.map((favorite) => { return <DestinationCard destination={favorite}/>; })
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
)(LinksScreen);
