// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	ChangePassword(ctx context.Context, arg ChangePasswordParams) error
	CheckEmail(ctx context.Context, email string) (Mahasiswa, error)
	CheckUsername(ctx context.Context, username string) (Mahasiswa, error)
	CreateMahasiswa(ctx context.Context, arg CreateMahasiswaParams) (Mahasiswa, error)
	CreatePractice(ctx context.Context, arg CreatePracticeParams) (Practice, error)
	GetMahasiswa(ctx context.Context, arg GetMahasiswaParams) (Mahasiswa, error)
	HomeDashboard(ctx context.Context, idMahasiswa int32) (HomeDashboardRow, error)
	ListPractice(ctx context.Context, arg ListPracticeParams) ([]Practice, error)
	ListPracticeByCategory(ctx context.Context, arg ListPracticeByCategoryParams) ([]Practice, error)
	UpdateMahasiswa(ctx context.Context, arg UpdateMahasiswaParams) error
	UpdateOTPInDB(ctx context.Context, arg UpdateOTPInDBParams) error
}

var _ Querier = (*Queries)(nil)