package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"h1Bot/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	searchHandle := flag.String("handle", "", "Get a specific program scoped assets by program handle")
	onlyWithBounty := flag.Bool("withbounty", false, "Gets scoped assets for programs who offers bounties")
	onlyUrls := flag.Bool("urls", false, "Gets only the scoped urls from programs")

	flag.Parse()

	if *searchHandle != "" {
		program := getProgram(*searchHandle)
		var scope []string
		if *onlyUrls {
			for _, data := range program.Relationships.StructuredScopes.Data {
				if data.Attributes.AssetType == "URL" {
					if *onlyWithBounty {
						if data.Attributes.EligibleForBounty {
							scope = append(scope, data.Attributes.AssetIdentifier)
						}
					} else {
						scope = append(scope, data.Attributes.AssetIdentifier)
					}
				}
			}
		} else {
			for _, data := range program.Relationships.StructuredScopes.Data {
				if *onlyWithBounty {
					if data.Attributes.EligibleForBounty {
						scope = append(scope, data.Attributes.AssetIdentifier)
					}
				} else {
					scope = append(scope, data.Attributes.AssetIdentifier)
				}
			}
		}

		for _, url := range scope {
			fmt.Println(url)
		}
		os.Exit(0)
	}

	if *onlyWithBounty {
		next := ""
		for {
			var scope []string
			result := getPrograms(next)
			next = result.Links.Next
			for _, program := range result.Data {
				if program.Attributes.OffersBounties {
					fullProgram := getProgram(program.Attributes.Handle)
					if *onlyUrls {
						for _, data := range fullProgram.Relationships.StructuredScopes.Data {
							if data.Attributes.AssetType == "URL" && data.Attributes.EligibleForBounty {
								scope = append(scope, data.Attributes.AssetIdentifier)
							}
						}
					} else {
						for _, data := range fullProgram.Relationships.StructuredScopes.Data {
							if data.Attributes.EligibleForBounty {
								scope = append(scope, data.Attributes.AssetIdentifier)
							}
						}
					}
				}
			}

			for _, url := range scope {
				fmt.Println(url)
			}

			if result.Links.Next == "" {
				break
			}
		}
		os.Exit(0)
	}

}

func getPrograms(next string) utils.HackeronePrograms {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	url := ""
	if next != "" {
		url = next
	} else {
		url = "https://api.hackerone.com/v1/hackers/programs"
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header = headers
	req.SetBasicAuth("username", "token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	result := utils.HackeronePrograms{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	return result
}

func getProgram(handle string) utils.HackeroneProgram {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	url := "https://api.hackerone.com/v1/hackers/programs/" + handle

	req, err := http.NewRequest("GET", url, nil)
	req.Header = headers
	req.SetBasicAuth("username", "token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	result := utils.HackeroneProgram{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	return result
}
