package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Trendyol/service/genproto"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func (r *UsersRepo) CreateUser(user *pb.CreateUserReq) (*pb.User, error) {

	query := `INSERT INTO users(id,first_name,last_name,username,phone,email,password,address,gender,role,postalcode,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id`

	fmt.Println(user)
	err := r.db.QueryRow(query, user.Id, user.FirstName, user.LastName, user.Username, user.Phone, user.Email, user.Password, user.Address, user.Gender, user.Role, user.Postalcode, time.Now().UTC()).Scan(
		&user.Id,
	)
	if err != nil {
		return nil, err
	}
	newuser, err := r.GetUserById(user.Id)
	if err != nil {
		return nil, err
	}
	return newuser, nil
}

func (r *UsersRepo) CheckField(field, value string) (bool, error) {
	var existClient int
	if field == "username" {
		row := r.db.QueryRow(`SELECT count(1) FROM users WHERE username = $1 AND deleted_at IS NULL`, field)
		if err := row.Scan(&existClient); err != nil {
			return false, err
		}
	} else if field == "email" {
		row := r.db.QueryRow(`SELECT count(1) FROM users WHERE email = $1 and deleted_at IS NULL`, value)
		if err := row.Scan(&existClient); err != nil {
			return false, err
		}

	} else {
		return false, nil
	}
	if existClient == 0 {
		return false, nil
	}
	return true, nil

}

func (r *UsersRepo) UpdateUser(user *pb.User) (*pb.User, error) {
	query := `UPDATE users SET first_name=$1,last_name=$2,username=$3,phone=$4,email=$5,gender=$6,role=$7,postalcode=$8,updated_at=$9 where id=$10
	RETURNING id,first_name,last_name,username,phone,email,password,gender,role,postalcode,created_at`
	var updUser pb.User
	fmt.Println(user)
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Username, user.Phone, user.Email, user.Gender, user.Role, user.Postalcode, time.Now().UTC(), user.Id).Scan(
		&updUser.Id,
		&updUser.FirstName,
		&updUser.LastName,
		&updUser.Username,
		&updUser.Phone,
		&updUser.Email,
		&updUser.Gender,
		&updUser.Role,
		&updUser.Postalcode,
		&updUser.CreatedAt,
	)
	updUser.UpdatedAt = time.Now().UTC().String()
	if err != nil {
		return nil, err
	}
	return &updUser, nil
}

func (r *UsersRepo) GetUserById(id string) (*pb.User, error) {

	var (
		user       pb.User
		updated_at sql.NullTime
	)
	query := `SELECT id,first_name,last_name,username,phone,email,password,address,gender,role,postalcode,created_at,updated_at from users where id=$1 and deleted_at is null`

	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.Gender,
		&user.Role,
		&user.Postalcode,
		&user.CreatedAt,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		user.UpdatedAt = updated_at.Time.String()
	}

	return &user, nil
}

func (r *UsersRepo) GetAllUsers() ([]*pb.User, error) {
	query := `SELECT id,first_name,last_name,username,phone,email,password,address,gender,role,postalcode,created_at,updated_at from users WHERE deleted_at is null`
	var users []*pb.User
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pb.User
		var updated_at sql.NullTime
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Phone,
			&user.Email,
			&user.Password,
			&user.Address,
			&user.Gender,
			&user.Role,
			&user.Postalcode,
			&user.CreatedAt,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			user.UpdatedAt = updated_at.Time.String()
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UsersRepo) DeleteUserById(id string) (*pb.User, error) {
	user, err := r.GetUserById(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE users SET deleted_at=$2 where id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UsersRepo) LoginUser(login *pb.LoginUserReq) (*pb.User, error) {
	var (
		user       pb.User
		updated_at sql.NullTime
	)

	query := `SELECT id,first_name,last_name,username,email,password,role,created_at,updated_at from users where email=$1 and deleted_at is null`
	err := r.db.QueryRow(query, login.Email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		user.UpdatedAt = updated_at.Time.String()
	}
	return &user, nil
}

func (r *UsersRepo) EmailValid(email string) (bool, error) {
	var count int = 0
	query := `SELECT count(*) from users where email=$1`
	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return true, err
	}
	if count > 1 {
		return false, nil
	}
	return true, nil

}

func (r *UsersRepo) ChangePassword(newPass *pb.ChangePassReq)(*pb.ChangePassRes,error){
	query:=`UPDATE users SET password=$1 where deleted_at is null and id =$2 returning id,password`
	var pass pb.ChangePassRes
	err:=r.db.QueryRow(query,newPass.NewPassword,newPass.Id).Scan(
		&pass.Id,
		&pass.NewPassword,
	)
	if err!=nil{
		return nil,err
	}
	return &pass,nil
}

func NewUserRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}
