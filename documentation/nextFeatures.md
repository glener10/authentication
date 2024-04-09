Update master and check badges
Create pre-commit triggers with actions, to check interesting things like code checking, formatting, security checks and commit message pattern
CI/CD securely (_snyk_, _app.codacy_)

E-mail verification
Password Recovery: use a unique token sended to email
2FA
Login with google

How to monitor and have observability
Security tests
Load Tests

Improvement feature: Notify user when a strange login occurs, necessary to save information such as IP address, geographic location and device used. When this happens, it is necessary to save the information from the new login device and send an email asking the user. If it wasn't, it should reset all known devices and inactivate the user
Improvement feature: Common middleware to block operation from another user if doesnt admin: BlockOperationFromAnotherUserIfNotAdmin, removing Jwt Middleware because just check signature and another middlewares already due this