package db

// Store is the minimal interface for the various repositories
type Store interface {
	Logins() LoginRepository
	CreditCards() CreditCardRepository
	Notes() NoteRepository
	Emails() EmailRepository
	Tokens() TokenRepository
	Users() UserRepository
	Servers() ServerRepository
	Subscriptions() SubscriptionRepository
	Ping() error
}
