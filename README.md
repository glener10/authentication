# **Authentication**

<div style="display: flex; justify-content: center; align-items: center; flex-wrap: wrap;">
  <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="3">Repository Informations</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://img.shields.io/github/repo-size/glener10/authentication" alt="GitHub Repo size"></td>
          <td><img src="https://img.shields.io/github/stars/glener10/authentication" alt="GitHub Stars"></td>
          <td><img src="https://img.shields.io/github/forks/glener10/authentication" alt="Forks"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="2">Open Tasks</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://img.shields.io/bitbucket/issues/glener10/authentication" alt="Open Issues"></td>
          <td><img src="https://img.shields.io/bitbucket/pr-raw/glener10/authentication" alt="Open Pull Requests"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="2">Current version</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://github.com/glener10/authentication/workflows/go/badge.svg" alt="Build Status"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="4">Last Updates</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://github.com/glener10/authentication/commits" alt="Last Commit"></td>
          <td><img src="https://github.com/glener10/authentication/tags" alt="Last Tag"></td>
          <td><img src="https://github.com/glener10/authentication/releases/latest" alt="Last Release"></td>
          <td><img src="https://somsubhra.github.io/github-release-stats/?username=glener10&repository=REPOSITORIONAME" alt="Last Release Stats"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="1">Docker</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://hub.docker.com/r/glener10/REPOSITORIONAME" alt="Docker"></td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- <div style="display: flex; justify-content: center; align-items: center;">
    <table>
      <thead>
        <tr>
          <th colspan="2">Security</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://snyk.io/test/github/glener10/REPOSITORIONAME?targetFile=app%2Fbuild.gradle" alt="Know Vulnerabilities"></td>
          <td><img src="https://app.codacy.com/gh/glener10/REPOSITORIONAME/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade" alt="Codacy"></td>
        </tr>
      </tbody>
    </table>
  </div> -->

  <div style="display: flex; justify-content: center; align-items: center;">
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

💡 [Technical Decisions](#technical)

📖 [Learn More](#learnmore)

🔒 [License](#license)

👷 [Author](#author)

<!--te-->

===================

<div id="features"></div>

## 🐱‍🏍 **Features**

🧾 **Documentation**

- Migrations
- BDD (Behavior Driven Development) to use cases
- The release and tag concept

⚙ **General**

-

🏗 **Use Cases**

-

<div id="dependenciesandenvironment"></div>

## 💻 **Dependencies and Environment**

My dependencies and versions

**Go**: go version go1.22.0 windows/amd64

**Docker**: Docker version 25.0.3, build 4debf41

**docker-compose**: Docker Compose version v2.24.5-desktop.1

<div id="installing"></div>

## 🚀 **Installing**

**1-** To install the dependencies you can run the following command in the root folder:

```
$ go mod download
```

**OBS**: We have the development .env file committed to the project, but you can change it as you see fit

**2-** (If you already have a PostgresSQL instance, you can skip this part) You will need a postgresSQL instance, we have a docker-compose ready to create a container, you can run the following command in the root folder

```
$ docker-compose up -d
```

**3-** Up the migrations with

```
$ migrate -database postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable -path src/db/migrations up
```

<div id="formatting"></div>

## 🧹 **Formatting the Code**

To check the code format you can run the following command in the root folder:

```
$ golangci-lint run
```

<div id="testing"></div>

## 🧪 **Testing**

To exec all the tests run the following command in the root folder:

```
$ go test -p 1 ./src/...
```

<div id="using"></div>

## ☕ **Using**

First, check the [dependencies](#dependenciesandenvironment) and the [installation](#installing) process:

Going to _root_ folder and exec:

```
$ go run .\main.go
```

Now you can open [http://localhost:8080](http://localhost:8080) with your browser to see the result.

You can see the routes documentation in '_rest_' folder, this files using de REST Client extension of VSCode, but you can export it any way you want

You can create new migrations using the command

```
migrate create -ext sql -dir src/db/migrations -seq MIGRATION_NAME
```

<div id="technical"></div>

## 💡 **Technical Decisions**

The project seeks to use some programming paradigms such as:

- Clean Code
- Scream Architecture
- Commit Lint

<div id="learnmore"></div>

## 📖 **Learn More**

To learn more about technologies used in the application:

- [Go](https://golang.org/) - learn about Go features and API.

- [Docker](https://www.docker.com/) - learn about Docker features and API.

- [Docker Compose](https://docs.docker.com/compose/) - learn about Docker Compose features and API.

<div id="license"></div>

## 🔒 **License**

Projeto contêm [GNU GENERAL PUBLIC LICENSE](LICENSE).

<div id="author"></div>

#### **👷 Author**

Made by Glener Pizzolato! 🙋

[![Linkedin Badge](https://img.shields.io/badge/-Glener-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/glener-pizzolato/)](https://www.linkedin.com/in/glener-pizzolato-6319821b0/)
[![Gmail Badge](https://img.shields.io/badge/-glenerpizzolato@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:glenerpizzolato@gmail.com)](mailto:glenerpizzolato@gmail.com)
