package main

import (
	"database/sql"
	"net/http"

	httpHandler "github.com/DanielH556/imersaofullcycle18/golang/internal/events/infra/http"
	"github.com/DanielH556/imersaofullcycle18/golang/internal/events/infra/service"
	"github.com/DanielH556/imersaofullcycle18/golang/internal/events/usecase"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://host.docker.internal:8000/partner1",
		2: "http://host.docker.internal:8000/partner2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventsUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventsUseCase,
		buyTicketUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListEvents)
	r.HandleFunc("POST /checkout", eventsHandler.ListEvents)

	http.ListenAndServe(":8080", r)
}
