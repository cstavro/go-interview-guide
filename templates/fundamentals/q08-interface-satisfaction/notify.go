package notify

// Notifier is the interface for sending notifications.
type Notifier interface {
	Send(to string, message string) error
}

// Compile-time interface checks.
// TODO: ensure Email and SMS satisfy Notifier

type Email struct{}

func (e *Email) Send(to, msg string) error {
	// TODO
	return nil
}

type SMS struct{}

func (s *SMS) Send(to, msg string) error {
	// TODO
	return nil
}
