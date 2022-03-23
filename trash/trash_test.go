package iwan

import (
	"fmt"
	"testing"
	"time"
)

func TestIwan(t *testing.T) {
	d := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	fmt.Println(d) // 1900-01-01 00:00:00 +0000 UTC

	d2 := d.AddDate(0, 0, 44562)

	fmt.Println(d2)

}
