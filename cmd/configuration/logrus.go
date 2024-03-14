package configuration

import "github.com/sirupsen/logrus"

func (cf ConfigurationSchema) NewLogrus() *logrus.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()
	return log
}
