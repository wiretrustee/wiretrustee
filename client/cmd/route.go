package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"

	"github.com/netbirdio/netbird/client/proto"
)

var appendFlag bool

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Manage network routes",
	Long:  `Commands to list, select, or deselect network routes.`,
}

var routesListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List routes",
	Example: "  netbird routes list",
	Long:    "List all available network routes.",
	RunE:    routesList,
}

var routesSelectCmd = &cobra.Command{
	Use:     "select route...|all",
	Short:   "Select routes",
	Long:    "Select a list of routes by identifiers or 'all' to clear all selections and to accept all (including new) routes.\nDefault mode is replace, use -a to append to already selected routes.",
	Example: "  netbird routes select all\n  netbird routes select route1 route2\n  netbird routes select -a route3",
	Args:    cobra.MinimumNArgs(1),
	RunE:    routesSelect,
}

var routesDeselectCmd = &cobra.Command{
	Use:     "deselect route...|all",
	Short:   "Deselect routes",
	Long:    "Deselect previously selected routes by identifiers or 'all' to disable accepting any routes.",
	Example: "  netbird routes deselect all\n  netbird routes deselect route1 route2",
	Args:    cobra.MinimumNArgs(1),
	RunE:    routesDeselect,
}

func init() {
	routesSelectCmd.PersistentFlags().BoolVarP(&appendFlag, "append", "a", false, "Append to current route selection instead of replacing")
}

func routesList(cmd *cobra.Command, _ []string) error {
	conn, err := getClient(cmd.Context())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewDaemonServiceClient(conn)
	resp, err := client.ListRoutes(cmd.Context(), &proto.ListRoutesRequest{})
	if err != nil {
		return fmt.Errorf("failed to list routes: %v", status.Convert(err).Message())
	}

	if len(resp.Routes) == 0 {
		cmd.Println("No routes available.")
		return nil
	}

	printRoutes(cmd, resp)

	return nil
}

func printRoutes(cmd *cobra.Command, resp *proto.ListRoutesResponse) {
	cmd.Println("Available Routes:")
	for _, route := range resp.Routes {
		printRoute(cmd, route)
	}
}

func printRoute(cmd *cobra.Command, route *proto.Route) {
	selectedStatus := getSelectedStatus(route)
	domains := route.GetDomains()

	if len(domains) > 0 {
		printDomainRoute(cmd, route, domains, selectedStatus)
	} else {
		printNetworkRoute(cmd, route, selectedStatus)
	}
}

func getSelectedStatus(route *proto.Route) string {
	if route.GetSelected() {
		return "Selected"
	}
	return "Not Selected"
}

func printDomainRoute(cmd *cobra.Command, route *proto.Route, domains []string, selectedStatus string) {
	cmd.Printf("\n  - ID: %s\n    Domains: %s\n    Status: %s\n", route.GetID(), strings.Join(domains, ", "), selectedStatus)
	resolvedIPs := route.GetResolvedIPs()

	if len(resolvedIPs) > 0 {
		printResolvedIPs(cmd, domains, resolvedIPs)
	} else {
		cmd.Printf("    Resolved IPs: -\n")
	}
}

func printNetworkRoute(cmd *cobra.Command, route *proto.Route, selectedStatus string) {
	cmd.Printf("\n  - ID: %s\n    Network: %s\n    Status: %s\n", route.GetID(), route.GetNetwork(), selectedStatus)
}

func printResolvedIPs(cmd *cobra.Command, domains []string, resolvedIPs map[string]*proto.IPList) {
	cmd.Printf("    Resolved IPs:\n")
	for _, domain := range domains {
		if ipList, exists := resolvedIPs[domain]; exists {
			cmd.Printf("      [%s]: %s\n", domain, strings.Join(ipList.GetIps(), ", "))
		}
	}
}

func routesSelect(cmd *cobra.Command, args []string) error {
	conn, err := getClient(cmd.Context())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewDaemonServiceClient(conn)
	req := &proto.SelectRoutesRequest{
		RouteIDs: args,
	}

	if len(args) == 1 && args[0] == "all" {
		req.All = true
	} else if appendFlag {
		req.Append = true
	}

	if _, err := client.SelectRoutes(cmd.Context(), req); err != nil {
		return fmt.Errorf("failed to select routes: %v", status.Convert(err).Message())
	}

	cmd.Println("Routes selected successfully.")

	return nil
}

func routesDeselect(cmd *cobra.Command, args []string) error {
	conn, err := getClient(cmd.Context())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewDaemonServiceClient(conn)
	req := &proto.SelectRoutesRequest{
		RouteIDs: args,
	}

	if len(args) == 1 && args[0] == "all" {
		req.All = true
	}

	if _, err := client.DeselectRoutes(cmd.Context(), req); err != nil {
		return fmt.Errorf("failed to deselect routes: %v", status.Convert(err).Message())
	}

	cmd.Println("Routes deselected successfully.")

	return nil
}
