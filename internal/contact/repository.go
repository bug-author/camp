package contact

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(c *Contact) (int64, error) {
	query := `
		INSERT INTO contacts (fname, lname, email, phone)
		VALUES(?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, c.FirstName, c.LastName, c.Email, c.Phone)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (r *Repository) GetByEmail(email string) (Contact, error) {
	query := `
		SELECT
			id,
			fname,
			lname,
			email,
			phone,
			created_at,
			updated_at
		FROM contacts
		WHERE email = ?
		LIMIT 1
	`
	var result Contact

	err := r.db.QueryRow(query, email).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Phone,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Contact{}, nil
		}

		return Contact{}, err
	}

	return result, nil
}
