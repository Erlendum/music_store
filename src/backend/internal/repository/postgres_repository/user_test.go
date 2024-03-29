package postgres_repository

import (
	"backend/internal/models"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"
)

var testUserPostgresRepositoryGetSuccess = []struct {
	TestName  string
	InputData struct {
		login string
	}
	CheckOutput func(t *testing.T, user *models.User, err error)
}{
	{
		TestName: "usual test",
		InputData: struct {
			login string
		}{login: "admin"},
		CheckOutput: func(t *testing.T, user *models.User, err error) {
			require.NoError(t, err)
			require.Equal(t, user, &models.User{
				UserId:    1,
				Login:     "admin",
				Password:  "$2a$10$1OOImwqdj8VCsC10WmVZZOqQqZ3roHRhFq69jZSwOdUQOZhzczv4S",
				Fio:       "admin",
				DateBirth: time.Date(2003, 1, 22, 0, 0, 0, 0, time.UTC),
				Gender:    "Male",
				IsAdmin:   true})
		},
	},
}

var testUserPostgresRepositoryGetFailed = []struct {
	TestName  string
	InputData struct {
		login string
	}
	CheckOutput func(t *testing.T, err error)
}{

	{
		TestName: "user does not exists",
		InputData: struct {
			login string
		}{login: "ha"},
		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
		},
	},
}

func TestUserPostgresRepositoryGet(t *testing.T) {
	for _, tt := range testUserPostgresRepositoryGetSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			dbContainer, db := SetupTestDatabase("./migrations/000001_create_init_tables.up.sql")
			defer func(dbContainer testcontainers.Container, ctx context.Context) {
				err := dbContainer.Terminate(ctx)
				if err != nil {
					return
				}
			}(dbContainer, context.Background())

			fields := new(PostgresRepositoryFields)
			fields.Db = db
			userRepository := CreateUserPostgresRepository(fields)

			user, err := userRepository.Get(tt.InputData.login)

			tt.CheckOutput(t, user, err)
		})
	}

	for _, tt := range testUserPostgresRepositoryGetFailed {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			dbContainer, db := SetupTestDatabase("./migrations/000001_create_init_tables.up.sql")
			defer func(dbContainer testcontainers.Container, ctx context.Context) {
				err := dbContainer.Terminate(ctx)
				if err != nil {
					return
				}
			}(dbContainer, context.Background())

			fields := new(PostgresRepositoryFields)
			fields.Db = db
			userRepository := CreateUserPostgresRepository(fields)

			_, err := userRepository.Get(tt.InputData.login)

			tt.CheckOutput(t, err)
		})
	}
}

var testUserPostgresRepositoryGetByIdSuccess = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	CheckOutput func(t *testing.T, user *models.User, err error)
}{
	{
		TestName: "usual test",
		InputData: struct {
			id uint64
		}{id: 1},
		CheckOutput: func(t *testing.T, user *models.User, err error) {
			require.NoError(t, err)
			require.Equal(t, user, &models.User{
				UserId:    1,
				Login:     "admin",
				Password:  "$2a$10$1OOImwqdj8VCsC10WmVZZOqQqZ3roHRhFq69jZSwOdUQOZhzczv4S",
				Fio:       "admin",
				DateBirth: time.Date(2003, 1, 22, 0, 0, 0, 0, time.UTC),
				Gender:    "Male",
				IsAdmin:   true})
		},
	},
}

var testUserPostgresRepositoryGetByIdFailed = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	CheckOutput func(t *testing.T, err error)
}{

	{
		TestName: "user does not exists",
		InputData: struct {
			id uint64
		}{id: 82828},
		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
		},
	},
}

func TestUserPostgresRepositoryGetById(t *testing.T) {
	for _, tt := range testUserPostgresRepositoryGetByIdSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			dbContainer, db := SetupTestDatabase("./migrations/000001_create_init_tables.up.sql")
			defer func(dbContainer testcontainers.Container, ctx context.Context) {
				err := dbContainer.Terminate(ctx)
				if err != nil {
					return
				}
			}(dbContainer, context.Background())

			fields := new(PostgresRepositoryFields)
			fields.Db = db
			userRepository := CreateUserPostgresRepository(fields)

			user, err := userRepository.GetById(tt.InputData.id)

			tt.CheckOutput(t, user, err)
		})
	}

	for _, tt := range testUserPostgresRepositoryGetByIdFailed {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			dbContainer, db := SetupTestDatabase("./migrations/000001_create_init_tables.up.sql")
			defer func(dbContainer testcontainers.Container, ctx context.Context) {
				err := dbContainer.Terminate(ctx)
				if err != nil {
					return
				}
			}(dbContainer, context.Background())

			fields := new(PostgresRepositoryFields)
			fields.Db = db
			userRepository := CreateUserPostgresRepository(fields)

			_, err := userRepository.GetById(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}
}
