SCENARIO: Admin active user with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /admin/users/active/:find with a valid search parameter (id or email) and valid JWT from a admin
THEN: It must return the success message and a 200 code

SCENARIO: Admin active user without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /admin/users/active/:find with a valid search parameter but without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Admin active user without Success because find param is invalid (even if the token is valid)
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /admin/users/active/:find with a invalid search parameter (not is email or id) and a valid admin token
THEN: It must return the a error message "wrong format, parameter need to be a id or a e-mail" and a 422 code

SCENARIO: Admin active user without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /admin/users/active/:find with a valid search parameter and invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Admin active user without Success because the token provided is not from admin user
GIVEN: The server is running and connected to the database without errors
WHEN: A POST request is made to /admin/users/active/:find with a valid search parameter but the token is from a non admin user
THEN: It must return a error message and a 401 code
