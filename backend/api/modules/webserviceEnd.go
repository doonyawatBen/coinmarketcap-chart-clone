package modules

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/nlp/infrastructure"

	"golang.org/x/crypto/acme/autocert"
)

func WebserviceEnd(app *fiber.App) {
	/*
		Start HTTP/s Server
	*/
	go func() {
		if os.Getenv("DEPLOYMENT_ENV") == "production" {
			// reference: https://github.com/gofiber/recipes/blob/master/https-tls/main.go

			// Certificate manager
			autocertManager := &autocert.Manager{
				// Folder to store the certificates
				Cache: autocert.DirCache("/var/www/.cache"),
			}

			// TLS Config
			tlsConfig := &tls.Config{
				// Get Certificate from Let's Encrypt
				GetCertificate: autocertManager.GetCertificate,
				// By default NextProtos contains the "h2"
				// This has to be removed since Fasthttp does not support HTTP/2
				// Or it will cause a flood of PRI method logs
				// http://webconcepts.info/concepts/http-method/PRI
				NextProtos: []string{
					"http/1.1", "acme-tls/1",
				},
			}

			// Create custom listener
			ln, err := tls.Listen("tcp", fmt.Sprintf(":%d", infrastructure.ConfigGlobal.ServerPort), tlsConfig)
			if err != nil {
				infrastructure.Log.Panicln(nil, "fiber", "", "Fiber port HTTPS invalid", err)
			}

			err = app.Listener(ln)
			if err != nil {
				infrastructure.Log.Fatalln(nil, "fiber", "", "Fiber port HTTP invalid", err)
			}
		} else {
			err := app.Listen(fmt.Sprintf(":%d", infrastructure.ConfigGlobal.ServerPort))
			if err != nil {
				infrastructure.Log.Fatalln(nil, "fiber", "", "Fiber port HTTP invalid", err)
			}
		}
	}()

	// Handler Shutdown Fiber
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// This blocks the main thread until an interrupt is received
	<-shutdown
	infrastructure.Log.Info("fiber", "", "Gracefully shutting down...")
	_ = app.Shutdown()

	// Your cleanup tasks go here
	infrastructure.Log.Info("fiber", "", "Running cleanup tasks...")
	infrastructure.ShutdownServices()

	infrastructure.Log.Info("fiber", "", "Successful Shutdown")
}
