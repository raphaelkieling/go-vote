# Go-Vote

A simple api for study. I have some goals:
- Allow a big concurrent requests with concurrent write/read database
- Allow the same ip make multiple votes but create a hash (like a blockchain algorithm) to make the process harder.
- Use basic auth to have some security, but i don't need login (jwt, auth0)
- Make unit tests
- Make integration tests with docker-compose and git ci/cd

## How run

```sh
# To create a mysql database
docker-compose up

# Create a env
cp env-example .env

go run .
```