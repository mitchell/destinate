import React from 'react';
import { TouchableOpacity, Text, Image } from 'react-native';
import styles from './LightButtonStyles.js';

export default class LightButton extends React.Component {
  render() {
    return (
      <TouchableOpacity onPress={this.props.onPress} style={styles.buttonContainer}>
        <Text>{this.props.value}</Text>
      </TouchableOpacity>
    );
  }
}
