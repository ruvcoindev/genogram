package symptoms

type Symptom struct {
    ID             string   `json:"id"`
    Name           string   `json:"name"`
    Category       string   `json:"category"`
    Description    string   `json:"description"`
    RelatedChakras []int    `json:"related_chakras"`
    RelatedHormones []string `json:"related_hormones"`
    WhenToSeeDoctor string   `json:"when_to_see_doctor"`
}

func GetSymptoms() []Symptom {
    return []Symptom{
        {ID: "PHYS_001", Name: "Хроническая усталость", Category: "physical",
            Description: "Постоянное чувство усталости", RelatedChakras: []int{0, 2, 6},
            RelatedHormones: []string{"Cortisol", "TSH"}, WhenToSeeDoctor: "Если длится более 2 недель"},
        {ID: "PHYS_002", Name: "Головные боли", Category: "physical",
            Description: "Боль в области головы", RelatedChakras: []int{5, 6},
            RelatedHormones: []string{"Cortisol", "Estrogen"}, WhenToSeeDoctor: "При внезапной сильной боли"},
        {ID: "EMO_001", Name: "Тревожность", Category: "emotional",
            Description: "Постоянное чувство беспокойства", RelatedChakras: []int{0, 1, 4},
            RelatedHormones: []string{"Cortisol", "GABA"}, WhenToSeeDoctor: "При панических атаках"},
        {ID: "EMO_002", Name: "Депрессивное состояние", Category: "emotional",
            Description: "Стойкое снижение настроения", RelatedChakras: []int{3, 6},
            RelatedHormones: []string{"Serotonin", "Dopamine"}, WhenToSeeDoctor: "При суицидальных мыслях"},
        {ID: "MENT_001", Name: "Туман в голове", Category: "mental",
            Description: "Нечёткость мышления", RelatedChakras: []int{5, 6},
            RelatedHormones: []string{"TSH", "Vitamin B12"}, WhenToSeeDoctor: "При внезапном появлении"},
    }
}
