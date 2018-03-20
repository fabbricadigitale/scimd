package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cenk/backoff"
	"github.com/spf13/viper"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/spf13/cobra"
)

var adapter storage.Storer

var scimd = &cobra.Command{
	Use: "scimd",
	// TraverseChildren: true,
	Short: "SCIMD is ...",
	Long: `...
	
	SCIM 2 RFC - published under the IETF - defines an open API for managing identities.
Complete documentation is available at ...`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Flags overrides (or merge with) configuration values
		// Thus we re-validate the configuration prior to execute and we collect errors
		if _, err := config.Valid(); err != nil {
			errors, _ := err.(validator.ValidationErrors)
			config.Errors = append(config.Errors, errors...)
		}

		// Printing errors
		if len(config.Errors) > 0 {
			fmt.Fprintln(os.Stderr, validation.Errors(config.Errors))
			os.Exit(1)
		}

		// Ensure eventual custom config is ok and load it
		if err := config.Custom(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// try to connect to storage, if not db up: notify error and os.Exit(1)
		// (todo) manage storage switching
		if err := dbConnect(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get unique attributes from loaded schemas
		keys, err := attr.GetUniqueAttributes()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// when db exists and it is up and running: get collection + ensure index SYNC (we need it to be blocking)
		// Note that the session executing EnsureIndex will be blocked for as long as it takes for the index to be built.
		// So I think the next line will be blocking
		if err := ensureIndexes(keys); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// verify if session needs to be closed
		// To avoid to leave an hanging session, I'll close it
		adapter.Close()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Start the server with the current service provider config
		spc := config.ServiceProviderConfig()
		if !config.Values.Debug {
			fmt.Printf("Serving on port %d ...\n", config.Values.Port)
		}
		server.Get(&spc).Run(":" + strconv.Itoa(config.Values.Port))
	},
	DisableAutoGenTag: true,
}

func init() {
	// Definition of flags
	scimd.Flags().StringVar(&config.Values.Storage.Type, "storage-type", config.Values.Storage.Type, "type of storage to use")
	scimd.Flags().StringVar(&config.Values.Storage.Host, "storage-host", config.Values.Storage.Host, "the storage's address")
	scimd.Flags().IntVar(&config.Values.Storage.Port, "storage-port", config.Values.Storage.Port, "the storage's port")
	scimd.Flags().StringVar(&config.Values.Storage.Name, "storage-name", config.Values.Storage.Name, "the storage's database name")
	scimd.Flags().StringVar(&config.Values.Storage.Coll, "storage-coll", config.Values.Storage.Coll, "the storage's collection name")

	scimd.PersistentFlags().BoolVar(&config.Values.Debug, "debug", config.Values.Debug, "wheter to enable or not the debug mode")

	scimd.Flags().IntVarP(&config.Values.Port, "port", "p", config.Values.Port, "port to run the server on")

	scimd.Flags().StringVarP(&config.Values.ServiceProviderConfig, "service-provider-config", "s", config.Values.ServiceProviderConfig, "the path of the service provider config to use")
	scimd.Flags().StringVarP(&config.Values.Config, "config", "c", config.Values.Config, "the path of directory containing the configuration resources")

	scimd.Flags().IntVar(&config.Values.PageSize, "page-size", config.Values.PageSize, "the page size the server has to use")

	scimd.Flags().BoolVar(&config.Values.Enable.Self, "enable-self", config.Values.Enable.Self, "whether to enable or not the creation of an endpoint for the authenticated entity")

	// Binding flags to configuration manager
	viper.BindPFlags(scimd.Flags())
}

// Execute is starting point for commands
func Execute() {
	scimd.SuggestionsMinimumDistance = 1
	if err := scimd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Root returns the root command
func Root() *cobra.Command {
	return scimd
}

func dbConnect() (err error) {
	endpoint := fmt.Sprintf("%s:%d", config.Values.Storage.Host, config.Values.Storage.Port)

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 1 * time.Millisecond

	err = backoff.Retry(func() error {
		fmt.Printf("Trying to reach storage at %s ...\n", endpoint)
		var err error

		adapter, err = mongo.New(endpoint, config.Values.Name, config.Values.Coll)
		if err != nil {
			return err
		}

		// Configuration step for ensure uniqueness attributes in the storage
		uniqueAttrs, err := attr.GetUniqueAttributes()
		if err != nil {
			return err
		}

		adapter.SetIndexes(uniqueAttrs)

		return adapter.Ping()
	}, b)

	return
}

func ensureIndexes(keys [][]string) (err error) {
	err = adapter.SetIndexes(keys)
	return
}
