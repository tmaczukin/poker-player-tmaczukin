package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func callRequest(action, game_state string) (writer *httptest.ResponseRecorder) {
	writer = httptest.NewRecorder()

	body := strings.NewReader("")
	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/?action=%s&game_state=%s", action, game_state), body)
	if err != nil {
		log.Fatal(err)
	}

	handleRequest(writer, r)

	//log.Printf("Request:  %s", r)
	//log.Printf("Response: %s", writer)
	//log.Printf("Response code: %d", writer.Code)
	log.Printf("Response body: %s", writer.Body)

	return
}

func assertIntEqual(t *testing.T, input, pattern int) {
	if input == pattern {
		return
	}

	t.Log(fmt.Sprintf("%d is not equal to %d", input, pattern))
	t.Fail()
}

func getGameState() string {
	return `{"small_blind": 0, "current_buy_in": 0}`
}

func TestUnknownAction(t *testing.T) {
	r := callRequest("unknown", "")
	assertIntEqual(t, r.Code, 400)
}

func TestVersionAction(t *testing.T) {
	r := callRequest("version", "")
	assertIntEqual(t, r.Code, 200)
}

func TestCheckAction(t *testing.T) {
	r := callRequest("check", "")
	assertIntEqual(t, r.Code, 200)
}

func TestInvalidBetRequestAction(t *testing.T) {
	r := callRequest("bet_request", "")
	assertIntEqual(t, r.Code, 500)
}

func TestBetReuqestAction(t *testing.T) {
	r := callRequest("bet_request", getGameState())
	assertIntEqual(t, r.Code, 200)
}

func TestInvalidShowndownAction(t *testing.T) {
	r := callRequest("showdown", "")
	assertIntEqual(t, r.Code, 500)
}

func TestShowndownAction(t *testing.T) {
	r := callRequest("showdown", getGameState())
	assertIntEqual(t, r.Code, 200)
}
