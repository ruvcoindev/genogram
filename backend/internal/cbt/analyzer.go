package cbt

import (
    "strings"
)

func DetectDistortions(text string) []CognitiveDistortion {
    var found []CognitiveDistortion
    textLower := strings.ToLower(text)

    if containsAny(textLower, []string{"всегда", "никогда", "все", "никто", "полностью", "совсем"}) {
        found = append(found, DistortionAllOrNothing)
    }
    if containsAny(textLower, []string{"опять", "снова", "постоянно", "каждый раз", "всё время"}) {
        found = append(found, DistortionOvergeneralization)
    }
    if containsAny(textLower, []string{"ужас", "катастрофа", "конец", "никогда не", "всё пропало"}) {
        found = append(found, DistortionMagnification)
    }
    if containsAny(textLower, []string{"должен", "должна", "надо", "следует", "обязан", "обязана"}) {
        found = append(found, DistortionShouldStatements)
    }
    if containsAny(textLower, []string{"чувствую, что", "мне кажется, что", "я знаю, что"}) {
        found = append(found, DistortionEmotionalReasoning)
    }
    if containsAny(textLower, []string{"я неудачник", "я слабак", "он манипулятор", "она истеричка"}) {
        found = append(found, DistortionLabeling)
    }
    if containsAny(textLower, []string{"это из-за меня", "я виноват", "моя вина"}) {
        found = append(found, DistortionPersonalization)
    }

    return found
}

func GenerateRationalResponse(automaticThought string, distortions []CognitiveDistortion) string {
    response := "Альтернативный взгляд: "

    for _, d := range distortions {
        switch d {
        case DistortionAllOrNothing:
            response += "Возможно, есть оттенки между «всё» и «ничего». "
        case DistortionMagnification:
            response += "Что самое реалистичное может произойти? "
        case DistortionShouldStatements:
            response += "Что будет, если заменить «должен» на «хочу»? "
        case DistortionLabeling:
            response += "Это поведение или вся личность? "
        case DistortionPersonalization:
            response += "Что из этого действительно зависит от вас? "
        }
    }

    if response == "Альтернативный взгляд: " {
        response += "Какие есть другие способы посмотреть на эту ситуацию?"
    }

    return response
}

func containsAny(s string, substrs []string) bool {
    for _, sub := range substrs {
        if strings.Contains(s, sub) {
            return true
        }
    }
    return false
}
