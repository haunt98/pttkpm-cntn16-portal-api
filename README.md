# Awesome Portal

API endpoints
```
GET     /students/:mssv
POST    /students
DELETE  /students/:mssv
POST    /auth/login

POST    /programs
POST    /faculties

GET     /subjects/:id
POST    /subjects
POST    /type_subs

GET     /students/:mssv/try_enrolls/:status
POST    /try_enrolls
POST    /try_enrolls/delete
POST    /real_enrolls
POST    /enrolls
POST    /scores
```

Sample `.env`
```
PORT=8080
GIN_MODE=debug
DATABASE_URL=postgres://postgres@localhost:54320/postgres?sslmode=disable
DATABASE_MODE=debug
```