package i18n

import (
	"encoding/json"
	"fmt"
	"os"
)

var data map[string]map[string]string
var currentLang = "en"

func Load(lang string) error {
	file, err := os.ReadFile("internal/i18n/lang.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}
	if _, ok := data[lang]; ok {
		currentLang = lang
	} else {
		currentLang = "en"
	}
	return nil
}

func T(key string) string {
	if val, ok := data[currentLang][key]; ok {
		return val
	}
	return fmt.Sprintf("??%s??", key)
}
