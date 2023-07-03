package config

// Gcp holds a GCP project configurations.
type Gcp struct {
	ProjectID   string `env:"GCP_PROJECT_ID"`
	Credentials string `env:"GCP_APPLICATION_CREDENTIALS"`

	// GCP logging client.
	LoggingProjectID          string `env:"GCP_LOGGING_PROJECT_ID,default=$GCP_PROJECT_ID"`
	LoggingProjectCredentials string `env:"GCP_LOGGING_APPLICATION_CREDENTIALS,default=$GCP_APPLICATION_CREDENTIALS"`

	// GCP Identity Platform.
	IdpProjectID          string `env:"GCP_IDP_PROJECT_ID,default=$GCP_PROJECT_ID"`
	IdpProjectCredentials string `env:"GCP_IDP_APPLICATION_CREDENTIALS,default=$GCP_APPLICATION_CREDENTIALS"`
}
