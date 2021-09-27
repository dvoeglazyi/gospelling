package templates

import "strings"

const (
	SpeechPartNoun         = "сущ"
	SpeechPartVerb         = "гл"
	SpeechPartAdjective    = "прил"
	SpeechPartPronoun      = "мест"
	SpeechPartNumeral      = "числ"
	SpeechPartParticiple   = "прич"
	SpeechPartPreposition  = "prep"   // предлог
	SpeechPartConjunction  = "conj"   // союз
	SpeechPartAdverb       = "adv"    // наречие
	SpeechPartPart         = "part"   // частица
	SpeechPartInterjection = "interj" // междометие
	SpeechPartSurname      = "Фам"
)

const (
	LanguagePrefixRU    = " ru"
	LanguagePrefixOldRU = " ru-old"
)

var SpeechParts = []string{SpeechPartNoun, SpeechPartVerb, SpeechPartAdjective, SpeechPartPronoun, SpeechPartNumeral, SpeechPartParticiple,
	SpeechPartPreposition, SpeechPartConjunction, SpeechPartAdverb, SpeechPartPart, SpeechPartInterjection, SpeechPartSurname}

var (
	speechParts      = []string{SpeechPartNumeral, SpeechPartParticiple, SpeechPartSurname}
	ruSpeechParts    = []string{SpeechPartVerb, SpeechPartAdjective, SpeechPartPronoun, SpeechPartPreposition, SpeechPartConjunction, SpeechPartAdverb, SpeechPartPart, SpeechPartInterjection, SpeechPartNoun}
	ruOldSpeechParts = []string{SpeechPartVerb}
)

func (t *Template) detectSpeechPart() {
	for _, part := range speechParts {
		if strings.HasPrefix(t.Title, part) {
			t.SpeechPart = part
		}
	}
	for _, part := range ruSpeechParts {
		if strings.HasPrefix(t.Title, part+LanguagePrefixRU) {
			t.SpeechPart = part
		}
	}
	for _, part := range ruOldSpeechParts {
		if strings.HasPrefix(t.Title, part+LanguagePrefixOldRU) {
			t.SpeechPart = part
		}
	}
}

func (t *Template) isPureSpeechPart() bool {
	return t.SpeechPart == SpeechPartPreposition || t.SpeechPart == SpeechPartConjunction || t.SpeechPart == SpeechPartAdverb ||
		t.SpeechPart == SpeechPartPart || t.SpeechPart == SpeechPartInterjection
}
