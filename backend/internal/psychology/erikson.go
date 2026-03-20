package psychology

type Stage struct {
    Number       int    `json:"number"`
    AgeRange     string `json:"age_range"`
    Crisis       string `json:"crisis"`
    Virtue       string `json:"virtue"`
    SignsOfBlock string `json:"signs_of_block"`
    HealingFocus string `json:"healing_focus"`
}

func GetAllStages() []Stage {
    return []Stage{
        {Number: 1, AgeRange: "0-1.5 года", Crisis: "Доверие vs Недоверие", Virtue: "Надежда", SignsOfBlock: "Базовое недоверие", HealingFocus: "Создание безопасности"},
        {Number: 2, AgeRange: "1.5-3 года", Crisis: "Автономия vs Стыд", Virtue: "Воля", SignsOfBlock: "Стыд за желания", HealingFocus: "Развитие воли"},
        {Number: 3, AgeRange: "3-6 лет", Crisis: "Инициатива vs Вина", Virtue: "Цель", SignsOfBlock: "Чувство вины", HealingFocus: "Здоровая инициатива"},
        {Number: 4, AgeRange: "6-12 лет", Crisis: "Трудолюбие vs Неполноценность", Virtue: "Компетентность", SignsOfBlock: "Чувство неполноценности", HealingFocus: "Развитие самоценности"},
        {Number: 5, AgeRange: "12-18 лет", Crisis: "Идентичность vs Смешение", Virtue: "Верность", SignsOfBlock: "Путаница идентичности", HealingFocus: "Поиск голоса"},
        {Number: 6, AgeRange: "18-40 лет", Crisis: "Близость vs Изоляция", Virtue: "Любовь", SignsOfBlock: "Страх близости", HealingFocus: "Развитие близости"},
        {Number: 7, AgeRange: "40-65 лет", Crisis: "Продуктивность vs Стагнация", Virtue: "Забота", SignsOfBlock: "Стагнация", HealingFocus: "Забота о других"},
        {Number: 8, AgeRange: "65+ лет", Crisis: "Целостность vs Отчаяние", Virtue: "Мудрость", SignsOfBlock: "Отчаяние", HealingFocus: "Принятие пути"},
    }
}

func GetStageByAge(age int) Stage {
    stages := GetAllStages()
    if age <= 1 {
        return stages[0]
    } else if age <= 3 {
        return stages[1]
    } else if age <= 6 {
        return stages[2]
    } else if age <= 12 {
        return stages[3]
    } else if age <= 18 {
        return stages[4]
    } else if age <= 40 {
        return stages[5]
    } else if age <= 65 {
        return stages[6]
    }
    return stages[7]
}
