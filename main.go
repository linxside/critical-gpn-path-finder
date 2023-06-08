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

	printTalkTable(pretalxSchedule, err)
}

func printTalkTable(pretalxSchedule *PretalxSchedule, err error) {
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

type PretalxSchedule struct {
	Schedule struct {
		Version    string `json:"version"`
		BaseUrl    string `json:"base_url"`
		Conference struct {
			Acronym          string `json:"acronym"`
			Title            string `json:"title"`
			Start            string `json:"start"`
			End              string `json:"end"`
			DaysCount        int    `json:"daysCount"`
			TimeslotDuration string `json:"timeslot_duration"`
			TimeZoneName     string `json:"time_zone_name"`
			Rooms            []Room `json:"rooms"`
			Days             []Day  `json:"days"`
		} `json:"conference"`
	}
}

type Room struct {
	Name        string  `json:"name"`
	Guid        *string `json:"guid"`
	Description *string `json:"description"`
	Capacity    *int    `json:"capacity"`
}

type Day struct {
	Index    int       `json:"index"`
	Date     string    `json:"date"`
	DayStart time.Time `json:"day_start"`
	DayEnd   time.Time `json:"day_end"`
	Rooms    RoomTalks `json:"rooms"`
}

type RoomTalks map[string][]Talk

type Talk struct {
	Id               int       `json:"id"`
	Guid             string    `json:"guid"`
	Logo             string    `json:"logo"`
	Date             time.Time `json:"date"`
	Start            string    `json:"start"`
	Duration         string    `json:"duration"`
	Room             string    `json:"room"`
	Slug             string    `json:"slug"`
	Url              string    `json:"url"`
	Title            string    `json:"title"`
	Subtitle         string    `json:"subtitle"`
	Track            string    `json:"track"`
	Type             string    `json:"type"`
	Language         string    `json:"language"`
	Abstract         string    `json:"abstract"`
	Description      string    `json:"description"`
	RecordingLicense string    `json:"recording_license"`
	DoNotRecord      bool      `json:"do_not_record"`
	Persons          []Person  `json:"persons"`
	Links            []string  `json:"links"`
	Attachments      []string  `json:"attachments"`
	Answers          []string  `json:"answers"`
}

type Person struct {
	Id         int      `json:"id"`
	Code       string   `json:"code"`
	PublicName string   `json:"public_name"`
	Biography  string   `json:"biography"`
	Answers    []string `json:"answers"`
}
