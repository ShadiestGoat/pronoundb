# PronounDB

[![Go Reference](https://pkg.go.dev/badge/github.com/ShadiestGoat/pronoundb.svg)](https://pkg.go.dev/github.com/ShadiestGoat/pronoundb)

This is an abstraction/wrapper to the [pronoundb](https://pronoundb.org) project. This repository currently supports V2 of the api

This package also includes several tools for grammar to do with these pronouns - a generalizes they/them/their/themself function, and tools for determining things like *is* vs *are*, or if an extra *s* is needed for verbs (think she *is*, she *prefers* vs they *are*, they *prefer*)

## Notes

- Some have 'meta' pronouns which must handled separately. These are `PR_ANY`, `PR_ASK`, `PR_AVOID`, `PR_OTHER`. `PR_ASK`, `PR_AVOID`, `PR_OTHER` return empty functions for all grammar functions. `PR_ANY` returns an empty string for `Abbreviation()`, and everything else would return results that are the same as `PR_THEY`. You can use `IsNominative()` to quickly determine this.
- You *should* use a user agent - it can help avoid rate limits

## Example

```go
package main

import "github.com/ShadiestGoat/pronoundb/v2"

func main() {
	// create new client
    c := pronoundb.NewClient(pronoundb.WithUserAgent(pronoundb.UserAgent{"ShadyBot", "v3", "githubSite"}))
	user := "654034575"

	pronouns, err := c.Lookup(pronoundb.PLATFORM_TWITCH, user)
	if err != nil {
		panic("Error when looking up pronouns: " + err.Error())
	}
	if _, ok := pronouns[user]; !ok {
		panic("Unknown user")
	}

	allPronouns := pronouns[user]
	bestPr := allPronouns[0]

	if bestPr.IsNominative() {
		fmt.Printf("User is %v - %v use %v pronouns!\n", bestPr.Gender(), pr.They(), pr.Abbreviation())
		if len(allPronouns) > 1 && allPronouns[1].IsNominative() {
			pr2 := allPronouns[1]
			fmt.Printf("%v also consider %v %v - you can use %v pronouns :3", pr2.They(), pr2.Themself(), pr2.Gender(), pr2.Abbreviation())
		}

		return
	}

	switch bestPr {
		case PR_ANY:
			fmt.Println("User doesn't have a preference in pronouns")
        case PR_ASK:
			fmt.Println("User wishes that you asked - don't use any pronouns/gendered language till you do so")
        case PR_AVOID:
			fmt.Println("User wishes to avoid all pronouns & gendered language. Please respect that :3")
        case PR_OTHER:
			fmt.Println("User does not fall under the traditional gender range - please ask before you use any gendered language or pronouns :3!")
	}
}
```

## Migrating from v1

Non-breaking changes (additions):
- Added `WithCustomHeaders`, `WithUserAgent` client options
- Added `IsNominative()` to UsefulGrammar
- Bulk Lookup now auto separates your IDs

Breaking Changes:
- Using `v2` of the pronoundb api.
- Got rid of the `PR_*_*` pronoun setup - thats not how v2 works
- `(*Pronoun).Default()` is no longer required (and is removed)
- `(*Pronoun).BestGender()` is no longer a thing - pronouns are in an array
- `GenderPronoun` is not a thing anymore - theres not difference between `Pronoun` and `GenderPronoun` in v2. `IsNominative()` is helpful here though!
  - Grammar methods have been mved to `Pronoun`
- `RawLookup` is removed (not a thing in `v2`)
- `Lookup` is now the 'bulk lookup' function
