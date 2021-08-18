package cmd

import (
	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"

	"github.com/walkergriggs/hellosb/api"
	"github.com/walkergriggs/hellosb/server"
)

type ServerOptions struct {
	Host        string
	Port        int
	CatalogPath string
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Host:        "localhost",
		Port:        3000,
		CatalogPath: "./mocks/catalog.json",
	}
}

func (o *ServerOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *ServerOptions) Run() {
	swaggerSpec, err := loads.Embedded(api.SwaggerJSON, api.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	defaultAPI := api.New(swaggerSpec)
	defaultAPI.Configure(&api.Config{
		CatalogPath: o.CatalogPath,
	})

	opts := server.Options{
		Host: o.Host,
		Port: o.Port,
		API:  defaultAPI,
	}

	if err := server.New(opts).Serve(); err != nil {
		panic(err)
	}
}

func NewCmdServer() *cobra.Command {
	o := NewServerOptions()

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Runs a Hellosb server",
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.Complete(cmd, args); err != nil {
				return
			}

			o.Run()
		},
	}

	cmd.Flags().StringVar(&o.Host, "host", o.Host, "The host to serve the REST API on")
	cmd.Flags().IntVar(&o.Port, "port", o.Port, "The port to serve the REST API on")
	cmd.Flags().StringVar(&o.CatalogPath, "catalog", o.CatalogPath, "The path to the catalog JSON file")

	return cmd
}
