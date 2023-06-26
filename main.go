package main

import (
	x "mo/db"
	m "mo/sendMail"
)

func main() {
	x.MoReport()
	m.Mail()
}
