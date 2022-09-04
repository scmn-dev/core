package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/scmn-dev/core/config"
	"github.com/scmn-dev/core/db/creditcard"
	"github.com/scmn-dev/core/db/email"
	"github.com/scmn-dev/core/db/login"
	"github.com/scmn-dev/core/db/note"
	"github.com/scmn-dev/core/db/server"
	"github.com/scmn-dev/core/db/subscription"
	"github.com/scmn-dev/core/db/token"
	"github.com/scmn-dev/core/db/user"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database is the concrete store provider.
type Database struct {
	db            *gorm.DB
	logins        LoginRepository
	cards         CreditCardRepository
	notes         NoteRepository
	emails        EmailRepository
	tokens        TokenRepository
	users         UserRepository
	servers       ServerRepository
	subscriptions SubscriptionRepository
}

// DBConn databese connection
func DBConn(cfg *config.DatabaseConfiguration) (*gorm.DB, error) {
	var _db *gorm.DB
	var err error

	logmode := viper.GetBool("database.logmode")
	loglevel := logger.Silent
	if logmode {
		loglevel = logger.Info
	}

	newDBLogger := logger.New(
		log.New(getWriter(), "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loglevel,    // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newDBLogger})

	if err != nil {
		return nil, fmt.Errorf("could not open postgresql connection: %v", err)
	}

	return _db, err
}

// New opens a database according to configuration.
func New(xdb *gorm.DB) *Database {
	return &Database{
		db:            xdb,
		logins:        login.NewRepository(xdb),
		cards:         creditcard.NewRepository(xdb),
		notes:         note.NewRepository(xdb),
		emails:        email.NewRepository(xdb),
		tokens:        token.NewRepository(xdb),
		users:         user.NewRepository(xdb),
		servers:       server.NewRepository(xdb),
		subscriptions: subscription.NewRepository(xdb),
	}
}

// Logins returns the LoginRepository.
func (xdb *Database) Logins() LoginRepository {
	return xdb.logins
}

// CreditCards returns the CreditCardRepository.
func (xdb *Database) CreditCards() CreditCardRepository {
	return xdb.cards
}

func (xdb *Database) Notes() NoteRepository {
	return xdb.notes
}

func (xdb *Database) Emails() EmailRepository {
	return xdb.emails
}

// Tokens returns the TokenRepository.
func (xdb *Database) Tokens() TokenRepository {
	return xdb.tokens
}

// Users returns the UserRepository.
func (xdb *Database) Users() UserRepository {
	return xdb.users
}

// Servers returns the UserRepository.
func (xdb *Database) Servers() ServerRepository {
	return xdb.servers
}

// Subscriptions returns the UserRepository.
func (xdb *Database) Subscriptions() SubscriptionRepository {
	return xdb.subscriptions
}

// Ping checks if database is up
func (xdb *Database) Ping() error {
	sqlDB, err := xdb.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func getWriter() io.Writer {
	file, err := os.OpenFile("sm-core.db.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}
