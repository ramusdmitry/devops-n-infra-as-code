package probes

import (
	commsApp "comms-app-service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Probes struct {
	isAlive bool
	isReady bool
}

func NewProbes() *Probes {
	return &Probes{
		isReady: false,
		isAlive: false,
	}
}

func (p *Probes) InitProbes(config *commsApp.Config) error {

	var err, probErr error

	r := gin.Default()

	r.GET("/liveness",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
			p.isAlive = true
		})

	r.GET("/readiness", func(c *gin.Context) {
		dbErr := PostgresHealthCheck(config.DB)
		if dbErr != nil {
			logrus.Warnf("failed check connection to db with probe: %s", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"status": "err",
			})
			p.isReady = false
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
			p.isReady = true
		}
	})

	go func() {
		err := r.Run(fmt.Sprintf(":%d", config.Probes.Port))
		if err != nil {
			logrus.Errorf("start probes failed: %s", err.Error())
		}
		if !p.isAlive {
			logrus.Errorf("connection unavailable")
			probErr = errors.New("status check failed")
		}
		if !p.isReady {
			logrus.Errorf("db unavailable")
			probErr = errors.New("status check failed")
		}
	}()

	return probErr
}

func PostgresHealthCheck(cfg commsApp.DBConfig) error {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode))

	defer db.Close()

	if err != nil {
		logrus.Fatalf("[%s] [DB] failed to connect to db\n cause: %s", time.Now().UTC().Format("2006-01-02 15:04:05"), err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("[%s] [DB] failed to connect to db\n cause: %s", time.Now().UTC().Format("2006-01-02 15:04:05"), err.Error())
		return err
	}

	return nil
}
