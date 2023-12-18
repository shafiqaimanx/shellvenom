package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/shafiqaimanx/shellvenom/src/core"
	"github.com/shafiqaimanx/shellvenom/src/menu"
)

func main() {
	var listFlag 		bool
	var outputFlag 		string
	var payloadFlag 	string
	var choosenIFace 	string
	var choosenLHost 	string
	var choosenLPort 	int
	var chosenPlatforms []string

	checkIPIntFace 	:= false
	checkIPRangeRgx := false

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
					nameCombined := fmt.Sprintf("cmd/%s/%s", output, item.Name)
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
			if maxLength < 35 {
				maxLength = 35
			}
			fmt.Printf("    %-*s    %s\n", maxLength, nameCombined, description)
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
			if maxLength < 35 {
				maxLength = 35
			}
			fmt.Printf("    %-*s    %s\n", maxLength, nameCombined, description)
		}
		fmt.Println("")
	}



	if payloadFlag != "" && *lhostFlag != "" && *lportFlag != "" && outputFlag != "" {
		checkIPIntFace, choosenIFace = core.CheckIPInterfaceName(*lhostFlag)
		checkIPRangeRgx = core.CheckIPrange(*lhostFlag)

		host, choosenIPRgx := core.CompareIPwithIntFaceAndRegex(checkIPIntFace, checkIPRangeRgx, lhostFlag)
		port, choosenPort := core.CheckPortRange(*lportFlag)
	
		menu.HostandPortErrorHandling(host, port)
		if choosenIFace != "" {
			{}
		} else if choosenIPRgx != "" {
			choosenIFace = choosenIPRgx
		}

		if (host && port) {
			choosenLHost = choosenIFace
			choosenLPort = choosenPort
		}

		for _, item := range core.ReadFileofShell().ShelListz {
			for _, platform := range core.ExtractPlatforms(item.Platform) {
				nameCombined := fmt.Sprintf("cmd/%s/%s", platform, item.Name)
				if payloadFlag == nameCombined {
					payload := core.DecodeBase64(item.Shell)
					result := core.FindAvailablePayloadIP(payload, choosenLHost, strconv.Itoa(choosenLPort))
					menu.OutputFile(payloadFlag, result)
					core.SavePayloadToFile(result, outputFlag)
					return
				}
			}
		}
		fmt.Printf("%s[-]%s %sPayload%s %s%s%s %snot found!%s\n", menu.CRIMSON, menu.RESET, menu.ITALIC, menu.RESET, menu.BOLD, payloadFlag, menu.RESET, menu.ITALIC, menu.RESET)
		os.Exit(0)
	}



	if payloadFlag != "" && *lhostFlag != "" && *lportFlag != "" {
		checkIPIntFace, choosenIFace = core.CheckIPInterfaceName(*lhostFlag)
		checkIPRangeRgx = core.CheckIPrange(*lhostFlag)

		host, choosenIPRgx := core.CompareIPwithIntFaceAndRegex(checkIPIntFace, checkIPRangeRgx, lhostFlag)
		port, choosenPort := core.CheckPortRange(*lportFlag)
	
		menu.HostandPortErrorHandling(host, port)
		if choosenIFace != "" {
			{}
		} else if choosenIPRgx != "" {
			choosenIFace = choosenIPRgx
		}

		if (host && port) {
			choosenLHost = choosenIFace
			choosenLPort = choosenPort
		}

		for _, item := range core.ReadFileofShell().ShelListz {
			for _, platform := range core.ExtractPlatforms(item.Platform) {
				nameCombined := fmt.Sprintf("cmd/%s/%s", platform, item.Name)
				if payloadFlag == nameCombined {
					payload := core.DecodeBase64(item.Shell)
					result := core.FindAvailablePayloadIP(payload, choosenLHost, strconv.Itoa(choosenLPort))
					menu.OutputRaw(payloadFlag, result)					
					fmt.Println(result)
					return
				}
			}
		}
		fmt.Printf("%s[-]%s %sPayload%s %s%s%s %snot found!%s\n", menu.CRIMSON, menu.RESET, menu.ITALIC, menu.RESET, menu.BOLD, payloadFlag, menu.RESET, menu.ITALIC, menu.RESET)
		os.Exit(0)
	}
}