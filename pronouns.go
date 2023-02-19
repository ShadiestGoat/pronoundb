package pronoundb

type Pronoun string

const (
	PR_UNSPECIFIED Pronoun = "unspecified"
	PR_ANY         Pronoun = "any"
	PR_ASK         Pronoun = "ask"
	PR_AVOID       Pronoun = "avoid"
	PR_HE_HIM      Pronoun = "hh"
	PR_HE_IT       Pronoun = "hi"
	PR_HE_SHE      Pronoun = "hs"
	PR_HE_THEY     Pronoun = "ht"
	PR_IT_HIM      Pronoun = "ih"
	PR_IT_ITS      Pronoun = "ii"
	PR_IT_SHE      Pronoun = "is"
	PR_IT_THEY     Pronoun = "it"
	PR_SHE_HE      Pronoun = "shh"
	PR_SHE_HER     Pronoun = "sh"
	PR_SHE_IT      Pronoun = "si"
	PR_SHE_THEY    Pronoun = "st"
	PR_THEY_HE     Pronoun = "th"
	PR_THEY_IT     Pronoun = "ti"
	PR_THEY_SHE    Pronoun = "ts"
	PR_THEY_THEM   Pronoun = "tt"
)


func (p *Pronoun) Default() {
	if *p == "" {
		*p = PR_UNSPECIFIED
	}
	if allPronouns[*p] == nil {
		*p = PR_UNSPECIFIED
	}
}

var allPronouns = map[Pronoun][]GenderPronoun{
	PR_AVOID: {GPR_AVOID},
	PR_ASK: {GPR_FUZZY},
	PR_UNSPECIFIED: {GPR_FUZZY},
	PR_ANY: {GPR_FUZZY, GPR_FEMALE, GPR_MALE, GPR_PERSONAL_IT},
	
	PR_THEY_THEM: {GPR_FUZZY},
	PR_THEY_HE:   {GPR_FUZZY, GPR_MALE},
	PR_THEY_SHE:  {GPR_FUZZY, GPR_FEMALE},
	PR_THEY_IT:   {GPR_FUZZY, GPR_PERSONAL_IT},

	PR_HE_HIM:  {GPR_MALE},
	PR_HE_THEY: {GPR_MALE, GPR_FUZZY},
	PR_HE_SHE:  {GPR_MALE, GPR_FEMALE},
	PR_HE_IT:   {GPR_MALE, GPR_PERSONAL_IT},
	
	PR_SHE_HER:  {GPR_FEMALE},
	PR_SHE_THEY: {GPR_FEMALE, GPR_FUZZY},
	PR_SHE_HE:   {GPR_FEMALE, GPR_MALE},
	PR_SHE_IT:   {GPR_FEMALE, GPR_PERSONAL_IT},

	PR_IT_ITS:  {GPR_PERSONAL_IT},
	PR_IT_THEY: {GPR_PERSONAL_IT, GPR_FUZZY},
	PR_IT_HIM:  {GPR_PERSONAL_IT, GPR_MALE},
	PR_IT_SHE:  {GPR_PERSONAL_IT, GPR_FEMALE},
}

func (p Pronoun) Abbreviation() string {
	info := allPronouns[p]

	switch len(info) {
	case 0:
		return "avoid"
	case 1:
		return info[0].Abbreviation()
	}

	item1 := info[0].They()
	item2 := info[1].They()

	return caps(item1) + "/" + caps(item2)
}

// Uses the default gender
func (p Pronoun) BestGender() GenderPronoun {
	return allPronouns[p][0]
} 

func (p Pronoun) They() string {
	return p.BestGender().They()
}

func (p Pronoun) Their() string {
	return p.BestGender().Their()
}

func (p Pronoun) Them() string {
	return p.BestGender().Them()
}

func (p Pronoun) Themself() string {
	return p.BestGender().Themself()
}

func (p Pronoun) Are() string {
	return p.BestGender().Are()
}

func (p Pronoun) ExtraS() string {
	return p.BestGender().ExtraS()
}

func (p Pronoun) Gender() string {
	return p.BestGender().Gender()
}

func (p Pronoun) Genders() []GenderPronoun {
	return allPronouns[p]
}
