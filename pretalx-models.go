package main

import "time"

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
