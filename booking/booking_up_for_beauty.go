package main

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate
}

func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	timeToCheck, _ := time.Parse(layout, date)
	timeNow := time.Now()
	return timeToCheck.Before(timeNow)
}

func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
	timeToCheck, _ := time.Parse(layout, date)
    hour := timeToCheck.Hour()
    return hour >= 12 && hour < 18
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	
	timeFormatted := fmt.Sprintf("You have an appointment on %s, at %s.", 
		parsedDate.Format("Monday, January 2, 2006"), 
		parsedDate.Format("15:04"))

	return timeFormatted
}

func AnniversaryDate() time.Time {
	anniversaryMonth := 9 
	anniversaryDay := 15  
    
	currentYear := time.Now().UTC().Year()

	anniversaryDate := time.Date(currentYear, time.Month(anniversaryMonth), anniversaryDay, 0, 0, 0, 0, time.UTC)

	return anniversaryDate
}