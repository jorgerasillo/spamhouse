# spamhouse

[![CircleCI](https://circleci.com/gh/jorgerasillo/spamhouse/tree/main.svg?style=shield&circle-token=2811c84fa07cbb92e78ffcc1eb54762c6ee8e4ad)](https://circleci.com/gh/jorgerasillo/spamhouse) 

Spamhouse is a graphql server that adds ips to a local storage and detemines whether the IPs are in spamhous block list.

![Spam](./resources/spam.jpeg)

## Usage

### Run the server

`make run`

Access the browser UI at [http://localhost:8080](http://localhost:8080), Credentials are bootstrapped at application startup. Default credentials are: secureworks/supersecret to authenticate. To modify default credentials, specifying them as environment variables: `DEFAULT_USER` and `DEFAULT_USER_PASSWORD`.

The application port default is 8080, to modify the port the application is running on, modify the `PORT` number in the [Makefile](./Makefile)

#### Example mutation

```
mutation{
  enqueue(input: ["1.2.3.3"]){
    message
    node{
      ip_address
      uuid
    }
  }
}
```

#### Example query

```
query{
  getIPDetails(input: "1.2.3.3"){
    node{
      ip_address
      response_code
      uuid
      created_at
      updated_at
    }
  }
}
```

### Start server from scratch

`make dev` stop/build/run the appplication

### Stop the server

`make stop`

### Shell into the database

`make db-shell`

Creates a shell into sqlite3. See instructions for sqlite3 [here](https://www.sqlite.org/)

### Test server

`make test`

Runs go tests locally, makes uses of build tags to only trigger local tests with `!integration`

### Test integration

`make test-integration`

Runs integration tests, makes uses of build tags to only trigger integration tests.

### Development

- To extend database access, modify [repo.go](repo/repo.go)
- To modify the graphql resolvers, modify [gqlgen.yml](gqlgen.yml), after the changes are saved, run `make regenerate`
- To update auth (e.g. add database user validation) update [auth.go](pkg/middleware/auth/auth.go)


## Requirements

- [X] Your API must be protected with basic authentication
  - Username: secureworks
  - Password: supersecret
- [X] You should have a GraphQL API that is available via a /graphql endpoint
  - We suggest using https://gqlgen.com/
    - Mutations
      - enqueue(ip: [“ip1”, “ip2”])
        - This should kick off a background job to do the DNS lookup and store it in the database for each IP passed in for future lookups
        - If the lookup has already happened, this should queue it up again and update the response and updated_at fields

      - Queries
        - getIPDetails(ip: “ip address here”)
          - It should look up the IP in the database
          - If it’s not there, it should respond with the appropriate response code

          - This should have:
            - uuid ID
            - created_at time
            - updated_at time
            - response_code string
            - ip_address string

- [X] You should use SQLite as your database to make this portable
- [X] Dependencies should be handled using go mod
- [X] Tests should be written for the key components of the system
- [X] A README is required and should explain how to develop and run the project as if it
were a new team member working on it
- [X] The application should be packaged as a Dockerfile, and should accept a PORT
environment variable to know which port to run on
- [X] You can use any external libraries you want, but you must document and explain why
you’re using them

## External dependencies

- [gorm](https://gorm.io/index.html) - golang orm, used it for migration simplicity mainly.
- [logrus](https://github.com/sirupsen/logrus) - structured logger utility, better than printf :)
- [testify](https://github.com/stretchr/testify) - testify, utility for test assertion
- [backoff](https://github.com/cenkalti/backoff) - backoff utility, used for ensuring db is up before staring up the application
- [envconfig](https://github.com/cenkalti/backoff) - utility library for environment variable configuration
- [chi](https://github.com/go-chi/chi) - Added chi for injecting auth middleware
- [gqlgen](https://github.com/99designs/gqlgen) - graphql, per the requirements suggestion 
- [uuid](github.com/google/uuid) - generate UUID, used for generating uuid before inserting ip entry into database. See [BeforeCreate](./repo/model/model.go) hook in model

## Additional comments

- Initially I had implemented the queuing process to return the ips added upon enqueueing. See [commit](https://github.com/jorgerasillo/spamhouse/blob/40201f7bb756c409d8a3d9f771c30269943c615f/graph/schema.resolvers.go#L19), however, I recognized that ultimately this was not the best approach since it would dpeend on each ip returning after it had been refreshed.
