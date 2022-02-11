# boilerplate-service

Awesome description here..

## Requirement
1. Go 1.15 or above (download [here](https://golang.org/dl/)).
2. Git (download [here](https://git-scm.com/downloads)).
3. Docker

## How To Setup
### Project
clone the source locally:
```shell
$ git clone https://github.com/verryp/cake-store-service.git <your-name-project>
$ cd <your-name-project>
```
Copy the example env file and make the required configuration changes in the `.env` file
```shell
$ cp .env.example .env
```
Start the local development server
```shell
$ make run
```
You can now access the server at http://localhost:<server-port>

### Database
You can run `docker-compose up -d` on root project to running the database (optional) and run command `make run ARGS=migrateup` to migrate database table.

## How To Use
We are encouraging makefile to operate this project. Each command has its own scope & roles.
Use `make help` to see complete command and details.

- Run service, `make run`
- Run a DB migration, `make run ARGS=migrateup`
- Rollback DB migration, `make run ARGS=migratedown`
- Run unit test, `make test`
- Manage container, `make compose.up`, `make compose.down`, `make compose.clean`

## How to contribute
1. Create your feature branch (`git checkout -b [feat/fix/ref]-foo-bar`)
2. Commit your changes (`git commit -am 'feat: add some fooBar'`)
3. Push to the branch (`git push origin [feat/fix/ref]-foo-bar`)
4. Create a new Merge Request, please use this template when submitting a MR