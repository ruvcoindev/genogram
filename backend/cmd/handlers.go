package main

import (
    "encoding/json"
    "net/http"

    "genogram-system/internal/cbt"
    "genogram-system/internal/diary"
    "genogram-system/internal/psychology"
    "genogram-system/internal/symptoms"
)

func handleCBTAnalyze(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Thought string `json:"thought"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    distortions := cbt.DetectDistortions(req.Thought)
    rationalResponse := cbt.GenerateRationalResponse(req.Thought, distortions)

    distortionStrings := make([]string, len(distortions))
    for i, d := range distortions {
        distortionStrings[i] = string(d)
    }

    response := map[string]interface{}{
        "distortions":       distortionStrings,
        "rational_response": rationalResponse,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func handleGetSymptoms(w http.ResponseWriter, r *http.Request) {
    symptomsList := symptoms.GetSymptoms()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(symptomsList)
}

func handleGetTraumas(w http.ResponseWriter, r *http.Request) {
    traumas := map[string]interface{}{
        "ptsd": map[string]interface{}{
            "name_ru":      "Посттравматическое стрессовое расстройство",
            "description":  "Реакция на травматическое событие",
            "symptoms":     []string{"Флешбеки", "Избегание", "Гипервозбуждение"},
            "healing_path": []string{"EMDR", "Соматическая терапия"},
        },
        "cptsd": map[string]interface{}{
            "name_ru":      "Комплексное ПТСР",
            "description":  "Длительная травма развития",
            "symptoms":     []string{"Нарушение регуляции эмоций", "Негативная самоконцепция"},
            "healing_path": []string{"Терапия привязанности", "Работа с внутренним ребенком"},
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(traumas)
}

func handleGetAffirmations(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")
    db := psychology.NewAffirmationDB()
    affirmations := db.FindByKeyword(keyword)

    result := make([]map[string]string, len(affirmations))
    for i, a := range affirmations {
        result[i] = map[string]string{
            "id":     a.ID,
            "text":   a.Text,
            "author": a.Author,
        }
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func handleGetStats(w http.ResponseWriter, r *http.Request) {
    stats := map[string]interface{}{
        "cbt_count":   0,
        "diary_count": 0,
        "version":     "1.1.0",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

func handleGetDiarySections(w http.ResponseWriter, r *http.Request) {
    sections := diary.GetSections()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sections)
}

func handleGetDiaryQuestions(w http.ResponseWriter, r *http.Request) {
    questions := diary.GetQuestions()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(questions)
}

func handleSaveDiaryEntry(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID  string   `json:"userId"`
        Section string   `json:"section"`
        Answer  string   `json:"answer"`
        Tags    []string `json:"tags"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "saved"})
}

// GetDailyAffirmation возвращает аффирмацию дня
func handleGetDailyAffirmation(w http.ResponseWriter, r *http.Request) {
    // Получаем день года (1-365) из query параметра, или используем текущий день
    dayStr := r.URL.Query().Get("day")
    var day int
    
    if dayStr != "" {
        day, _ = strconv.Atoi(dayStr)
    } else {
        day = time.Now().YearDay()
    }
    
    db := psychology.NewAffirmationDB()
    affirmation := db.GetDailyAffirmation(day)
    
    if affirmation == nil {
        http.Error(w, "Affirmation not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(affirmation)
}

// GetAllAffirmations возвращает все аффирмации (с пагинацией)
func handleGetAllAffirmations(w http.ResponseWriter, r *http.Request) {
    limitStr := r.URL.Query().Get("limit")
    offsetStr := r.URL.Query().Get("offset")
    
    limit := 50 // по умолчанию 50
    offset := 0
    
    if limitStr != "" {
        limit, _ = strconv.Atoi(limitStr)
        if limit > 365 {
            limit = 365
        }
    }
    if offsetStr != "" {
        offset, _ = strconv.Atoi(offsetStr)
    }
    
    db := psychology.NewAffirmationDB()
    all := db.GetAll()
    
    start := offset
    end := offset + limit
    if start > len(all) {
        start = len(all)
    }
    if end > len(all) {
        end = len(all)
    }
    
    result := all[start:end]
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "total": len(all),
        "limit": limit,
        "offset": offset,
        "affirmations": result,
    })
}

// GetAffirmationByDay возвращает аффирмацию по дню года
func handleGetAffirmationByDay(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    dayStr := vars["day"]
    day, err := strconv.Atoi(dayStr)
    if err != nil || day < 1 || day > 365 {
        http.Error(w, "Invalid day (1-365)", http.StatusBadRequest)
        return
    }
    
    db := psychology.NewAffirmationDB()
    affirmation := db.GetDailyAffirmation(day)
    
    if affirmation == nil {
        http.Error(w, "Affirmation not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(affirmation)
}
