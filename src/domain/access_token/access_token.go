package access_token

import "time"

const (
	expirationTime = 24
)

// AccessToken ...
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (a *AccessToken) IsExpired() bool {
	return time.Unix(a.Expires, 0).Before(time.Now().UTC())
}
