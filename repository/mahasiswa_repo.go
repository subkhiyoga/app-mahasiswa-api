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
	Create(newMahasiswa *model.Mahasiswa) string
	Update(mahasiswa *model.Mahasiswa) string
	Delete(id int) string
}

type mahasiswaRepo struct {
	db *sql.DB
}

func (r *mahasiswaRepo) GetAll() any {
	var msiswa []model.Mahasiswa

	query := "SELECT m.name, m.age, m.major, c.username FROM mahasiswa m JOIN credentials c ON m.user_name = c.username"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}
	if rows == nil {
		return "no data"
	}
	defer rows.Close()

	for rows.Next() {
		var mahasiswa model.Mahasiswa

		err := rows.Scan(&mahasiswa.Name, &mahasiswa.Age, &mahasiswa.Major, &mahasiswa.UserName)
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

func (r *mahasiswaRepo) Create(newMahasiswa *model.Mahasiswa) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "failed to create data"
	}

	// insert data to credentials table
	query2 := "INSERT INTO credentials (username, password) VALUES ($1, $2) RETURNING username"
	_, err = tx.Exec(query2, newMahasiswa.Username, newMahasiswa.Password)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create data"
	}

	// insert data to mahasiswa table
	query1 := "INSERT INTO mahasiswa (name, age, major, user_name) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(query1, newMahasiswa.Name, newMahasiswa.Age, newMahasiswa.Major, newMahasiswa.UserName)

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

	// jika id tidak ada, return message
	if result == "data not found" {
		return result.(string)
	}

	// menggunakan transaksi agar jika data salah bisa di rollback
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "failed to update data"
	}

	// update credentials table
	query1 := "UPDATE credentials SET password = $1 WHERE username = $2"
	_, err = tx.Exec(query1, mahasiswa.Password, mahasiswa.Username)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to update data"
	}

	// update mahasiswa table
	query2 := "UPDATE mahasiswa SET name = $1, age = $2, major = $3 WHERE id = $4"
	_, err = tx.Exec(query2, mahasiswa.Name, mahasiswa.Age, mahasiswa.Major, mahasiswa.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to update data"
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "failed to update data"
	}

	// return success message
	return fmt.Sprintf("Data with id %d updated successfully", mahasiswa.ID)
}

func (r *mahasiswaRepo) Delete(id int) string {
	// mengambil user_name mahasiswa dengan id yang diberikan
	var userName string
	query1 := "SELECT user_name FROM mahasiswa WHERE id = %1"
	err := r.db.QueryRow(query1, id).Scan(&userName)
	if err == sql.ErrNoRows {
		return "data not found"
	} else if err != nil {
		log.Println(err)
		return "failed to delete data"
	}

	// delete data by id
	query2 := "DELETE FROM mahasiswa WHERE id = $1"
	_, err = r.db.Exec(query2, id)
	if err != nil {
		log.Println(err)
		return "failed to delete data"
	}

	// delete data in credentials table with user_name value
	query3 := "DELETE FROM credentials WHERE username = $1"
	_, err = r.db.Exec(query3, userName)
	if err != nil {
		log.Println(err)
		return "failed to delete mahasiswa's credentials"
	}

	return fmt.Sprintf("Data with id %d and credential deleted successfully", id)
}

func NewMahasiswaRepo(db *sql.DB) MahasiswaRepo {
	repo := new(mahasiswaRepo)
	repo.db = db

	return repo
}
