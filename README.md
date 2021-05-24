# spamhouse

[![CircleCI](https://circleci.com/gh/jorgerasillo/spamhouse/tree/main.svg?style=shield&circle-token=2811c84fa07cbb92e78ffcc1eb54762c6ee8e4ad)](https://circleci.com/gh/jorgerasillo/spamhouse) 

## Usage

### Run server

`make run`

### Test server

`make test`


## API


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
- [ ] Tests should be written for the key components of the system
- [ ] A README is required and should explain how to develop and run the project as if it
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

