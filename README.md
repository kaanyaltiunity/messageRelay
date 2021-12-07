# Dev Overview

## How to start the app
In order to start the app first run `docker-compose up` from the root of the project
Then within the `hub` directory run `go run main.go`

## Improvements
- I didn't have time to implement tests, so that would be the first that I would do. I would add unit tests for each of the components included in the service (eg UserRepositroy, RedisClientCache, MessageService).
  - I would first create unit tests. I paid attention to create a loosely coupled service with clear separations between layers. Application and Domain layer are at the core, they define the models and the business logic.
- I would have taken care of tracking users in a better way. Currently I am relying on a `userid` cookie to determine who's who. I would either keep track of sessions, or provide a jwt with an expiry date.
- I wanted to also dockerize the service and the client; however didn't have time to do so. This also means that I would have to implement a config management system to properly handle environment variables.