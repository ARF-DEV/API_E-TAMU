package models

type Visit struct {
	VisitId            int    `db:"visit_id"`
	UserVisitedId      int    `db:"user_visited_id"`
	GuestName          string `db:"guest_name"`
	GuestEmail         string `db:"guest_email"`
	VisitIntention     string `db:"visit_intention"`
	VaccineCertificate string `db:"vaccine_certificate"`
	VisitStatus        string `db:"visit_is_done"`
	GuestFeedback      string `db:"guest_feedback"`
	GuestCount         int    `db:"guest_count"`
	VisitDate          string `db:"visit_date"`
	VisitHour          string `db:"visit_hour"`
	Transportation     string `db:"transpotation"`
}
