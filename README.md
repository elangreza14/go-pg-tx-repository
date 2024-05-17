# go pg tx repository

this project demonstrate easy way to handle transaction in the postgres repo. 

this project also include
- [X] separate the repo with every single entity
- [X] unit test for each repo
- [X] extending every repo with basic functionality (ex, get, create, edit)
- [X] combine transaction between entities
- [X] log with zap

to run the project 

1. install docker-compose, go, makefile and create .env file from sample `./example.env` 

2. run docker compose 
```
docker compose up -d
```

3. run migration 
```
make migrate 
```

4. run the program 
```
go run .
```


