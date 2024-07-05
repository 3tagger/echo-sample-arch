package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/3tagger/echo-sample-arch/internal/config"
	"github.com/3tagger/echo-sample-arch/internal/database"
)

const (
	userEntity = "user"
)

var availableEntity = []string{
	userEntity,
}

func main() {
	ctx := context.Background()

	// flags
	var (
		n      int
		entity string
	)
	flag.IntVar(&n, "n", 10, "the number of fake data to be generated")
	flag.StringVar(&entity, "entity", "", fmt.Sprintf("The entity to be seeded, supported entities: \n%s",
		strings.Join(availableEntity, "\n")))
	flag.Parse()

	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	// database
	db, err := database.InitPostgreSQL(cfg.Database.Primary)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	seederMap := initSeederMap(db)

	seederExec, ok := seederMap[entity]
	if !ok {
		log.Println("please choose an entity (entity) name to seed")
		return
	}

	now := time.Now()
	log.Println("running the seeder...")
	seederExec.Run(ctx, n)

	log.Printf("seeding finished in %v...\n", time.Since(now))
}
