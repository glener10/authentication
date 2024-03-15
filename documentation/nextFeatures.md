- Setup tests (integration) see the https://www.youtube.com/watch?v=x2C-n5ybyqQ&t=973s&ab_channel=ErickWendel
- All tests create_user (repository, usecase)

# UseCases

_All useCases: Documentation BDD, Route Documentation, Tests, Security Documentation_

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

- Docker Ambient
- Security documentation
- Create pre-commit triggers with actions, to check interesting things like code checking, formatting, security checks and commit message pattern
- Security Documentation of CreateUser
- Route documentation
- Response Format Standard
- Exception/Error Format Standard
- Good logging of application (Success/Error)
- How to monitor and have observability
- CI/CD securely (_snyk_, _app.codacy_)
- Security tests
- Load Tests
