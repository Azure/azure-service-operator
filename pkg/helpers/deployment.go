package helpers

import (
	"bytes"
	"text/template"
)

// IsDeploymentComplete will dtermine if the deployment is complete
func IsDeploymentComplete(status string) bool {
	switch status {
	case "Succeeded":
		return true
	case "Failed":
		return true
	case "Canceled":
		return true
	}
	return false
}

// Templatize returns the proper values based on the templating
func Templatize(tempStr string, data interface{}) (resp string, err error) {
	t := template.New("templating")
	t, err = t.Parse(string(tempStr))
	if err != nil {
		return
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	return tpl.String(), err
}

func GetOutput(outputs interface{}, key string) string {
	return outputs.(map[string]interface{})[key].(map[string]interface{})["value"].(string)
}
