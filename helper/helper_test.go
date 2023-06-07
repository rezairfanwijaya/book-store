package helper

import "testing"

func TestGetEnv(t *testing.T) {
	t.Log(GetENV("../.env"))
}
