package service

import (
	"nunu-rcc-stack/internal/repository"
	"nunu-rcc-stack/pkg/jwt"
	"nunu-rcc-stack/pkg/log"
	"nunu-rcc-stack/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
