SCENARIO: Create user with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request was sent to /user with valid values
THEN: It must return a message and a 201 code, without the password in the body and save in the database

SCENARIO: Dont create a user because the e-mail is already in use
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request was sent to /user with valid values
THEN: It must return a error message and a 422 code

SCENARIO: Dont create a user because the e-mail is in wrong format
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request was sent to /user with a invalid e-mail
THEN: It must return a error message and a 422 code

SCENARIO: Dont create a user because the password is weak
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request was sent to /user with a weak password
THEN: It must return a error message and a 422 code
