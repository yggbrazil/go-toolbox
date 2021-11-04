// Package log in as super set of the standart golang "log" with somes new methods
package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var log *Log
var lock = &sync.Mutex{}

type Log struct {
	*logrus.Logger
	debug bool
}

func New() *Log {
	lock.Lock()
	defer lock.Unlock()

	if log == nil {
		l := logrus.New()
		l.SetFormatter(&logrus.JSONFormatter{})

		log = &Log{
			Logger: l,
		}
	}

	return log
}

// EnableDebug active the debug mode
func (l *Log) EnableDebug(enable bool) {
	l.debug = enable
}

// Error print in terminal the error message and the line of code with de error
func (l *Log) Error(e error) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	hostname, err := os.Hostname()

	if err != nil {
		log.Println(err)
	}

	file_path := fmt.Sprintf("%s:%d", file, line)

	l.WithField("file_path", file_path).
		WithField("hostname", hostname).
		Error(e.Error())
}

// Debug is a function for print in terminal if the variable debug it's true
func (l *Log) Debug(a ...interface{}) {
	if l.debug {
		log.Println(a...)
	}
}

// Debugf is a function for print formated in terminal if the variable debug it's true
func (l *Log) Debugf(format string, v ...interface{}) {
	if l.debug {
		log.Printf(format, v...)
	}
}

// File save ou create a new log file with errors
func (l *Log) File(file, text string) error {
	if len(file) < 1 {
		return errors.New("invalid file name")
	}

	path := "logs/"
	if file[0:1] == "/" {
		path += file[1:]
	} else {
		path += file
	}

	pathSplited := strings.Split(path, "/")
	folders := strings.Join(pathSplited[0:len(pathSplited)-1], "/")

	_, err := os.Stat(folders)

	if err != nil {

		err = os.MkdirAll(folders, os.ModePerm)

		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		return err
	}

	f.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + text + "\n")
	f.Close()
	return nil
}

func createFileIfNotExists(filePath string) (f *os.File) {
	dirName := filepath.Dir(filePath)

	err := os.MkdirAll(dirName, os.ModePerm)

	if err != nil {
		log.Fatalf("Error create path: %v", err)
	}

	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("Error create file: %v", err)
	}

	return
}

//SetOutputFiles set a new path for logs
func (l *Log) SetOutputFiles(outfilePath, errfilePath string) {
	outFile := createFileIfNotExists(outfilePath)
	errFile := createFileIfNotExists(errfilePath)

	defer outFile.Close()
	defer errFile.Close()

	syscall.Dup2(int(outFile.Fd()), 1) /* -- stdout */
	syscall.Dup2(int(errFile.Fd()), 2) /* -- stderr */
}
