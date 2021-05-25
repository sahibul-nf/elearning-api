# API Spec

## Create Comment

Request 🔥

- Method : POST
- Endpoint : `/api/v1/comments`
- Header :
  - Content_Type : application/json
  - Accept : application/json
- Params :
    <!-- - comment -->
- Body :

```json
{
  "comment": "string"
}
```

Response 🚀

```json
{
  "meta": {
    "message": "string",
    "code": "int",
    "status": "string"
  },
  "data": {
    "id": "int",
    "email": "string",
    "username": "string",
    "avatar_file_name": "string",
    "comment": "string",
    "created_at": "date"
  }
}
```

\
.

## Get Comments

Request 🔥

- Method : GET
- Endpoint : `/api/v1/comments`
- Header :
  - Accept : application/json

Response 🚀

```json
{
  "meta": {
    "message": "string",
    "code": "int",
    "status": "string"
  },
  "data": {
    "total": "int",
    "list": [
      {
        "id": "int",
        "email": "string",
        "username": "string",
        "avatar_file_name": "string",
        "comment": "string",
        "created_at": "date"
      },
      {
        "id": "int",
        "email": "string",
        "username": "string",
        "avatar_file_name": "string",
        "comment": "string",
        "created_at": "date"
      }
    ]
  }
}
```

\
.

## Delete Comment

Request 🔥

- Method : DELETE
- Endpoint : `/api/v1/comments/{id_comment}`
- Header :
  - Accept : application/json

Response 🚀

```json
{
  "meta": {
    "message": "string",
    "code": "int",
    "status": "string"
  },
  "data": {
    "is_deleted": "bool"
  }
}
```
