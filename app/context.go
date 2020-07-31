package app

import (
	"github.com/sirupsen/logrus"

	"leggett.dev/devmarks/api/db"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}
