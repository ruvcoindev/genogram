package psychology

type Affirmation struct {
    ID          string   `json:"id"`
    Text        string   `json:"text"`
    Author      string   `json:"author"`
    Keywords    []string `json:"keywords"`
    ChakraIndex int      `json:"chakra_index"`
    DayOfYear   int      `json:"day_of_year,omitempty"`
}

type AffirmationDB struct {
    affirmations []Affirmation
    medicalMap   map[string]string
}

func NewAffirmationDB() *AffirmationDB {
    db := &AffirmationDB{
        affirmations: GetAllAffirmations(),
        medicalMap:   make(map[string]string),
    }
    db.loadMedicalMapping()
    return db
}

func (db *AffirmationDB) loadMedicalMapping() {
    db.medicalMap["alopecia"] = "облысение"
    db.medicalMap["цефалгия"] = "головная боль"
    db.medicalMap["гастрит"] = "воспаление желудка"
    db.medicalMap["бессонница"] = "нарушение сна"
    db.medicalMap["депрессия"] = "подавленное состояние"
    db.medicalMap["тревожность"] = "беспокойство"
}

func (db *AffirmationDB) FindByKeyword(keyword string) []Affirmation {
    var results []Affirmation
    for _, aff := range db.affirmations {
        for _, kw := range aff.Keywords {
            if containsKeyword(kw, keyword) {
                results = append(results, aff)
                break
            }
        }
    }
    return results
}

func (db *AffirmationDB) GetDailyAffirmation(day int) *Affirmation {
    for _, aff := range db.affirmations {
        if aff.DayOfYear == day {
            return &aff
        }
    }
    // Если не найдено, возвращаем первую
    if len(db.affirmations) > 0 {
        return &db.affirmations[0]
    }
    return nil
}

func (db *AffirmationDB) GetAll() []Affirmation {
    return db.affirmations
}

func containsKeyword(s, substr string) bool {
    return len(s) >= len(substr) && (s == substr ||
        (len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr)))
}
