package psychology

type Affirmation struct {
    ID          string   `json:"id"`
    Text        string   `json:"text"`
    Author      string   `json:"author"`
    Keywords    []string `json:"keywords"`
    ChakraIndex int      `json:"chakra_index"`
}

type AffirmationDB struct {
    affirmations map[string]Affirmation
    medicalMap   map[string]string
}

func NewAffirmationDB() *AffirmationDB {
    db := &AffirmationDB{
        affirmations: make(map[string]Affirmation),
        medicalMap:   make(map[string]string),
    }
    db.load()
    return db
}

func (db *AffirmationDB) load() {
    db.affirmations["lh_001"] = Affirmation{
        ID: "lh_001", Text: "Я люблю и одобряю себя", Author: "Луиза Хей",
        Keywords: []string{"self-love", "самолюбие"}, ChakraIndex: 4,
    }
    db.affirmations["lh_002"] = Affirmation{
        ID: "lh_002", Text: "Я в безопасности. Всё хорошо в моём мире", Author: "Луиза Хей",
        Keywords: []string{"safety", "безопасность"}, ChakraIndex: 0,
    }
    db.affirmations["lh_003"] = Affirmation{
        ID: "lh_003", Text: "Я достойна любви и уважения", Author: "Луиза Хей",
        Keywords: []string{"self-worth", "достоинство"}, ChakraIndex: 3,
    }
    db.medicalMap["alopecia"] = "облысение"
    db.medicalMap["цефалгия"] = "головная боль"
    db.medicalMap["гастрит"] = "воспаление желудка"
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

func containsKeyword(s, substr string) bool {
    return len(s) >= len(substr) && (s == substr || 
        (len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr)))
}
