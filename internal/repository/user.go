package repository

import (
	"context"
	"database/sql"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func User(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (u userRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("user").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("user").Where(goqu.Ex{
		"username": username,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) InsertUser(ctx context.Context, req dto.UserRegisterReq) (interface{}, error) {
	dataset := u.db.Insert("user").Cols("username", "password", "phone", "fullname").Vals(goqu.Vals{req.Username, req.Password, req.Phone, req.Fullname})

	// Menjalankan query dan menyimpan hasilnya
	sql, _, err := dataset.ToSQL()
	if err != nil {
		return nil, err
	}

	// Eksekusi query INSERT
	_, err = u.db.ExecContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	// Menyediakan hasil query jika diperlukan, misalnya ID user yang baru dimasukkan
	// bisa menggunakan cara yang sesuai untuk pengembalian data
	return nil, nil // Bisa dikembangkan jika ingin mengembalikan data lebih lanjut

}
