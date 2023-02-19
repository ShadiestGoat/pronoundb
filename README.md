# PronounDB

[![Go Reference](https://pkg.go.dev/badge/github.com/shadiestgoat/pronoundb.svg)](https://pkg.go.dev/github.com/shadiestgoat/pronoundb)

This is an abstraction/wrapper to the [pronoundb](https://pronoundb.org) project. This repository currently supports V1 of the api (V2 is not yet written, see [this issue](https://github.com/cyyynthia/pronoundb.org/issues/68))

This package also includes several tools for grammar to do with these pronouns - a generalizes they/them/their/themself function, and tools for determining things like *is* vs *are*, or if an extra *s* is needed for verbs (think she *is*, she *prefers* vs they *are*, they *prefer*)

## Notes

- Some people prefer that pronouns are avoided all together. For that, most of the time the grammar should be handled separately! This package will return empty strings for all the 'grammar' functions for the avoid pronouns option (both `GPR_AVOID` and `PR_AVOID`)
- If you choose to parse a `Pronoun` yourself, you should call `*Pronoun.Default()` on it, so that it can get be set into correct form for the internal `allPronouns` map. Just remember that `Pronoun` is a string of the pronoundb abbreviations internally! [Check them out here](https://pronoundb.org/docs)

## Example

```go
package main

import "github.com/shadiestgoat/pronoundb"

func main() {
	// create new client
    c := pronoundb.NewClient()

	pr, err := c.Lookup(pronoundb.PLATFORM_TWITCH, "654034575")

	if err != nil {
		panic("Error when looking up pronouns: " + err.Error())
	}

	// Avoid pronouns should always be handled separately!
	if pr == pronoundb.PR_AVOID {
		fmt.Println("Twitch streamer Shadiest Goat asks you to avoid pronouns!")
		fmt.Println("I think Shadiest Goat is the best streamer and you should totally donate on Shadiest Goat's donation page")
		fmt.Println("Who knows, Shady might even thank you personally (maybe do a little skirt speen if the bank account is running low....)")
		return
	}

	// pronoundb supports up to 2 pronouns
	if len(pr.Genders()) == 2 {
		fmt.Printf("Twitch streamer Shadiest Goat has %s pronouns, meaning %s prefer%s %s pronouns, but also fully accept%s %s pronouns!\n", pr.Abbreviation(), pr.They(), pr.ExtraS(), pr.BestGender().Gender(), pr.ExtraS(), pr.Genders()[1].Gender())
	} else {
		fmt.Printf("Twitch streamer Shadiest Goat is %s and therefor has the following pronouns: %s\n", pr.BestGender(), pr.Abbreviation())
	}

	fmt.Printf("I think %s %s the best streamer, and you should donate to %s\n", pr.They(), pr.Are(), pr.Them())
	fmt.Printf("Who knows, %s might even thank you %s (maybe even do a little skirt speen if %s bank account is running low....)\n", pr.They(), pr.Themself(), pr.Their())
}
```
