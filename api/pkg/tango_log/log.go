package tango_log

import (
	"fmt"
	"log"
)

// const colorRed = "\033[31m"
const txtGrenn = "\033[32m"
const txtYellow = "\033[33m"
const txtReset = "\033[0m"
const txtBold = "\033[1m"
const msgBegin = "TANGO > "
const msgEnd = "\n"

func getMsg(msg string) string {
	return fmt.Sprintf("%s %v %s", txtReset, msg, msgEnd)

}

func Print(msg string) {
	log.Printf("%s%s %s", txtBold, msgBegin, getMsg(msg))

}

func PrintOk(msg string) {
	log.Printf("%s%s%s %s%s", txtBold, txtGrenn, msgBegin, txtGrenn, getMsg(msg))

}

func PrintWarning(msg string) {
	log.Printf("%s%s%s %s", txtBold, txtYellow, msgBegin, getMsg(msg))

}
