package auth

type RedirectUrl = struct {
	RedirectUrl		string	`json:"redirect_url"`
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)