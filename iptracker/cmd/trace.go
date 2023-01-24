package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

//data showing
// {
//     "ip": "103.124.132.64",
//     "city": "Belmont",
//     "region": "Auckland",
//     "country": "NZ",
//     "loc": "-36.8075,174.7864",
//     "postal": "0622",
//     "timezone": "Pacific/Auckland",
//     "readme": "https://ipinfo.io/missingauth"
// }

type IP struct {
	Ip string `json:"ip"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	Location string `json:"loc"`
	Postal string `json:"postal"`
	Timezone string `json:"timezone"`
}




// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long: `Tracing the IP addresses.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please Provide an IP address to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

func showData(i string) {
	url := "http://ipinfo.io/"+ i +"/geo"

	res := getData(url)

	data := IP{}

	err := json.Unmarshal(res, &data)
	if err != nil {
		log.Println("Unable to unmarshal data")
	}

	fmt.Println("DATA:")

	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nPOSTAL:%s\nTIMEZONE:%s\n", data.Ip, data.City, data.Region, data.Country, data.Location, data.Postal, data.Timezone)
}

func getData(url string) []byte{
	response, err := http.Get(url)

	if err != nil {
		log.Println("Unable to Response data")
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Println("Unable to Read response data")
	}

	return res
}

