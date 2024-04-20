SCENARIO: Send password recovery code with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/sendPasswordRecoveryCode/:find with a valid search parameter (id or email)
THEN: It must return a 200 code

SCENARIO: Send password recovery code without Success because find param is invalid (even if the token is valid)
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/sendPasswordRecoveryCode/:find with a invalid search parameter (not is email or id) but with valid token
THEN: It must return the a error message "wrong format, parameter need to be a id or a e-mail" and a 422 code

SCENARIO: Send password recovery code without Success because dont exists a user with de find param
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/sendPasswordRecoveryCode/:find with a valid search parameter and invalid authorization header
THEN: It must return the a error message "invalid token" and a 404 code