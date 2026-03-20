#!/bin/bash
echo "🚀 Genogram System v1.1"
echo "======================="

# Останавливаем старые процессы
pkill -f "genogram-server" 2>/dev/null || true
pkill -f "python3 -m http.server" 2>/dev/null || true

# Запускаем backend
cd ~/genogram/backend
echo "🔨 Компилируем backend..."
go build -o genogram-server ./cmd/

echo "🚀 Запускаем backend на порту 8082..."
export SERVER_PORT=8082
./genogram-server &
BACKEND_PID=$!
sleep 2

# Проверяем
if curl -s http://localhost:8082/api/v1/health > /dev/null 2>&1; then
    echo "✅ Backend работает на http://localhost:8082"
else
    echo "❌ Ошибка запуска backend"
    exit 1
fi

# Запускаем фронтенд
cd ~/genogram/frontend
echo "🚀 Запускаем фронтенд на порту 8000..."
python3 -m http.server 8000 &
FRONTEND_PID=$!

echo ""
echo "✅ Система запущена!"
echo "   🌐 Тестовая страница: http://localhost:8000/test.html"
echo "   📡 API: http://localhost:8082/api/v1/health"
echo ""
echo "🛑 Для остановки: kill $BACKEND_PID $FRONTEND_PID"

wait
