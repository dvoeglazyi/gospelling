package templates

// Form суммарный набор характеристик формы слова для всех частей речи.
type Form struct {
	Name         string // название (код) формы
	Identified   bool   // опознана
	Gender       byte   // род
	Plural       bool   // множественное
	Case         byte   // падеж
	Compare      bool   // сравнительная степень (прилагательные)
	Short        bool   // краткая форма (прилагательные)
	Tense        byte   // время (глагол)
	Person       byte   // лицо (глагол)
	FromVerbForm byte   // отглагольная форма (глагол)
	Imperative   bool   // повелительное наклонение (глагол)
	Animated     bool   // одушевлённое (прилагательные)
	Pure         bool   // исходная и единственная форма (например, для предолгов)
}

var pureForm = Form{
	Name:       "pure",
	Identified: true,
	Pure:       true,
}

func (f Form) String() string {
	if !f.Identified {
		return "undetected"
	}
	return f.Name
}

// Роды.
const (
	GenderMale = 1 + iota
	GenderFemale
	GenderNeuter
)

// Лица.
const (
	PersonInfinitive = 1 + iota // инфинитив
	PersonFirst                 // первое лицо: я, мы
	PersonSecond                // второе лицо: ты,  вы
	PersonThird                 // третье лицо: он, она, оно, они
	PersonParticiple = 1 + iota // причастие
	PersonContinuous            // деепричастие
)

// Времена.
const (
	TensePast    = 1 + iota // прошедшее
	TensePresent            // настоящее
	TenseFuture             // будущее
)

// Падежи.
const (
	CaseNominative    = 1 + iota // именительный
	CaseGenitive                 // родительный
	CaseDative                   // дательный
	CaseAccusative               // винительный
	CaseInstrumental             // творительный
	CasePrepositional            // предложный
	CaseVocative                 // звательный (прим. "господи", "боже")
	CaseLocative                 // местный (когда ответ на вопрос "где?" не в предложном падеже)
	CaseParticular               // разделительный, частичный (при употреблении по отношению к части используется вместо родительного)
	CaseCounting                 // счётный (когда используется винительный вместо родительного, напр. 5 грамм вместо 5 граммов)
)

var forms = map[string]Form{
	"nom":         {Case: CaseNominative},
	"gen":         {Case: CaseGenitive},
	"dat":         {Case: CaseDative},
	"acc":         {Case: CaseAccusative},
	"ins":         {Case: CaseInstrumental},
	"prp":         {Case: CasePrepositional},
	"nom-sg-m":    {Gender: GenderMale, Case: CaseNominative},
	"nom-sg-n":    {Gender: GenderNeuter, Case: CaseNominative},
	"nom-sg-f":    {Gender: GenderFemale, Case: CaseNominative},
	"nom-pl":      {Plural: true, Case: CaseNominative},
	"gen-sg-m":    {Gender: GenderMale, Case: CaseGenitive},
	"gen-sg-n":    {Gender: GenderNeuter, Case: CaseGenitive},
	"gen-sg-f":    {Gender: GenderFemale, Case: CaseGenitive},
	"gen-pl":      {Plural: true, Case: CaseGenitive},
	"dat-sg-m":    {Gender: GenderMale, Case: CaseDative},
	"dat-sg-n":    {Gender: GenderNeuter, Case: CaseDative},
	"dat-sg-f":    {Gender: GenderFemale, Case: CaseDative},
	"dat-pl":      {Plural: true, Case: CaseDative},
	"acc-sg-m-a":  {Gender: GenderMale, Case: CaseAccusative, Animated: true},
	"acc-sg-m-n":  {Gender: GenderMale, Case: CaseAccusative},
	"acc-sg-n":    {Gender: GenderNeuter, Case: CaseAccusative},
	"acc-sg-f":    {Gender: GenderFemale, Case: CaseAccusative},
	"acc-pl-a":    {Plural: true, Case: CaseAccusative, Animated: true},
	"acc-pl-n":    {Plural: true, Case: CaseAccusative},
	"ins-sg-m":    {Gender: GenderMale, Case: CaseInstrumental},
	"ins-sg-n":    {Gender: GenderNeuter, Case: CaseInstrumental},
	"ins-sg-f":    {Gender: GenderFemale, Case: CaseInstrumental},
	"ins-pl":      {Plural: true, Case: CaseInstrumental},
	"prp-sg-m":    {Gender: GenderMale, Case: CasePrepositional},
	"prp-sg-n":    {Gender: GenderNeuter, Case: CasePrepositional},
	"prp-sg-f":    {Gender: GenderFemale, Case: CasePrepositional},
	"prp-pl":      {Plural: true, Case: CasePrepositional},
	"srt-sg-m":    {Gender: GenderMale, Short: true},
	"srt-sg-n":    {Gender: GenderNeuter, Short: true},
	"srt-sg-f":    {Gender: GenderFemale, Short: true},
	"srt-pl":      {Plural: true, Short: true},
	"им-м":        {Gender: GenderMale, Case: CaseNominative},
	"им-ж":        {Gender: GenderFemale, Case: CaseNominative},
	"им-мн":       {Plural: true, Case: CaseNominative},
	"род-м":       {Gender: GenderMale, Case: CaseGenitive},
	"род-ж":       {Gender: GenderFemale, Case: CaseGenitive},
	"род-мн":      {Plural: true, Case: CaseGenitive},
	"дат-м":       {Gender: GenderMale, Case: CaseDative},
	"дат-ж":       {Gender: GenderFemale, Case: CaseDative},
	"дат-мн":      {Plural: true, Case: CaseDative},
	"вин-м":       {Gender: GenderMale, Case: CaseAccusative},
	"вин-ж":       {Gender: GenderFemale, Case: CaseAccusative},
	"вин-мн":      {Plural: true, Case: CaseAccusative},
	"тв-м":        {Gender: GenderMale, Case: CaseInstrumental},
	"тв-ж":        {Gender: GenderFemale, Case: CaseInstrumental},
	"тв-мн":       {Plural: true, Case: CaseInstrumental},
	"пр-м":        {Gender: GenderMale, Case: CasePrepositional},
	"пр-ж":        {Gender: GenderFemale, Case: CasePrepositional},
	"пр-мн":       {Plural: true, Case: CasePrepositional},
	"Я":           {Person: PersonFirst, Tense: TensePresent}, // глаголы
	"Я (прош.)":   {Person: PersonFirst, Tense: TensePast},
	"Мы":          {Person: PersonFirst, Tense: TensePresent, Plural: true},
	"Мы (прош.)":  {Person: PersonFirst, Tense: TensePast, Plural: true},
	"Ты":          {Person: PersonSecond, Tense: TensePresent},
	"Ты (повел.)": {Person: PersonSecond, Tense: TensePresent, Imperative: true},
	"Ты (прош.)":  {Person: PersonSecond, Tense: TensePast},
	"Вы":          {Person: PersonSecond, Tense: TensePresent, Plural: true},
	"Вы (повел.)": {Person: PersonSecond, Tense: TensePresent, Plural: true, Imperative: true},
	"Вы (прош.)":  {Person: PersonSecond, Tense: TensePast, Plural: true},
	"Они":         {Person: PersonThird, Tense: TensePresent, Plural: true},
	"Они (прош.)": {Person: PersonThird, Tense: TensePast, Plural: true},
	"Прич":        {Person: PersonParticiple, Tense: TensePresent},
	"ПричНаст":    {Person: PersonParticiple, Tense: TensePresent},
	"ПричПрош":    {Person: PersonParticiple, Tense: TensePast},
	"ПричСтрад":   {Person: PersonParticiple},
	"ДеепрНаст":   {Person: PersonContinuous, Tense: TensePresent},
	"ДеепрПрош":   {Person: PersonContinuous, Tense: TensePast},
	"Инфинитив":   {Person: PersonInfinitive, Tense: TensePresent},
	"Будущее":     {Tense: TenseFuture},
}

func (t *Template) parseForm(name string) Form {
	var form Form
	switch t.SpeechPart {
	case SpeechPartNoun, SpeechPartVerb, SpeechPartAdjective, SpeechPartPronoun, SpeechPartNumeral, SpeechPartParticiple, SpeechPartSurname:
		var ok bool
		form, ok = forms[name]
		form.Identified = ok
	default:
		form.Identified = true
	}
	form.Name = name
	return form
}
