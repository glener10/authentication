SCENARIO: Admin Find All Logs with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs with valid JWT from a admin
THEN: It must return the list of all logs and a 200 code

SCENARIO: Admin Find All Logs without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Admin Find All Logs without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs with a invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Admin Find All Logs without Success because the token provided is not from admin user
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs with a token of a non admin user
THEN: It must return a error message and a 401 code
