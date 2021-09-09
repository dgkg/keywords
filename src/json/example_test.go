package json

import (
	"os"
	"time"
)

func ExampleString() {
	var u User = User{
		Name:      "Bob",
		BirthDate: BirthDate(time.Date(2007, 1, 2, 0, 0, 0, 0, time.FixedZone("UTC+1", 1*60*60))),
	}
	os.Stdout.Write([]byte(u.BirthDate.String()))
	// Output:
	// 2007-01-02
}
