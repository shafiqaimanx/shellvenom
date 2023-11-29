package core

import (
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

func DecodeBase64(encodedTemplatePayload string) string{
	sd, err := base64.StdEncoding.DecodeString(encodedTemplatePayload)
	if err != nil {
		fmt.Println(err)
	}
	return string(sd)
}

func FindAvailablePayloadInterface(payload string, lport string, iteration int, ipList []net.IP) string {
	c := strings.Replace(payload, "CHANGE_LHOST", IPToStringSlice(ipList)[iteration], 1)
	result := strings.Replace(c, "CHANGE_LPORT", lport, 1)
	return result
}

func FindAvailablePayloadIP(payload string, lhost string, lport string) string {
	c := strings.Replace(payload, "CHANGE_LHOST", lhost, 1)
	result := strings.Replace(c, "CHANGE_LPORT", lport, 1)
	return result
}

//func FindAvailablePayload(payloadFlag string) string {
//	var result string
//	payloadFound := false
//
//	for _, item := range ReadFileofShell().ShelListz {
//		payloadName := fmt.Sprintf("%s/%s", item.Platform, item.Name)
//		if payloadFlag == payloadName {
//			result = DecodeBase64(item.Shell)
//			payloadFound = true
//			break
//		}
//	}
//	if !payloadFound {
//		fmt.Printf("%s payload not found (%s%s%s)", colors.ERROR, colors.RED, payloadFlag, colors.RESET)
//	}
//	return result
//}