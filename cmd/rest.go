package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arvinpaundra/sesen-api/application/rest/router"
	"github.com/arvinpaundra/sesen-api/config"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/database/memorydb"
	"github.com/arvinpaundra/sesen-api/database/relationaldb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var port string

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		relationaldb.NewConnection(relationaldb.NewPostgres())

		memorydb.NewInMemoryConnection(memorydb.NewRedisDB())

		g := gin.New()

		router.Register(g, memorydb.GetInMemoryConnection(), relationaldb.GetConnection())

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: g,
		}

		go func() {
			log.Println("Starting REST server...")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("failed to start server: %s", err.Error())
			}
		}()

		wait := util.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"rest-server": func(_ context.Context) error {
				return srv.Close()
			},
			"postgres": func(_ context.Context) error {
				return relationaldb.Close()
			},
		})

		<-wait
	},
}

func init() {
	restCmd.Flags().StringVarP(&port, "port", "p", "8000", "bind rest server to port")
	rootCmd.AddCommand(restCmd)
}
