SCENARIO: Admin Find User All Logs with Success
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs/:find with valid JWT from a admin
THEN: It must return the list of all logs of the user and a 200 code

SCENARIO: Admin Find User All Logs without Success because token is not informed
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs/:find without authorization header
THEN: It must return the a error message "token not provided" and a 401 code

SCENARIO: Admin Find User All Logs without Success because token is not valid
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs/:find with a invalid authorization header
THEN: It must return the a error message "invalid token" and a 401 code

SCENARIO: Admin Find User All Logs without Success because the token provided is not from admin user
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs/:find with a token of a non admin user
THEN: It must return a error message and a 401 code

SCENARIO: Admin Find User All Logs without Success because the find param is not a id
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /admin/logs/:find with a token of a admin user but the find param is in wrong format
THEN: It must return a error message and a 422 code