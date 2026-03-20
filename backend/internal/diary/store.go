package diary

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "time"
)

type Entry struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Section   string    `json:"section"`
    Question  string    `json:"question"`
    Answer    string    `json:"answer"`
    Tags      []string  `json:"tags"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Store struct {
    dataDir string
    entries map[string]map[string]*Entry
    mu      sync.RWMutex
}

func NewStore(dataDir string) (*Store, error) {
    if dataDir == "" {
        home, _ := os.UserHomeDir()
        dataDir = filepath.Join(home, ".genogram", "diary")
    }
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        return nil, err
    }
    store := &Store{
        dataDir: dataDir,
        entries: make(map[string]map[string]*Entry),
    }
    store.load()
    return store, nil
}

func (s *Store) SaveEntry(userID, section, answer string, tags []string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    questions := GetQuestionsBySection(section)
    if len(questions) == 0 {
        return fmt.Errorf("no questions for section: %s", section)
    }

    if s.entries[userID] == nil {
        s.entries[userID] = make(map[string]*Entry)
    }

    now := time.Now()
    for _, q := range questions {
        if _, ok := s.entries[userID][q.ID]; !ok {
            s.entries[userID][q.ID] = &Entry{
                ID:        fmt.Sprintf("%s:%s:%s", userID, section, q.ID),
                UserID:    userID,
                Section:   section,
                Question:  q.Text,
                Answer:    answer,
                Tags:      tags,
                CreatedAt: now,
                UpdatedAt: now,
            }
        } else {
            s.entries[userID][q.ID].Answer = answer
            s.entries[userID][q.ID].Tags = tags
            s.entries[userID][q.ID].UpdatedAt = now
        }
    }
    return s.save(userID)
}

func GetQuestionsBySection(sectionID string) []Question {
    all := GetQuestions()
    var result []Question
    for _, q := range all {
        if q.Section == sectionID {
            result = append(result, q)
        }
    }
    return result
}

func (s *Store) load() error {
    files, _ := filepath.Glob(filepath.Join(s.dataDir, "*.json"))
    for _, file := range files {
        data, _ := os.ReadFile(file)
        var entries map[string]*Entry
        if err := json.Unmarshal(data, &entries); err != nil {
            continue
        }
        userID := filepath.Base(file)
        userID = userID[:len(userID)-5]
        s.entries[userID] = entries
    }
    return nil
}

func (s *Store) save(userID string) error {
    file := filepath.Join(s.dataDir, fmt.Sprintf("%s.json", userID))
    data, _ := json.MarshalIndent(s.entries[userID], "", "  ")
    return os.WriteFile(file, data, 0600)
}
