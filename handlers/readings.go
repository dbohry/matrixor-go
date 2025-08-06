package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"matrixor/config"
	"matrixor/models"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReadingHandler struct {
	Collection *mongo.Collection
}

func NewReadingHandler(client *mongo.Client, dbCfg config.DatabaseConfig) *ReadingHandler {
	col := client.Database(dbCfg.Name).Collection(dbCfg.Collection)
	return &ReadingHandler{Collection: col}
}

func (h *ReadingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReadingHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var reading models.Reading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	reading.CreatedAt = time.Now()

	_, err := h.Collection.InsertOne(context.Background(), reading)
	if err != nil {
		http.Error(w, "DB insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reading)
}

func (h *ReadingHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		http.Error(w, "Missing 'limit' query parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, "'limit' must be a positive integer", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := h.Collection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(limit)).SetSort(bson.M{"createdat": -1}))
	if err != nil {
		http.Error(w, "Database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var readings []models.Reading
	if err := cursor.All(ctx, &readings); err != nil {
		http.Error(w, "Cursor decode failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}
