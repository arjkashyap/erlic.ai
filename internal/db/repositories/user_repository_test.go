package repositories

import (
	"database/sql"
	"testing"

	"github.com/arjkashyap/erlic.ai/internal/models"
)

func TestUserRepository_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepository{
				db: tt.fields.db,
			}
			if err := ur.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
