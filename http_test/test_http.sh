#!/bin/bash

# Test script for HTTP API
BASE_URL="http://localhost:8080"

echo "🧪 Testing HTTP API endpoints..."
echo "=================================="

# Health check
echo -e "\n1️⃣ Health Check:"
curl -s "$BASE_URL/health" | jq '.'

# Create user
echo -e "\n2️⃣ Create User:"
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bob Wilson",
    "email": "bob@example.com",
    "age": 35
  }')
echo $CREATE_RESPONSE | jq '.'

# Extract user ID
USER_ID=$(echo $CREATE_RESPONSE | jq -r '.data.id')
echo "User ID: $USER_ID"

# Get all users
echo -e "\n3️⃣ Get All Users:"
curl -s "$BASE_URL/users" | jq '.'

# Get user by ID
echo -e "\n4️⃣ Get User by ID:"
curl -s "$BASE_URL/users/$USER_ID" | jq '.'

# Update user
echo -e "\n5️⃣ Update User:"
curl -s -X PUT "$BASE_URL/users/$USER_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bob Wilson Updated",
    "age": 36
  }' | jq '.'

# Get updated user
echo -e "\n6️⃣ Get Updated User:"
curl -s "$BASE_URL/users/$USER_ID" | jq '.'

# Delete user
echo -e "\n7️⃣ Delete User:"
curl -s -X DELETE "$BASE_URL/users/$USER_ID" | jq '.'

# Verify deletion
echo -e "\n8️⃣ Verify Deletion:"
curl -s "$BASE_URL/users" | jq '.'

echo -e "\n✅ HTTP API testing completed!"

