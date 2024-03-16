Route documentation of create_user
Security Documentation of CreateUser

Login: Use JWT
Interceptor/Guard of request for private routes to check JWT
Account Update (Logged): Do not allow invalid email, do not allow weak password, do not allow repeated email
Delete Account (Logged)
User information (logger)
Password Recovery: use a unique token
Login with google
E-mail verification
2FA
User Control: Log out, block, delete (Only users with permission can due)
Database backups, rules and administration
Use notification when strange login ocurred
Operation history: User, location, IP, success or not, date, time

Docker Ambient
Security documentation
Create pre-commit triggers with actions, to check interesting things like code checking, formatting, security checks and commit message pattern
Response Format Standard
Exception/Error Format Standard
Good logging of application (Success/Error)
How to monitor and have observability
CI/CD securely (_snyk_, _app.codacy_)
Security tests
Load Tests
