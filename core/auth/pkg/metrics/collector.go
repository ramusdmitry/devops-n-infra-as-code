package metrics

import (
	authApp "auth-app-service"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

const (
	Namespace       = "business_metrics"
	RegisterUser    = "/register" // Путь для регистрации пользователей
	MetricsEndpoint = "/metrics"
)

type Metrics struct {
	registeredUsers prometheus.Counter // Метрика для подсчета успешно зарегистрированных пользователей
	mu              sync.Mutex
}

func NewMetrics() *Metrics {
	m := &Metrics{
		registeredUsers: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "registered_users",
			Help:      "The number of registered users",
		}),
	}
	prometheus.MustRegister(m.registeredUsers)
	return m
}

func (m *Metrics) Serve(config *authApp.Config) error {
	http.Handle(MetricsEndpoint, promhttp.Handler())
	http.Handle(RegisterUser, http.HandlerFunc(m.RegisterUserHandlerHTTP))
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Metrics.Port), nil)
}

func writeResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
	_, _ = w.Write([]byte("\n"))
}

func (m *Metrics) RegisterUserHandlerHTTP(w http.ResponseWriter, r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.registeredUsers.Inc()
	writeResponse(w, http.StatusOK, "ok")
}

func (m *Metrics) RegisterUserHandler() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.registeredUsers.Inc()
}
