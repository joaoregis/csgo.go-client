package logs

import (
	"gosource/internal/global"
	"log"
)

func Fatal(fmt string, args ...interface{}) {
	log.SetPrefix("[Fatal] ")
	m := global.SPRINT_F(fmt, args...)
	log.Fatal(m)
}

func Debug(fmt string, args ...interface{}) {
	if global.LOG_DEBUG {
		log.SetPrefix("[Debug] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Info(fmt string, args ...interface{}) {
	if global.VERBOSE {
		log.SetPrefix("[Info] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Error(fmt string, args ...interface{}) {
	if global.LOG_ERRORS {
		log.SetPrefix("[Error] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Warn(fmt string, args ...interface{}) {
	if global.LOG_WARNINGS {
		log.SetPrefix("[Warn] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}
