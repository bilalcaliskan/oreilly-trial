package options

import (
	"github.com/spf13/pflag"
)

var oreillyTrialOptions = &OreillyTrialOptions{}

func init() {
	oreillyTrialOptions.addFlags(pflag.CommandLine)
	pflag.Parse()
}

// GetOreillyTrialOptions returns the pointer of OreillyTrialOptions
func GetOreillyTrialOptions() *OreillyTrialOptions {
	return oreillyTrialOptions
}

// OreillyTrialOptions contains frequent command line and application options.
type OreillyTrialOptions struct {
	// CreateUserUrl is the url of the user creation on Oreilly API
	CreateUserUrl string
	// EmailDomains is the comma seperated list of usable domain for creating trial account, it should be a valid domain
	EmailDomains []string
	// RandomLength is the length of the random generated username and password
	RandomLength int
}

func (oto *OreillyTrialOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&oto.CreateUserUrl, "createUserUrl", "https://learning.oreilly.com/api/v1/registration/individual/",
		"url of the user creation on Oreilly API")
	// for more usable domains, check https://temp-mail.org/
	fs.StringSliceVar(&oto.EmailDomains, "emailDomains", []string{"jentrix.com", "geekale.com", "64ge.com", "frnla.com"},
		"comma seperated list of usable domain for creating trial account, it should be a valid domain")
	fs.IntVar(&oto.RandomLength, "randomLength", 16, "length of the random generated username and password")
}
