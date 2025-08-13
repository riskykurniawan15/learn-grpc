# Go RPC CRUD Example

Project sederhana untuk belajar gRPC dengan Go yang mengimplementasikan CRUD operations pada tabel user.

## Fitur

- **gRPC Server**: Server yang menyediakan service untuk operasi CRUD user
- **SQLite Database**: Database sederhana menggunakan SQLite dengan GORM
- **CRUD Operations**: Create, Read, Update, Delete user
- **Protobuf**: Definisi service dan message menggunakan Protocol Buffers
- **Client Example**: Contoh client untuk testing service

## Struktur Project

```
gorpc/
├── proto/           # Definisi protobuf
├── models/          # Model database
├── database/        # Database connection
├── repository/      # Data access layer
├── service/         # gRPC service implementation
├── server/          # gRPC server
├── client/          # gRPC client untuk testing
├── go.mod          # Go module dependencies
├── generate.sh     # Script untuk generate protobuf
└── README.md       # Dokumentasi project
```

## Prerequisites

Sebelum menjalankan project, pastikan sudah menginstall:

1. **Go** (versi 1.21 atau lebih baru)
2. **Protocol Buffers Compiler** (protoc)
3. **Go protobuf plugins**:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Setup dan Running

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Generate Protobuf Code

```bash
# Di Windows
generate.sh

# Atau manual
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
```

### 3. Run Server

```bash
go run server/server.go
```

Server akan berjalan di port 50051.

### 4. Run Client (Testing)

```bash
go run client/client.go
```

## API Endpoints

Service menyediakan 5 operasi utama:

1. **CreateUser** - Membuat user baru
2. **GetUser** - Mengambil user berdasarkan ID
3. **GetAllUsers** - Mengambil semua user
4. **UpdateUser** - Update data user
5. **DeleteUser** - Hapus user berdasarkan ID

## Database Schema

Tabel `users` memiliki struktur:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    age INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
);
```

## Contoh Penggunaan

### Create User
```go
resp, err := client.CreateUser(ctx, &proto.CreateUserRequest{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   30,
})
```

### Get User by ID
```go
resp, err := client.GetUser(ctx, &proto.GetUserRequest{
    Id: 1,
})
```

### Update User
```go
resp, err := client.UpdateUser(ctx, &proto.UpdateUserRequest{
    Id:   1,
    Name: "John Doe Updated",
    Age:  31,
})
```

### Delete User
```go
resp, err := client.DeleteUser(ctx, &proto.DeleteUserRequest{
    Id: 1,
})
```

## Error Handling

Service menggunakan gRPC status codes untuk error handling:

- `codes.InvalidArgument` - Input tidak valid
- `codes.NotFound` - User tidak ditemukan
- `codes.Internal` - Error database

## Testing

Client example akan menjalankan test sequence:

1. Create 2 users
2. Get all users
3. Get user by ID
4. Update user
5. Delete user
6. Verify deletion

## Troubleshooting

### Port Already in Use
Jika port 50051 sudah digunakan, ubah port di `server/server.go` dan `client/client.go`.

### Database Error
Pastikan SQLite dapat diakses dan folder memiliki permission write.

### Protobuf Generation Error
Pastikan protoc dan plugins sudah terinstall dengan benar.

## Learning Resources

- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [GORM Documentation](https://gorm.io/docs/)
- [Go Modules](https://go.dev/doc/modules/)

## License

Project ini dibuat untuk tujuan pembelajaran.

