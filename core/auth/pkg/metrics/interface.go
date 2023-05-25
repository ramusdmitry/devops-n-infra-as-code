package metrics

import authApp "auth-app-service"

type MetricsCollector interface {
	Serve(config *authApp.Config) error
	RegisterUserHandler()
}
