<p align="center">
  <img src="https://user-images.githubusercontent.com/25181517/192149581-88194d20-1a37-4be8-8801-5dc0017ffbbe.png" width="100">
</p>
<h1 align="center">AvitoBTA2025-merchshop</h1>
<h3 align="center">2025 Entrance <a href="https://github.com/avito-tech/tech-internship/blob/main/Tech%20Internships/Backend/Backend-trainee-assignment-winter-2025/Backend-trainee-assignment-winter-2025.md">Tests</a> for future <a href="https://avito.tech/">«Avito.Tech»</a> interns at <a href="https://www.avito.ru/company">«Avito»</a></h3>
<p align="center">The service provides an HTTP API for shopping in Avito's branded merch store. You can also give coins to friends or other buyers. Also, at the first registration we give everyone 1000 coins!</p>

---

### — _Technology:_
![PostgreSQL](https://img.shields.io/badge/postgreSQL-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

---

### — _How to Install and Use:_
#### 🟢 **Git Clone:**
```
git clone https://github.com/DigiRon4ik/AvitoBTA2025-merchshop.git
cd AvitoBTA2025-merchshop
```

> [!TIP]
> If you don't change the source code, when the service is raised, the builder container will work with cache, skipping capacitive processes.

#### 🟢 **Docker | Start/Stop:**
- 🚀
  ```
  docker-compose up -d
  ```
- ⛔
  ```
  docker-compose down
  ```
#### 🟢 **Make | Start/Stop:**
- 🚀
  ```
  make d-up
  ```
- ⛔
  ```
  make d-down
  ```
<details>
    <summary>🟢 <b>Make | Other:</b></summary>

- Rolling migrations to a real DB:
  ```
  make m-up
  ```
- Rollback migrations in the real DB:
  ```
  make m-down
  ```
- Checking the migration version in the actual DB:
  ```
  make m-status
  ```
- Rebuilds and service uplifts:
  ```
  make d-up-b
  ```
- Stopping the service and deleting all data::
  ```
  make d-down-v
  ```
- Raising the service if there is data left in the database:
  ```
  make d-up-app
  ```
- Starting the linter:
  ```
  make lint
  ```
- Running only unit tests:
  ```
  make tests
  ```
- Running unit and integration tests:
  ```
  make tests-integration
  ```
- Run only unit tests with results in HTML:
  ```
  make cover
  ```
- Running unit and integration tests with results in HTML:
  ```
  make cover-integration
  ```
</details>

---

### — _APIs:_
The developed API is implemented according to the REST API design.

- **[Postman](https://www.postman.com/downloads/)**
  - You can import a [JSON-file](https://github.com/avito-tech/tech-internship/blob/main/Tech%20Internships/Backend/Backend-trainee-assignment-winter-2025/schema.json) from the API into Postman to run queries.

<div align="center">

| Name           |     Method | API              |                      Body                      |
|:---------------|-----------:|:-----------------|:----------------------------------------------:|
| Authentication |   **POST** | `/api/auth`      | `{"username"}: <string>, "password": <string>` |
| Info           |    **GET** | `/api/info`      |                       -                        |
| Send Coin      |   **POST** | `/api/sendCoin`  |  `{"toUser": <string>, "amount": <integer>}`   |
| Buy Item       |    **GET** | `/api/buy/:item` |                       -                        |

</div>

---

### — _DataBase Schema:_
<p align="center">
  <a href="https://i.imgur.com/OCN8iJv.png">
    <img width=800 src="https://i.imgur.com/OCN8iJv.png" >
  </a>
</p>

---

### — _Progress in the fulfillment of assigned tasks:_

<div align="center">

| Task                                                                                                                                                                                         | Progress | Comment                                                                      |
|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------:|:-----------------------------------------------------------------------------|
| **Use this [API](https://github.com/avito-tech/tech-internship/blob/main/Tech%20Internships/Backend/Backend-trainee-assignment-winter-2025/schema.json)**                                    | ✅ | Implemented the given API                                                    |
| **Employees can be up to 100k, RPS is 1k, response time SLI is 50ms, response success SLI is 99.99%**                                                                                        | ✅ | That seems about right                                                       |
| **JWT must be used for access authorization. User token of API access is issued after user authorization/registration. At the first authorization the user should be created automatically** | ✅ | No RefreshToken for now                                                      |
| **Implement unit test coverage of business scenarios. Total test coverage of the project should exceed 40%**                                                                                 | ✅ | Business logic ≈ 90% and the whole project including integration tests ≈ 45% |
| **Implement an integration or E2E test for a merch purchase scenario**                                                                                                                       | ✅ | Realized in integration form with mocks                                      |
| **Implement an integration or E2E test for the scenario of passing coins to other employees**                                                                                                | ✅ | Realized in integration form with mocks                                      |
| **AddTask No. 1 (*Perform load testing of the obtained solution and attach the results of testing*)**                                                                                        | ❌ | In the process...                                                            |
| **AddTask No. 2 (*Implement integration or E2E testing for the remaining scenarios*)**                                                                                                       | 🙈 | Not really. Also in process...                                               |
| **AddTask No. 3 (*Describe the linter configuration (.golangci.yaml in the root of the project for go)*)**                                                                                   | ✅ | -                                                                            |

</div>

---

### — _Video:_
In the process...

[//]: # (<p align="center">)

[//]: # (  <a href="https://youtu.be/27ToZvGJTVY">)

[//]: # (    <img width=800 src="https://i.imgur.com/bqTOsir.png" >)

[//]: # (  </a>)

[//]: # (</p>)

---

### — _Further development:_
1. Integrate Swagger and Swagger UI
2. Implement RefreshToken and a handle to refresh AccessToken
3. Realize CI/CD pipelines for internet deployment

---

### — _Q&A:_
- Which tests should I choose for the scenarios, integration or E2E?
  - I chose integration ones as they are simpler and cheaper.
- Which libraries/packages to choose?
  - Did a comparative analysis between the popular ones, and chose the best ones. For example, for the server - Gin: because, easier to implement and wanted something new).
- What validations should I do for username, for example?
  - Made less standardized requirements for incoming fields.

---
