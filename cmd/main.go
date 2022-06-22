package main

import (
	"github.com/glebnaz/go-platform/metrics"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

func main() {
	bot, err := tg.NewBotAPI("5499171458:AAHGbmFUTjdMIm4RvtI4NO9Jjz_E_KRd1CU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	e := echo.New()

	e.GET("/metrics", echo.WrapHandler(metrics.Handler()))

	go func() {
		err := e.Start(":8084")
		if err != nil {
			return
		}

	}()

	c := controller{bot}

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			start := time.Now()
			c.handleMSG(update.Message)
			end := time.Now()
			timeHandleMetric.WithLabelValues().Observe(end.Sub(start).Seconds())
		}
	}
}
