package usecase

import (
	"github.com/subkhiyoga/app-mahasiswa-api/model"
	"github.com/subkhiyoga/app-mahasiswa-api/repository"
)

type MahasiswaUsecase interface {
	FindData() any
	FindDataById(id int) any
	Register(newMahasiswa *model.Mahasiswa, newCredential *model.Credential) string
	Edit(mahasiswa *model.Mahasiswa) string
	Unreg(id int) string
}

type mahasiswaUsecase struct {
	mahasiswaRepo repository.MahasiswaRepo
}

func (u *mahasiswaUsecase) FindData() any {
	return u.mahasiswaRepo.GetAll()
}

func (u *mahasiswaUsecase) FindDataById(id int) any {
	return u.mahasiswaRepo.GetById(id)
}

func (u *mahasiswaUsecase) Register(newMahasiswa *model.Mahasiswa, newCredential *model.Credential) string {
	return u.mahasiswaRepo.Create(newMahasiswa, newCredential)
}

func (u *mahasiswaUsecase) Edit(mahasiswa *model.Mahasiswa) string {
	return u.mahasiswaRepo.Update(mahasiswa)
}

func (u *mahasiswaUsecase) Unreg(id int) string {
	return u.mahasiswaRepo.Delete(id)
}

func NewMahasiswaUsecase(mahasiswaRepo repository.MahasiswaRepo) MahasiswaUsecase {
	return &mahasiswaUsecase{
		mahasiswaRepo: mahasiswaRepo,
	}
}
