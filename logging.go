package main

import (
	"github.com/withmandala/go-log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr).WithColor()
}
