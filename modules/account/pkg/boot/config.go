package boot

import (
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func Load[T any](path string, file string, section string) *T {

	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error on read: %s\n", err)
	}

	log.Printf("Config loaded: %s\n", file)

	var config T
	if err := viper.UnmarshalKey(section, &config); err != nil {
		log.Panicf("Error on map: %s\n", err)
	}

	replaceEnvVar(&config)

	log.Printf("Config section: %s\n", section)

	return &config
}

// TODO: url do postgresql
// ${VAR:default}
func getEnvVar(value string) string {
	re := regexp.MustCompile(`\${([^:}]+):([^}]+)}`)

	return re.ReplaceAllStringFunc(value, func(match string) string {
		parts := strings.Split(match[2:len(match)-1], ":")
		envVar := parts[0]
		defaultVal := parts[1]

		if value, exists := os.LookupEnv(envVar); exists {
			return value
		}

		return defaultVal
	})
}

func replaceEnvVar[T any](config *T) {
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.String {
			currentValue := field.String()
			field.SetString(getEnvVar(currentValue))
		}
	}
}
