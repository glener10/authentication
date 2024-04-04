SCENARIO: Admin Find All Users with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/users with valid JWT from a admin
THEN: It must return the list of all users and a 200 code

SCENARIO: Admin Find All Users without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/users without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Admin Find All Users without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/users with a invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Admin Find All Users without Success because the token provided is not from admin user
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/users with a token of a non admin user
THEN: It must return a error message and a 401 code
