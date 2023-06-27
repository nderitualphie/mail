package sendMail

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

// Mail sends an email with an attached file
func Mail() {
	// Sender data
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email addresses
	to := strings.Split(os.Getenv("EMAIL_RECIPIENTS"), ",")

	// Create a new email message
	subject := "Daily MO Report"
	message := "Find attached MO report."

	// Create a new email client
	email := gomail.NewMessage()
	email.SetHeader("From", from)
	email.SetHeader("To", to...)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", message)

	// Find the latest CSV file in the folder

	// Get yesterday's date
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayStr := yesterday.Format("2006-01-02")
	csvFile := filepath.Join(os.Getenv("DIR"), fmt.Sprintf("%s.csv", yesterdayStr))

	// Attach the CSV file
	email.Attach(csvFile)

	// Create a new email client with SMTP configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	// Send the email
	err := d.DialAndSend(email)
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err)
		return
	}

	fmt.Println("Email Sent!")
}
