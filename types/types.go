package types

type Tournament struct { // "STL Locals"
	name      string
	date      string
	events    []Event
	attendees []Attendee
}

type Event struct { // "Super Smash Doubles"
	name     string
	kind     string
	brackets []Bracket
}

type Bracket struct { // "Winners/Losers/Consolation"
	name    string
	kind    string
	matches []Match
}

type Match struct { // ""
	round string // "one, two, semifinal, final"
	teams []Team
}

type Attendee struct {
	name string
	seed int
	user *User
}

type User struct {
	name     string
	userName string
}

type Team struct {
	attendees []Attendee
	state     int
}
