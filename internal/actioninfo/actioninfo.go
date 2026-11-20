package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим данные

		err := dp.Parse(data)
		if err != nil {
			log.Printf("Error parsing '%s' : %v\n", data, err)
			continue
		}
		// Форматируем и выводим информацию

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Error calculating action info: %v\n", err)
			continue
		}

		fmt.Println(info)
	}
}
