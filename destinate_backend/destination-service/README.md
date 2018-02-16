# destination-service
Here you'll find one of destinate's micro-services. The destination-service's
role is to list destinations based on user preferences using machine-learning.
This is the 'gravy' of the application, so to say. It's a set of AWS Lambda
functions (or will be), that is hooked up to API Gateway, using the Serverless framework.


Current Endpoints:
- POST - https://nre1rvg4y1.execute-api.us-west-1.amazonaws.com/dev/destination
