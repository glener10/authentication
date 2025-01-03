# [v1.0.0] - 03/05/2024 (dd/mm/YYYY)

- login_2fa: use case, all tests, bdd and route documentation
- middleware to check 2fa when user has the 2fa activated
- active_2fa: use case, all tests, bdd and route documentation
- desactive_2fa: use case, all tests, bdd and route documentation
- notify the user by email: when your password is changed and when your email is verified
- create gateway to send email
- verify_change_email_code: Verifies that the code is correct and not expired
- send_change_email_code: use case, all tests, bdd and route documentation
- ALTERING USECASE change_email: Now it is necessary to use a unique code that is sent to the current email
- change_password_in_recovery: Verifies that the code is correct and not expired and change the password to the new password
- verify_password_recovery_code: Verifies that the code is correct and not expired
- send_password_recovery_code: use case, all tests, bdd and route documentation
- verify_email (need to be logged in): use case, all tests, bdd and route documentation
- send_email_verification_code (need to be logged in): use case, all tests, bdd and route documentation

# [v0.0.1] - 12/04/2024

- change_email (need to be logged in) use case, all tests, bdd and route documentation
- change_password (need to be logged in) use case, all tests, bdd and route documentation
- jwt middleware to check if the token signature is valid
- find_user (need to be logged in) use case, all tests, bdd and route documentation
- login use case, all tests, bdd and route documentation
- create_user use case, all tests, bdd and route documentation
- Setup integration tests with testcontainers (https://testcontainers.com/), up postgres container for tests, run migrations and clean database before each test
- CI/CD to exec lint, build and automated tests
- middlewares: rate limiter, timeout and only https
- Enable all possible security features on the GitHub platform
- Fisrt version of code documentation
- Migration to create user table
