SCENARIO: Find User
GIVEN: The server is running and connected to the database without errors
WHEN: A GET request is made to /ser with a valid search parameter and valid JWT from the owner of the search parameter
THEN: It must return the user information without password and a 200 code
