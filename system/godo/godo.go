// Ties in the digital ocean API

package godo

import (
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type GoDO struct {
	Client *godo.Client
	Token  *TokenSource
}

// Create a new digital ocean client
func New(t string) (*GoDO, error) {
	var err error

	g := new(GoDO)

	g.Token, err = NewToken(t)

	if err != nil {
		return nil, err
	}

	return g, nil
}

// Connect to the DO client
func (g *GoDO) Connect() {
	oc := oauth2.NewClient(oauth2.NoContext, g.Token)
	g.Client = godo.NewClient(oc)
}

type TokenSource struct {
	AccessToken string
}

// Create a oauth token
func NewToken(t string) (*TokenSource, error) {
	return &TokenSource{AccessToken: t}, nil
}

// Access the oauth token
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
