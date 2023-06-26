package sendMail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"path/filepath"
)

// Mail sends an email with an attached file
func Mail() {
	// Sender data
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email addresses
	to := []string{
		"mercymaina567@gmail.com",
		"aliphonzanderitu@gmail.com",
	}

	// Create a new email message
	subject := "Daily MO Report"
	message := "Find attached MO report."

	// Create a new email client
	email := gomail.NewMessage()
	email.SetHeader("From", from)
	email.SetHeader("to", to...)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", message)

	// Find the latest CSV file in the folder
	files, err := filepath.Glob(os.Getenv("DIR")) // Update with the correct folder path
	if err != nil {
		fmt.Printf("Error finding CSV files: %s\n", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No CSV files found in the folder.")
		return
	}

	latestFile := getLatestFile(files)
	if latestFile == "" {
		fmt.Println("Unable to determine the latest CSV file.")
		return
	}

	// Attach the CSV file
	email.Attach(latestFile)

	// Create a new email client with SMTP configuration
	email.Attach(latestFile)

	// Create a new email client with SMTP configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	// Send the email
	err = d.DialAndSend(email)
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err)
		return
	}

	fmt.Println("Email Sent!")
}

// getLatestFile returns the latest file from the list of file paths
func getLatestFile(files []string) string {
	var latestTime int64
	var latestFile string

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			continue
		}

		modTime := fileInfo.ModTime().Unix()
		if modTime > latestTime {
			latestTime = modTime
			latestFile = file
		}
	}

	return latestFile
}
