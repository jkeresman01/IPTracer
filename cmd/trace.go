package cmd

import (
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
	"io/ioutil"
	"encoding/json"
	"log"
)

type IpInfo struct{
	Country 	string		`json:"country"`
	Region 		string		`json:"region"`
	City		string		`json:"city"`
	Timezone	string		`json:"timezone"`
}

type IPInfoPrinter interface{
	printIPInfo()
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP address",
	Long: `Trace the provided IP adrress, and display information regarding, country, region, city and timezone`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			displayIpInfo(args[0])	
		} else {
			fmt.Println("Please provide IP address to trace!")
		}
	},
}

func (ip *IpInfo) printIPInfo(){
	fmt.Println("--------------- Location -------------------")
	fmt.Printf("Country: %s \nCity: %s\nRegion: %s\nTimezone: %s\n", ip.Country, ip.City, ip.Region, ip.Timezone)
	fmt.Println("--------------------------------------------")
}

func displayIpInfo(ipAdrress string){
	var ipInfo IpInfo
	err := json.Unmarshal(getIpInfo(ipAdrress), &ipInfo)

	if err != nil{
		log.Printf("Unable to unmarshall data for IP: %s", ipAdrress)
	}

	ipInfo.printIPInfo()
}


func getIpInfo(ipAdrress string) []byte{
	url := fmt.Sprintf("http://ipinfo.io/%s", ipAdrress)

	res, err := http.Get(url)
	if err != nil{
		log.Printf("Unable to get IP data for: %s", ipAdrress)
	}
	defer res.Body.Close()


	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Printf("Unable to read the response for ip adrress: %s", ipAdrress)
	}

	return body
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
