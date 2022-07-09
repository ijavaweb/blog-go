package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
)

var (
	InfoLogger *logrus.Logger
	ErrorLogger *logrus.Logger
)

func init() {
	InfoLogger=logrus.New()
	ErrorLogger = logrus.New()
	errFile, err := os.OpenFile("../log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	file, err := os.OpenFile("../log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	InfoLogger.SetOutput(file)
	InfoLogger.SetLevel(4)
	InfoLogger.SetReportCaller(true)
	InfoLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		TimestampFormat: "2006-01-02 15:03:04",
	})

	ErrorLogger.SetOutput(errFile)
	ErrorLogger.SetLevel(2)
	ErrorLogger.SetReportCaller(true)
	ErrorLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		TimestampFormat: "2006-01-02 15:03:04",
	})
}
func Log(f gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx:=context
		f(ctx)
		bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
		go loggingQuery(bodyBytes, ctx.Request.RequestURI)
	}
}
func loggingQuery(bodyBytes []byte, api string) {
	InfoLogger.WithField("api",api).WithField("param",string(bodyBytes)).Info("request")
}
