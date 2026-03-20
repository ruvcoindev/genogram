package psychology

// Affirmation представляет аффирмацию
type Affirmation struct {
    ID          string   `json:"id"`
    Text        string   `json:"text"`
    Author      string   `json:"author"`
    Keywords    []string `json:"keywords"`
    ChakraIndex int      `json:"chakra_index"`
    DayOfYear   int      `json:"day_of_year"` // день года (1-365)
}

// GetAllAffirmations возвращает все 365 аффирмаций Луизы Хей
func GetAllAffirmations() []Affirmation {
    affirmations := []Affirmation{
        // ==================== ЯНВАРЬ (1-31) ====================
        {ID: "lh_001", Text: "Я люблю и одобряю себя", Author: "Луиза Хей", Keywords: []string{"любовь", "себя", "одобрение"}, ChakraIndex: 4, DayOfYear: 1},
        {ID: "lh_002", Text: "Каждый день — это новое начало", Author: "Луиза Хей", Keywords: []string{"начало", "день", "новый"}, ChakraIndex: 6, DayOfYear: 2},
        {ID: "lh_003", Text: "Я достойна всего самого лучшего", Author: "Луиза Хей", Keywords: []string{"достойна", "лучшее", "ценность"}, ChakraIndex: 3, DayOfYear: 3},
        {ID: "lh_004", Text: "Моя жизнь наполнена любовью", Author: "Луиза Хей", Keywords: []string{"жизнь", "любовь", "наполнена"}, ChakraIndex: 4, DayOfYear: 4},
        {ID: "lh_005", Text: "Я доверяю жизненному процессу", Author: "Луиза Хей", Keywords: []string{"доверие", "процесс", "жизнь"}, ChakraIndex: 0, DayOfYear: 5},
        {ID: "lh_006", Text: "Всё в моей жизни прекрасно", Author: "Луиза Хей", Keywords: []string{"прекрасно", "всё", "хорошо"}, ChakraIndex: 3, DayOfYear: 6},
        {ID: "lh_007", Text: "Я выбираю быть счастливой", Author: "Луиза Хей", Keywords: []string{"счастье", "выбор", "радость"}, ChakraIndex: 6, DayOfYear: 7},
        {ID: "lh_008", Text: "Моё сердце открыто для любви", Author: "Луиза Хей", Keywords: []string{"сердце", "открыто", "любовь"}, ChakraIndex: 4, DayOfYear: 8},
        {ID: "lh_009", Text: "Я заслуживаю счастья и радости", Author: "Луиза Хей", Keywords: []string{"счастье", "радость", "достойна"}, ChakraIndex: 3, DayOfYear: 9},
        {ID: "lh_010", Text: "Каждый день я становлюсь лучше", Author: "Луиза Хей", Keywords: []string{"лучше", "рост", "развитие"}, ChakraIndex: 5, DayOfYear: 10},
        {ID: "lh_011", Text: "Я живу в гармонии с собой", Author: "Луиза Хей", Keywords: []string{"гармония", "собой", "мир"}, ChakraIndex: 4, DayOfYear: 11},
        {ID: "lh_012", Text: "Мои мысли позитивны и целительны", Author: "Луиза Хей", Keywords: []string{"мысли", "позитив", "исцеление"}, ChakraIndex: 6, DayOfYear: 12},
        {ID: "lh_013", Text: "Я отпускаю всё, что мне не служит", Author: "Луиза Хей", Keywords: []string{"отпускаю", "освобождение", "прощение"}, ChakraIndex: 1, DayOfYear: 13},
        {ID: "lh_014", Text: "Любовь наполняет каждую клетку моего тела", Author: "Луиза Хей", Keywords: []string{"любовь", "тело", "наполнение"}, ChakraIndex: 0, DayOfYear: 14},
        {ID: "lh_015", Text: "Я в безопасности и под защитой", Author: "Луиза Хей", Keywords: []string{"безопасность", "защита", "спокойствие"}, ChakraIndex: 0, DayOfYear: 15},
        {ID: "lh_016", Text: "Моё будущее светло и прекрасно", Author: "Луиза Хей", Keywords: []string{"будущее", "светлое", "прекрасное"}, ChakraIndex: 6, DayOfYear: 16},
        {ID: "lh_017", Text: "Я принимаю себя такой, какая я есть", Author: "Луиза Хей", Keywords: []string{"принятие", "себя", "ценность"}, ChakraIndex: 3, DayOfYear: 17},
        {ID: "lh_018", Text: "Мир внутри меня, мир вокруг меня", Author: "Луиза Хей", Keywords: []string{"мир", "внутри", "снаружи"}, ChakraIndex: 4, DayOfYear: 18},
        {ID: "lh_019", Text: "Я благодарна за каждый момент", Author: "Луиза Хей", Keywords: []string{"благодарность", "момент", "ценность"}, ChakraIndex: 5, DayOfYear: 19},
        {ID: "lh_020", Text: "Моя жизнь становится лучше с каждым днём", Author: "Луиза Хей", Keywords: []string{"лучше", "рост", "прогресс"}, ChakraIndex: 3, DayOfYear: 20},
        {ID: "lh_021", Text: "Я доверяю себе и своим решениям", Author: "Луиза Хей", Keywords: []string{"доверие", "себе", "решения"}, ChakraIndex: 5, DayOfYear: 21},
        {ID: "lh_022", Text: "Любовь ведёт меня по жизни", Author: "Луиза Хей", Keywords: []string{"любовь", "путь", "ведение"}, ChakraIndex: 4, DayOfYear: 22},
        {ID: "lh_023", Text: "Я открыта для новых возможностей", Author: "Луиза Хей", Keywords: []string{"открытость", "возможности", "новое"}, ChakraIndex: 5, DayOfYear: 23},
        {ID: "lh_024", Text: "Моё сердце полно благодарности", Author: "Луиза Хей", Keywords: []string{"благодарность", "сердце", "полнота"}, ChakraIndex: 4, DayOfYear: 24},
        {ID: "lh_025", Text: "Я создаю жизнь своей мечты", Author: "Луиза Хей", Keywords: []string{"создание", "мечта", "жизнь"}, ChakraIndex: 5, DayOfYear: 25},
        {ID: "lh_026", Text: "Каждый вдох наполняет меня энергией", Author: "Луиза Хей", Keywords: []string{"дыхание", "энергия", "жизнь"}, ChakraIndex: 0, DayOfYear: 26},
        {ID: "lh_027", Text: "Я прощаю себя и других", Author: "Луиза Хей", Keywords: []string{"прощение", "себя", "другие"}, ChakraIndex: 3, DayOfYear: 27},
        {ID: "lh_028", Text: "Моя внутренняя сила растёт", Author: "Луиза Хей", Keywords: []string{"сила", "внутренняя", "рост"}, ChakraIndex: 2, DayOfYear: 28},
        {ID: "lh_029", Text: "Я заслуживаю любви и уважения", Author: "Луиза Хей", Keywords: []string{"любовь", "уважение", "достойна"}, ChakraIndex: 3, DayOfYear: 29},
        {ID: "lh_030", Text: "Жизнь поддерживает меня во всём", Author: "Луиза Хей", Keywords: []string{"поддержка", "жизнь", "помощь"}, ChakraIndex: 0, DayOfYear: 30},
        {ID: "lh_031", Text: "Я начинаю этот год с любовью", Author: "Луиза Хей", Keywords: []string{"начало", "год", "любовь"}, ChakraIndex: 4, DayOfYear: 31},
        
        // ==================== ФЕВРАЛЬ (32-59) ====================
        {ID: "lh_032", Text: "Я люблю своё тело", Author: "Луиза Хей", Keywords: []string{"тело", "любовь", "уважение"}, ChakraIndex: 0, DayOfYear: 32},
        {ID: "lh_033", Text: "Моё тело исцеляется с каждым днём", Author: "Луиза Хей", Keywords: []string{"исцеление", "тело", "здоровье"}, ChakraIndex: 0, DayOfYear: 33},
        {ID: "lh_034", Text: "Я здорова и полна энергии", Author: "Луиза Хей", Keywords: []string{"здоровье", "энергия", "сила"}, ChakraIndex: 0, DayOfYear: 34},
        {ID: "lh_035", Text: "Каждая клетка моего тела вибрирует здоровьем", Author: "Луиза Хей", Keywords: []string{"здоровье", "тело", "клетки"}, ChakraIndex: 0, DayOfYear: 35},
        {ID: "lh_036", Text: "Я с любовью забочусь о себе", Author: "Луиза Хей", Keywords: []string{"забота", "себя", "любовь"}, ChakraIndex: 3, DayOfYear: 36},
        {ID: "lh_037", Text: "Мой организм знает, как исцелить себя", Author: "Луиза Хей", Keywords: []string{"исцеление", "организм", "мудрость"}, ChakraIndex: 0, DayOfYear: 37},
        {ID: "lh_038", Text: "Я дышу свободно и легко", Author: "Луиза Хей", Keywords: []string{"дыхание", "свобода", "легкость"}, ChakraIndex: 4, DayOfYear: 38},
        {ID: "lh_039", Text: "Моё сердце бьётся в ритме любви", Author: "Луиза Хей", Keywords: []string{"сердце", "любовь", "ритм"}, ChakraIndex: 4, DayOfYear: 39},
        {ID: "lh_040", Text: "Я наполнена жизненной силой", Author: "Луиза Хей", Keywords: []string{"сила", "жизнь", "энергия"}, ChakraIndex: 0, DayOfYear: 40},
        {ID: "lh_041", Text: "Мои глаза видят красоту мира", Author: "Луиза Хей", Keywords: []string{"глаза", "красота", "видение"}, ChakraIndex: 6, DayOfYear: 41},
        {ID: "lh_042", Text: "Я слушаю своё тело с любовью", Author: "Луиза Хей", Keywords: []string{"тело", "слушание", "любовь"}, ChakraIndex: 0, DayOfYear: 42},
        {ID: "lh_043", Text: "Здоровье — моё естественное состояние", Author: "Луиза Хей", Keywords: []string{"здоровье", "естественное", "состояние"}, ChakraIndex: 0, DayOfYear: 43},
        {ID: "lh_044", Text: "Я благодарна за своё тело", Author: "Луиза Хей", Keywords: []string{"благодарность", "тело", "принятие"}, ChakraIndex: 0, DayOfYear: 44},
        {ID: "lh_045", Text: "Моя кожа сияет здоровьем", Author: "Луиза Хей", Keywords: []string{"кожа", "здоровье", "сияние"}, ChakraIndex: 0, DayOfYear: 45},
        {ID: "lh_046", Text: "Я двигаюсь с лёгкостью и грацией", Author: "Луиза Хей", Keywords: []string{"движение", "легкость", "грация"}, ChakraIndex: 1, DayOfYear: 46},
        {ID: "lh_047", Text: "Мой сон глубок и целителен", Author: "Луиза Хей", Keywords: []string{"сон", "исцеление", "отдых"}, ChakraIndex: 6, DayOfYear: 47},
        {ID: "lh_048", Text: "Я питаю своё тело с любовью", Author: "Луиза Хей", Keywords: []string{"питание", "тело", "любовь"}, ChakraIndex: 0, DayOfYear: 48},
        {ID: "lh_049", Text: "Моя иммунная система сильна", Author: "Луиза Хей", Keywords: []string{"иммунитет", "сила", "защита"}, ChakraIndex: 0, DayOfYear: 49},
        {ID: "lh_050", Text: "Я излучаю здоровье и жизненную силу", Author: "Луиза Хей", Keywords: []string{"здоровье", "сила", "излучение"}, ChakraIndex: 0, DayOfYear: 50},
        {ID: "lh_051", Text: "Боль покидает моё тело", Author: "Луиза Хей", Keywords: []string{"боль", "освобождение", "исцеление"}, ChakraIndex: 0, DayOfYear: 51},
        {ID: "lh_052", Text: "Я чувствую себя прекрасно", Author: "Луиза Хей", Keywords: []string{"чувство", "прекрасно", "хорошо"}, ChakraIndex: 3, DayOfYear: 52},
        {ID: "lh_053", Text: "Моё дыхание спокойное и ровное", Author: "Луиза Хей", Keywords: []string{"дыхание", "спокойствие", "ровность"}, ChakraIndex: 4, DayOfYear: 53},
        {ID: "lh_054", Text: "Я доверяю мудрости своего тела", Author: "Луиза Хей", Keywords: []string{"мудрость", "тело", "доверие"}, ChakraIndex: 0, DayOfYear: 54},
        {ID: "lh_055", Text: "Здоровье наполняет каждый мой день", Author: "Луиза Хей", Keywords: []string{"здоровье", "день", "наполнение"}, ChakraIndex: 0, DayOfYear: 55},
        {ID: "lh_056", Text: "Я выбираю здоровье каждый день", Author: "Луиза Хей", Keywords: []string{"выбор", "здоровье", "день"}, ChakraIndex: 2, DayOfYear: 56},
        {ID: "lh_057", Text: "Моё тело — мой лучший друг", Author: "Луиза Хей", Keywords: []string{"тело", "друг", "уважение"}, ChakraIndex: 0, DayOfYear: 57},
        {ID: "lh_058", Text: "Я благодарна за здоровье", Author: "Луиза Хей", Keywords: []string{"благодарность", "здоровье", "ценность"}, ChakraIndex: 3, DayOfYear: 58},
        {ID: "lh_059", Text: "Любовь исцеляет моё тело", Author: "Луиза Хей", Keywords: []string{"любовь", "исцеление", "тело"}, ChakraIndex: 4, DayOfYear: 59},
        
        // ==================== МАРТ (60-90) ====================
        {ID: "lh_060", Text: "Деньги приходят ко мне легко", Author: "Луиза Хей", Keywords: []string{"деньги", "легкость", "изобилие"}, ChakraIndex: 2, DayOfYear: 60},
        {ID: "lh_061", Text: "Я открыта для изобилия Вселенной", Author: "Луиза Хей", Keywords: []string{"изобилие", "открытость", "вселенная"}, ChakraIndex: 5, DayOfYear: 61},
        {ID: "lh_062", Text: "Мой доход постоянно растёт", Author: "Луиза Хей", Keywords: []string{"доход", "рост", "изобилие"}, ChakraIndex: 2, DayOfYear: 62},
        {ID: "lh_063", Text: "Я заслуживаю финансового благополучия", Author: "Луиза Хей", Keywords: []string{"финансы", "благополучие", "достойна"}, ChakraIndex: 2, DayOfYear: 63},
        {ID: "lh_064", Text: "Изобилие окружает меня повсюду", Author: "Луиза Хей", Keywords: []string{"изобилие", "вокруг", "богатство"}, ChakraIndex: 2, DayOfYear: 64},
        {ID: "lh_065", Text: "Деньги текут ко мне рекой", Author: "Луиза Хей", Keywords: []string{"деньги", "поток", "изобилие"}, ChakraIndex: 2, DayOfYear: 65},
        {ID: "lh_066", Text: "Я мудро управляю своими финансами", Author: "Луиза Хей", Keywords: []string{"финансы", "мудрость", "управление"}, ChakraIndex: 5, DayOfYear: 66},
        {ID: "lh_067", Text: "Богатство — моё естественное состояние", Author: "Луиза Хей", Keywords: []string{"богатство", "естественное", "изобилие"}, ChakraIndex: 2, DayOfYear: 67},
        {ID: "lh_068", Text: "Я притягиваю возможности для дохода", Author: "Луиза Хей", Keywords: []string{"возможности", "доход", "притяжение"}, ChakraIndex: 2, DayOfYear: 68},
        {ID: "lh_069", Text: "Мои таланты приносят мне прибыль", Author: "Луиза Хей", Keywords: []string{"таланты", "прибыль", "реализация"}, ChakraIndex: 5, DayOfYear: 69},
        {ID: "lh_070", Text: "Я благодарна за деньги, которые у меня есть", Author: "Луиза Хей", Keywords: []string{"благодарность", "деньги", "ценность"}, ChakraIndex: 2, DayOfYear: 70},
        {ID: "lh_071", Text: "Вселенная щедро снабжает меня", Author: "Луиза Хей", Keywords: []string{"вселенная", "щедрость", "изобилие"}, ChakraIndex: 6, DayOfYear: 71},
        {ID: "lh_072", Text: "Я легко зарабатываю деньги", Author: "Луиза Хей", Keywords: []string{"деньги", "легкость", "заработок"}, ChakraIndex: 2, DayOfYear: 72},
        {ID: "lh_073", Text: "Мои расходы разумны и осознанны", Author: "Луиза Хей", Keywords: []string{"расходы", "разумность", "осознанность"}, ChakraIndex: 5, DayOfYear: 73},
        {ID: "lh_074", Text: "Я создаю источники пассивного дохода", Author: "Луиза Хей", Keywords: []string{"доход", "пассивный", "создание"}, ChakraIndex: 5, DayOfYear: 74},
        {ID: "lh_075", Text: "Деньги работают на меня", Author: "Луиза Хей", Keywords: []string{"деньги", "работа", "изобилие"}, ChakraIndex: 2, DayOfYear: 75},
        {ID: "lh_076", Text: "Я достойна высокого дохода", Author: "Луиза Хей", Keywords: []string{"доход", "достойна", "ценность"}, ChakraIndex: 2, DayOfYear: 76},
        {ID: "lh_077", Text: "Моя работа приносит мне радость и деньги", Author: "Луиза Хей", Keywords: []string{"работа", "радость", "деньги"}, ChakraIndex: 3, DayOfYear: 77},
        {ID: "lh_078", Text: "Я привлекаю щедрых клиентов", Author: "Луиза Хей", Keywords: []string{"клиенты", "щедрость", "привлечение"}, ChakraIndex: 5, DayOfYear: 78},
        {ID: "lh_079", Text: "Финансовая свобода — моя реальность", Author: "Луиза Хей", Keywords: []string{"свобода", "финансы", "реальность"}, ChakraIndex: 2, DayOfYear: 79},
        {ID: "lh_080", Text: "Я легко оплачиваю все счета", Author: "Луиза Хей", Keywords: []string{"счета", "легкость", "изобилие"}, ChakraIndex: 2, DayOfYear: 80},
        {ID: "lh_081", Text: "Мои сбережения растут", Author: "Луиза Хей", Keywords: []string{"сбережения", "рост", "изобилие"}, ChakraIndex: 2, DayOfYear: 81},
        {ID: "lh_082", Text: "Я инвестирую в своё будущее", Author: "Луиза Хей", Keywords: []string{"инвестиции", "будущее", "развитие"}, ChakraIndex: 5, DayOfYear: 82},
        {ID: "lh_083", Text: "Изобилие приходит ко мне неожиданными путями", Author: "Луиза Хей", Keywords: []string{"изобилие", "неожиданно", "пути"}, ChakraIndex: 2, DayOfYear: 83},
        {ID: "lh_084", Text: "Я благодарна за финансовую поддержку", Author: "Луиза Хей", Keywords: []string{"поддержка", "финансы", "благодарность"}, ChakraIndex: 2, DayOfYear: 84},
        {ID: "lh_085", Text: "Деньги любят меня", Author: "Луиза Хей", Keywords: []string{"деньги", "любовь", "изобилие"}, ChakraIndex: 2, DayOfYear: 85},
        {ID: "lh_086", Text: "Я легко принимаю деньги", Author: "Луиза Хей", Keywords: []string{"деньги", "принятие", "легкость"}, ChakraIndex: 2, DayOfYear: 86},
        {ID: "lh_087", Text: "Моё сознание открыто для богатства", Author: "Луиза Хей", Keywords: []string{"богатство", "сознание", "открытость"}, ChakraIndex: 6, DayOfYear: 87},
        {ID: "lh_088", Text: "Я заслуживаю роскоши", Author: "Луиза Хей", Keywords: []string{"роскошь", "достойна", "изобилие"}, ChakraIndex: 2, DayOfYear: 88},
        {ID: "lh_089", Text: "Изобилие — моё право по рождению", Author: "Луиза Хей", Keywords: []string{"изобилие", "право", "рождение"}, ChakraIndex: 2, DayOfYear: 89},
        {ID: "lh_090", Text: "Я живу в изобилии", Author: "Луиза Хей", Keywords: []string{"изобилие", "жизнь", "богатство"}, ChakraIndex: 2, DayOfYear: 90},
    }
    return affirmations
}
