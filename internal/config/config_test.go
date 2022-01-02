package config

import "testing"

func TestDatabase_GetDsn(t *testing.T) {
	type fields struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simple test",
			fields{
				Host: "localhost",
				Port: "8080",
				User: "admin",
				Password: "password",
				DBName: "db",
			},
			"admin:password@tcp(localhost:8080)/db?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				DBName:   tt.fields.DBName,
			}
			if got := db.GetDsn(); got != tt.want {
				t.Errorf("Database.GetDsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
