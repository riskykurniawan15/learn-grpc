@echo off
echo Testing HTTP API endpoints...
echo ==================================

set BASE_URL=http://localhost:8080

REM Health check
echo.
echo 1. Health Check:
curl -s "%BASE_URL%/health"

REM Create user
echo.
echo.
echo 2. Create User:
curl -s -X POST "%BASE_URL%/users" -H "Content-Type: application/json" -d "{\"name\": \"Bob Wilson\", \"email\": \"bob@example.com\", \"age\": 35}"

echo.
echo.
echo Testing completed!
echo Note: For better JSON formatting, install jq and use the shell script version
pause

