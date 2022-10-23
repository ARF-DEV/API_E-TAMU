package models

type Visit struct {
	VisitId            int    `db:"visit_id" json:"visit_id" validate:"required"`
	UserVisitedId      int    `db:"user_visited_id" json:"user_visited_id" validate:"required"`
	GuestName          string `db:"guest_name" json:"guest_name" validate:"required"`
	GuestEmail         string `db:"guest_email" json:"guest_email" validate:"required"`
	VisitIntention     string `db:"visit_intention" json:"visit_intention" validate:"required"`
	VaccineCertificate string `db:"vaccine_certificate" json:"vaccine_certificate" validate:"required"`
	VisitStatus        string `db:"visit_status" json:"visit_status" validate:"required"`
	GuestCount         int    `db:"guest_count" json:"guest_count" validate:"required"`
	VisitDate          string `db:"visit_date" json:"visit_date" validate:"required"`
	VisitHour          string `db:"visit_hour" json:"visit_hour" validate:"required"`
	Transportation     string `db:"transportation" json:"transportation" validate:"required"`
}
