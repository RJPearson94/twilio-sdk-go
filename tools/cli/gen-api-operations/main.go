package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {
	fixture, _ := ioutil.ReadFile("../../../service/studio/v2/temp/apiOperations.json")
	apiOperations := make([]apiOperation, 0)
	json.Unmarshal(fixture, &apiOperations)

	r := strings.NewReplacer("{ifFormEncodedDataAddTags}", ifFormEncodedDataAddTags, "{ifJSONResponseAddTags}", ifJSONResponseAddTags)
	parsedTemplate := template.Must(template.New("generateAPIOperations").Parse(r.Replace(apiOperationContent)))

	for _, value := range apiOperations {
		file, err := os.Create(fmt.Sprintf("../../../service/studio/v2/temp/api_op_%s_%s.go", strings.ToLower(value.Service), strings.ToLower(value.Name)))
		defer file.Close()
		if err != nil {
			fmt.Printf("Unable to create file on disk. %s", err)
			return
		}

		if err := parsedTemplate.Execute(file, value); err != nil {
			fmt.Printf("Unable to write to file. %s", err)
			return
		}
	}

}
