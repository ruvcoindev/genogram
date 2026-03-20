# 🧬 Genogram System v1.1

# 🧬 Genogram System v1.1

**Полная система анализа семейных систем по Мюррею Боуэну с интеграцией HyperCube, CBT, дневника "Кто я", психологии симптомов и AI.**

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-1.1.0-purple.svg)]()

---

## 🌟 Уникальные возможности

### 🔬 HyperCube (4D анализ)
- **4D координаты** из даты рождения (X, Y, Z, W)
- **Комнаты-ловушки** (111, 222, 333...) — повторяющиеся паттерны
- **Комнаты-мосты** (999) — точки исцеления
- **Векторы движения** — жизненные траектории

### 🧠 CBT (Когнитивно-поведенческая терапия)
- **Автоматическое обнаружение** 11 когнитивных искажений
- **Генерация рациональных ответов**
- **Интеграция с травматическими ролями**

### 📔 Дневник "Кто я"
- **5 разделов** × 5 вопросов (Мотивация, Границы, Ресурс, Паттерны, Выбор)
- **Стоп-сигналы** для саморефлексии
- **Автосохранение** ответов

### 🤒 Симптомы
- **Физические** симптомы (усталость, головные боли)
- **Эмоциональные** симптомы (тревога, депрессия)
- **Связь с чакрами и гормонами**

### 🧬 Психология
- **Травмы**: ПТСР, КПТСР, травма нарциссического родителя
- **Стадии Эриксона**: 8 стадий психосоциального развития
- **Аффирмации Луизы Хей**: 100+ утверждений

### 📊 Векторное хранилище
- **Pure-Go реализация** (без CGO)
- **Поиск по косинусному сходству**
- **768-мерные эмбеддинги** для текстов
- **9-мерные векторы** для отношений HyperCube

### 🗄️ Базы данных
- **PostgreSQL + PgVector** (векторный поиск)
- **Neo4j** (графовые связи)
- **Redis** (кэширование)

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

📚 Модули

HyperCube Core

```go

// Создание 4D координат из даты рождения
coords := core.ParseDateToCoords(birthDate)

// Вычисление векторов движения
vectors := core.CalculateVectors(coords)

// Проверка комнаты-ловушки
isTrap := core.IsTrapRoom(coords)

// Проверка комнаты-моста
isBridge := core.IsBridgeRoom(coords)

CBT Analyzer

```go

// Обнаружение когнитивных искажений
distortions := cbt.DetectDistortions(thought)

// Генерация рационального ответа
response := cbt.GenerateRationalResponse(thought, distortions)

Diary Store

```go

// Сохранение записи дневника
store.SaveEntry(userID, section, answer, tags)

// Получение всех записей
entries, _ := store.GetEntries(userID)


Vector Store

```go

// Создание хранилища
store := vector.NewStore(768)

// Добавление вектора
store.Add(id, embedding, metadata)

// Поиск похожих
results, _ := store.Search(query, 10)


🐳 Docker Compose

```bash

# Запуск всех сервисов
docker-compose up -d

# Остановка
docker-compose down

# Просмотр логов
docker-compose logs -f backend

# Очистка данных
docker-compose down -v


📊 API Endpoints

   Endpoint	                   Метод	Описание
/api/v1/health	                    GET	        Проверка здоровья
/api/v1/cbt/analyze	            POST	Анализ мысли
/api/v1/diary/sections	            GET	        Разделы дневника
/api/v1/diary/entries	            POST	Сохранить запись
/api/v1/symptoms	            GET	        Список симптомов
/api/v1/psychology/traumas	    GET	        База травм
/api/v1/affirmations	            GET	        Поиск аффирмаций
/api/v1/stats	                    GET	        Статистика

🧪 Тестирование

```bash

# Запуск тестов бэкенда
cd backend && go test ./...

# Запуск тестов фронтенда
cd frontend && npm test  # если есть тесты


🙏 Благодарности
Мюррей Боуэн — теория семейных систем

Арон Бек — когнитивная терапия

Луиза Хей — аффирмации

Эрик Эриксон — стадии развития

Анодея Джудит — чакра-психология

🔗 Ссылки
GitHub Repository

Документация

Примеры использования

Genogram System v1.1 — инструмент для самопознания и понимания семейной динамики.


📄 Лицензия
MIT License
