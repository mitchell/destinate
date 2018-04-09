import React from 'react';
import { View, Text, Image } from 'react-native';
import styles from './DestinationCardStyles.js';

export default class DestinationCard extends React.Component {
  render() {
    return (
      <View style={styles.cardContainer}>
        <Text style={styles.name}>{this.props.destination.name}</Text>
        {this.props.destination.opening_hours !== undefined && <Text style={styles.open}>{this.props.destination.opening_hours.open_now ? 'Open Now' : 'Closed'}</Text> }
        <Image style={styles.image} source={{ uri: this.props.destination.icon }} />
        {this.props.destination.vicinity && <Text style={styles.vicinity}>Vicinity: {this.props.destination.vicinity}</Text>}
      </View>
    );
  }
}
