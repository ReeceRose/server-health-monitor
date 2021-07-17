package repository

import "github.com/PR-Developers/server-health-monitor/internal/database"

type baseRepository struct {
	db *database.MongoDB
}
