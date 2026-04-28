#!/bin/bash

echo "DevBoard deploy başlıyor..."

echo "Eski containerlar kaldırılıyor..."
docker compose down -v

echo "Build & up işlemi başlıyor..."
docker compose up -d --build

echo "Container durumu:"
docker ps

echo "Deploy tamamlandı!"
