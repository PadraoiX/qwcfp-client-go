# [QWCFP](https://qwcfp.pix.com.br) GOLANG CLIENT 

Welcome to [QWCFP](https://qwcfp.pix.com.br) GOLANG CLIENT 


Ensure that this folder is at the following location:

`${GOPATH}/src/github.com/yourepena/qwcfp-client-go`

## Getting Started with [QWCFP](https://qwcfp.pix.com.br) GOLANG CLIENT 

### Requirements

* [Golang](https://golang.org/dl/) 1.7


To push QWCFP GOLANG CLIENT  in the git repository, run the following commands:
```
git remote set-url origin https://github.com/yourepena/qwcfp-client-go.git
git push origin master
```

### Exemple
```go
package main

import (
	"fmt"
	"github.com/yourepena/qwcfp-client-go"
	"strconv"
)

func main() {

	dnsServer := "http://172.16.253.108:8080"

	rootConfig := "/home/youre/workspaceGo/src/github.com/yourepena/sysoutjobbeat/soap/"

	groupName := "GFNS"

	loginKey, err := soap.Login(dnsServer, rootConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	groupid, errG := soap.GetGroup("GFNS_PROCESSADO", loginKey, dnsServer, rootConfig)

	if errG != nil {
		fmt.Println(errG)
	}

	fileVersionArray, errF := soap.GetFilesFromQWCFP(loginKey, groupName, dnsServer, rootConfig)

	if errF != nil {
		fmt.Println(errF)
		return
	}

	for u := 0; u < len(fileVersionArray); u++ {

		FileName := fileVersionArray[u].FileName
		Path := fileVersionArray[u].Path
		FileVersionId := fileVersionArray[u].FileVersionId
		FileId := fileVersionArray[u].FileId

		fmt.Printf("FileName: %s\nPath: %s\nFileVersionId: %d\nFileId: %d\n\n\n", FileName, Path, FileVersionId, FileId)

		//TODO Processamento

		_, errC := soap.MoveFile(strconv.Itoa(FileVersionId), strconv.Itoa(groupid), loginKey, dnsServer, rootConfig)

		/*if sucesso {
			fmt.Printf("%s: [%s] FileID:%s\n ", k, v, id)
		}*/

		if errC != nil {
			fmt.Println(errC)
		}
	}
}
```
