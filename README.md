# RUN

Please run main.go file and listen `8000` port.

```bash
$ go get -v
$ go run mian.go
```
Next Endpoints are available:
```text
POST   /login                 -> Create token
GET    /auth/refresh_token
GET    /auth/events           -> Get Events (Event location comes from JWT)
```

Download and install [httpie](https://github.com/jkbrzt/httpie) CLI HTTP client. Your can use Curl or PostMan, but examples below with [httpie](https://github.com/jkbrzt/httpie)

## Generate token:
```bash
$ http -v --json POST localhost:8000/login username=admin password=admin
```
response example:
```
POST /login HTTP/1.1
Accept: application/json
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 42
Content-Type: application/json
Host: localhost:8000
User-Agent: HTTPie/0.9.4

{
    "password": "admin",
    "username": "admin"
}

HTTP/1.1 200 OK
Content-Length: 311
Content-Type: application/json; charset=utf-8
Date: Sun, 27 May 2018 11:48:34 GMT

{
    "code": 200,
    "expire": "2018-05-28T21:48:34+10:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTI3NTA4MTE0LCJpZCI6ImFkbWluIiwibG9jYXRpb24iOiJDb3JuaW5nIiwibmFtZSI6IlRpbmEgVHVybmVyIiwib3JpZ19pYXQiOjE1Mjc0MjE3MTQsInN1YiI6IjBwd21hNmdqMTAifQ.MiWijlItZw8J2Ti9mnGJv6eU9mepsn0vySpwkymD6Cs"                                                
}
```
Please not that payload for token randomly generated


## Get Events:
Use token form step above `"Authorization:Bearer <TOKEN>"` to get events. Events Randomly generates, with `location` from JWT payload.
```bash
$ http -f GET localhost:8000/auth/events "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTI3NTAxNzE3LCJpZCI6ImFkbWluIiwibG9jYXRpb24iOiJZdWJhIENpdHkiLCJuYW1lIjoiRG9yb3RoeSBCdXRsZXIiLCJvcmlnX2lhdCI6MTUyNzQxNTMxNywic3ViIjoiZzNwZXdjanRzMyJ9.rcdBH7NTP8zvLpmsHSb_cveJ6ECPCkFko4fMtITo658"
```

Response example:
```json
{
    "events": [
        {
            "Title": "SEQUI QUIDEM NON",
            "availableSeats": 7278511,
            "date": "2018-08-29T12:00:21Z",
            "image": "8si3lwj1j0.jpg",
            "location": "Corning"
        },
        {
            "Title": "OMNIS REPREHENDERIT",
            "availableSeats": 7455089,
            "date": "2018-07-28T12:00:21Z",
            "image": "s6v0vz7sme.jpg",
            "location": "Corning"
        },
        {
            "Title": "ID AUT",
            "availableSeats": 6933274,
            "date": "2018-06-24T12:00:21Z",
            "image": "mz2y89by8j.jpg",
            "location": "Corning"
        },
        {
            "Title": "DEBITIS REPREHENDERIT",
            "availableSeats": 9431445,
            "date": "2018-06-07T12:00:21Z",
            "image": "qg0fqrj2pk.jpg",
            "location": "Corning"
        },
        {
            "Title": "ET AUT OMNIS NEQUE MINIMA",
            "availableSeats": 9339106,
            "date": "2018-07-03T12:00:21Z",
            "image": "co2aaznohr.jpg",
            "location": "Corning"
        }
    ]
}
```

# Test
To run test 
```bash
$ go get -v
$ go test -v
```
#### Expected result
```bash
  === RUN   TestGetUserEvents
  
    Test GetUserEvents 
      Check correct number of events ✔
      Check correct model ✔
      Check that all events match Location from JWT ✔✔✔✔✔
  
  
  7 total assertions
  
  --- PASS: TestGetUserEvents (0.00s)
  PASS
  ok      go-jwt  0.006s
```

#### Test Cover
```bash
$ go test -covermode=count 
.......
7 total assertions

PASS
coverage: 52.6% of statements
ok      go-jwt  0.006s
```

# Swagger
`swagger.json` - build from annotations with go-swagger
```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
~/go/bin/swagger generate spec -m -o ./swagger.json
```
