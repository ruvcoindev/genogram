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
        Description: "Реакция на травматическое событие",
        Symptoms:    []string{"Флешбеки", "Избегание", "Гипервозбуждение"},
        Origins:     []string{"Насилие", "Катастрофы", "Потеря близких"},
        Defenses:    []string{"Отрицание", "Смещение"},
        HealingPath: []string{"EMDR", "Соматическая терапия"},
        RelatedChakras: []int{0, 3},
    }
    db.Entries["cptsd"] = TraumaInfo{
        Type:        "cptsd",
        NameRU:      "Комплексное ПТСР",
        Description: "Длительная травма развития",
        Symptoms:    []string{"Нарушение регуляции эмоций", "Негативная самоконцепция"},
        Origins:     []string{"Нарциссическое воспитание", "Эмоциональное пренебрежение"},
        Defenses:    []string{"Проекция", "Рационализация"},
        HealingPath: []string{"Терапия привязанности", "Работа с внутренним ребенком"},
        RelatedChakras: []int{0, 1, 3},
    }
}
