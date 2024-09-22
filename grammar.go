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

	IsNominative() bool
}

func (p Pronoun) They() string {
	switch p {
	case PR_SHE:
		return "she"
	case PR_HE:
		return "he"
	case PR_IT:
		return "it"
	case PR_THEY, PR_ANY:
		return "they"
	}

	return ""
}

func (p Pronoun) Their() string {
	switch p {
	case PR_SHE:
		return "her"
	case PR_HE:
		return "his"
	case PR_IT:
		return "its"
	case PR_THEY, PR_ANY:
		return "their"
	}
	return ""
}

func (p Pronoun) Them() string {
	switch p {
	case PR_SHE:
		return "her"
	case PR_HE:
		return "him"
	case PR_IT:
		return "it"
	case PR_THEY, PR_ANY:
		return "them"
	}

	return ""
}

func (p Pronoun) Themself() string {
	switch p {
	case PR_SHE:
		return "herself"
	case PR_HE:
		return "himself"
	case PR_IT:
		return "itself"
	case PR_THEY, PR_ANY:
		return "themself"
	}

	return ""
}

// used for gender neutral 'they' vs gender 'he/she':
//
// She *is*, they *are*
func (p Pronoun) Are() string {
	if p == PR_THEY || p == PR_ANY {
		return "are"
	}

	return "is"
}

// used for verb conjugation:
//
// She prefer*s*, they *prefer**
func (p Pronoun) ExtraS() string {
	switch p {
	case PR_THEY, PR_IT, PR_ANY:
		return ""
	}

	return "s"
}

func (p Pronoun) Gender() string {
	switch p {
	case PR_SHE:
		return "female"
	case PR_HE:
		return "male"
	case PR_THEY, PR_IT:
		return "gender neutral"
	}

	return ""
}

func (p Pronoun) Abbreviation() string {
	if p == PR_ANY {
		return ""
	}

	they := p.They()
	them := p.Them()

	return they + "/" + them
}

func (p Pronoun) IsNominative() bool {
	switch p {
	case PR_SHE, PR_HE, PR_THEY, PR_IT:
		return true
	}

	return false
}
