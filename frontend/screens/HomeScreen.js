import React from 'react';
import {
  Image,
  Platform,
  ScrollView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native';
import { WebBrowser } from 'expo';
import { connect } from 'react-redux';
import { getDestinations } from '../services/destinations/actions.js';
import { addFavorite } from '../services/favorites/actions.js';

import DestinationCard from '../components/DestinationCard/DestinationCard';
import LightButton from '../components/LightButton/LightButton';
import { MonoText } from '../components/StyledText';

class HomeScreen extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      destinationIndex: 0,
    };
  }

  addDestination = () => {
    this.props.addFavorite(this.props.destinations[this.state.destinationIndex]);
    this.setState({ destinationIndex: this.state.destinationIndex + 1 });
  }

  skipDestination = () => {
    this.setState({ destinationIndex: this.state.destinationIndex + 1 });
  }

  static navigationOptions = {
    header: null,
  };

  componentWillMount() {
    this.props.getDestinations();
  }

  render() {
    if (this.props.destinations[this.state.destinationIndex]) {
      return (
        <View style={styles.container}>
          <View style={styles.getStartedContainer}>
            <DestinationCard destination={this.props.destinations[this.state.destinationIndex]}/>
            <LightButton value='Interested' onPress={this.addDestination}/>
            <LightButton value='Not Interested' onPress={this.skipDestination}/>
          </View>


          <View style={styles.tabBarInfoContainer}>
            <Text style={styles.tabBarInfoText}>Welcome to Destinate, check out your other tabs:</Text>
          </View>
        </View>
      );
    }
    return <View/>

  }

  _maybeRenderDevelopmentModeWarning() {
    if (__DEV__) {
      const learnMoreButton = (
        <Text onPress={this._handleLearnMorePress} style={styles.helpLikeText}>
          Learn more
        </Text>
      );

      return (
        <Text style={styles.developmentModeText}>
          Development mode is enabled, your app will be slower but you can use useful development
          tools. {learnMoreButton}
        </Text>
      );
    } else {
      return (
        <Text style={styles.developmentModeText}>
          You are not in development mode, your app will run at full speed.
        </Text>
      );
    }
  }

  _handleLearnMorePress = () => {
    WebBrowser.openBrowserAsync('https://docs.expo.io/versions/latest/guides/development-mode');
  };

  _handleHelpPress = () => {
    WebBrowser.openBrowserAsync(
      'https://docs.expo.io/versions/latest/guides/up-and-running.html#can-t-see-your-changes'
    );
  };
}

const mapStateToProps = state => ({
  destinations: state.destinations,
});

const mapDispatchToProps = dispatch => ({
  getDestinations: () => { dispatch(getDestinations()); },
  addFavorite: (favorite) => { dispatch(addFavorite(favorite)); },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(HomeScreen);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    paddingTop: 30,
  },
  developmentModeText: {
    marginBottom: 20,
    color: 'rgba(0,0,0,0.4)',
    fontSize: 14,
    lineHeight: 19,
    textAlign: 'center',
  },
  contentContainer: {
    paddingTop: 30,
  },
  welcomeContainer: {
    alignItems: 'center',
    marginTop: 10,
    marginBottom: 20,
  },
  welcomeImage: {
    width: 100,
    height: 80,
    resizeMode: 'contain',
    marginTop: 3,
    marginLeft: -10,
  },
  getStartedContainer: {
    alignItems: 'center',
    marginHorizontal: 50,
    marginTop: 10,
    marginBottom: 20,
  },
  homeScreenFilename: {
    marginVertical: 7,
  },
  codeHighlightText: {
    color: 'rgba(96,100,109, 0.8)',
  },
  codeHighlightContainer: {
    backgroundColor: 'rgba(0,0,0,0.05)',
    borderRadius: 3,
    paddingHorizontal: 4,
  },
  getStartedText: {
    fontSize: 17,
    color: 'rgba(96,100,109, 1)',
    lineHeight: 24,
    textAlign: 'center',
  },
  tabBarInfoContainer: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
    ...Platform.select({
      ios: {
        shadowColor: 'black',
        shadowOffset: { height: -3 },
        shadowOpacity: 0.1,
        shadowRadius: 3,
      },
      android: {
        elevation: 20,
      },
    }),
    alignItems: 'center',
    backgroundColor: '#fbfbfb',
    paddingVertical: 20,
  },
  tabBarInfoText: {
    fontSize: 14,
    color: 'rgba(96,100,109, 1)',
    textAlign: 'center',
  },
  navigationFilename: {
    marginTop: 5,
  },
  helpContainer: {
    marginTop: 15,
    alignItems: 'center',
  },
  helpLike: {
    paddingVertical: 15,
  },
  helpLikeText: {
    fontSize: 14,
    color: '#2e78b7',
  },
});
