package logger

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

var General *GeneralLogger

type GeneralLogger struct {
	*logrus.Logger
}

type Entry struct {
	*logrus.Entry
}

func (l *GeneralLogger) Component(t string) *Entry {
	entry := l.WithField("component", t)
	return &Entry{entry}
}

func (l *GeneralLogger) Type(t string) *Entry {
	entry := l.WithField("type", t)
	return &Entry{entry}
}

func Init() {
	General = &GeneralLogger{
		Logger: logrus.New(),
	}

	General.SetFormatter(&TextFormatter{
		FieldsOrder: []string{"component", "type", "name"},
	})

	General.SetReportCaller(true)
}

func LogObject(name string, ob interface{}) {
	lg := General.Component("ObjectLogger").WithField("type", name)
	res, _ := json.MarshalIndent(ob, "", "  ")

	lg.Println(string(res))
}
