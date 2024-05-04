SCENARIO: Login 2FA with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /login/2fa with valid JWT from the owner of the search parameter and a valid code in the body
THEN: It must return a JWT and a 200 code

SCENARIO: Login 2FA without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /login/2fa without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Login 2FA without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /login/2fa with invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Login 2FA without Success because the 2FA code is invalid
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /login/2fa with a valid JWT but the code in the body is invalid
THEN: It must return the a error message "invalid 2FA code" and a 401 code