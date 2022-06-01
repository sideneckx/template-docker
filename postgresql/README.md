# Instruction for newbie
- By default, when you put the environment:
```Docker
environment:
	- POSTGRES_USER=postgres
	- POSTGRES_PASSWORD=postgres
	- POSTGRES_DB=postgres
```
It's will create a user, password, schema with the given name "postgres"
- Start postgres docker with command:
```bash
make up
```
- And end it with command:
```bash
make down
```

*Note that when the compose is downed, all the data store in this instance will be export to `./db` folder*