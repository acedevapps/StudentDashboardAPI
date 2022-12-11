<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">ACEDev Student Dashboard REST API v0.1</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> REST API written in Go 1.19 for the AICS Student Dashboard. Stores userdata in sqlite database and has 2 endpoints to request and post users.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

Write about 1-2 paragraphs describing the purpose of your project.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a basic idea of what the API entails and some example requests in python. To see the installations for prequisities use links in [built using](#⛏️-built-using)

### Prerequisites

The program requires:
- Go 1.19
- SQLite 3
- Gin (see built using to find link to installation)

### Example Requests

Example of get request in python:

```python
import requests

x = requests.get("http://localhost:8080/users")

print(x.text)
print(x.status_code)
```

Example of post request in python:

```python
import requests

x = requests.post("http://localhost:8080/users", json = {"id": "0", "name": "John Doe", "mail": "j.doe@gmail.com", "cred": "cred", "token": "token", "link": ""})

print(x.text)
print(x.status_code)
```

## 🎈 Usage <a name="usage"></a>

There are 2 main endpoints in this API both in /users.

The GET request for the /users header will return a json file with all users. The POST request for /users requires a json object to be passed in matching an instance of the go struct user. The properties which need to be passed in are:

- "id", string value but it does not matter as it will be overwritten
- "name", full name of student
- "mail", school mail of student
- "cred", jwt encoded Google API signon credential
- "token", active session token

There is also an optional property "link" representing the profile picture link

## 🚀 Deployment <a name = "deployment"></a>

The API is deployed on a server and currently is only accessed through localhost.

## ⛏️ Built Using <a name = "built_using"></a>

- [SQLite 3](https://www.sqlite.org/index.html) - Database
- [Golang](https://go.dev/) - Language
- [Gin](https://github.com/gin-gonic/gin) - Web Framework

## ✍️ Authors <a name = "authors"></a>

- [@priyacoding](https://github.com/priyacoding) - The whole damn thing

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hank's laptop
