package commands

import (
	"fmt"

	"github.com/devopsdays/migrator/migrate"
	"github.com/spf13/cobra"
)

// eventCmd represents the "event" command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Migrate an event",
	Long: `Migrate an event.
`,
	Example: `  migrator event
	migrator event -e 2019-chicago`,

	Run: func(cmd *cobra.Command, args []string) {
		migrateEvent(Event)
	},
}

// eventsCmd represents the "events" command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Migrate events",
	Long: `Migrate events.
`,
	Example: `  migrator events
	`,

	Run: func(cmd *cobra.Command, args []string) {
		migrateEvents()
	},
}

// sponsorsCmd represents the "sponsors" command
var sponsorsCmd = &cobra.Command{
	Use:   "sponsors",
	Short: "Migrate sponsors",
	Long: `Migrate sponsors.
`,
	Example: `  migrator sponsors
	`,

	Run: func(cmd *cobra.Command, args []string) {
		migrateSponsors()
	},
}

func init() {
	RootCmd.AddCommand(eventCmd)
	RootCmd.AddCommand(sponsorsCmd)
	RootCmd.AddCommand(eventsCmd)

	eventCmd.Flags().StringVarP(&Event, "event", "e", "", "event to use")
	eventCmd.MarkFlagRequired("event")
}

func migrateEvent(event string) {
	fmt.Println("This would migrate the ", event, " event.")
	migrate.MigrateEvent(migrate.CityStrip(event), migrate.YearStrip(event))
}

func migrateSponsors() {
	migrate.ConvertSponsors(migrate.GetOldSponsorsPath(), migrate.GetNewSponsorsPath())

	fmt.Println("This would migrate sponsors")
}

func migrateEvents() {
	fmt.Println("This would migrate all events")
	migrate.MigrateAllEvents()

}
