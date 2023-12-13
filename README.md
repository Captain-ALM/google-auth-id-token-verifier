# google-auth-id-token-verifier

Golang port of [OAuth2Client.prototype.verifyIdToken](https://github.com/google/google-auth-library-nodejs/blob/master/lib/auth/oauth2client.js) from [google-auth-library-nodejs](https://github.com/google/google-auth-library-nodejs)

Verify idtoken without making http request to tokeninfo API.

## Usage

```go

import (
    "golang.local/google-auth-id-token-verifier"
)

v := googleAuthIDTokenVerifier.Verifier{}
aud := "xxxxxx-yyyyyyy.apps.googleusercontent.com"
claimSet ,err := v.ClaimIDToken(TOKEN, []string{
    aud,
})
if err == nil {
    // claimSet.Email ... (See claimset.go)
}
```

## Features

- Fetch public key from www.googleapis.com/oauth2/v3/certs
- Respect cache-control in response from www.googleapis.com/oauth2/v3/certs
- JWT Parser
- Check Signature 
- Check IssueTime, ExpirationTime with ClockSkew
- Check Issuer
- Check Audience

## See also

- http://stackoverflow.com/questions/36716117/validating-google-sign-in-id-token-in-go#
- https://github.com/GoogleIdTokenVerifier/GoogleIdTokenVerifier

## Patches applied by Captain ALM
