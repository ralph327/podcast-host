// Implements the Minio Client

package minio

import (
	"github.com/minio/minio-go"
	"log"
	"net/url"
)

type Minio struct {
	Client *minio.Client
	Conf   *Conf
}

func New(c *Conf) (*Minio, error) {
	client, err := minio.New(c.Endpoint,
		c.AccessKeyID,
		c.SecretAccessKey,
		isSecure(c.Endpoint),
	)
	if err != nil {
		return nil, err
	}

	conf := c

	m := &Minio{Client: client, Conf: conf}

	return m, nil
}

type Conf struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewConf(url string, accessKey string, secretKey string) (*Conf, error) {
	c := &Conf{
		Endpoint:        findHost(url),
		AccessKeyID:     accessKey,
		SecretAccessKey: secretKey,
		UseSSL:          isSecure(url),
	}

	return c, nil
}

// Find the Host of the given url.
func findHost(urlStr string) string {
	u, err := url.Parse("//" + urlStr)
	if err != nil {
		panic(err)
	}
	return u.Host
}

// Finds out whether the url is http(insecure) or https(secure).
func isSecure(urlStr string) bool {
	u, err := url.Parse("//" + urlStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return u.Scheme == "https"
}
