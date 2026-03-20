package diary

type Section struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Order       int    `json:"order"`
}

type Question struct {
    ID          string   `json:"id"`
    Section     string   `json:"section"`
    Text        string   `json:"text"`
    Placeholder string   `json:"placeholder"`
    MinLength   int      `json:"min_length"`
    MaxLength   int      `json:"max_length"`
    Tags        []string `json:"tags"`
    Order       int      `json:"order"`
}

func GetSections() []Section {
    return []Section{
        {ID: "motivation", Title: "Мотивация", Description: "Зачем я это делаю? Ради кого или чего?", Order: 1},
        {ID: "boundaries", Title: "Границы", Description: "Где заканчиваюсь «я» и начинается «другой»?", Order: 2},
        {ID: "resources", Title: "Ресурс", Description: "Работа, деньги, энергия — что я вкладываю и что получаю?", Order: 3},
        {ID: "patterns", Title: "Паттерны", Description: "Узнаю ли я это? Я уже проходил этот сценарий?", Order: 4},
        {ID: "choice", Title: "Выбор", Description: "Что я выбираю сегодня — ради себя?", Order: 5},
    }
}

func GetQuestions() []Question {
    return []Question{
        {ID: "motivation_1", Section: "motivation", Text: "Ради кого я это делаю?", Placeholder: "Напиши имя или «для себя»", MinLength: 5, MaxLength: 500, Tags: []string{"motivation"}, Order: 1},
        {ID: "motivation_2", Section: "motivation", Text: "Если этот человек исчезнет — буду ли я делать это всё равно?", Placeholder: "Честно: да или нет?", MinLength: 10, MaxLength: 1000, Tags: []string{"motivation"}, Order: 2},
        {ID: "boundaries_1", Section: "boundaries", Text: "Где заканчиваюсь «я» и начинается «другой»?", Placeholder: "Что моё, а что не моё?", MinLength: 10, MaxLength: 1000, Tags: []string{"boundaries"}, Order: 1},
        {ID: "boundaries_2", Section: "boundaries", Text: "Я беру на себя ответственность за чужие чувства?", Placeholder: "Что именно?", MinLength: 10, MaxLength: 1000, Tags: []string{"boundaries"}, Order: 2},
        {ID: "resources_1", Section: "resources", Text: "Я работаю ради потока или ради доказательства?", Placeholder: "Деньги или подтверждение ценности?", MinLength: 10, MaxLength: 1000, Tags: []string{"resources"}, Order: 1},
        {ID: "patterns_1", Section: "patterns", Text: "Я уже проходил этот сценарий? Когда? Чем закончилось?", Placeholder: "Вспомни похожие ситуации", MinLength: 20, MaxLength: 2000, Tags: []string{"patterns"}, Order: 1},
        {ID: "choice_1", Section: "choice", Text: "Я действую из страха или из интереса?", Placeholder: "Что движет тобой?", MinLength: 10, MaxLength: 1000, Tags: []string{"choice"}, Order: 1},
        {ID: "choice_5", Section: "choice", Text: "Что я выбираю сегодня — ради себя?", Placeholder: "Не «надо». Не «правильно». А сейчас.", MinLength: 10, MaxLength: 1000, Tags: []string{"choice", "final"}, Order: 5},
    }
}

func GetStopSignals() []string {
    return []string{
        "Я проверяю статус/онлайн человека, которому обещал «не писать»",
        "Я отменяю свои планы, потому что «вдруг она напишет»",
        "Я объясняю одно и то же человеку, который уже сказал «нет»",
        "Я чувствую, что «должен» доказать, что я хороший",
    }
}
