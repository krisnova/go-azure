package arm

type Authentication struct {
	clientID       string
	clientSecret   string
	subscriptionID string
	tenantID       string
}

func NewAuthentication(clientId, clientSecret, subscriptionId, tenantId string) *Authentication {
	return &Authentication{
		clientID:       clientId,
		clientSecret:   clientSecret,
		subscriptionID: subscriptionId,
		tenantID:       tenantId,
	}
}
