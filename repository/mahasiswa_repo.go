package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/subkhiyoga/app-mahasiswa-api/model"
)

type MahasiswaRepo interface {
	GetAll() any
	GetById(id int) any
	Create(newMahasiswa *model.Mahasiswa, newCredential *model.Credential) string
	Update(mahasiswa *model.Mahasiswa) string
	Delete(id int) string
}

type mahasiswaRepo struct {
	db *sql.DB
}

func (r *mahasiswaRepo) GetAll() any {
	var msiswa []model.Mahasiswa

	query := "SELECT * FROM mahasiswa"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var mahasiswa model.Mahasiswa

		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Name, &mahasiswa.Age, &mahasiswa.Major, &mahasiswa.UserName)
		if err != nil {
			log.Println(err)
		}

		msiswa = append(msiswa, mahasiswa)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(msiswa) == 0 {
		return "no data"
	}

	return msiswa
}

func (r *mahasiswaRepo) GetById(id int) any {
	var mInDb model.Mahasiswa

	query := "SELECT id, name, age, major, user_name FROM mahasiswa WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&mInDb.ID, &mInDb.Name, &mInDb.Age, &mInDb.Major, &mInDb.UserName)

	if err != nil {
		log.Println(err)
	}

	if mInDb.ID == 0 {
		return "data not found"
	}

	return mInDb
}

func (r *mahasiswaRepo) Create(newMahasiswa *model.Mahasiswa, newCredential *model.Credential) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "failed to create data"
	}

	// insert data to mahasiswa table
	query1 := "INSERT INTO mahasiswa (name, age, major, user_name) VALUES ($1, $2, $3, $4)"
	_, err = r.db.Exec(query1, newMahasiswa.Name, newMahasiswa.Age, newMahasiswa.Major, newMahasiswa.UserName)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create data"
	}

	// insert data to credentials table
	query2 := "INSERT INTO credential (username, password) VALUES ($1, $2)"
	_, err = r.db.Exec(query2, newCredential.Username, newCredential.Password)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create data"
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "failed to create data"
	}

	return "data created successfully"
}

func (r *mahasiswaRepo) Update(mahasiswa *model.Mahasiswa) string {
	result := r.GetById(mahasiswa.ID)

	if result == "data not found" {
		return result.(string)
	}

	query := "UPDATE mahasiswa SET name = $1, age = $2, major = $3, user_name = $4 WHERE id = $5"
	_, err := r.db.Exec(query, mahasiswa.Name, mahasiswa.Age, mahasiswa.Major, mahasiswa.UserName, mahasiswa.ID)

	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("data with id %d updated succesfully", mahasiswa.ID)
}

func (r *mahasiswaRepo) Delete(id int) string {
	result := r.GetById(id)

	if result == "data not found" {
		return result.(string)
	}

	query := "DELETE FROM mahasiswa WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete data"
	}

	return fmt.Sprintf("data with id %d deleted successfully", id)
}

func NewMahasiswaRepo(db *sql.DB) MahasiswaRepo {
	repo := new(mahasiswaRepo)
	repo.db = db

	return repo
}
