package config

import "testing"

func TestGetWebPort(t *testing.T) {
	t.Log(Get("key"))
}
