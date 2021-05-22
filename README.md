# spamhouse

## Usage

### Run server

`make run`

### Test server

`make test`


## API


## Requirements

- Your API must be protected with basic authentication
  - Username: secureworks
  - Password: supersecret
- You should have a GraphQL API that is available via a /graphql endpoint
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

- You should use SQLite as your database to make this portable
-  Dependencies should be handled using go mod
- Tests should be written for the key components of the system
- A README is required and should explain how to develop and run the project as if it
were a new team member working on it
- The application should be packaged as a Dockerfile, and should accept a PORT
environment variable to know which port to run on
- You can use any external libraries you want, but you must document and explain why
you’re using them

## Questions

1. Can the API query multiple IP addresses at a time?