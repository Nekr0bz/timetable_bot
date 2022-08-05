package entity

type University struct {
	Id         int64       `json:"id"`
	Name       string      `json:"name"`
	GroupTypes []GroupType `json:"group_types"`
}

type GroupType struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Faculties []Faculty `json:"faculties"`
}

type Faculty struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Courses []Course `json:"curses"`
}

type Course struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
