Route to inative a user (admin), implementing in repository to evite code duplication. update test of use cases "change_email", "change_password", "delete_user", "find_user" and "login"
Route to active a user (admin), implementing in repository to evite code duplication. update test of use cases "change_email", "change_password", "delete_user", "find_user" and "login"
Create pre-commit triggers with actions, to check interesting things like code checking, formatting, security checks and commit message pattern
CI/CD securely (_snyk_, _app.codacy_)
How to monitor and have observability
Security tests
Load Tests

E-mail verification
Password Recovery: use a unique token sended to email
2FA
Login with google


Notify user when a strange login occurs, necessary to save information such as IP address, geographic location and device used. When this happens, it is necessary to save the information from the new login device and send an email asking the user. If it wasn't, it should reset all known devices and inactivate the user
