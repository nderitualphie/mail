package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MoReport() {
	// Database credentials
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	// CSV file path
	// Get yesterday's date
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayStr := yesterday.Format("2006-01-02")
	outputFolder := os.Getenv("DIR") // Specify the desired folder path
	csvFile := filepath.Join(outputFolder, yesterdayStr+".csv")

	// Generate the date string for today
	orgID := 209

	// SQL query to retrieve data created today
	query := fmt.Sprintf("SELECT network,  src_address, DATE(created_on) AS DATE, count(*) AS total FROM ucm.tbl_campaign_messages WHERE flow = 'MO' AND org_id = '%v'  AND DATE(LEFT(created_on, 10)) = '%v'  and network <> 'Unknown' GROUP BY network,src_address, DATE(created_on)\nORDER BY network ASC;\n", orgID, yesterdayStr)

	// Connect to the database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName))
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}
	defer db.Close()
	// Execute the SQL query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}

	defer rows.Close()

	// Open the CSV file in write mode
	file, err := os.Create(csvFile)
	if err != nil {
		log.Fatal("Error creating CSV file: ", err)
	}
	defer file.Close()

	// Write the CSV header
	headers := []string{"Network", "Date", "src_address", "Total"} // Replace with actual column names
	writer := csv.NewWriter(file)
	writer.Write(headers)

	// Write the query results to the CSV file
	for rows.Next() {
		var network, srcAddress, createdOn string
		err := rows.Scan(&network, &srcAddress, &createdOn)
		if err != nil {
			log.Fatal("Error scanning rows: ", err)
		}

		row := []string{network, srcAddress, createdOn}
		writer.Write(row)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal("Error writing CSV: ", err)
	}

	log.Println("CSV file created successfully!")
}
