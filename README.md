# APP

## Specs

### Endpoints

#### GET session

call:
```sh
curl -X POST localhost:8080/login -d '{ "login":"casper", "pass":"Tatata" }'
curl -X POST localhost:8080/login -d '{ "login":"boss", "pass":"bibibi" }'
```

returns:
```json
200.OK
{
    "jwt":"jwt_value_xxxxxxxxxxxxx"
}
```

```json
401.Unauthorized
error:
{
    "error":"user not authorized."
}
```

### Healthcheck

call:
```sh
curl -X GET localhost:8081/health-check -H 'Authorization: Baerer jwt_value'
```

returns:
```json
200.OK
{
    "status":"ok" // ok, altered, down.
}
```

```json
401.Unauthorized
error:
{
    "error":"user not authorized."
}
```

### Authenticate with JWT

```json
{
    "access_level":"admin", // admin, operator, supervisor.
    "user_name": "Rob Pike",
    "uuid_user": "227c6c63-70a8-4528-a327-09509a2c9613",
    "exp": "1631005742"
}
```