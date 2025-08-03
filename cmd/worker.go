package cmd

import (
	"context"
	"log"
	"time"

	"github.com/arvinpaundra/sesen-api/application/worker/factory"
	"github.com/arvinpaundra/sesen-api/config"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/database/memorydb"
	"github.com/arvinpaundra/sesen-api/database/relationaldb"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start domain event worker",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		relationaldb.NewConnection(relationaldb.NewPostgres())
		memorydb.NewInMemoryConnection(memorydb.NewRedisDB())

		// Create and configure the domain event subscriber
		subscriber := factory.CreateDomainEventSubscriber(
			relationaldb.GetConnection(),
			memorydb.GetInMemoryConnection(),
		)

		// Start the worker
		go func() {
			log.Println("Starting domain event worker...")
			if err := subscriber.Start(context.Background()); err != nil {
				log.Fatalf("failed to start domain event worker: %s", err.Error())
			}
		}()

		// Setup graceful shutdown
		wait := util.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"domain-event-worker": func(_ context.Context) error {
				subscriber.Shutdown()
				return nil
			},
			"redis": func(_ context.Context) error {
				return memorydb.Close()
			},
		})

		<-wait
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
