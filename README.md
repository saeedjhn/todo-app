## Architecture Layers of the project

___

- **Router**
- **Controller**
- **Service(Use-case)**
- **Repository**
- **Domain(Entities)**

![backend-arch-diagram.png](assets%2Fbackend-arch-diagram.png "simorgh-backend-arch-diagram")

- ### Domain(Entities)

    - ***AdminUser***\
      <sub>**Properties:**</sub>\
      <sub>- Id</sub>\
      <sub>- Name</sub>\
      <sub>- Family</sub>\
      <sub>- Username</sub>\
      <sub>- Password</sub>\
      <sub>- BirthDate</sub>\
      <sub>- NationalCode</sub>\
      <sub>- Mobile</sub>\
      <sub>- Phone</sub>\
      <sub>- Status</sub>\
      <sub>- Creator</sub>\
      <sub>**Behaviors:**</sub>\
      <sub>- Register a user</sub>\
      <sub>- Log in user</sub>

## Major Packages used in this project

___

- **gin**: Gin is an HTTP web framework written in Go (Golang). It features a Martini-like API with much better
  performance -- up to 40 times faster. If you need a smashing performance, get yourself some Gin.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two
  parties. Used for Access Token and Refresh Token.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.

## How to run this project?

___

## The Complete Project Folder Structure

___

```
.
+-- assets
|   +-- backend-arch-diagram.png
+-- README.md

```

## API documentation of Simorgh Backend

___
