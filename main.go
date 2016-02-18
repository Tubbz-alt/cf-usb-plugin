package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	swaggerclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/commands"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

var target string

type UsbPlugin struct {
	ui         terminal.UI
	httpClient *operations.Client
}

func main() {
	plugin.Start(new(UsbPlugin))
}

func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	argLength := len(args)

	c.ui = terminal.NewUI(os.Stdin, terminal.NewTeePrinter())

	// except command to set target
	if !(args[1] == "target" && argLength == 3) {
		target, err := config.GetTarget()
		if target == "" {
			fmt.Println("Usb management target not set. Use cf usb target <usb-mgmt-endpoint> to set the target")
			return
		}
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		/*sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}*/
		u, err := url.Parse(target)
		if err != nil {
			fmt.Println("ERROR :", err)
			return
		}
		transport := httptransport.New(u.Host, "/", []string{u.Scheme})

		debug, _ := strconv.ParseBool(os.Getenv("CF_TRACE"))

		transport.Debug = debug

		c.httpClient = operations.New(transport, strfmt.Default)
	}

	switch args[1] {
	case "target":
		if argLength == 2 {
			target, err := config.GetTarget()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Println("Usb management target: " + target)
		} else if argLength == 3 {
			target = args[2]

			err := config.SetTarget(target)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Println("Usb management target set to: " + target)
		}
	case "info":
		bearer, err := commands.GetBearerToken(cliConnection)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		infoResp, err := commands.NewInfoCommands(c.httpClient).GetInfo(bearer)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Broker API version: " + infoResp.BrokerAPIVersion)

		fmt.Println("USB version: " + infoResp.UsbVersion)
	case "create-driver":
		if argLength == 4 || argLength == 5 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			createdDriverId, err := commands.NewDriverCommands(c.httpClient).Create(bearer, args[2:5])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Println("Driver created with ID:", createdDriverId)
		} else {
			fmt.Println("ERROR: Invalid number of arguments")
			fmt.Println("Usage: cf usb create-driver [driver-type] [driver-name] [driver-bits-path]")
			return
		}
	case "delete-driver":
		if argLength == 3 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			deletedDriverId, err := commands.NewDriverCommands(c.httpClient).Delete(bearer, args[2])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if deletedDriverId == "" {
				fmt.Println("Driver not found")
			} else {
				fmt.Println("Driver deleted:", deletedDriverId)
			}
		} else {
			fmt.Println("Usage: cf usb delete-driver [driver-name]")
		}

	case "create-instance":
		if argLength == 6 || argLength == 4 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			createdInstanceId, err := commands.NewInstanceCommands(c.httpClient, c.ui).Create(bearer, args[2:6])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if createdInstanceId != "" {
				fmt.Println("New driver instance created. ID:" + createdInstanceId)
			}

		} else {
			fmt.Println("Usage cf usb create-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
			return
		}
	case "delete-instance":
		if argLength == 3 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			deletedInstanceId, err := commands.NewInstanceCommands(c.httpClient, c.ui).Delete(bearer, args[2])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if deletedInstanceId == "" {
				fmt.Println("Driver instance not found")
			} else {
				fmt.Println("Deleted driver instance:", deletedInstanceId)
			}
		} else {
			fmt.Println("Usage cf usb delete-instance [instanceName]")
		}
	case "update-driver":
		if argLength == 4 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			updatedDriverName, err := commands.NewDriverCommands(c.httpClient).Update(bearer, args[2:4])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if updatedDriverName == "" {
				fmt.Println("Driver not found")
			} else {
				fmt.Println("Driver updated:", updatedDriverName)
			}
		} else {
			fmt.Println("Usage: cf usb update-driver [old-driver-name] [new-driver-name]")
		}
	case "update-instance":
		if argLength == 6 || argLength == 4 {
			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			updateInstanceName, err := commands.NewInstanceCommands(c.httpClient, c.ui).Update(bearer, args[2:6])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if updateInstanceName != "" {
				fmt.Println("Driver instance updated:" + updateInstanceName)
			}
		} else {
			fmt.Println("Usage: cf usb update-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
			return
		}
	case "update-service":
		if argLength == 3 {
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			instanceName := args[2]

			instance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
			if instance == nil {
				fmt.Println("Driver instance not found")
				return
			}

			params := operations.NewUpdateServiceParams()
			params.ServiceID = *instance.Service

			var service models.Service
			service.DriverInstanceID = *instance.ID
			bindable := true
			bind := c.ui.Ask("Is service bindable?[Y/n]")
			if strings.ToLower(strings.Trim(bind, " ")) == "n" {
				bindable = false
			}
			service.Bindable = &bindable

			service.Name = c.ui.Ask("Service name")
			if service.Name == "" {
				fmt.Println("ERROR: Empty service name provided")
				return
			}
			desc := c.ui.Ask("Service description")
			if desc == "" {
				fmt.Println("ERROR: Empty service description provided")
				return
			}
			service.Description = &desc
			tags := c.ui.Ask("Tags (comma separated)")
			if tags == "" {
				fmt.Println("ERROR: Empty tags array provided")
				return
			}
			service.Tags = strings.Split(tags, ",")

			params.Service = &service

			response, err := c.httpClient.UpdateService(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Updated service with ID:", *response.Payload.ID)
		} else {
			fmt.Println("Usage: cf usb update-service [instanceName]")
			return
		}
	case "dials":
		if argLength == 3 {
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
			instance := getDriverInstanceByName(c.httpClient, bearer, args[2])
			if instance == nil {
				fmt.Println("Driver instance not found")
				return
			}
			params := operations.NewGetAllDialsParams()
			params.DriverInstanceID = instance.ID

			dials, err := c.httpClient.GetAllDials(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			for _, dial := range dials.Payload {
				fmt.Println("Dial configuration:\t", dial.Configuration)
				fmt.Println("Dial ID:\t\t", *dial.ID)
				fmt.Println("Plan ID:\t\t", *dial.Plan)

				plan, err := c.httpClient.GetServicePlan(&operations.GetServicePlanParams{PlanID: *dial.Plan}, bearer)
				if err != nil {
					fmt.Println("ERROR - getting plan", err)
				}
				fmt.Println("Plan:\t\t\t Name:", plan.Payload.Name, "; Description:", *plan.Payload.Description)
				fmt.Println("")
			}
		} else {
			fmt.Println("Usage: cf usb dials [instanceName]")
			return
		}
	case "instances":
		if argLength == 3 {

			bearer, err := commands.GetBearerToken(cliConnection)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			instances, err := commands.NewInstanceCommands(c.httpClient, c.ui).List(bearer, args[2])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if instances != nil {
				for _, di := range instances {
					fmt.Println("Driver Instance Name:\t", di.Name)
					fmt.Println("Driver Instance Id:\t", *di.ID)
					fmt.Println("Configuration:\t\t", di.Configuration)
					fmt.Println("Dials:\t\t\t", len(di.Dials))
					service, err := c.httpClient.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: *di.ID}, bearer)
					if err != nil {
						fmt.Println("ERROR:", err)
					}

					fmt.Println("Service:\t\t", "Name:", service.Payload.Name, "; Bindable:", *service.Payload.Bindable, "; Tags:", service.Payload.Tags)
					fmt.Println("")
				}
			}

		} else {
			fmt.Println("Usage cf usb instances [driverName]")
		}
	case "drivers":
		bearer, err := commands.GetBearerToken(cliConnection)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		drivers, err := commands.NewDriverCommands(c.httpClient).List(bearer)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if drivers != nil {
			table := terminal.NewTable(c.ui, []string{"Id", "Name", "Type"})
			for _, driver := range drivers {
				table.Add(*driver.ID, driver.Name, driver.DriverType)
			}
			table.Print()
		}
	}
}

func (c *UsbPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-plugin-usb",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "usb",
				HelpText: "Usb plugin command's help text",
				UsageDetails: plugin.Usage{
					Usage: "cf usb",
				},
			},
			plugin.Command{
				Name:     "usb target",
				HelpText: "Gets or sets usb management endpoint",

				UsageDetails: plugin.Usage{
					Usage: "usb target <usb-mgmt-endpoint>\n   cf usb target <usb-mgmt-endpoint>",
				},
			},
			plugin.Command{
				Name:     "usb info",
				HelpText: "Usb plugin token command text",

				UsageDetails: plugin.Usage{
					Usage: "usb token\n   cf usb token",
				},
			},
			plugin.Command{
				Name:     "usb create-instance",
				HelpText: "Usb plugin create driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb create-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]",
				},
			},
			plugin.Command{
				Name:     "usb delete-instance",
				HelpText: "Usb plugin delete driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb delete-instance [instanceName]",
				},
			},
			plugin.Command{
				Name:     "usb create-driver",
				HelpText: "Usb plugin create driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb create-driver [driverType] [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb update-driver",
				HelpText: "Usb plugin update driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-driver [oldDriverName] [newDriverName]",
				},
			},
			plugin.Command{
				Name:     "usb update-instance",
				HelpText: "Usb plugin update driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-instance [driverName] [instanceName]  configValue/configFile [jsonValue/filePath]",
				},
			},
			plugin.Command{
				Name:     "usb update-service",
				HelpText: "Usb plugin update service command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-service [instanceName]",
				},
			},
			plugin.Command{
				Name:     "usb delete-driver",
				HelpText: "Usb plugin delete driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb delete-driver [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb drivers",
				HelpText: "List existing drivers",

				UsageDetails: plugin.Usage{
					Usage: "usb drivers\n   cf usb drivers",
				},
			},
			plugin.Command{
				Name:     "usb instances",
				HelpText: "List existing driver instances",

				UsageDetails: plugin.Usage{
					Usage: "usb instances  [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb dials",
				HelpText: "List existing dials for instance",

				UsageDetails: plugin.Usage{
					Usage: "usb dials  [instanceName]",
				},
			},
		},
	}
}
