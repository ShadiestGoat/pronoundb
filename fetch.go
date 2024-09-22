package pronoundb

import "strings"

type Platform string

const (
	PLATFORM_DISCORD   Platform = "discord"
	PLATFORM_GITHUB    Platform = "github"
	PLATFORM_MINECRAFT Platform = "minecraft"
	PLATFORM_TWITCH    Platform = "twitch"
	PLATFORM_TWITTER   Platform = "twitter"
)

type Pronoun string

// Nominative Pronouns
const (
	PR_HE   Pronoun = "he"
	PR_IT   Pronoun = "it"
	PR_SHE  Pronoun = "she"
	PR_THEY Pronoun = "they"
)

// Meta Pronouns
const (
	PR_ANY   Pronoun = "any"
	PR_ASK   Pronoun = "ask"
	PR_AVOID Pronoun = "avoid"
	PR_OTHER Pronoun = "other"
)

type RespRawBulk = map[string]*RespRawUser
type RespRawUser struct {
	Sets map[string][]string `json:"sets"`
}

// Looks up the data saved in PronounDB for one or more account for a given platform.
// The response is a map of IDs to the corresponding data. If an ID is not in our database, it will not be present in the response.
// It is strongly recommended to fetch IDs in bulk when possible and applicable, to help prevent hitting and potential rate limits.\
//
// Note - this package auto-paginates your IDs. pronoundb caps the ids at 50, we will do multiple requests if len(ids) > 50
// Note -
func (c *Client) RawLookupBulk(platform Platform, ids []string) (RespRawBulk, error) {
	if len(ids) == 0 {
		return RespRawBulk{}, nil
	}

	realResp := RespRawBulk{}

	for i := 0; i <= (len(ids)-1)/BULK_LOOKUP_ID_LIMIT; i++ {
		batch := ids[i*BULK_LOOKUP_ID_LIMIT:]
		if len(batch) > BULK_LOOKUP_ID_LIMIT {
			batch = batch[:BULK_LOOKUP_ID_LIMIT]
		}

		resp := RespRawBulk{}
		err := httpFetch(`GET`, `/api/v2/lookup?platform=`+string(platform)+`&ids=`+strings.Join(batch, ","), c, resp)
		if err != nil {
			return realResp, err
		}

		for k, v := range resp {
			realResp[k] = v
		}
	}

	return realResp, nil
}

// Bulk lookup pronouns for users given a platform. IDs are based on the platform you specify
// Note that while pronouns db only supports up 50 ids, you can specify as many as you want (we do multiple requests)
// If any error happens between fetches, then the already fetched IDs get returned, and no further requests are done
// If an ID is not present in the response (and no error happened in that fetch), then it means that ID is not saved
// Note that only 'en' locale is supported. If en set is not available, the user is skipped.
// Also - each user will have at least 1 pronoun
func (c *Client) Lookup(platform Platform, ids ...string) (map[string][]Pronoun, error) {
	raw, err := c.RawLookupBulk(platform, ids)

	resp := map[string][]Pronoun{}
	for u, p := range raw {
		if p == nil || len(p.Sets) == 0 || len(p.Sets["en"]) == 0 {
			continue
		}
		pronouns := p.Sets["en"]
		arr := make([]Pronoun, len(pronouns))
		for i, v := range pronouns {
			arr[i] = Pronoun(v)
		}

		resp[u] = arr
	}

	return resp, err
}
