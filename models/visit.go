package models

type Visit struct {
	VisitId            int    `db:"visit_id" json:"visit_id"`
	UserVisitedId      int    `db:"user_visited_id" json:"user_visited_id"`
	GuestName          string `db:"guest_name" json:"guest_name"`
	GuestEmail         string `db:"guest_email" json:"guest_email"`
	VisitIntention     string `db:"visit_intention" json:"visit_intention"`
	VaccineCertificate string `db:"vaccine_certificate" json:"vaccine_certificate"`
	VisitStatus        string `db:"visit_status" json:"visit_status"`
	GuestCount         int    `db:"guest_count" json:"guest_count"`
	VisitDate          string `db:"visit_date" json:"visit_date"`
	VisitHour          string `db:"visit_hour" json:"visit_hour"`
	Transportation     string `db:"transportation" json:"transportation"`
}
