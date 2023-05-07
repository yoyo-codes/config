package config

// Gcp holds a GCP project configurations.
type Gcp struct {
	ProjectID   string `env:"GOOGLE_PROJECT_ID"`
	Credentials string `env:"GOOGLE_APPLICATION_CREDENTIALS"`

	// GCP logging client.
	LoggingProjectID          string `env:"LOGGING_PROJECT_ID,default=$GOOGLE_PROJECT_ID"`
	LoggingProjectCredentials string `env:"LOGGING_APPLICATION_CREDENTIALS,default=$GOOGLE_APPLICATION_CREDENTIALS"`

	// GCP Identity Platform.
	IdpProjectID          string `env:"IDP_PROJECT_ID,default=$GOOGLE_PROJECT_ID"`
	IdpProjectCredentials string `env:"IDP_APPLICATION_CREDENTIALS,default=$GOOGLE_APPLICATION_CREDENTIALS"`
}
