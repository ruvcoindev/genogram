# 🧬 Genogram System v1.1

**Полная система анализа семейных систем по Мюррею Боуэну с интеграцией HyperCube, CBT, дневника "Кто я" и AI.**

## 🚀 Быстрый старт

```bash
# Клонировать репозиторий
git clone https://github.com/ruvcoindev/genogram.git
cd genogram

# Запустить все сервисы
make start

# Применить миграции
make migrate-up

# Открыть в браузере
open http://localhost

📦 Структура проекта

genogram/
├── backend/           # Go бэкенд
│   ├── cmd/          # Точка входа
│   ├── internal/     # Внутренние модули
│   │   ├── hypercube/ # 4D анализ
│   │   ├── cbt/      # Когнитивная терапия
│   │   ├── diary/    # Дневник "Кто я"
│   │   ├── symptoms/ # Симптоматика
│   │   └── psychology/ # Психология
│   └── migrations/   # SQL миграции
├── frontend/         # Веб-интерфейс
│   ├── css/         # Стили
│   ├── js/          # JavaScript
│   └── assets/      # Изображения
└── docker-compose.yml

🔧 Технологии
Backend: Go 1.21, Gorilla Mux, WebSocket

Database: PostgreSQL + PgVector, Neo4j, Redis

AI: Ollama (Qwen3.5, Llama3.2, BGE-M3)

Frontend: HTML5, CSS3, JavaScript, Cytoscape.js, Three.js

📄 Лицензия
MIT License
