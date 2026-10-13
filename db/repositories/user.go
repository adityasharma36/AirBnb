
package db

import (
    "database/sql"
    "errors"
    "fmt"

    "goMod/model"

    mysql "github.com/go-sql-driver/mysql"
)

var ErrUserAlreadyExists = errors.New("email or username already exists")

type UserRepository interface {
    GetUserById(id int64) (*model.User, error)
    Create(username, email, password string) (*model.User, error)
    GetAll() ([]*model.User, error)
    DeleteById(id int64) (bool, error)
    GetUserByEmail(email string) (*model.EmailFetch, error)
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &UserRepositoryImpl{
        db: db,
    }
}

func (u *UserRepositoryImpl) Create(
    username string,
    email string,
    password string,
) (*model.User, error) {

    query := `
        INSERT INTO users (username, email, password)
        VALUES (?, ?, ?)
    `

    result, err := u.db.Exec(
        query,
        username,
        email,
        password,
    )

    if err != nil {
        var mysqlErr *mysql.MySQLError

        if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
            return nil, ErrUserAlreadyExists
        }

        return nil, fmt.Errorf("insert user: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, fmt.Errorf("get created user ID: %w", err)
    }

    user, err := u.GetUserById(id)
    if err != nil {
        return nil, fmt.Errorf("fetch created user: %w", err)
    }

    return user, nil
}

func (u *UserRepositoryImpl) GetUserById(
    id int64,
) (*model.User, error) {

    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
        WHERE id = ?
    `

    user := &model.User{}

    err := u.db.QueryRow(
        query,
        id,
    ).Scan(
        &user.Id,
        &user.Username,
        &user.Email,
        &user.Created_at,
        &user.Updated_at,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, sql.ErrNoRows
        }

        return nil, fmt.Errorf("fetch user by ID: %w", err)
    }

    return user, nil
}

func (u *UserRepositoryImpl) GetUserByEmail(
    email string,
) (*model.EmailFetch, error) {

    query := `
        SELECT id, email, password
        FROM users
        WHERE email = ?
    `

    user := &model.EmailFetch{}

    err := u.db.QueryRow(
        query,
        email,
    ).Scan(
        &user.Id,
        &user.Email,
        &user.Password,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, sql.ErrNoRows
        }

        return nil, fmt.Errorf("fetch user by email: %w", err)
    }

    return user, nil
}

func (u *UserRepositoryImpl) GetAll() ([]*model.User, error) {

    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
    `

    rows, err := u.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("fetch users: %w", err)
    }

    defer rows.Close()

    users := make([]*model.User, 0)

    for rows.Next() {
        user := &model.User{}

        err := rows.Scan(
            &user.Id,
            &user.Username,
            &user.Email,
            &user.Created_at,
            &user.Updated_at,
        )

        if err != nil {
            return nil, fmt.Errorf("scan user: %w", err)
        }

        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("iterate users: %w", err)
    }

    return users, nil
}

func (u *UserRepositoryImpl) DeleteById(
    id int64,
) (bool, error) {

    query := `DELETE FROM users WHERE id = ?`

    result, err := u.db.Exec(query, id)
    if err != nil {
        return false, fmt.Errorf("delete user: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, fmt.Errorf("get affected rows: %w", err)
    }

    return rowsAffected > 0, nil
}
