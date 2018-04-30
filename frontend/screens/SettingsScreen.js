import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

import SettingsButton from '../components/SettingsButton/SettingsButton';

export default class SettingsScreen extends React.Component {
  static navigationOptions = {
    title: 'Settings',
  };

  render() {
    return (
      <View>
        <SettingsButton value='Location'/> 
        <SettingsButton value='Radius'/> 
        <SettingsButton value='Likes'/> 
        <SettingsButton value='Dislikes'/> 
        <SettingsButton value='Account'/> 
        <SettingsButton value='Logout'/> 
      </View>
    );
  }
}
