package pronoundb

type UsefulGrammar interface {
	They() string
	Their() string
	Them() string
	Themself() string
	Are() string
	ExtraS() string
	Abbreviation() string
	Gender() string
}

type GenderPronoun string

const (
	GPR_FEMALE      GenderPronoun = "she"
	GPR_MALE        GenderPronoun = "he"
	GPR_PERSONAL_IT GenderPronoun = "it"
	GPR_FUZZY       GenderPronoun = "they"

	// Please note, this should be handled yourself 
	GPR_AVOID       GenderPronoun = "avoid"
)

func (g GenderPronoun) They() string {
	switch g {
	case GPR_FEMALE:
		return "she"
	case GPR_MALE:
		return "he"
	case GPR_PERSONAL_IT:
		return "it"
	case GPR_FUZZY:
		return "they"
	}

	return ""
}

func (g GenderPronoun) Their() string {
	switch g {
	case GPR_FEMALE:
		return "her"
	case GPR_MALE:
		return "his"
	case GPR_PERSONAL_IT:
		return "its"
	case GPR_FUZZY:
		return "their"
	}
	return ""
}

func (g GenderPronoun) Them() string {
	switch g {
	case GPR_FEMALE:
		return "her"
	case GPR_MALE:
		return "him"
	case GPR_PERSONAL_IT:
		return "it"
	case GPR_FUZZY:
		return "them"
	}

	return ""
}

func (g GenderPronoun) Themself() string {
	switch g {
	case GPR_FEMALE:
		return "herself"
	case GPR_MALE:
		return "himself"
	case GPR_PERSONAL_IT:
		return "itself"
	case GPR_FUZZY:
		return "themself"
	}

	return ""
}

// used for gender neutral 'they' vs gender 'he/she':
// 
// She *is*, they *are* 
func (g GenderPronoun) Are() string {
	if g == GPR_AVOID {
		return ""
	}

	if g == GPR_FUZZY {
		return "are"
	}

	return "is"
}

// used for verb conjugation:
// 
// She prefer*s*, they *prefer**
func (g GenderPronoun) ExtraS() string {
	switch g {
	case GPR_FUZZY, GPR_AVOID, GPR_PERSONAL_IT:
		return ""
	}
	
	return "s"
}

func (g GenderPronoun) Gender() string {
	switch g {
	case GPR_FEMALE:
		return "female"
	case GPR_MALE:
		return "male"
	case GPR_FUZZY:
		return "gender neutral"
	case GPR_PERSONAL_IT:
		return "gender neutral (personal 'it')"
	}

	return "avoid"
}

func (g GenderPronoun) Abbreviation() string {
	if g == GPR_AVOID {
		return "avoid"
	}

	they := g.They()
	them := g.Them()

	return they + "/" + them
}
