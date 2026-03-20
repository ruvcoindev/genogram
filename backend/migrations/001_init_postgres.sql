-- =====================================================
-- ГЕНОГРАММА СИСТЕМА v1.1
-- Миграция 001: Инициализация PostgreSQL + PgVector
-- =====================================================

-- Включаем расширения
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- =====================================================
-- ТАБЛИЦА: ПОЛЬЗОВАТЕЛИ
-- =====================================================
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    role VARCHAR(50) DEFAULT 'therapist',
    organization VARCHAR(255),
    subscription_tier VARCHAR(50) DEFAULT 'free',
    subscription_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);

-- =====================================================
-- ТАБЛИЦА: ГЕНОГРАММЫ
-- =====================================================
CREATE TABLE IF NOT EXISTS genograms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    therapist_id UUID NOT NULL REFERENCES users(id),
    client_id UUID REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_genograms_therapist ON genograms(therapist_id);
CREATE INDEX IF NOT EXISTS idx_genograms_client ON genograms(client_id);

-- =====================================================
-- ТАБЛИЦА: ЧЛЕНЫ СЕМЬИ
-- =====================================================
CREATE TABLE IF NOT EXISTS family_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genogram_id UUID NOT NULL REFERENCES genograms(id) ON DELETE CASCADE,
    external_id VARCHAR(100),
    name VARCHAR(255) NOT NULL,
    gender VARCHAR(20),
    birth_date DATE,
    death_date DATE,
    
    -- HyperCube координаты
    hc_x INTEGER,
    hc_y INTEGER,
    hc_z INTEGER,
    hc_w INTEGER,
    
    -- Травматическая роль
    trauma_role VARCHAR(50),
    trauma_role_confidence FLOAT,
    
    -- Метаданные
    emotional_state VARCHAR(50),
    differentiation_level FLOAT,
    sibling_position VARCHAR(20),
    occupation TEXT,
    education TEXT,
    health TEXT,
    
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_family_members_genogram ON family_members(genogram_id);
CREATE INDEX IF NOT EXISTS idx_family_members_trauma_role ON family_members(trauma_role);

-- =====================================================
-- ТАБЛИЦА: ОТНОШЕНИЯ (с векторами)
-- =====================================================
CREATE TABLE IF NOT EXISTS relationships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genogram_id UUID NOT NULL REFERENCES genograms(id) ON DELETE CASCADE,
    from_member_id UUID NOT NULL REFERENCES family_members(id) ON DELETE CASCADE,
    to_member_id UUID NOT NULL REFERENCES family_members(id) ON DELETE CASCADE,
    relationship_type VARCHAR(50) NOT NULL,
    
    emotional_distance FLOAT DEFAULT 5.0,
    conflict_level FLOAT DEFAULT 5.0,
    intimacy_level FLOAT DEFAULT 5.0,
    stability FLOAT DEFAULT 5.0,
    
    hc_vectors VECTOR(9),
    
    is_triangulated BOOLEAN DEFAULT FALSE,
    triangulated_by UUID[] DEFAULT '{}',
    history JSONB DEFAULT '[]',
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT check_not_self CHECK (from_member_id != to_member_id)
);

CREATE INDEX IF NOT EXISTS idx_relationships_genogram ON relationships(genogram_id);
CREATE INDEX IF NOT EXISTS idx_relationships_vectors ON relationships USING ivfflat (hc_vectors vector_cosine_ops);

-- =====================================================
-- ТАБЛИЦА: CBT ЗАПИСИ
-- =====================================================
CREATE TABLE IF NOT EXISTS cbt_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    genogram_id UUID REFERENCES genograms(id),
    family_member_id UUID REFERENCES family_members(id),
    
    situation TEXT NOT NULL,
    automatic_thought TEXT NOT NULL,
    emotions TEXT[] DEFAULT '{}',
    intensity INTEGER CHECK (intensity >= 0 AND intensity <= 100),
    
    distortions TEXT[] DEFAULT '{}',
    rational_response TEXT,
    new_intensity INTEGER CHECK (new_intensity >= 0 AND new_intensity <= 100),
    
    thought_vector VECTOR(768),
    
    completed BOOLEAN DEFAULT FALSE,
    notes TEXT,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_cbt_records_user ON cbt_records(user_id);
CREATE INDEX IF NOT EXISTS idx_cbt_records_vectors ON cbt_records USING ivfflat (thought_vector vector_cosine_ops);
CREATE INDEX IF NOT EXISTS idx_cbt_thought_gin ON cbt_records USING gin(to_tsvector('russian', automatic_thought));

-- =====================================================
-- ТАБЛИЦА: ЗАПИСИ ДНЕВНИКА
-- =====================================================
CREATE TABLE IF NOT EXISTS diary_entries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    section VARCHAR(50) NOT NULL,
    question_id VARCHAR(100) NOT NULL,
    question_text TEXT NOT NULL,
    answer TEXT NOT NULL,
    tags TEXT[] DEFAULT '{}',
    
    answer_vector VECTOR(768),
    
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_diary_entries_user ON diary_entries(user_id);
CREATE INDEX IF NOT EXISTS idx_diary_entries_section ON diary_entries(section);
CREATE INDEX IF NOT EXISTS idx_diary_entries_tags ON diary_entries USING gin(tags);
CREATE INDEX IF NOT EXISTS idx_diary_entries_vectors ON diary_entries USING ivfflat (answer_vector vector_cosine_ops);
CREATE INDEX IF NOT EXISTS idx_diary_answer_gin ON diary_entries USING gin(to_tsvector('russian', answer));

-- =====================================================
-- ТАБЛИЦА: СИМПТОМЫ
-- =====================================================
CREATE TABLE IF NOT EXISTS symptom_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symptom_id VARCHAR(50) NOT NULL,
    symptom_name VARCHAR(255) NOT NULL,
    severity INTEGER CHECK (severity >= 0 AND severity <= 10),
    frequency VARCHAR(50),
    onset_date DATE,
    duration_days INTEGER,
    triggers TEXT[] DEFAULT '{}',
    notes TEXT,
    
    symptom_vector VECTOR(768),
    
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_symptom_records_user ON symptom_records(user_id);
CREATE INDEX IF NOT EXISTS idx_symptom_records_vectors ON symptom_records USING ivfflat (symptom_vector vector_cosine_ops);

-- =====================================================
-- ТАБЛИЦА: ТРАВМЫ
-- =====================================================
CREATE TABLE IF NOT EXISTS trauma_assessments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    trauma_type VARCHAR(50) NOT NULL,
    trauma_name VARCHAR(255) NOT NULL,
    severity FLOAT CHECK (severity >= 0 AND severity <= 1),
    confidence FLOAT CHECK (confidence >= 0 AND confidence <= 1),
    
    symptoms TEXT[] DEFAULT '{}',
    defenses TEXT[] DEFAULT '{}',
    
    trauma_vector VECTOR(768),
    
    assessment_data JSONB DEFAULT '{}',
    assessed_by UUID REFERENCES users(id),
    assessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_trauma_assessments_user ON trauma_assessments(user_id);
CREATE INDEX IF NOT EXISTS idx_trauma_assessments_vectors ON trauma_assessments USING ivfflat (trauma_vector vector_cosine_ops);

-- =====================================================
-- ТАБЛИЦА: АФФИРМАЦИИ (использование)
-- =====================================================
CREATE TABLE IF NOT EXISTS affirmation_usage (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    affirmation_id VARCHAR(50) NOT NULL,
    affirmation_text TEXT NOT NULL,
    context VARCHAR(100),
    effectiveness INTEGER CHECK (effectiveness >= 1 AND effectiveness <= 5),
    used_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_affirmation_usage_user ON affirmation_usage(user_id);

-- =====================================================
-- ТАБЛИЦА: СТАДИИ ЭРИКСОНА
-- =====================================================
CREATE TABLE IF NOT EXISTS erikson_stages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_number INTEGER CHECK (stage_number >= 1 AND stage_number <= 8),
    stage_name VARCHAR(100) NOT NULL,
    crisis_resolved BOOLEAN DEFAULT FALSE,
    resolution_level FLOAT CHECK (resolution_level >= 0 AND resolution_level <= 1),
    notes TEXT,
    assessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_erikson_stages_user ON erikson_stages(user_id);

-- =====================================================
-- ТАБЛИЦА: ЧАКРА-ПСИХОЛОГИЯ
-- =====================================================
CREATE TABLE IF NOT EXISTS chakra_patterns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chakra_index INTEGER CHECK (chakra_index >= 0 AND chakra_index <= 6),
    chakra_name VARCHAR(50) NOT NULL,
    childhood_pattern TEXT,
    core_wound TEXT,
    defense_mechanism TEXT,
    healing_progress FLOAT CHECK (healing_progress >= 0 AND healing_progress <= 1),
    notes TEXT,
    assessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_chakra_patterns_user ON chakra_patterns(user_id);

-- =====================================================
-- ТАБЛИЦА: ОТЧЕТЫ
-- =====================================================
CREATE TABLE IF NOT EXISTS analysis_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genogram_id UUID NOT NULL REFERENCES genograms(id) ON DELETE CASCADE,
    report_type VARCHAR(50) NOT NULL,
    content JSONB NOT NULL,
    summary TEXT,
    recommendations TEXT[] DEFAULT '{}',
    risks JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_reports_genogram ON analysis_reports(genogram_id);

-- =====================================================
-- ТАБЛИЦА: АУДИТ ЛОГИ
-- =====================================================
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    action VARCHAR(100) NOT NULL,
    entity_type VARCHAR(50),
    entity_id UUID,
    old_data JSONB,
    new_data JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_audit_user ON audit_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_audit_entity ON audit_logs(entity_type, entity_id);

-- =====================================================
-- ФУНКЦИЯ ОБНОВЛЕНИЯ updated_at
-- =====================================================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_genograms_updated_at BEFORE UPDATE ON genograms FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_family_members_updated_at BEFORE UPDATE ON family_members FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_relationships_updated_at BEFORE UPDATE ON relationships FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_cbt_records_updated_at BEFORE UPDATE ON cbt_records FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_diary_entries_updated_at BEFORE UPDATE ON diary_entries FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_symptom_records_updated_at BEFORE UPDATE ON symptom_records FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_trauma_assessments_updated_at BEFORE UPDATE ON trauma_assessments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_erikson_stages_updated_at BEFORE UPDATE ON erikson_stages FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_chakra_patterns_updated_at BEFORE UPDATE ON chakra_patterns FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- =====================================================
-- ФУНКЦИИ ГИБРИДНОГО ПОИСКА
-- =====================================================
CREATE OR REPLACE FUNCTION hybrid_search_thoughts(
    dense_vector VECTOR(768),
    query_text TEXT,
    dense_weight FLOAT DEFAULT 0.7,
    similarity_threshold FLOAT DEFAULT 0.5,
    limit_count INTEGER DEFAULT 10
)
RETURNS TABLE (
    thought_id UUID,
    user_id UUID,
    thought_text TEXT,
    distortions TEXT[],
    hybrid_score FLOAT
) AS $$
BEGIN
    RETURN QUERY
    WITH dense_scores AS (
        SELECT 
            id,
            user_id,
            automatic_thought,
            distortions,
            1 - (thought_vector <=> dense_vector) as dense_score
        FROM cbt_records
        WHERE thought_vector IS NOT NULL
    ),
    sparse_scores AS (
        SELECT 
            id,
            ts_rank_cd(to_tsvector('russian', automatic_thought), plainto_tsquery('russian', query_text)) as sparse_score
        FROM cbt_records
        WHERE to_tsvector('russian', automatic_thought) @@ plainto_tsquery('russian', query_text)
    )
    SELECT 
        d.id,
        d.user_id,
        d.automatic_thought,
        d.distortions,
        (dense_weight * d.dense_score + (1 - dense_weight) * COALESCE(s.sparse_score, 0)) as hybrid_score
    FROM dense_scores d
    LEFT JOIN sparse_scores s ON d.id = s.id
    WHERE (dense_weight * d.dense_score + (1 - dense_weight) * COALESCE(s.sparse_score, 0)) > similarity_threshold
    ORDER BY hybrid_score DESC
    LIMIT limit_count;
END;
$$ LANGUAGE plpgsql;
