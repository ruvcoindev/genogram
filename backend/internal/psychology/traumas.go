package psychology

type TraumaInfo struct {
    Type           string   `json:"type"`
    NameRU         string   `json:"name_ru"`
    Description    string   `json:"description"`
    Symptoms       []string `json:"symptoms"`
    Origins        []string `json:"origins"`
    Defenses       []string `json:"defenses"`
    HealingPath    []string `json:"healing_path"`
    RelatedChakras []int    `json:"related_chakras"`
}

type TraumaDB struct {
    Entries map[string]TraumaInfo
}

func NewTraumaDB() *TraumaDB {
    db := &TraumaDB{Entries: make(map[string]TraumaInfo)}
    db.load()
    return db
}

func (db *TraumaDB) load() {
    db.Entries["ptsd"] = TraumaInfo{
        Type:        "ptsd",
        NameRU:      "Посттравматическое стрессовое расстройство",
        Description: "Реакция на травматическое событие, угрожающее жизни или целостности",
        Symptoms:    []string{"Флешбеки, навязчивые воспоминания", "Избегание триггеров", "Гипервозбуждение, тревожность", "Эмоциональное онемение"},
        Origins:     []string{"Насилие (физическое, сексуальное, эмоциональное)", "Военные действия, катастрофы", "Внезапная потеря близкого"},
        Defenses:    []string{"Отрицание", "Смещение", "Интеллектуализация"},
        HealingPath: []string{"EMDR-терапия", "Соматическая проработка", "Безопасное повторное переживание"},
        RelatedChakras: []int{0, 3},
    }
    
    db.Entries["cptsd"] = TraumaInfo{
        Type:        "cptsd",
        NameRU:      "Комплексное ПТСР",
        Description: "Длительная травма развития, часто в детстве, в условиях отсутствия безопасной привязанности",
        Symptoms:    []string{"Нарушение регуляции эмоций", "Негативная самоконцепция ('я дефектный')", "Трудности в отношениях", "Диссоциация, деперсонализация"},
        Origins:     []string{"Нарциссическое воспитание", "Эмоциональное пренебрежение", "Хроническое насилие в семье"},
        Defenses:    []string{"Проекция", "Рационализация", "Триангуляция"},
        HealingPath: []string{"Терапия привязанности", "Работа с внутренним ребенком", "Развитие самосострадания"},
        RelatedChakras: []int{0, 1, 3},
    }
    
    db.Entries["narcissistic_parent"] = TraumaInfo{
        Type:        "narcissistic_parent",
        NameRU:      "Травма от нарциссического родителя",
        Description: "Воспитание родителем с нарциссическими чертами, где ребенок используется для регуляции самооценки родителя",
        Symptoms:    []string{"Чувство, что любовь нужно заслужить", "Размытые границы, трудности с 'нет'", "Гиперответственность за чувства других", "Страх быть 'эгоистичным'"},
        Origins:     []string{"Условная любовь", "Триангуляция", "Обесценивание потребностей ребенка", "Проекция стыда"},
        Defenses:    []string{"Идеализация", "Обесценивание", "Триангуляция"},
        HealingPath: []string{"Сепарация", "Работа со стыдом и виной", "Разрешение на здоровый эгоизм", "Построение здоровых границ"},
        RelatedChakras: []int{2, 3, 4},
    }
    
    db.Entries["attachment"] = TraumaInfo{
        Type:        "attachment",
        NameRU:      "Травма привязанности",
        Description: "Нарушение формирования безопасной привязанности в раннем детстве",
        Symptoms:    []string{"Страх брошенности или поглощения", "Трудности с доверием", "Циклы сближения-отдаления", "Проекция родительских фигур на партнеров"},
        Origins:     []string{"Непредсказуемость опекуна", "Эмоциональная недоступность родителей", "Ранняя сепарация, потеря"},
        Defenses:    []string{"Проекция", "Реактивное образование", "Регрессия"},
        HealingPath: []string{"Корректирующий эмоциональный опыт", "Осознание паттернов привязанности", "Практика безопасной уязвимости"},
        RelatedChakras: []int{0, 3, 4},
    }
}
