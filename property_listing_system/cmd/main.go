package main

import (
	"fmt"
	"log"
	"property_listing_system/infrastructure/db"
)

func main() {
	fmt.Println("🚀 Starting Property Listing Backend")

	// initialize DB connection
	db.InitDB()

	
	log.Println("Server is running...")
}