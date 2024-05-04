SCENARIO: Verify Email with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyEmail/:find with a valid search parameter (id or email) and valid JWT from the owner of the search parameter and a valid code in the body (not expired)
THEN: It must return a 200 code

SCENARIO: Verify Email without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyEmail/:find with a valid search parameter but without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Verify Email without Success because find param is invalid (even if the token is valid)
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyEmail/:find with a invalid search parameter (not is email or id) but with valid token
THEN: It must return the a error message "wrong format, parameter need to be a id or a e-mail" and a 422 code

SCENARIO: Verify Email without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyEmail/:find with a valid search parameter and invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Verify Email without Success because the token provided is from a different user than the search parameter
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /users/verifyEmail/:find with a valid search parameter but the token is from a different user than the search parameter
THEN: It must return the a error message "you do not have permission to perform this operation" and a 401 code
