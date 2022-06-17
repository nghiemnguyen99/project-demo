package entity

type Users struct {
	ID        int     `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	MSSV      *string `json:"mssv"`
	SubjectID *int    `json:"subject_id"`
}
