package repository

import (
	"E-TamuAPI/models"
	"database/sql"
	"fmt"
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
	fmt.Println(visit)
	sqlStatement := `
	insert into visit (user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_hour, transportation, visit_date, confirmation) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, '')
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status, guest_count, visit_hour, transportation, visit_date, confirmation;`

	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visit.UserVisitedId, visit.GuestName, visit.GuestEmail, visit.VisitIntention, visit.VaccineCertificate, visit.VisitStatus, visit.GuestCount, visit.VisitHour, visit.Transportation, visit.VisitDate).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (v *VisitRepository) GetAllVisit() ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit ORDER BY visit_date DESC;`

	rows, err := v.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.Confirmation)

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

	err := v.db.QueryRow(sqlStatement, visitId).Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.Confirmation)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (v *VisitRepository) GetVisitByStaffID(staffId int) ([]models.Visit, error) {
	fmt.Println(staffId)
	sqlStatement := `SELECT * FROM visit WHERE user_visited_id = ? ORDER BY visit_date DESC;`

	rows, err := v.db.Query(sqlStatement, staffId)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit
	fmt.Println(rows)

	for rows.Next() {
		var result models.Visit
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.Confirmation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) GetVisitByStatus(status string) ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE visit_status LIKE ? ORDER BY visit_date DESC;`

	rows, err := v.db.Query(sqlStatement, status)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		// err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestFeedback, &result.GuestCount, &result.VisitHour, &result.Transportation)
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.Confirmation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) GetVisitByDate(startDate string, endDate string) ([]models.Visit, error) {
	sqlStatement := `SELECT * FROM visit WHERE visit_date BETWEEN ? and ? ORDER BY visit_date DESC;`

	rows, err := v.db.Query(sqlStatement, startDate, endDate)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit

	for rows.Next() {
		var result models.Visit
		err = rows.Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitDate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.Confirmation)

		if err != nil {
			return nil, err
		}
		visits = append(visits, result)
	}
	return visits, nil
}

func (v *VisitRepository) ConfirmFinish(visitId int) (*models.Visit, error) {
	sqlStatement := `
	UPDATE visit
	SET
		visit_status = 'selesai'
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status,guest_count, visit_hour, transportation, visit_date , confirmation;`

	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (v *VisitRepository) ConfirmArrival(visitId int) (*models.Visit, error) {
	sqlStatement := `
	UPDATE visit
	SET
		visit_status = 'sedang berlangsung'
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status,guest_count, visit_hour, transportation, visit_date, confirmation;`

	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VisitRepository) DeleteVisit(visitId int) (*models.Visit, error) {
	sqlStatement := `
	DELETE FROM visit 
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status,guest_count, visit_hour, transportation, visit_date, confirmation;`
	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VisitRepository) ConfirmVisitProposal(visitId int) (*models.Visit, error) {
	sqlStatement := `
	UPDATE visit
	SET
		confirmation = '1'
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status,guest_count, visit_hour, transportation, visit_date, confirmation;`

	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (v *VisitRepository) CancelVisitProposal(visitId int) (*models.Visit, error) {
	fmt.Println("EHHHLL")
	sqlStatement := `
	UPDATE visit
	SET
		confirmation = '0'
	WHERE visit_id = ?
	RETURNING visit_id, user_visited_id, guest_name, guest_email, visit_intention, vaccine_certificate, visit_status,guest_count, visit_hour, transportation, visit_date, confirmation;`
	var result models.Visit
	err := v.db.QueryRow(sqlStatement, visitId).
		Scan(&result.VisitId, &result.UserVisitedId, &result.GuestName, &result.GuestEmail, &result.VisitIntention, &result.VaccineCertificate, &result.VisitStatus, &result.GuestCount, &result.VisitHour, &result.Transportation, &result.VisitDate, &result.Confirmation)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
