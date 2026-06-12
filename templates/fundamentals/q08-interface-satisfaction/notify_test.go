package notify

import "testing"

func TestNotifier(t *testing.T) {
	var _ Notifier = (*Email)(nil)
	var _ Notifier = (*SMS)(nil)
}
