- CI/CD process (Lint, Build, Tests)
- Setup tests (integration)
- Formatting Badges of README.md
- Database creation (basic access)
- Docker Ambient
- Security documentation
- Active all GitHub security features

# UseCases

_All useCases: Documentation BDD, Route Documentation, Tests, Security Documentation_

**Account Creation**

- Do not allow invalid email
- Do not allow weak password
- Do not allow repeated email

**Login**

- Use JWT

**Interceptor/Guard of request for private routes to check JWT**

**Account Update (Logged)**

- Do not allow invalid email
- Do not allow weak password
- Do not allow repeated email

**Delete Account (Logged)**

**Private Route (Logged)**

- Return use information

**Password Recovery**

- Use a unique token

**Login with google**

**E-mail verification**

**2FA**

**Control with use**

- Log out, block, delete (Only users with permission can due)

**Database backups, rules and administration**

**Use notification when strange login ocurred**

**Operation history**

- User, location, IP, success or not, date, time

-> Response Format Standard
-> Exception/Error Format Standard
-> Good logging of application (Success/Error)
-> How to monitor and have observability
-> CI/CD securely (_snyk_, _app.codacy_)
-> Security tests
-> Load Tests
