package command

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type ivCommandExecutor struct {
	ivs map[string]*irregularVerb
}

func (h *ivCommandExecutor) Execute(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || strings.ContainsAny(value, " \n\r\t") {
		return "🤦‍♀️ Invalid irregular verb request."
	}

	iv, ok := h.ivs[value]
	if !ok {
		return fmt.Sprintf("🤔 Seems _\"%s\"_ is not irregular verb.", value)
	}

	return fmt.Sprintf(
		"💡\nBase form: _%s_\nPast Simple: _%s_\nPast Participle: _%s_",
		iv.v1,
		iv.v2,
		iv.v3,
	)
}

type irregularVerb struct {
	v1 string
	v2 string
	v3 string
}

func getIrregularVerbs() map[string]*irregularVerb {
	csvFile, err := os.Open("./data/iv.csv")
	if err != nil {
		log.Fatalln("can not load irregular verbs database")
	}

	result := make(map[string]*irregularVerb)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can't read CSV: %s", err.Error())
		}

		iv := irregularVerb{
			v1: values[0],
			v2: values[1],
			v3: values[2],
		}

		addSplittedValues := func(value string) {
			for _, forms := range strings.Split(value, "/") {
				for _, val := range strings.Split(forms, " ") {
					if !strings.ContainsAny(val, "[]") {
						result[strings.TrimSpace(val)] = &iv
					}
				}
			}
		}

		addSplittedValues(values[0])
		addSplittedValues(values[1])
		addSplittedValues(values[2])
	}

	return result
}
