#!/bin/bash
set -e

echo "🚀 Запуск миграций Genogram System v1.1"

# Загрузка .env
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

# 1. PostgreSQL миграции
echo "📦 Миграция PostgreSQL..."
for migration in backend/migrations/*.sql; do
    echo "  → Применение: $(basename $migration)"
    PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -U $POSTGRES_USER -d $POSTGRES_DB -f $migration 2>/dev/null || true
done

# 2. Neo4j миграции
echo "🔷 Миграция Neo4j..."
for migration in backend/migrations/*.cypher; do
    if [ -f "$migration" ]; then
        echo "  → Применение: $(basename $migration)"
        cat $migration | cypher-shell -u $NEO4J_USER -p $NEO4J_PASSWORD -a $NEO4J_URI 2>/dev/null || true
    fi
done

echo "✅ Миграции завершены!"
