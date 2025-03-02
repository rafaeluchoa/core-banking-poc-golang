package boot

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"slices"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var commonHeaders = []string{
	"accept", "accept-encoding", "accept-language", "cache-control",
	"connection", "content-length", "content-type", "host", "pragma",
}

type LogEntry struct {
	Timestamp time.Time              `bson:"timestamp"`
	Method    string                 `bson:"method"`
	URL       string                 `bson:"url"`
	Query     string                 `bson:"query"`
	Body      map[string]interface{} `bson:"body"`
	Headers   map[string]string      `bson:"headers"`
	Response  string                 `bson:"response"`
	Status    int                    `bson:"status"`
}

type AuditMid struct {
	app   *fiber.App
	mongo *mongo.Database
}

func NewAudit(app *fiber.App, mongo *mongo.Database) *AuditMid {
	audit := &AuditMid{
		app:   app,
		mongo: mongo,
	}

	app.Use("/api/*", audit.handle)

	return audit
}

func (s *AuditMid) handle(c *fiber.Ctx) error {
	err := c.Next()

	method := c.Method()
	url := c.OriginalURL()
	query := string(c.Context().QueryArgs().QueryString())

	var body map[string]any
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		body = nil
	}

	headers := make(map[string]string)

	c.Request().Header.VisitAll(func(key, value []byte) {
		headerKey := strings.ToLower(string(key))
		if !slices.Contains(commonHeaders, headerKey) {
			headers[headerKey] = string(value)
		}
	})

	status := c.Response().StatusCode()
	response := string(c.Response().Body())

	logEntry := LogEntry{
		Timestamp: time.Now(),
		Method:    method,
		URL:       url,
		Query:     query,
		Body:      body,
		Headers:   headers,
		Response:  response,
		Status:    status,
	}

	go func() {
		s.insertLog(logEntry)
	}()

	return err
}

func (s *AuditMid) insertLog(logEntry LogEntry) {
	tracelog := s.mongo.Collection("tracelog")
	_, errMongo := tracelog.InsertOne(context.Background(), logEntry)

	if errMongo != nil {
		log.Println("Error on insert mongo:", errMongo)
	}
}
