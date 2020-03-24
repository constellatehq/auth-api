package auth

type RedirectUrl = struct {
	Url		string	`json:"url"`
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)