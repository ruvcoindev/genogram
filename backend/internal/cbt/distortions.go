package cbt

type CognitiveDistortion string

const (
    DistortionAllOrNothing      CognitiveDistortion = "all_or_nothing"
    DistortionOvergeneralization CognitiveDistortion = "overgeneralization"
    DistortionMentalFilter      CognitiveDistortion = "mental_filter"
    DistortionDisqualifying     CognitiveDistortion = "disqualifying"
    DistortionJumpingConclusions CognitiveDistortion = "jumping_conclusions"
    DistortionMagnification     CognitiveDistortion = "magnification"
    DistortionEmotionalReasoning CognitiveDistortion = "emotional_reasoning"
    DistortionShouldStatements  CognitiveDistortion = "should_statements"
    DistortionLabeling          CognitiveDistortion = "labeling"
    DistortionPersonalization   CognitiveDistortion = "personalization"
    DistortionBlaming           CognitiveDistortion = "blaming"
)

type ThoughtRecord struct {
    ID               string                `json:"id"`
    PersonID         string                `json:"person_id"`
    Situation        string                `json:"situation"`
    AutomaticThought string                `json:"automatic_thought"`
    Emotions         []string              `json:"emotions"`
    Intensity        int                   `json:"intensity"`
    Distortions      []CognitiveDistortion `json:"distortions"`
    RationalResponse string                `json:"rational_response"`
    NewIntensity     int                   `json:"new_intensity"`
    Completed        bool                  `json:"completed"`
}
