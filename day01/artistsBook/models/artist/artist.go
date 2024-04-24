package artist

import (
	"fmt"
	"time"
)

// Artist structure
type Artist struct {
	Id   				 string			`json:"id"`
	Name 				 string 		`json:"name"`
	Top  				 string 		`json:"top"`
	Fans 				 int 				`json:"fans"`
	ListenedTime time.Time 	`json:"listenedTime"`
}

func (a Artist) String() string {
	return fmt.Sprintf("-- %s --\n", a.Name)
}