APP_RUNNING=http://localhost:3000

### URI/auth/register
method: POST
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```
if success
response 200
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJqb2huLmRvZUBleGFtcGxlLmNvbSIsImV4cCI6MTcwMTExMTExMX0.abc123xyz456",
  "message": "Registered with Google successfully"
}
```

if failed
response 500
```json
{
  "status": "error",
  "message": "Failed to check if user exists: database connection timeout"
}
```

