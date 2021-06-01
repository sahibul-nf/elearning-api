# Entity

- Users
- Comments
- Classes
- Forums
- Taks
- SubmittedTaks

\`

# Entity Relationchip Design

```plantuml
@startuml
    entity Users {
        *id : INT
        --
        name : VARCHAR
        email : VARCHAR
        role : TINYINT
        token : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Comments {
        *id : INT
        --
        ' class_id : INT
        user_id : INT
        comment : TEXT
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Classes {
        *id : VARCHAR, unique
        --
        user_id : INT
        name : VARCHAR
        course : VARCHAR
        unit : VARCHAR
        part : VARCHAR
        banner_image_file_name : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Forums {
        *id : INT
        --
        user_id : INT
        description : TEXT
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Forum_Files {
        *id : INT
        --
        forum_id : INT
        file_name : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Tasks {
        *id : INT
        --
        ' class_id : INT
        user_id : INT
        title : VARCHAR
        description : TEXT
        deadline: DATETIME
        goal_point : INT
        status : ["submitted", "given", "rated"]
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Task_Files {
        *id : INT
        --
        task_id : INT
        file_name : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity User_Submitted_Tasks {
        *id : INT
        --
        user_submitted_class_id : INT
        ' mission_id : INT
        score : INT
        status : ["assignmed", "submitted"]
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity User_Submitted_Task_Files {
        *id : INT
        --
        user_submitted_task_id : INT
        file_name : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Users_Classes {
        *id : INT
        --
        user_id : INT
        class_id : VARCHAR
    }

    entity Users_Forums_Classes {
        *id : INT
        --
        user_id : INT
        forum_id : INT
        class_id : VARCHAR
        comment_id : INT
    }

    entity Users_Tasks_Classes {
        *id : INT
        --
        user_id : INT
        task_id : INT
        class_id : VARCHAR
        comment_id : INT
    }

    Users ||--o{ Classes : Create
    Users ||--o{ Users_Classes : Has Classes
    Users_Classes }o--|| Classes : Has Users

    Users ||--o{ Forums : Create
    Users ||--|{ Users_Forums_Classes : Has Forums Classes
    Classes ||--|{ Users_Forums_Classes : Has Users Forums
    Users_Forums_Classes }--|| Forums : Has Users Classes

    Users ||--o{ Comments : Create
    Comments ||--o{ Users_Forums_Classes : Has 
    Comments ||--o{ Users_Tasks_Classes : Has

    Users ||--o{ Tasks : Create
    Users ||--|{ Users_Tasks_Classes : Has
    Classes ||--|{ Users_Tasks_Classes : Has
    Tasks ||--|{ Users_Tasks_Classes : Has

    Users_Tasks_Classes ||--o{ User_Submitted_Tasks : Submit

    Forums ||--o{ Forum_Files : Has
    Tasks ||--o{ Task_Files : Has
    User_Submitted_Tasks ||--o{ User_Submitted_Task_Files : Has

@enduml
```
