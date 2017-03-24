package controllers

import (
	"github.com/yashsriv/dashboard-http/config"
	"gopkg.in/kataras/iris.v5"
)

// TimeTableRequest is the request format for the server
type TimeTableRequest struct {
	Username string `json:"username"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

// TimeTableEvent is the response from the server
type TimeTableEvent struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	End         string `json:"end"`
	Start       string `json:"start"`
}

// CalendarEvent is the format required by FullCalendar
type CalendarEvent struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

// CalendarAddEvent is the event received from calendar
type CalendarAddEvent struct {
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"name"`
}

// TimeTableAddEvent is the event format required for backend
type TimeTableAddEvent struct {
	Username string           `json:"username"`
	Event    CalendarAddEvent `json:"event"`
}

// GetFromTimeTable gets entry from time table
func GetFromTimeTable(ctx *iris.Context) {
	username := ctx.RequestHeader("X-Username-Header")
	start := ctx.URLParam("start")
	end := ctx.URLParam("end")
	response, err := config.Timetable.Call("Timetable.get", TimeTableRequest{
		Username: username,
		Start:    start,
		End:      end,
	})
	if err != nil {
		SendInternalServer(err, ctx)
	}
	var events []TimeTableEvent
	response.GetObject(&events)

	calendarEvents := make([]CalendarEvent, len(events))
	for i, event := range events {
		calendarEvents[i] = CalendarEvent{
			ID:    event.ID,
			Title: event.Description,
			Start: event.Start,
			End:   event.End,
		}
	}

	_ = ctx.JSON(iris.StatusOK, calendarEvents)
}

type status struct {
	Successful bool `json:"successful"`
}

// AddToTimeTable adds event to the timetable
func AddToTimeTable(ctx *iris.Context) {
	username := ctx.RequestHeader("X-Username-Header")

	event := CalendarAddEvent{}
	err := ctx.ReadJSON(&event)
	if err != nil {
		_ = ctx.Text(iris.StatusBadRequest, err.Error())
		return
	}

	timetableEvent := TimeTableAddEvent{Username: username, Event: event}

	response, err := config.Timetable.Call("Timetable.add", timetableEvent)
	if err != nil {
		SendInternalServer(err, ctx)
		return
	}
	var success status
	response.GetObject(&success)

	if success.Successful {
		_ = ctx.Text(iris.StatusOK, "")
	} else {
		_ = ctx.Text(iris.StatusInternalServerError, "")
	}

}
