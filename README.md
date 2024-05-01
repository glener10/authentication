# **Authentication**

<div>
  <div>
    <table>
      <thead>
        <tr>
          <th colspan="4">Repository Informations</th>
          <th colspan="2">Open Tasks</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://img.shields.io/github/repo-size/glener10/authentication" alt="GitHub Repo size"></td>
          <td><img src="https://img.shields.io/github/stars/glener10/authentication" alt="GitHub Stars"></td>
          <td><img src="https://img.shields.io/github/forks/glener10/authentication" alt="Forks"></td>
          <td><img src="https://github.com/glener10/authentication/workflows/go/badge.svg" alt="Build Status"></td>
          <td><img src="https://img.shields.io/bitbucket/issues/glener10/authentication" alt="Open Issues"></td>
          <td><img src="https://img.shields.io/bitbucket/pr-raw/glener10/authentication" alt="Open Pull Requests"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div>
    <table>
      <thead>
        <tr>
          <th colspan="3">Last Updates</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://github.com/glener10/authentication/tags" alt="Last Tag"></td>
          <td><img src="https://github.com/glener10/authentication/releases/latest" alt="Last Release"></td>
          <td><img src="https://somsubhra.github.io/github-release-stats/?username=glener10&repository=authentication" alt="Last Release Stats"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div>
    <table>
      <thead>
        <tr>
          <th colspan="2">Copyright</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://img.shields.io/github/license/glener10/authentication" alt="License"></td>
          <td><img src="https://img.shields.io/github/contributors/glener10/authentication.svg" alt="Contributors"></td>
        </tr>
      </tbody>
    </table>
  </div>
</div>

<p align="center"> 🚀 Project created to train authentication issues, password recovery, login with third parties, permissions, etc. </p>

<h3>🏁 Table of Contents</h3>

<br>

===================

<!--ts-->

🐱‍🏍 [Features](#features)

💻 [Dependencies and Environment](#dependenciesandenvironment)

🚀 [Installing](#installing)

🧹 [Formatting the Code](#formatting)

🧪 [Testing](#testing)

☕ [Using](#using)

🔒 [License](#license)

👷 [Author](#author)

<!--te-->

===================

<div id="features"></div>

## 🐱‍🏍 **Features**

🧾 **Documentation**

- Migrations
- BDD (Behavior Driven Development) to use cases
- Gin Swagger to routes
- Concept of semantic versioning with tags and releases

⚙ **General**

- CI/CD process with github actions to perform code formatting check (golangci-lint), build and run automated tests
- Test setup with [TestContainers](https://testcontainers.com/):

  1- For each test switch/file that uses the database, a Postgres container is created just for testing

  2- Then all migrations are run in this container

  3- Before each test, a script is run to clean all records from the tables

  4- After executing the switch, the container is terminated

- Common middlewares to routes: block inactives users, rate limiter, timeout, only https, jwt signature checker for some routes, admin only for some routes, check 2fa when user has the 2fa activated

🏗 **Use Cases**

- active_2fa (need to be logged in): returns qrcode to synchronize with google authenticator
- desactive_2fa (need to be logged in)
- notify the user by email: when your password is changed and when your email is verified
- verify_change_email_code (need to be logged in): Verifies that the code is correct and not expired
- send_change_email_code (need to be logged in): Saves a code and an expiration time (5 minutes) in the database and sends an email with the code
- change_email (need to be logged in): It is necessary to use a unique code that is sent to the current email
- change_password_in_recovery: Verifies that the code is correct and not expired and change the password to the new password
- verify_password_recovery_code: Verifies that the code is correct and not expired
- send_password_recovery_code: Saves a code and an expiration time (5 minutes) in the database and sends an email with the code
- verify_email: Verifies that the code is correct and not expired and updates the email as verified
- send_email_verification_code: Saves a code and an expiration time (5 minutes) in the database and sends an email with the code
- admin elevation: you can promote anothers users to admin, delete users, inative user, find user information, list all users, list all logs, list all logs of a user
- log: all operations have log persistence with information such as: user id, operation code, method, route, success (true/false), ip and timestamp
- delete_user (need to be logged in): delete by id or e-mail
- find_user (need to be logged in): find by id or e-mail
- change_password (need to be logged in)
- login: With JWT
- create_user: Do not allow repeated emails and weak passwords

💡 **Technical Decisions**

- Clean Code
- Scream Architecture
- Commit Lint
- SOLID
- Clean Architecture

<div id="dependenciesandenvironment"></div>

## 💻 **Dependencies and Environment**

My dependencies and versions

[**Go**](https://golang.org/): go version go1.22.0 windows/amd64

[**Docker**](https://www.docker.com/): Docker version 25.0.3, build 4debf41

[**docker-compose**](https://docs.docker.com/compose/): Docker Compose version v2.24.5-desktop.1

<div id="installing"></div>

## 🚀 **Installing**

**1-** To install the dependencies you can run the following command in the root folder:

```
$ go mod tidy
$ go mod download
```

**OBS**: We have the development [.env](.env) file committed to the project, but you can change it as you see fit

**2-** (If you already have a PostgresSQL instance, you can skip this part) You will need a postgresSQL instance, we have a docker-compose ready to create a container, you can run the following command in the root folder

```
$ docker-compose up -d
```

**3-** Up the migrations: Naturally, when [running the server](#☕-using) it will execute the migrations, but they can be executed by code with (change pg url to yours):

```
$ migrate -database postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable -path src/db/migrations up
```

<div id="formatting"></div>

## 🧹 **Formatting the Code**

To check the code format you will need [instal golangci-lint](https://golangci-lint.run/welcome/install/) and run the following command in the root folder:

```
$ golangci-lint run
```

<div id="testing"></div>

## 🧪 **Testing**

To exec all the tests run the following command in the root folder:

```
$ go test -p 1 ./src/...
```

You can add the "**-v**" flag to see detailed output

```
$ go test -v -p 1 ./src/...
```

<div id="using"></div>

## ☕ **Using**

First, check the [dependencies](#dependenciesandenvironment) and the [installation](#installing) process:

Going to _root_ folder and exec:

```
$ go run .\main.go
```

Now you can open [http://localhost:8080](http://localhost:8080) with your browser to see the result.

You can see the routes in [Local Swagger Documentation](http://localhost:8080/swagger/index.html#) or you can see the routes documentation in '_rest_' folder, this files using de REST Client extension of VSCode, but you can export it any way you want

You can create new migrations using the command

```
migrate create -ext sql -dir src/db/migrations -seq MIGRATION_NAME
```

<div id="license"></div>

## 🔒 **License**

Projeto contêm [GNU GENERAL PUBLIC LICENSE](LICENSE).

<div id="author"></div>

#### **👷 Author**

Made by Glener Pizzolato! 🙋

[![Linkedin Badge](https://img.shields.io/badge/-Glener-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/glener-pizzolato/)](https://www.linkedin.com/in/glener-pizzolato-6319821b0/)
[![Gmail Badge](https://img.shields.io/badge/-glenerpizzolato@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:glenerpizzolato@gmail.com)](mailto:glenerpizzolato@gmail.com)
