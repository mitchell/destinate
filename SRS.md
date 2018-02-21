# Destinate (Working Title) software requirements
## 5.1 Introduction
This application is a multi(two)-tier application, consisting of a
component-based native frontend application, for iOS, and a micro-service based
serverless, FaaS (Functions as a Service), backend application. The frontend
application will consist of a [Meteor](https://meteor.com/) native
application, which is a component-based native application development
framework. This will communicate with the AWS Lambda (a FaaS provider) backend
applicaiton, which will use the [Serverless](https://serverless.com)
framework, which is a micro-service based, serverless, FaaS, development
framework.
The remainder of this document is structured as follows:
## 5.2 Functional Requirements
* 5.2.1 Destinate Account Management: Users shall be presented with option to login or sign up upon initial usage of the application.
  * 5.2.1.1 The login option shall validate the users input and provide access to the application if the information is correct.
  * 5.2.1.2 The Sign Up option shall ask users for an email, password and confirm password, and will validate that the password has been inputted correctly.
  * 5.2.1.3 Upon account creation the user will be given a series of simple questions relating to preference.
    * 5.2.1.3.1 Once the user has answered the series of questions, they will be given a ranking within our algorithm for providing suggestions.
* 5.2.2 Destinate shall have two primary usage modes: "Now" mode and "Later" mode
  * 5.2.2.1 "Now" Mode shall provide users with a list of suggestions within their vicinity based on their preferences.
  * 5.2.2.2 "Later" Mode will ask the users for dates and locations and will then generate an itinerary for the user.
    * 5.2.2.2.1 After the itinerary has been generated, the user will be given the option to change it to their liking.

## 5.3 Performance Requirements
* 5.3.1 Activities around the user
  * 5.3.1.1 The app shall return activities around the users current location, this should take no longer than 20 seconds.
* 5.3.2 Searching Activities
  * 5.3.2.1 The app shall return activities based on the users search query, this should take no longer than 10 seconds.
## 5.4 Environment Requirements
* 5.4.1 Development Environment Requirements
  * 5.4.1.1 This Application's frontend will use javascript with the Meteor framework
  * 5.4.1.2 This Application's backend will use the serverless framework
* 5.4.2 Execution Environment Requirements
  * 5.4.2.1 This application shall run natively on iOS
  * 5.4.2.2 This application requires an internet connection
