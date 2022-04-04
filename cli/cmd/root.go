package cmd

import (
	"fmt"
	"sync"

	"github.com/blackcrw/wprecon/internal/banner"
	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/models"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
	"github.com/blackcrw/wprecon/internal/text"
	"github.com/blackcrw/wprecon/internal/views"
	finders_plugins "github.com/blackcrw/wprecon/tools/finders/plugins"
	finders_themes "github.com/blackcrw/wprecon/tools/finders/themes"
	finders_users "github.com/blackcrw/wprecon/tools/finders/users"

	"github.com/blackcrw/wprecon/tools/interesting"
	"github.com/spf13/cobra"
)

func RootOptionsPreRun(cmd *cobra.Command, args []string) {
	net.ThisIsHostValid(database.Memory.GetString("Options URL"))
	var is_url = net.ThisIsURL(database.Memory.GetString("Options URL"))

	if is_url {
		banner.SBanner()
	} else {
		banner.Banner()
	}

	var response = net.SimpleRequest(database.Memory.GetString("Options URL"))
	
	database.Memory.SetString("HTTP Index Raw", response.Raw)
	database.Memory.SetString("HTTP PHP Version", response.Response.Header.Get("x-powered-by"))
	database.Memory.SetString("HTTP Server", response.Response.Header.Get("Server"))
	database.Memory.SetString("HTTP Index Cookie", response.Response.Header.Get("Set-Cookie"))
}

func RootOptionsRun(cmd *cobra.Command, args []string) {
	var (
		flag_target_url, _ = cmd.Flags().GetString("url")
		flag_aggressive_mode, _ = cmd.Flags().GetBool("aggressive-mode")
		flag_detection_waf, _ = cmd.Flags().GetBool("detection-waf")
	)

	var (
		wordpress_version = interesting.WordPressVersion()
		wordpress_confidence = interesting.WordpressCheck()
		wordpress_confidence_string = fmt.Sprintf("%.2f%%", wordpress_confidence)
	)

	if wordpress_confidence >= 40.0 {
		printer.Done("WordPress confirmed with", wordpress_confidence_string, "confidence!\n")
	} else if wordpress_confidence < 40.0 && wordpress_confidence > 15.0 && !database.Memory.GetBool("Options Force") {
		if q := printer.ScanQ("I'm not absolutely sure that this target is using wordpress!", wordpress_confidence_string, "chance. do you wish to continue ? [Y]es | [n]o : "); q != "y" && q != "\n" {
			printer.Fatal("Exiting...")
		}
		printer.Println()
	} else if wordpress_confidence < 15.0 && !database.Memory.GetBool("Options Force") {
		printer.Fatal("This target is not running wordpress!")
	}

	if wordpress_version != "" {
		printer.Done("WordPress Version:")
		printer.NewTopics("Version:", wordpress_version, "\n").Default()
	}

	var wordpress_waf = func() *models.InterestingModel { if flag_detection_waf { var waf, err = interesting.WordpressFirewall(); if waf.Name != "" { return waf } else if err != nil { printer.Danger(err) } }; return &models.InterestingModel{} }()

	if wordpress_waf.Name != "" {
		views.RootWAF(wordpress_waf)
	}


	switch flag_aggressive_mode {
	case true:
		printer.Info("User Enumerate:\n")
		if finders_user_slice := *finders_users.Run(flag_target_url); len(finders_user_slice) == 0 {
			printer.Danger("Unfortunately no user was found.\n")
		} else {
			for _, finders_user := range finders_user_slice { views.RootFindersUser(finders_user) }; printer.Println()
		}
	
		printer.Info("Plugin Enumerate:\n")
		if finders_plugins_slice := *finders_plugins.Aggressive(flag_target_url); len(finders_plugins_slice) == 0 {
			printer.Danger("I couldn't find any plugins\n")
		} else { for _, finders_plugin := range finders_plugins_slice { views.RootFindersPluginsAndThemes(finders_plugin) } }
	
		printer.Info("Theme Enumerate:\n")
		if finders_themes_slice := *finders_themes.Aggressive(flag_target_url); len(finders_themes_slice) == 0 {
			printer.Danger("I couldn't find any themes\n")
		} else { for _, finders_theme := range finders_themes_slice { views.RootFindersPluginsAndThemes(finders_theme) } }

	case false:
		printer.Info("Plugin Enumerate:\n")
		if finders_plugins_slice := *finders_plugins.Passive(flag_target_url); len(finders_plugins_slice) == 0 {
			printer.Danger("I couldn't find any plugins\n")
		} else { for _, finders_plugin := range finders_plugins_slice { views.RootFindersPluginsAndThemes(finders_plugin) } }
	
		printer.Info("Theme Enumerate:\n")
		if finders_themes_slice := *finders_themes.Passive(flag_target_url); len(finders_themes_slice) == 0 {
			printer.Danger("I couldn't find any themes\n")
		} else { for _, finders_theme := range finders_themes_slice { views.RootFindersPluginsAndThemes(finders_theme) } }
	}
}

func RootOptionsPostRun(cmd *cobra.Command, args []string) {
	var wait_group sync.WaitGroup

	wait_group.Add(5)

	printer.Info("Other interesting information:\n")

	if database.Memory.GetString("HTTP Server") != "" || database.Memory.GetString("HTTP PHP Version") != "" {
		printer.Done("Target information(s):")

		if server := database.Memory.GetString("HTTP Server"); server != "" { printer.NewTopics("Server:", server).Default() }
		if version := database.Memory.GetString("HTTP PHP Version"); version != "" { printer.NewTopics("PHP Version:", version).Default() }
		if version := database.Memory.GetString("HTTP WordPress Version"); version != "" { printer.NewTopics("WordPress Version:", version).Default() }

		printer.Println()
	}

	go func(){
		if len(database.Memory.GetSlice("HTTP Index Of's")) > 0 {
			printer.Done("Index Of's:")
			for _, indexofs := range database.Memory.GetSlice("HTTP Index Of's") {
				printer.NewTopics(indexofs).Default()
			}
			printer.Println()
		}
		
		defer wait_group.Done()
	}()

	go func(){
		if response, err := interesting.XMLRPC(); response.Confidence > 0 {
			if response.Confidence <= 10 {
				printer.Done("XML-RPC Possibly enabled:")
			} else {
				printer.Done("XML-RPC Enabled:")
				printer.NewTopics("Status:", response.Status).Default()
			}
	
			printer.NewTopics("Confidence:", fmt.Sprint(response.Confidence)+"%").Default()
			printer.NewTopics("Found By:", response.FoundBy).Default()
			printer.NewTopics("Location:", database.Memory.GetString("Options URL")+"xmlrpc.php", "\n").Default()
		} else if err != nil { printer.Danger("%s", err) }

		defer wait_group.Done()
	}()

	go func(){
		if URL := database.Memory.GetString("HTTP Admin Page"); URL != "" {
			printer.Done("Admin Page Found:")
			printer.NewTopics("Found by: Access").Default()
			printer.NewTopics("Location:", URL, "\n").Default()
		}

		defer wait_group.Done()
	}()

	go func(){	
		if response, err := interesting.ReadmePage(); response.Status == 200 {
			printer.Done("WordPress Readme:")
			printer.NewTopics("Found by:", response.FoundBy).Default()
			printer.NewTopics("Location:", response.Url, "\n").Default()
		} else if err != nil { printer.Danger(err) }

		defer wait_group.Done()
	}()

	go func(){
		if raw := database.Memory.GetString("HTTP wp-content/uploads Index Of Raw"); raw != "" {
			var list_backup_paths = text.FindBackupFileOrPath(raw)
			
			if len(list_backup_paths) > 0 {
				printer.Done("File or Path backup:")
				for _, backup_path := range list_backup_paths {
					printer.NewTopics(database.Memory.GetString("Options URL") + database.Memory.GetString("HTTP wp-content") + "/uploads/" + backup_path[1]).Default()
				}
				printer.Println()
			}
		}

		defer wait_group.Done()
	}()

	wait_group.Wait()

	printer.Done("Total requests:", database.Memory.GetInt("HTTP Total"))
}