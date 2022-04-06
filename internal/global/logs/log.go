package logs

import (
	"gosource/internal/global"
	"log"
)

func Fatal(args ...any) {
	log.SetPrefix("[Fatal] ")
	log.Fatal(args...)
}

func Fatalf(fmt string, args ...any) {
	log.SetPrefix("[Fatal] ")
	m := global.SPRINT_F(fmt, args...)
	log.Fatal(m)
}

func Debug(args ...any) {
	if global.LOG_DEBUG {
		log.SetPrefix("[Debug] ")
		log.Println(args...)
	}
}

func Debugf(fmt string, args ...any) {
	if global.LOG_DEBUG {
		log.SetPrefix("[Debug] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Info(args ...any) {
	if global.VERBOSE {
		log.SetPrefix("[Info] ")
		log.Println(args...)
	}
}

func Infof(fmt string, args ...any) {
	if global.VERBOSE {
		log.SetPrefix("[Info] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Error(args ...any) {
	if global.LOG_ERRORS {
		log.SetPrefix("[Error] ")
		log.Println(args...)
	}
}

func Errorf(fmt string, args ...any) {
	if global.LOG_ERRORS {
		log.SetPrefix("[Error] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}

func Warn(args ...any) {
	if global.LOG_WARNINGS {
		log.SetPrefix("[Warn] ")
		log.Println(args...)
	}
}

func Warnf(fmt string, args ...any) {
	if global.LOG_WARNINGS {
		log.SetPrefix("[Warn] ")
		m := global.SPRINT_F(fmt, args...)
		log.Println(m)
	}
}
