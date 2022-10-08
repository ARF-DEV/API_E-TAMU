package repository

import (
	"E-TamuAPI/models"
	"database/sql"
)

type VisitRepository struct {
	db *sql.DB
}

func NewVisitRepository(db *sql.DB) *VisitRepository {
	return &VisitRepository{
		db: db,
	}
}

func (v *VisitRepository) CreateVisit(visit models.Visit) (*models.Visit, error) {
	sqlStatement := `
	insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_feedback, guest_count, visit_hour, transportation, visit_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_feedback, guest_count, visit_hour, transportation, visit_date;`

	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visit.VisitId, visit.UserVisitedId, visit.GuestName, visit.GuestEmail, visit.VisitIntention, visit.VaccineCertificate, visit.VisitStatus, visit.GuestFeedback, visit.GuestCount, visit.VisitHour, visit.Transportation, visit.VisitDate).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (v *VisitRepository) GetAllVisit() ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit;`

	rows, err := v.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) GetVisitByID(visitId int) (*models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE visit_id = ?;`

	var result models.Visit

	err := v.db.QueryRow(sqlStatement, visitId).Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (v *VisitRepository) GetVisitByStaffID(staffId int) ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE user_visited_id = ?;`

	rows, err := v.db.Query(sqlStatement, staffId)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		// err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) GetVisitByStatus(status int) ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE visit_status = ?;`

	rows, err := v.db.Query(sqlStatement, status)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		// err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) GetVisitByDate(startDate string, endDate string) ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE visit_status BETWEEN ? and ?;`

	rows, err := v.db.Query(sqlStatement, startDate, endDate)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) DeleteVisit(visitId int) (*models.Visit, error) {
	sqlStatement := `
	DELETE FROM visit 
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_feedback, guest_count, visit_hour, transportation, visit_date;`
	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
