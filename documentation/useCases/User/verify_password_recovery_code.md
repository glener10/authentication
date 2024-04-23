SCENARIO: Verify password Recovery Code with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyPasswordRecoveryCode/:find with a valid search parameter (id or email) and a valid code in the body (not expired)
THEN: It must return a 200 code

SCENARIO: Verify password Recovery Code without Success because find param is invalid (even if the token is valid)
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyPasswordRecoveryCode/:find with a invalid search parameter (not is email or id) but with valid code
THEN: It must return the a error message "wrong format, parameter need to be a id or a e-mail" and a 422 code

SCENARIO: Verify password Recovery Code without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyPasswordRecoveryCode/:find with a valid search parameter and invalid code
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Verify password Recovery Code without Success because the token provided is expired
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyPasswordRecoveryCode/:find with a valid search parameter but the code is expired
THEN: It must return the a error message "your code has expired" and a 401 code
