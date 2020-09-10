package config

type Configuration struct {
	Server     ServerConfiguration
	Database   DatabaseConfiguration
	Staticfile StaticfileConfiguration
	Gcloud     GcloudConfiguration
}
