package harness

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleGet(t *testing.T) {
	setup()
	defer teardown()

	fmt.Println("DOCKERUP")
	time.Sleep(1 * time.Second)
}
