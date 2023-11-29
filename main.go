package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/shafiqaimanx/shellvenom/src/core"
	"github.com/shafiqaimanx/shellvenom/src/menu"
)

func main() {
	var listFlag bool
	var payloadFlag string
	var outputFlag string
	var chosenPlatforms []string

	flag.StringVar(&payloadFlag, "payload", "", "")
	flag.BoolVar(&listFlag, "list", false, "")
	flag.StringVar(&outputFlag, "output", "", "")

	flag.StringVar(&payloadFlag, "p", "", "")
	flag.BoolVar(&listFlag, "l", false, "")
	flag.StringVar(&outputFlag, "o", "", "")

	lhostFlag := flag.String("lhost", "", "")
	lportFlag := flag.String("lport", "", "")

	flag.Usage = func() {
		menu.HelpMenu()
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if listFlag {
		args := flag.Args()
		if len(args) != 1 || (args[0] != "windows" && args[0] != "unix" && args[0] != "all") {
			menu.ListRequiredFlag()
			os.Exit(0)
		}

		switch args[0] {
		case "all":
			chosenPlatforms = append(chosenPlatforms, "windows", "unix")
		case "windows":
			chosenPlatforms = append(chosenPlatforms, "windows")
		case "unix":
			chosenPlatforms = append(chosenPlatforms, "unix")
		default:
			return
		}
	}

	var unixPlatforms = make(map[string]string)
	var windowsPlatforms = make(map[string]string)

	for _, item := range core.ReadFileofShell().ShelListz {
		for _, platform := range core.ExtractPlatforms(item.Platform) {
			for _, chosen := range chosenPlatforms {
				if platform == chosen {
					var output string
					output += platform
					nameCombined := fmt.Sprintf("%s/%s", output, item.Name)
					if chosen == "unix" {
						unixPlatforms[nameCombined] = item.Description
					} else if chosen == "windows" {
						windowsPlatforms[nameCombined] = item.Description
					}
				}
			}
		}
	}

	if core.Contains(chosenPlatforms, "unix") {
		menu.ListingBanner("UNIX", unixPlatforms)
		var keys []string
		for nameCombined := range unixPlatforms {
			keys = append(keys, nameCombined)
		}
		sort.Strings(keys)
	
		for _, nameCombined := range keys {
			description := unixPlatforms[nameCombined]
			maxLength := len(nameCombined)
			if maxLength < 30 {
				maxLength = 28
			}
			fmt.Printf("%-*s    %s\n", maxLength, nameCombined, description)
		}
		fmt.Println("")
	}

	if core.Contains(chosenPlatforms, "windows") {
		menu.ListingBanner("WINDOWS", windowsPlatforms)
		var keys []string
		for nameCombined := range windowsPlatforms {
			keys = append(keys, nameCombined)
		}
		sort.Strings(keys)

		for _, nameCombined := range keys {
			description := windowsPlatforms[nameCombined]
			maxLength := len(nameCombined)
			if maxLength < 30 {
				maxLength = 28
			}
			fmt.Printf("%-*s    %s\n", maxLength, nameCombined, description)
		}
		fmt.Println("")
	}

	check := false
	for _, ip := range core.GettingIPAddresses() {
		if core.GettingIPInterfaceName(ip) != *lhostFlag {
			ipPattern := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
			ipRegex := regexp.MustCompile(ipPattern)
			if ipRegex.MatchString(*lhostFlag) {
				check = true
				break
			}
		}
	}

	if check {
		if payloadFlag != "" && *lhostFlag != "" && *lportFlag != "" && outputFlag != "" {
			for _, item := range core.ReadFileofShell().ShelListz {
				for _, platform := range core.ExtractPlatforms(item.Platform) {
					nameCombined := fmt.Sprintf("%s/%s", platform, item.Name)
					if payloadFlag == nameCombined {
						payload := core.DecodeBase64(item.Shell)
						result := core.FindAvailablePayloadIP(payload, *lhostFlag, *lportFlag)
						menu.OutputFile(payloadFlag, result)
						core.SavePayloadToFile(result, outputFlag)
						return
					}
				}
			}
		}

		if payloadFlag != "" && *lhostFlag != "" && *lportFlag != "" {
			for _, item := range core.ReadFileofShell().ShelListz {
				for _, platform := range core.ExtractPlatforms(item.Platform) {
					nameCombined := fmt.Sprintf("%s/%s", platform, item.Name)
					if payloadFlag == nameCombined {
						payload := core.DecodeBase64(item.Shell)
						result := core.FindAvailablePayloadIP(payload, *lhostFlag, *lportFlag)
						menu.OutputRaw(payloadFlag, result)					
						fmt.Println("\n"+result)
						return
					}
				}
			}
		}
	} else {
		if outputFlag != "" {
			ipList := core.GettingIPAddresses()
			for i:=0; i<len(ipList); i++ {
				nameListIndex := core.GettingIPInterfaceName(ipList[i])
				if nameListIndex == *lhostFlag {
					for _, item := range core.ReadFileofShell().ShelListz {
						for _, platform := range core.ExtractPlatforms(item.Platform) {
							nameCombined := fmt.Sprintf("%s/%s", platform, item.Name)
							if payloadFlag == nameCombined {
								payload := core.DecodeBase64(item.Shell)
								result := core.FindAvailablePayloadInterface(payload, *lportFlag, i, ipList)
								menu.OutputFile(payloadFlag, result)
								core.SavePayloadToFile(result, outputFlag)
								break
							}
						}
					}
				}
			}
		} else {
			ipList := core.GettingIPAddresses()
			for i:=0; i<len(ipList); i++ {
				nameListIndex := core.GettingIPInterfaceName(ipList[i])
				if nameListIndex == *lhostFlag {
					for _, item := range core.ReadFileofShell().ShelListz {
						for _, platform := range core.ExtractPlatforms(item.Platform) {
							nameCombined := fmt.Sprintf("%s/%s", platform, item.Name)
							if payloadFlag == nameCombined {
								payload := core.DecodeBase64(item.Shell)
								result := core.FindAvailablePayloadInterface(payload, *lportFlag, i, ipList)
								menu.OutputRaw(payloadFlag, result)
								fmt.Println("\n"+result)
								break
							}
						}
					}
				}
			}
		}
	}

	if !check {
		return
	}
}