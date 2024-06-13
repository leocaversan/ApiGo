<h1 align="center" 
    style="font-weight: bold;"
    >
    Rest with cache 💻
</h1>
<div align="center" >
    <img align="center" alt="GoLang" height="60" width="70" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg">
    <img align="center" alt="Mongo" height="40" width="70" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/redis/redis-original-wordmark.svg">
    <img align="center" alt="docker" height="60" width="80" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original-wordmark.svg">
</div>

<p align="center">
    • <a href="#started">Getting Started</a> 
    • <a href="#routes">API Endpoints</a> 
</p>

<p align="center">
    <b>
        This API is a simple payment simulation, that you can transfer credit between two users.
    </b>
</p>

<h2 id="started" align="center" >
    🚀 Getting started
</h2>

To run this project localy you will need to install `docker` and `docker compose`. After install thats components just follow this steps.

<h3> 
    Prerequisites 
</h3>

Here are all prerequisites necessary to runner the project.

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)


<h3>
    Starting
</h3>


- To run the app, run the command.
````bash
docker-compose up 
````
<!-- 
<h2 id="routes">
    📍 API Authentication
</h2>

To access the functions is required the userId and username to realize requests, those informations are required to pass for the cookies of the browser.

Was created a fake database to simulate the users, you can use this information to send requests.

````
session_id=22222
username=alice
````
 -->

<h2 id="routes">
    📍 API Endpoints
</h2>

Here you can list the main routes of your API, and what are their expected request.
​
| route               | description                                          
|----------------------|-----------------------------------------------------
| <kbd>GET /auth/</kbd>     | Authenticate user into the API, see [request details](#post-auth-detail)
| <kbd>GET /amount</kbd>     | Return amount aboult specific user on database, see [request details](#get-amount)
| <kbd>POST /sendMoney</kbd>     | Transfer money betwwen two users, see [request details](#post-sendMoney)
| <kbd>POST /create/</kbd>     | Register user in database, see [request details](#post-create)

<h3 id="post-auth-detail">
    POST /auth
</h3>

**REQUEST**
```json
{
  "username": "string",
  "password": "string"
}
```
**RESPONSE**
```json
{
  "id": "string"
}
```

<h3 id="get-amount">
    GET /amount/
</h3>

**REQUEST**
```json
{
  "cpf":"string"
}
```

**RESPONSE**
```json
{
  "Amount": float,
  "Name": "string",
  "status code": 200
}

```
<h3 id="post-sendMoney">
    POST /sendMoney/
</h3>

**REQUEST**
```json
{
  "value": float,
  "payer":"string-CPF",
  "payee":"string-CPF"
}

```
**RESPONSE**
```json
{
  "Sucess": "Sucess",
  "status code": 200
}
```
<h3 id="post-createUser">
    POST /createUser/
</h3>

**REQUEST**
```json
{
  "Name":"string",
  "cpf":"string",
  "Mail":"string",
  "Password":"string",
  "Amount": float,
  "Shopper": boolean
}
```

**RESPONSE**
```json
{
  "CPF": "string",
  "Name": "string",
  "Sucess": "User created successfully",
  "status code": 200
}