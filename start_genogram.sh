#!/bin/bash
echo "🚀 Genogram System v1.1"
echo "======================="

# Останавливаем старые процессы
pkill -f genogram-server 2>/dev/null || true
pkill -f "python3 -m http.server" 2>/dev/null || true
sudo fuser -k 8082/tcp 2>/dev/null || true

# Запускаем backend
cd ~/genogram/backend
echo "🔨 Компилируем backend..."
go build -o genogram-server ./cmd/ 2>/dev/null

echo "🚀 Запускаем backend на порту 8082..."
export SERVER_PORT=8082
./genogram-server > /tmp/genogram.log 2>&1 &
BACKEND_PID=$!
sleep 2

# Проверяем
if curl -s http://localhost:8082/api/v1/health > /dev/null 2>&1; then
    echo "✅ Backend работает (PID: $BACKEND_PID)"
else
    echo "❌ Ошибка запуска backend"
    tail -20 /tmp/genogram.log
    exit 1
fi

# Запускаем фронтенд
cd ~/genogram/frontend
echo "🚀 Запускаем фронтенд на порту 8000..."
python3 -m http.server 8000 > /tmp/frontend.log 2>&1 &
FRONTEND_PID=$!
echo "✅ Frontend работает (PID: $FRONTEND_PID)"

echo ""
echo "✅ Система запущена!"
echo "   🌐 http://localhost:8000/index.html"
echo "   🧪 http://localhost:8000/test_cors.html"
echo ""
echo "📋 Логи: tail -f /tmp/genogram.log"
echo "🛑 Остановить: pkill -f genogram-server && pkill -f 'python3 -m http.server'"

wait
