package calendar

import (
	"fmt"
	"strings"
)

type Event struct {
	title string
	Date
}

func (e *Event)Title() string {
	return e.title
}
func (e *Event) SetEvent(event string)error{
	var wordcount int
	for i := 0; i < len(event); i++ {
		alphabet := strings.Replace(event, " ", "", -1)
		wordcount =  len(strings.Split(alphabet, ""))
	}
	if wordcount > 30{
		return fmt.Errorf("invalid title")
	}
	e.title = event
	return nil
}