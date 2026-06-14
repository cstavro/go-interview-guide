package notify

// Notifier is the interface for sending notifications.
type Notifier interface {
	Send(to string, message string) error
}

// Compile-time interface checks.
// TODO: ensure Email and SMS satisfy Notifier

type Email struct{}

// TODO: implement Notifier for Email

type SMS struct{}

// TODO: implement Notifier for SMS
