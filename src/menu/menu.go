package menu

import (
	"fmt"
	"os"
)

const (
	RESET 					= "\033[0m"
	BOLD 					= "\033[1m"
	ITALIC 					= "\033[3m"
	UNDERSCORE 				= "\033[4m"
	INDIANRED 				= "\033[38;5;167m"
	LIGHTCORAL 				= "\033[38;5;174m"
	SALMON 					= "\033[38;5;174m"
	DARKSALMON 				= "\033[38;5;174m"
	LIGHTSALMON 			= "\033[38;5;216m"
	CRIMSON 				= "\033[38;5;161m"
	RED 					= "\033[38;5;196m"
	FIREBRICK 				= "\033[38;5;124m"
	DARKRED					= "\033[38;5;88m"
	PINK 					= "\033[38;5;217m"
	LIGHTPINK 				= "\033[38;5;217m"
	HOTPINK 				= "\033[38;5;211m"
	DEEPPINK 				= "\033[38;5;198m"
	MEDIUMVIOLETRED 		= "\033[38;5;126m"
	PALEVIOLETRED 			= "\033[38;5;174m"
	CORAL 					= "\033[38;5;209m"
	TOMATO 					= "\033[38;5;203m"
	ORANGERED 				= "\033[38;5;202m"
	DARKORANGE 				= "\033[38;5;208m"
	ORANGE 					= "\033[38;5;214m"
	GOLD 					= "\033[38;5;220m"
	YELLOW 					= "\033[38;5;226m"
	LIGHTYELLOW 			= "\033[38;5;230m"
	LEMONCHIFFON 			= "\033[38;5;224m"
	LIGHTGOLDENRODYELLOW 	= "\033[38;5;188m"
	PAPAYAWHIP 				= "\033[38;5;224m"
	MOCCASIN 				= "\033[38;5;223m"
	PEACHPUFF 				= "\033[38;5;223m"
	PALEGOLDENROD 			= "\033[38;5;187m"
	KHAKI 					= "\033[38;5;186m"
	DARKKHAKI 				= "\033[38;5;144m"
	LAVENDER 				= "\033[38;5;188m"
	THISTLE 				= "\033[38;5;182m"
	PLUM 					= "\033[38;5;182m"
	VIOLET 					= "\033[38;5;176m"
	ORCHID 					= "\033[38;5;176m"
	FUCHSIA 				= "\033[38;5;201m"
	MAGENTA 				= "\033[38;5;201m"
	MEDIUMORCHID 			= "\033[38;5;134m"
	MEDIUMPURPLE 			= "\033[38;5;104m"
	BLUEVIOLET				= "\033[38;5;92m"
	DARKVIOLET				= "\033[38;5;92m"
	DARKORCHID				= "\033[38;5;128m"
	DARKMAGENTA				= "\033[38;5;90m"
	PURPLE					= "\033[38;5;90m"
	INDIGO					= "\033[38;5;54m"
	SLATEBLUE				= "\033[38;5;98m"
	DARKSLATEBLUE			= "\033[38;5;60m"
	MEDIUMSLATEBLUE 		= "\033[38;5;104m"
	GREENYELLOW 			= "\033[38;5;154m"
	CHARTREUSE 				= "\033[38;5;118m"
	LAWNGREEN 				= "\033[38;5;112m"
	LIME					= "\033[38;5;46m"
	LIMEGREEN				= "\033[38;5;40m"
	PALEGREEN 				= "\033[38;5;114m"
	LIGHTGREEN 				= "\033[38;5;114m"
	MEDIUMSPRINGGREEN		= "\033[38;5;43m"
	SPRINGGREEN				= "\033[38;5;48m"
	MEDIUMSEAGREEN			= "\033[38;5;72m"
	SEAGREEN				= "\033[38;5;29m"
	FORESTGREEN				= "\033[38;5;28m"
	GREEN					= "\033[38;5;28m"
	DARKGREEN				= "\033[38;5;22m"
	YELLOWGREEN 			= "\033[38;5;148m"
	OLIVEDRAB 				= "\033[38;5;100m"
	OLIVE 					= "\033[38;5;100m"
	DARKOLIVEGREEN			= "\033[38;5;64m"
	MEDIUMAQUAMARINE 		= "\033[38;5;115m"
	DARKSEAGREEN 			= "\033[38;5;108m"
	LIGHTSEAGREEN			= "\033[38;5;37m"
	DARKCYAN				= "\033[38;5;30m"
	TEAL					= "\033[38;5;30m"
	AQUA					= "\033[38;5;51m"
	CYAN					= "\033[38;5;51m"
	LIGHTCYAN 				= "\033[38;5;195m"
	PALETURQUOISE 			= "\033[38;5;152m"
	AQUAMARINE 				= "\033[38;5;122m"
	TURQUOISE				= "\033[38;5;80m"
	MEDIUMTURQUOISE			= "\033[38;5;80m"
	DARKTURQUOISE			= "\033[38;5;44m"
	CADETBLUE				= "\033[38;5;73m"
	STEELBLUE				= "\033[38;5;67m"
	LIGHTSTEELBLUE 			= "\033[38;5;146m"
	POWDERBLUE 				= "\033[38;5;152m"
	LIGHTBLUE 				= "\033[38;5;152m"
	SKYBLUE 				= "\033[38;5;116m"
	LIGHTSKYBLUE 			= "\033[38;5;116m"
	DEEPSKYBLUE				= "\033[38;5;39m"
	DODGERBLUE				= "\033[38;5;33m"
	CORNFLOWERBLUE			= "\033[38;5;68m"
	ROYALBLUE				= "\033[38;5;68m"
	BLUE					= "\033[38;5;21m"
	MEDIUMBLUE				= "\033[38;5;20m"
	DARKBLUE				= "\033[38;5;18m"
	NAVY					= "\033[38;5;18m"
	MIDNIGHTBLUE			= "\033[38;5;18m"
	CORNSILK 				= "\033[38;5;224m"
	BLANCHEDALMOND 			= "\033[38;5;224m"
	BISQUE 					= "\033[38;5;223m"
	NAVAJOWHITE 			= "\033[38;5;223m"
	WHEAT 					= "\033[38;5;187m"
	BURLYWOOD 				= "\033[38;5;180m"
	TAN 					= "\033[38;5;180m"
	ROSYBROWN 				= "\033[38;5;138m"
	SANDYBROWN 				= "\033[38;5;179m"
	GOLDENROD 				= "\033[38;5;178m"
	DARKGOLDENROD 			= "\033[38;5;136m"
	PERU 					= "\033[38;5;173m"
	CHOCOLATE 				= "\033[38;5;172m"
	SADDLEBROWN				= "\033[38;5;94m"
	SIENNA 					= "\033[38;5;130m"
	BROWN 					= "\033[38;5;124m"
	MAROON					= "\033[38;5;88m"
	WHITE 					= "\033[38;5;231m"
	SNOW 					= "\033[38;5;224m"
	HONEYDEW 				= "\033[38;5;194m"
	MINTCREAM 				= "\033[38;5;194m"
	AZURE 					= "\033[38;5;195m"
	ALICEBLUE 				= "\033[38;5;189m"
	GHOSTWHITE 				= "\033[38;5;189m"
	WHITESMOKE 				= "\033[38;5;188m"
	SEASHELL 				= "\033[38;5;224m"
	BEIGE 					= "\033[38;5;188m"
	OLDLACE 				= "\033[38;5;188m"
	FLORALWHITE 			= "\033[38;5;224m"
	IVORY 					= "\033[38;5;230m"
	ANTIQUEWHITE 			= "\033[38;5;188m"
	LINEN 					= "\033[38;5;188m"
	LAVENDERBLUSH 			= "\033[38;5;224m"
	MISTYROSE 				= "\033[38;5;224m"
	GAINSBORO 				= "\033[38;5;188m"
	LIGHTGREY 				= "\033[38;5;188m"
	SILVER 					= "\033[38;5;145m"
	DARKGRAY 				= "\033[38;5;145m"
	GRAY 					= "\033[38;5;102m"
	DIMGRAY 				= "\033[38;5;102m"
	LIGHTSLATEGRAY 			= "\033[38;5;103m"
	SLATEGRAY 				= "\033[38;5;102m"
	DARKSLATEGRAY			= "\033[38;5;23m"
	BLACK 					= "\033[38;5;16m"
)

func ShellVenomBanner() {
	fmt.Printf("%s╔═╗╦ ╦╔═╗╦  ╦ ╦  ╦╔═╗╔╗╔╔═╗╔╦╗%s\n", MEDIUMORCHID, RESET)
	fmt.Printf("%s╚═╗╠═╣║╣ ║  ║ ╚╗╔╝║╣ ║║║║ ║║║║%s\n", MEDIUMORCHID, RESET)
	fmt.Printf("%s╚═╝╩ ╩╚═╝╩═╝╩═╝╚╝ ╚═╝╝╚╝╚═╝╩ ╩%s\n", MEDIUMORCHID, RESET)
}


func HelpMenu() {
	shell := fmt.Sprintf("%s%srevshells.com%s", UNDERSCORE, ITALIC, RESET)
	venom := fmt.Sprintf("%s%smsfvenom%s", UNDERSCORE, ITALIC, RESET)

	ShellVenomBanner()
	fmt.Printf("%sshellVenom - reverse shell payload generator.%s\n", DARKGRAY, RESET)
	fmt.Printf("%sThis project was inspire by %s %sand %s%s%s.\n\n", DARKGRAY, shell, DARKGRAY, venom, DARKGRAY, RESET)
	fmt.Printf("%s%sUsage%s: %sshellvenom%s [flags]\n", BOLD, UNDERSCORE, RESET, BOLD, RESET)
	fmt.Printf("%s%sExample%s: shellvenom -p cmd/unix/reverse_bash -lhost=[ip] -lport=[port]\n\n", BOLD, UNDERSCORE, RESET)
	fmt.Printf("%s%sFlags%s:\n", BOLD, UNDERSCORE, RESET)
	fmt.Printf("  -l, -list       <platform>      List all payloads for [platform]. PLatform are: %s%sall%s, %s%sunix%s, %s%swindows%s.\n", ITALIC, UNDERSCORE, RESET, ITALIC, UNDERSCORE, RESET, ITALIC, UNDERSCORE, RESET)
	fmt.Printf("  -p, -payload    <payloads>      Payload to use (-list %s%sall%s to list).\n", ITALIC, UNDERSCORE, RESET)
	fmt.Println("      -lhost      <ip address>    Listening IP of attacker.")
	fmt.Println("      -lport      <port>          Listening PORT of attacker.")
	fmt.Println("  -o, -output     <path>          Save the payload to a file.")
	fmt.Println("  -h, -help                       Show help menu and exit.")
}

func ListRequiredFlag() {
	fmt.Printf("%s[-]%s -list flag requires a valid argument: %sall%s, %sunix%s, %swindows%s\n", CRIMSON, RESET, UNDERSCORE, RESET, UNDERSCORE, RESET, UNDERSCORE, RESET)
}

func ListingBanner(platformName string, platforms map[string]string) {
	fmt.Printf("\n[%s%s%s] Listing of unix payloads (%s%d%s total) [-payload <value>]\n", CHARTREUSE, platformName, RESET, CHARTREUSE, len(platforms), RESET)
	fmt.Printf("============================================================\n\n")
	fmt.Printf("  %s%sName%s\t\t\t\t\t%s%sDescription%s\n\n", MEDIUMORCHID, UNDERSCORE, RESET, MEDIUMORCHID, UNDERSCORE, RESET)
}

func OutputRaw(payloadFlag string, result string) {
	ShellVenomBanner()
	fmt.Printf("%s[-]%s Selecting %s%s%s from the payload.\n", CHARTREUSE, RESET, BOLD, payloadFlag, RESET)
	fmt.Printf("%s[-]%s No file specified, outputting raw payload.\n", CHARTREUSE, RESET)
	fmt.Printf("%s[-]%s Payload size: %d bytes\n", CHARTREUSE, RESET, len(result))
}

func OutputFile(payloadFlag string, result string) {
	ShellVenomBanner()
	fmt.Printf("%s[-]%s Selecting %s%s%s from the payload.\n", CHARTREUSE, RESET, BOLD, payloadFlag, RESET)
	fmt.Printf("%s[-]%s Payload size: %d bytes\n", CHARTREUSE, RESET, len(result))
}

func HostandPortErrorHandling(host, port bool) {
	if (!host && !port) {
		fmt.Printf("%s[-]%s %sLHOST%s %s&%s %sLPORT%s %svalue is not valid!%s\n", CRIMSON, RESET, BOLD, RESET, ITALIC, RESET, BOLD, RESET, ITALIC, RESET)
		os.Exit(0)
	} else if (!host && port) {
		fmt.Printf("%s[-]%s %sLHOST%s %svalue is not valid!%s\n", CRIMSON, RESET, BOLD, RESET, ITALIC, RESET)
		os.Exit(0)
	} else if (!port && host) {
		fmt.Printf("%s[-]%s %sLPORT%s %svalue is not valid!%s\n", CRIMSON, RESET, BOLD, RESET, ITALIC, RESET)
		os.Exit(0)
	}
}