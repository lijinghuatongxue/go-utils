package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestForFileLines(t *testing.T) {
	var err error
	var array []string
	array, err = utils.ForFileLines("./../README.md", false)
	if err != nil {
		logrus.Error(err)
	}
	for _, value := range array {
		logrus.Info(value)
	}
}
func PrintTestForFileLines(str string) {
	logrus.Warn(str)
}
func TestForFileLinesFunc(t *testing.T) {
	utils.ForFileLinesExecFunc("./data/tpl.txt", PrintTestForFileLines)
}
