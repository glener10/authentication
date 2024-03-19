SCENARIO: Login with success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is sent to /login with the email and password of the registered user
THEN: It must return a JWT and a 200 code

SCENARIO: Login without success because body is in wrong format
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is sent to /login with invalid format of email or password
THEN: It must return a error message and a 422 code

SCENARIO: Login without success because is no user with the data provided
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is sent to /login with email and password in valid format but the email or password is incorrect
THEN: It must return a error message and a 401 code
