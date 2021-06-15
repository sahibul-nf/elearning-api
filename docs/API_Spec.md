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

## Get Comments by Forum

Request 🔥

- Method : GET
- Endpoint : `/api/v1/comments/forum?forum_id=1`
- Header :
  - Accept : application/json
- Params :
  - forum_id
  
Response 🚀

```json
{
  "meta": {
    "message": "Successfuly get comments by that forumID",
    "code": 200,
    "status": "success"
  },
  "data": [
    {
      "id": 12,
      "user_id": 2,
      "forum_id": 1,
      "task_id": 0,
      "comment": "Hello worlsd",
      "created_at": "2021-06-14T00:52:25+07:00"
    },
    {
      "id": 14,
      "user_id": 1,
      "forum_id": 1,
      "task_id": 0,
      "comment": "Golang is awesome",
      "created_at": "2021-06-14T01:44:43+07:00"
    }
  ]
}
```

\
.

## Get Comments by Task

Request 🔥

- Method : GET
- Endpoint : `/api/v1/comments/task?task_id=1`
- Header :
  - Accept : application/json
- Params :
  - task_id

Response 🚀

```json
{
  "meta": {
    "message": "Successfuly get comments by that taskID",
    "code": 200,
    "status": "success"
  },
  "data": [
    {
      "id": 13,
      "user_id": 2,
      "forum_id": 0,
      "task_id": 1,
      "comment": "Keep learning",
      "created_at": "2021-06-14T00:53:19+07:00"
    }
  ]
}
```


\
.

## Delete Comment

Request 🔥

- Method : DELETE
- Endpoint : `/api/v1/comments/:id`
- Header :
  - Accept : application/json
- Params :
  - id

Response 🚀

```json
{
  "meta": {
    "message": "Successfuly to get campaign detail",
    "code": 200,
    "status": "success"
  },
  "data": {
    "is_deleted": true
  }
}
```
