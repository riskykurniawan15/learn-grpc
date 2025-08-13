# 🚀 Quick Start Guide

## Langkah Cepat untuk Testing

### 1. Install Dependencies
```bash
# Install protobuf plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install project dependencies
go mod tidy
```

### 2. Generate Protobuf Code
```bash
# Windows
generate.bat

# Linux/Mac
./generate.sh
```

### 3. Run Server
```bash
go run server/server.go
```

### 4. Test dengan Client
```bash
# Di terminal baru
go run client/client.go
```

## 🧪 Testing dengan HTTP API

### 1. Run HTTP Server
```bash
go run http_server/main.go
```

### 2. Test API Endpoints
```bash
# Health check
curl http://localhost:8080/health

# Create user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","age":25}'

# Get all users
curl http://localhost:8080/users
```

## 📱 Postman Collection

Import file `postman/User_Service_API.postman_collection.json` ke Postman untuk testing yang lebih mudah.

## 🔧 Troubleshooting

### Port Already in Use
- Server gRPC: Port 50051
- HTTP Server: Port 8080
- Ubah port di file yang sesuai jika diperlukan

### Database Error
- Pastikan folder memiliki permission write
- SQLite akan dibuat otomatis

### Protobuf Error
- Pastikan protoc terinstall
- Jalankan `go mod tidy` setelah generate

## 📚 Next Steps

1. **Explore Code**: Lihat struktur project dan implementasi
2. **Modify Schema**: Ubah model User di `proto/user.proto`
3. **Add Validation**: Tambahkan validasi di service layer
4. **Add Tests**: Buat unit tests untuk setiap layer
5. **Deploy**: Deploy ke production environment

## 🎯 Learning Points

- **gRPC**: Protocol Buffers, service definition, streaming
- **GORM**: Database ORM, migrations, relationships
- **Clean Architecture**: Repository pattern, service layer
- **Error Handling**: gRPC status codes, proper error responses
- **Testing**: Client testing, HTTP wrapper, Postman collection

