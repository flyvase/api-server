package gateway

import (
	"api-server/src/config"
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() *firebase.App {
	if config.Mode == "debug" && config.Environment == "dev" {
		opt := option.WithCredentialsJSON([]byte(`{
			"type": "service_account",
			"project_id": "flyvase-dev",
			"private_key_id": "ae8a8cedb78e227cd873d9bb7add3485dfd995be",
			"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDwH7xrdqVkYNVE\nEwhGO0Bfck8BZeLuGxkioB3kVbSfd7Js1PAdcMHZoZmRtwoE8Umsz+IsMDHA/ENg\nBs3I8+NQ6HgJHVyTc4jwDIFGnTckAM+SGrHjEnB7l9Y44gCCBafu+WqSgA3IU6ZB\nnrNom0PCHepaDVnJvUCRqFZ/pNzjBwSj+KshGP0VgL0C07Vm3icTjwUSvmt3APNT\nEvKc0Epq7UggyK9F5fdKtGZU2BNzjwx/KWtJiQva65EECP4k1lR/6zOTtg2qO63B\njFa6kFtI4luPKoBMeEfBOvPAp3i/8PNUN6Obd3CldUjjfHAZBz5LkeIxxzf+Nbmb\nqYy3eTVJAgMBAAECggEAB9lduRyMfPq+VeaSXWtX6BbONtM8BwD7Qhq+R2W2zQSm\nYDDYtXmnWIQUB0OWsAEsIZo9AClGIxML58JpLSGtw0oKDHA5g7dh2UjrW+HNON/O\n6S12tZ+3bg7RSZQ3pIhnTCQ6g9XsPWlglKIJpsh2f0ynibCIiBPn3asFQtBiiQg0\nPJgjBZWEjYfyi0g1G/QaOJHvYqJxAT6AlXHvWHjBqoC0XtRO2tJVpRZGt7KGP9/6\nS2dSkrAMHn9Fpi9/7c1QiZv9E5eeU3vo7xz1UayfUz1qqY318QbIw36OwViHMwIZ\nLCh+P+URriD7fyyKipD9FFj8AxQ9eH81xyWWdsAXZQKBgQD6vu1R95yTms7Ky12c\ny6BQ4S/bHozsFdm2ky3qjO64JVPvJGP9UyhUgttSka9PoW77MrivnKVYpMaOs+Ii\nd6TWWODVtZSJ/rfW90zswp5D0zMFPUTNGtLQqJq3EGpzeF2pokmIC1KDfShhObBF\nxeKqyI9CaqoUfH7qTMEIFeOPLQKBgQD1J9SSw8lLVFXcrds04QErWaONqzMBKi6j\nFAVSV7+d5Gaa20U279H8Hd4t5/8JJKSiFnEndq58U1h3ZhebrYXRz9ybCtSnxto3\nww+IPv7oe4qaOsJbsPJ6f6B54l1mXpLRcmVSCe/b9KB7BMNERf6HGnCNMBhDWuoh\n3s2Ot+ywDQKBgQCyhp+nnXlwqmcTginbiitRipMtGppRONlyoWWuurr4hzM0cB90\n7PEpAqYvKqS8OH8xAJL4Dbq321G3pqGZ0r+dEVi7L/ZbLe0sc5TkeUh/l+6ai13h\n/ngcsf692kVAg3GEZpSd9RnBnqnSV2WGt/bDi9pb+l+wCFfAb/d7z0He3QKBgHMK\nx+Q/YByrbMdsBmvgEo5nOCro5bxsMQpyALVNXCwZ3FRbBTXsgFuRIIvmTHXb7i3p\n3huow+Svr9oH3Jg/a4azxlAzWuuIkyqLrKRA+5nZL4eEf2RLut3lOkc4BKBfTQJg\nbk97PkW3m/9ekstdFT4NGDgJE7m31hsi1M9eXAGZAoGAJzZp9heN0iF3DAWRB5An\nn3+Ix/Ef0/6deGmPFfYthdUPUYo4mDDalnygJY7obKwifqkqrIctIZ+k6FVNjd1m\n6Egxe4R2m60WlnFiuw9FPOPDFfpprCU54CKwCXI6nGDk1SkEn9I8shhhkzHYZhA5\nZovAltt2e64o3VTbMpii39g=\n-----END PRIVATE KEY-----\n",
			"client_email": "api-server-firebase-admin@flyvase-dev.iam.gserviceaccount.com",
			"client_id": "107543502615991839732",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/api-server-firebase-admin%40flyvase-dev.iam.gserviceaccount.com"
		}`))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(err)
		}

		return app
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return app
}

func InitializeFirebaseAuth(app *firebase.App) *auth.Client {
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return auth
}
