package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

const ScheduleURL = "https://cfp.gulas.ch/gpn21/schedule/export/schedule.json"

func main() {
	schedule, err := retrieveJSONSchedule(ScheduleURL)
	if err != nil {
		log.Fatalln(err)
	}

	pretalxSchedule, err := parseSchedule(schedule)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("GPN Schedule Version: %v\n", pretalxSchedule.Schedule.Version)
	fmt.Printf("GPN Schedule Critical-Path (Non recorded talks):\n")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Watch?\tDay\tTime\tRoom\tTalk\tRecorded\n")

	for _, day := range pretalxSchedule.Schedule.Conference.Days {
		for roomName, talks := range day.Rooms {
			for _, talk := range talks {
				if talk.DoNotRecord {
					fmt.Fprintf(w, "[ ]\t%v\t%v\t%v\t%v\t%v\n", day.Date, talk.Date.Format(time.TimeOnly), roomName, talk.Title, !talk.DoNotRecord)
				}
			}
		}
	}

	err = w.Flush()
	if err != nil {
		log.Fatalln(err)
	}
}

func parseSchedule(schedule []byte) (pretalxSchedule *PretalxSchedule, err error) {
	err = json.Unmarshal(schedule, &pretalxSchedule)
	return
}

func retrieveJSONSchedule(url string) (rawJSON []byte, err error) {
	var response *http.Response

	response, err = http.Get(url)
	if err != nil {
		return
	}

	rawJSON, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}
