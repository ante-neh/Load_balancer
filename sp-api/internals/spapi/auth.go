package spapi

import (
	"context"
	"net/http"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"golang.org/x/oauth2"
)


type LWAClient struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	Token        *oauth2.Token
}

func (l *LWAClient) GetAccessToken() (string, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     l.ClientID,
		ClientSecret: l.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://api.amazon.com/auth/o2/token",
		},
	}

	token, err := conf.TokenSource(ctx, &oauth2.Token{
		RefreshToken: l.RefreshToken,
		Expiry:       time.Now().Add(-24 * time.Hour), 
	}).Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func SignRequest(req *http.Request, region string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return err
	}

	credentials, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		return err
	}

	signer := v4.NewSigner()
	return signer.SignHTTP(context.TODO(), credentials, req, "execute-api", cfg.Region, time.Now())
}