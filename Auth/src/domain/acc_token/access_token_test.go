package acc_token

import "testing"

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Errorf("Brand new access token should not expired")
	}
}
