# 🛠️ Development Environment Setup

## Prerequisites

### 1. Go Installation
- **Download**: [golang.org/dl](https://golang.org/dl/)
- **Version**: 1.21 atau lebih baru
- **Verify**: `go version`

### 2. Protocol Buffers Compiler
- **Windows**: Download dari [GitHub releases](https://github.com/protocolbuffers/protobuf/releases)
- **Linux**: `sudo apt install protobuf-compiler`
- **macOS**: `brew install protobuf`
- **Verify**: `protoc --version`

### 3. Go Protobuf Plugins
```bash
# Install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Verify installation
which protoc-gen-go
which protoc-gen-go-grpc
```

## Project Setup

### 1. Clone/Download Project
```bash
git clone <repository-url>
cd gorpc
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Generate Protobuf Code
```bash
# Windows
generate.bat

# Linux/macOS
chmod +x generate.sh
./generate.sh
```

### 4. Verify Setup
```bash
# Check if protobuf files were generated
ls proto/*.pb.go
ls proto/*_grpc.pb.go

# Check Go modules
go mod verify
```

## IDE Setup

### VS Code
1. **Install Extensions**:
   - Go (official)
   - Protocol Buffers
   - gRPC

2. **Go Tools**:
   ```bash
   go install golang.org/x/tools/gopls@latest
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

### GoLand
- Built-in Go support
- Protobuf support included
- gRPC debugging tools

## Database Setup

### SQLite
- **Auto-created**: Database akan dibuat otomatis
- **Location**: `users.db` di root project
- **Schema**: Auto-migrated dengan GORM

### Alternative Databases
Untuk menggunakan database lain:

1. **PostgreSQL**:
   ```go
   import "gorm.io/driver/postgres"
   
   dsn := "host=localhost user=postgres password=password dbname=gorpc port=5432"
   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   ```

2. **MySQL**:
   ```go
   import "gorm.io/driver/mysql"
   
   dsn := "user:password@tcp(127.0.0.1:3306)/gorpc?charset=utf8mb4&parseTime=True&loc=Local"
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   ```

## Testing Tools

### 1. HTTP Testing
- **curl**: Command line testing
- **Postman**: GUI testing (import collection)
- **Insomnia**: Alternative to Postman

### 2. gRPC Testing
- **grpcurl**: Command line gRPC testing
- **BloomRPC**: GUI gRPC client
- **gRPC UI**: Web-based gRPC testing

### 3. Database Testing
- **SQLite Browser**: GUI untuk SQLite
- **DBeaver**: Universal database tool

## Common Issues & Solutions

### 1. Protobuf Generation Fails
```bash
# Check protoc version
protoc --version

# Check plugin installation
which protoc-gen-go
which protoc-gen-go-grpc

# Reinstall plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 2. Import Errors
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod download
go mod tidy
```

### 3. Port Conflicts
```bash
# Check port usage
netstat -an | grep :50051
netstat -an | grep :8080

# Kill process using port
taskkill /PID <process_id> /F  # Windows
kill -9 <process_id>           # Linux/macOS
```

### 4. Permission Issues
```bash
# Windows: Run as Administrator
# Linux/macOS: Check folder permissions
ls -la
chmod 755 .
```

## Development Workflow

### 1. Code Changes
1. Edit `.proto` files
2. Run `make generate` or `generate.bat`
3. Update Go code
4. Test changes

### 2. Testing
1. Start gRPC server: `make server`
2. Test with client: `make client`
3. Test HTTP API: `go run http_server/main.go`
4. Use Postman collection

### 3. Debugging
1. Use `fmt.Printf` for logging
2. Use Go debugger (delve)
3. Check gRPC server logs
4. Monitor database changes

## Performance Tuning

### 1. gRPC
- Enable compression
- Use connection pooling
- Implement streaming for large data

### 2. Database
- Add indexes
- Use transactions
- Implement connection pooling

### 3. HTTP
- Enable gzip compression
- Add caching headers
- Use connection pooling

## Security Considerations

### 1. Production
- Use TLS for gRPC
- Implement authentication
- Add rate limiting
- Validate all inputs

### 2. Development
- Use environment variables
- Don't commit secrets
- Use local databases only

## Next Steps

1. **Explore Code**: Understand the architecture
2. **Add Features**: Implement new endpoints
3. **Add Tests**: Write unit and integration tests
4. **Add Logging**: Implement structured logging
5. **Add Monitoring**: Health checks and metrics
6. **Deploy**: Containerize and deploy

