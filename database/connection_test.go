package database

import "testing"

func TestNewConnection(t *testing.T) {
	t.Log(NewConnection("../.env"))
}
