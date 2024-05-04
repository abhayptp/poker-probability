package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/strategy"
	"github.com/gorilla/mux"
)

type Card struct {
	Rank string `json:"rank"`
	Suit string `json:"suit"`
}

type CommunityCards struct {
	Cards []Card `json:"cards"`
}

type PlayerCard struct {
	name  string `json:"name"`
	Cards []Card `json:"cards"`
}

type GetProbabilityRequest struct {
	CommunityCards   CommunityCards `json:"community_cards"`
	PlayerCards      []PlayerCard   `json:"player_cards"`
	SimulationRounds int            `json:"simulation_rounds"`
	PlayersCount     int            `json:"players_count"`
}

func AdaptToModelsCard(c Card) models.Card {
	card, _ := models.NewCard(models.Rank(c.Rank), models.Suit(c.Suit))
	return card
}

func AdaptToModelsCardList(cardList []Card) []models.Card {
	res := make([]models.Card, 0)
	for _, c := range cardList {
		res = append(res, AdaptToModelsCard(c))
	}
	return res
}

type PlayerResult struct {
	Name           string  `json:"name"`
	WinProbability float64 `json:"win_probability"`
	TieProbability float64 `json:"tie_probability"`
}

type GetProbabilityResponse struct {
	PlayerResults []PlayerResult `json:"player_results"`
}

func getProbabilityHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request GetProbabilityRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	pCards := make([]models.DealtCards, 0)

	for _, playerCards := range request.PlayerCards {
		cardList := AdaptToModelsCardList(playerCards.Cards)
		pCards = append(pCards, models.NewDealtCards(cardList...))
	}

	cardList := AdaptToModelsCardList(request.CommunityCards.Cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cCards := models.NewCommunityCards(cardList...)

	strategy := strategy.NewApproximate(
		cCards, pCards, request.PlayersCount, request.SimulationRounds, 5,
	)

	strategyRes := strategy.Run()

	playerResults := make([]PlayerResult, request.PlayersCount)
	for i, result := range strategyRes.PlayerResult {
		playerResults[i] = PlayerResult{
			Name:           request.PlayerCards[i].name,
			WinProbability: result.WinProbability,
			TieProbability: result.TieProbability,
		}
	}
	response := GetProbabilityResponse{
		PlayerResults: playerResults,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get_probability", getProbabilityHandler).Methods("GET")
	return r
}

func main() {
	log.Print("starting server...")

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
