package ca

import (
	"testing"
)

//
// TestResolveUploadKey
//  @Description:
//  @param t
//
func TestResolveUploadKey(t *testing.T) {
	key := "-----BEGIN EC PRIVATE KEY-----\n" +
		"MHcCAQEEIGTw58Huh6IYtNq49mpzxO2V/HTToAxw9GAEy116HyaCoAoGCCqBHM9V\n" +
		"AYItoUQDQgAEWsluIWKiyR0f0rL1ZOpR2MRl2bNRTbMb0m6gKo2oq+9MpAOykQxj\n" +
		"9MDmy/ZwTwbWFyyFBAzhFRX/Lxflmht3dA==\n" +
		"-----END EC PRIVATE KEY-----"
	ResolveUploadKey(key)
}
