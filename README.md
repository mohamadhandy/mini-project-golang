# Mini Project Golang Celerates CAP 2022
1 domain (foods), 1 domain for register and login (member).

Creating a 2 domain REST API with gin, gorm, PostgreSQL.

## Inside mini project

- CRUD FOODS
- Login Member
- Register Member
- JWT Token
- Hexagonal architecture

## Postgre SQL

SQL used in this project

### members
```sql
CREATE TABLE "members" (
  "member_id" serial primary key not null,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
```

### foods
```sql
CREATE TABLE "foods" (
  "food_id" SERIAL PRIMARY KEY NOT NULL,
  "food_name" varchar(100) NOT NULL,
  "food_price" int NOT NULL,
  "food_stock" int NOT NULL,
  "food_status" smallint NOT NULL DEFAULT '1'
);

INSERT INTO "foods" (food_name, food_price, food_stock, food_status)
VALUES
	('Kopiko',1000,90,1),
	('Astor',2000,80,1),
	('Bengbeng',3000,70,1),
	('Choki-choki',1500,60,0);
```

## REST API
There are 7 endpoints that can be called. That is:

### GET FOODS
```
/api/v1/foods?page=1&limit=3
```

### EXAMPLE RESULT GET FOODS
```
{
    "meta": {
        "message": "Success get all food!",
        "code": 200,
        "status": "success"
    },
    "data": {
        "limit": 3,
        "page": 1,
        "total_rows": 5,
        "from_row": 4,
        "to_row": 0,
        "rows": [
            {
                "food_id": 4,
                "food_name": "Choki-choki",
                "food_price": 1500,
                "food_stock": 60,
                "food_status": "0"
            },
            {
                "food_id": 5,
                "food_name": "test oreo lll",
                "food_price": 6000,
                "food_stock": 40,
                "food_status": "1"
            }
        ]
    }
}
```

### GET SINGLE FOODS
```
/api/v1/foods/4
```

### EXAMPLE RESULT GET SINGLE FOODS
```
{
    "meta": {
        "message": "Success get single food with id 4",
        "code": 200,
        "status": "success"
    },
    "data": {
        "food_id": 4,
        "food_name": "Choki-choki",
        "food_price": 1500,
        "food_stock": 60,
        "food_status": "0"
    }
}
```

### PUT FOODS
```
/api/v1/foods/5
```

### BODY PUT FOODS
```
{
    "food_name": "menu baru",
    "food_price": 100,
    "food_status": "1",
    "food_stock": 250
}
```

### EXAMPLE RESULT PUT FOODS
```
{
    "meta": {
        "message": "Success update food with id 5",
        "code": 200,
        "status": "success"
    },
    "data": {
        "food_id": 5,
        "food_name": "menu baru",
        "food_price": 100,
        "food_stock": 250,
        "food_status": "1"
    }
}
```

### POST FOODS
```
/api/v1/foods/5
```

### BODY POST FOODS
```
{
    "food_name": "test oreo lll",
    "food_price": 6000,
    "food_status": "1",
    "food_stock": 40
}
```

### EXAMPLE RESULT POST FOODS
```
{
    "meta": {
        "message": "Your Food has been created",
        "code": 200,
        "status": "success"
    },
    "data": {
        "food_id": 5,
        "food_name": "test oreo lll",
        "food_price": 6000,
        "food_stock": 40,
        "food_status": "1"
    }
}
```

### DELETE SINGLE FOODS
```
/api/v1/foods/5
```

### EXAMPLE RESULT DELETE SINGLE FOODS
```
{
    "meta": {
        "message": "Success delete food with id 5",
        "code": 200,
        "status": "success"
    },
    "data": null
}
```

### REGISTER MEMBER
```
/api/v1/members
```

### BODY POST MEMBERS
```
{
    "email": "kuy@handy.com",
    "password": "pass",
    "role": "user"
}
```

### EXAMPLE RESULT POST FOODS
```
{
    "meta": {
        "message": "Register member success!",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 8,
        "email": "kuy@handy.com",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4fQ.UXQdR3dLOvVSgBHhLU8YSGcHeJF7TNrOFMewS6AkWYE"
    }
}
```

### REGISTER MEMBER
```
/api/v1/sessions
```

### BODY POST LOGIN
```
{
    "email": "yuk@handynugraha.com",
    "password": "pass"
}
```

### EXAMPLE RESULT LOGIN
```
{
    "meta": {
        "message": "Login member success!",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 6,
        "email": "yuk@handynugraha.com",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2fQ.F-8Wkdr-84JAw0CGhloMpbKFi2QwpXvSbvXBCtoYrHs"
    }
}
```

## License
Mohamad Handy Nugraha
