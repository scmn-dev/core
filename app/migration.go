package app

import (
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/logger"
)

// MigrateSystemTables runs auto migration for the system models (Token and User),
// will only add missing fields won't delete/change current data in the store.
func MigrateSystemTables(s db.Store) {
	if err := s.Tokens().Migrate(); err != nil {
		logger.Errorf("failed to migrate tokens: %v", err)
	}

	if err := s.Users().Migrate(); err != nil {
		logger.Errorf("failed to migrate users: %v", err)
	}

	if err := s.Subscriptions().Migrate(); err != nil {
		logger.Errorf("failed to migrate subscriptions: %v", err)
	}
}

// MigrateUserTables runs auto migration for user models in user schema,
// will only add missing fields won't delete/change current data in the store.
func MigrateUserTables(s db.Store, schema string) error {
	if err := s.Logins().Migrate(schema); err != nil {
		logger.Errorf("failed to migrate logins: %v", err)
		return err
	}

	if err := s.CreditCards().Migrate(schema); err != nil {
		logger.Errorf("failed to migrate credit cards: %v", err)
		return err
	}

	if err := s.Notes().Migrate(schema); err != nil {
		logger.Errorf("failed to migrate notes: %v", err)
		return err
	}

	if err := s.Emails().Migrate(schema); err != nil {
		logger.Errorf("failed to migrate emails: %v", err)
		return err
	}

	if err := s.Servers().Migrate(schema); err != nil {
		logger.Errorf("failed to migrate servers: %v", err)
		return err
	}

	return nil
}
