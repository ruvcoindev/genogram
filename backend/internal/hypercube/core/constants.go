package core

const CubeSize = 26

type RelationshipType string

const (
    RelationshipMotherDaughter     RelationshipType = "mother-daughter"
    RelationshipMotherSon          RelationshipType = "mother-son"
    RelationshipFatherDaughter     RelationshipType = "father-daughter"
    RelationshipFatherSon          RelationshipType = "father-son"
    RelationshipSisterSister       RelationshipType = "sister-sister"
    RelationshipBrotherBrother     RelationshipType = "brother-brother"
    RelationshipSisterBrother      RelationshipType = "sister-brother"
    RelationshipHusbandWife        RelationshipType = "husband-wife"
    RelationshipCivilPartnership   RelationshipType = "civil-partnership"
    RelationshipRomantic           RelationshipType = "romantic"
    RelationshipFriendly           RelationshipType = "friendly"
    RelationshipWork               RelationshipType = "work"
)

type Gender string

const (
    GenderMale   Gender = "male"
    GenderFemale Gender = "female"
    GenderOther  Gender = "other"
)

type FamilyRole string

const (
    RoleMother       FamilyRole = "mother"
    RoleFather       FamilyRole = "father"
    RoleDaughter     FamilyRole = "daughter"
    RoleSon          FamilyRole = "son"
    RoleGrandmother  FamilyRole = "grandmother"
    RoleGrandfather  FamilyRole = "grandfather"
    RoleGranddaughter FamilyRole = "granddaughter"
    RoleGrandson     FamilyRole = "grandson"
    RoleHusband      FamilyRole = "husband"
    RoleWife         FamilyRole = "wife"
    RolePartner      FamilyRole = "partner"
)

type TraumaRole string

const (
    TraumaRoleGoldenChild     TraumaRole = "golden_child"
    TraumaRoleScapegoat       TraumaRole = "scapegoat"
    TraumaRoleLostChild       TraumaRole = "lost_child"
    TraumaRoleInvisible       TraumaRole = "invisible"
    TraumaRoleGlassChild      TraumaRole = "glass_child"
    TraumaRoleParentified     TraumaRole = "parentified"
    TraumaRoleMascot          TraumaRole = "mascot"
    TraumaRoleEmotionalSpouse TraumaRole = "emotional_spouse"
    TraumaRoleReplacement     TraumaRole = "replacement_child"
    TraumaRoleGhost           TraumaRole = "ghost"
    TraumaRoleAncestor        TraumaRole = "ancestor"
    TraumaRoleHealer          TraumaRole = "healer"
)

type LossType string

const (
    LossTypeAbortion    LossType = "abortion"
    LossTypeMiscarriage LossType = "miscarriage"
    LossTypeStillborn   LossType = "stillborn"
    LossTypeInfantDeath LossType = "infant_death"
    LossTypeChildDeath  LossType = "child_death"
    LossTypeAdultDeath  LossType = "adult_death"
)

type HypercubeCoords struct {
    X int32
    Y int32
    Z int32
    W int32
}

type PersonVectors struct {
    X    []int32
    Y    []int32
    Z    []int32
    Full [3][]int32
}

type IdentifiedTraumaRoles struct {
    PrimaryRole   TraumaRole
    SecondaryRoles []TraumaRole
    Confidence    float64
    Evidence      []string
}

var TraumaThresholds = struct {
    ErraticMin      int
    FrozenMax       int
    RepetitiveCount int
}{
    ErraticMin:      3,
    FrozenMax:       1,
    RepetitiveCount: 2,
}
