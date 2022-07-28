package middleware

import (
	"encoding/json"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func Logger(log *zap.Logger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			log.Info(string(data))
			err := next(c)
			if err != nil {
				log.Error("Telebot error", zap.Error(err))
			}
			return err
		}
	}
}
