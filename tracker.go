package tracker

import (
	"bufio"
	"bytes"
	/*"encoding/json"*/
	"fmt"
	"github.com/vjeantet/jodaTime"
	"log"
	"os"
	"strconv"
	"time"
)

type OutPut struct {
	Identificacao string
	Versao        int
	Jobname       string
	Jobid         string
	Procname      string
	Stepname      string
	DDname        string
	SysoutClass   string
	Destination   string
	Form          string
	Userid        string
	Lines         int
	Dataextracao  time.Time
	Dataexecucao  time.Time
	Account       string
}

func tracker(fullPath string) OutPut {

	// fullPath := "/home/youre/workspaceGo/src/github.com/yourepena/sysoutjobbeat/SYS2.FDS.SYSTKOUT.G0001V00"

	file, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if len(line) > 8 {
			if line[0:8] == "FDSSYSTK" {

				var (
					dataextracao time.Time
					dataexecucao time.Time
					versao       int
					qtdLine      int
					account      string
				)

				if len(line) > 95 {

					var buffer bytes.Buffer

					buffer.WriteString(line[87:89]) //dia
					buffer.WriteString("/")
					buffer.WriteString(line[85:87]) // mes
					buffer.WriteString("/")
					buffer.WriteString(line[81:85]) // ano
					buffer.WriteString(":")
					buffer.WriteString(line[89:91]) // hora
					buffer.WriteString(":")
					buffer.WriteString(line[91:93]) // minuto
					buffer.WriteString(":")
					buffer.WriteString(line[93:95]) // segundo

					dataextracao, _ = jodaTime.Parse("dd/MM/yyyy:HH:mm:ss", buffer.String())

				}

				if len(line) > 109 {

					var buffer bytes.Buffer

					buffer.WriteString(line[101:103]) //dia
					buffer.WriteString("/")
					buffer.WriteString(line[99:101]) // mes
					buffer.WriteString("/")
					buffer.WriteString(line[97:99]) // ano
					buffer.WriteString(":")
					buffer.WriteString(line[103:105]) // hora
					buffer.WriteString(":")
					buffer.WriteString(line[105:107]) // minuto
					buffer.WriteString(":")
					buffer.WriteString(line[107:109]) // segundo

					dataexecucao, _ = jodaTime.Parse("dd/MM/yyyy:HH:mm:ss", buffer.String())

				}

				identicacao := getValue(line, 0, 8)

				if getValue(line, 8, 10) != "" {
					versao, _ = strconv.Atoi(getValue(line, 8, 10))
				}

				jobname := getValue(line, 10, 18)
				jobdid := getValue(line, 18, 26)
				procname := getValue(line, 26, 34)
				stepname := getValue(line, 34, 42)
				ddname := getValue(line, 42, 50)
				sysout := getValue(line, 50, 51)
				destination := getValue(line, 51, 59)
				form := getValue(line, 59, 63)
				userid := getValue(line, 63, 71)

				if getValue(line, 71, 81) != "" {
					qtdLine, _ = strconv.Atoi(getValue(line, 71, 81))
				}

				if len(line) > 113 {
					account = getValue(line, 113, len(line))
				}

				m := OutPut{
					Identificacao: identicacao,
					Versao:        versao,
					Jobname:       jobname,
					Jobid:         jobdid,
					Procname:      procname,
					Stepname:      stepname,
					DDname:        ddname,
					SysoutClass:   sysout,
					Destination:   destination,
					Form:          form,
					Userid:        userid,
					Lines:         qtdLine,
					Dataextracao:  dataextracao,
					Dataexecucao:  dataexecucao,
					Account:       account,
				}

				return m

				/*				as := time.Time{}

													if m.Dataextracao == as {

														fmt.Println(m.Dataextracao)
													}

									b, err := json.Marshal(m.Dataextracao)

									if err != nil {
										log.Fatal(err)
								 }

								fmt.Println(string(b)) */
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getValue(line string, positionStart int, positionEnd int) string {

	if len(line) >= positionEnd {

		return line[positionStart:positionEnd]
	} else {

		return ""
	}

}
