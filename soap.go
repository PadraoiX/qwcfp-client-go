package soap

// package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*func main() {

	username := "98765432100"
	password := "33C3109AAA028CCB"
	dnsServer := "http://172.16.253.166:8080"
	rootConfig := "/home/youre/workspaceGo/src/github.com/yourepena/qwcfp-client-go/xml-conf/"

	str, erro := Login(username, password, dnsServer, rootConfig)

	if erro != nil {
		fmt.Printf("%s", erro)
	} else {
		fmt.Printf("%+v", str)
	}

}*/

/*func main() {

	username := "98765432100"
	password := "33C3109AAA028CCB"
	dnsServer := "http://172.16.253.60:8080"
	groupName := "SYSTK"
	rootConfig := "/home/youre/workspaceGo/src/github.com/yourepena/qwcfp-client-go/xml-conf/"

	cs, erroCs := New(username, password, dnsServer, rootConfig)

	if erroCs != nil {
		fmt.Printf("Aconteceu um erro realizando login, parando o programa.... %+v", erroCs)
	}

	page := 0
	rows := 10

	listFilesArray, errListFiles := cs.ListFiles(groupName, page, 1)

	if errListFiles != nil {
		fmt.Printf("Aconteceu um erro realizando login, parando o programa.... %+v", errListFiles)
	}

	if len(listFilesArray) > 0 {
		total := listFilesArray[0].TotalRecords

		fmt.Printf("Temos um total de %d arquivos por isso vamos ter %d lotes", total, total/rows)

		for j := 0; j < total/rows; j++ {

			fileVersionArray, errorR := cs.GetFilesFromQWCFP(groupName, j, rows)

			if errorR != nil {
				fmt.Printf("Aconteceu um erro realizando login, parando o programa.... %+v", errorR)
			}

			for i := 0; i < len(fileVersionArray); i++ {
				fmt.Printf("Lote %d -> Arquivo %d: %+v", (j + 1), (i + 1), fileVersionArray[i])
			}
		}
	}

}*/

/*func main() {

	loginKey := "ec9a60a0536ba3548225f71a3395649f"
	cs.dnsServer := "http://172.16.253.60:8080"
	rootConfig := "/home/youre/workspaceGo/src/github.com/yourepena/qwcfp-client-go/xml-conf/"
	systemGroup := "P001"
	groupSystemId := 0

	FileVersionId := "1098"

	//verificando se o sistema do arquivo possui um grupo para ele
	groupId, errG := GetGroup(systemGroup, loginKey, dnsServer, rootConfig)

	if errG != nil { //nao foi encontrado um grupo para esse sistema, procedendo com o cadastro do mesmo
		managerGroup, erroCrG := CreateNewGroup(1195, systemGroup, systemGroup, loginKey, dnsServer, rootConfig)
		if erroCrG != nil {
			fmt.Printf("Aconteceu criando o grupo do sistema %s, parando o Programa... %s", systemGroup, erroCrG)
		}
		groupSystemId = managerGroup.GroupId
	} else {
		groupSystemId = groupId
	}

	if groupSystemId > 0 { //Vamos tentar encontrar o subgrupo

		subgroupJobId := 0

		// for j := 0; j < len(outputarraylet); j++ {

		nome := "FDS00661_P001"

		subgroupId, errSG := GetGroup(nome, loginKey, dnsServer, rootConfig)

		if errSG != nil { // Nao encontramos, vamos criar um sub-grupo com o nome do Job

			jobGroup, erroCrSG := CreateNewGroup(groupSystemId, nome, nome, loginKey, dnsServer, rootConfig)

			if erroCrSG != nil {
				fmt.Printf("Aconteceu criando o grupo do sistema %s, parando o Programa... %s", nome, erroCrSG)

			}

			subgroupJobId = jobGroup.GroupId

		} else {
			subgroupJobId = subgroupId
		}

		// }

		if subgroupJobId > 0 {
			//Movendo o arquivo para o Grupo do Job Correspondente
			_, errorM := MoveFile(FileVersionId, strconv.Itoa(subgroupJobId), loginKey, dnsServer, rootConfig)

			if errorM != nil {
				fmt.Printf("Aconteu um erro movendo o arquivo para o grupo %d, parando o Programa... %s", subgroupJobId, errorM)

			}

			fmt.Printf("Sucesso muleque")
		}

	}

}*/

/* COMECO DO ENVELOPE DE NEWGROUP*/
type EnvelopeNewGroup struct {
	Body ManagerGroup `xml:"Body"`
}
type ManagerGroup struct {
	ManagerGroupResponse ManagerGroupResponse `xml:"ManagerGroupResponse>return"`
}

type ManagerGroupResponse struct {
	GroupId                 int       `xml:"groupId"`
	Apelido                 string    `xml:"apelido"`
	Nome                    string    `xml:"nome"`
	Description             string    `xml:"Description"`
	OwnerCreator            string    `xml:"ownerCreator"`
	CreationDate            time.Time `xml:"creationDate"`
	InputDateLimit          time.Time `xml:"inputDateLimit"`
	OutputDateLimit         time.Time `xml:"outputDateLimit"`
	SizeInBytes             int       `xml:"sizeInBytes"`
	DaysLimitDiscart        int       `xml:"daysLimitDiscart"`
	AceptVersion            string    `xml:"aceptVersion"`
	Status                  int       `xml:"status"`
	SubordinateGroup        int       `xml:"subordinateGroup"`
	ManagerGroupId          int       `xml:"managerGroupId"`
	ErrorCode               int       `xml:"errorCode"`
	ErrorMsg                string    `xml:"errorMsg"`
	Suporte1Id              int       `xml:"suporte1Id"`
	AreaId                  int       `xml:"areaId"`
	Suporte2Id              int       `xml:"suporte2Id"`
	FileSystemStorageDomain int       `xml:"fileSystemStorageDomain"`
	NotificationTypeDomain  int       `xml:"notificationTypeDomain"`
}

/*FIM DO ENVELOPE DE NEWGROUP*/

/* COMECO DO ENVELOPE DE MYGROUPS */
type EnvelopeListVersions struct {
	Body BodyListVersions `xml:"Body"`
}
type BodyListVersions struct {
	ListVersionsResponse []ListVersionsResponse `xml:"ListVersionsResponse>return"`
}
type ListVersionsResponse struct {
	Id                       int       `xml:"id"`
	SenderMemberIdFk         int       `xml:"senderMemberIdFk"`
	SenderMemberName         string    `xml:"senderMemberName"`
	AddInformation           string    `xml:"addInformation"`
	FileManagedIdFk          int       `xml:"fileManagedIdFk"`
	FileStatusDomIdFk        int       `xml:"fileStatusDomIdFk"`
	Version                  int       `xml:"version"`
	SizeInBytes              float32   `xml:"sizeInBytes"`
	ErrorCode                int       `xml:"errorCode"`
	FileStatusDomStringValue string    `xml:"fileStatusDomStringValue"`
	SourceFileFullPat        string    `xml:"sourceFileFullPat"`
	DateStatusChanged        time.Time `xml:"dateStatusChanged"`
	CreationDate             time.Time `xml:"creationDate"`
	QueueDate                time.Time `xml:"queueDate"`
	TransferDate             time.Time `xml:"transferDate"`
	ErrorMsg                 string    `xml:"errorMsg"`
	QwareObjId               string    `xml:"qwareObjId"`
}

/* FIM DO ENVELOPE DE MYGROUPS */

/* COMECO DO ENVELOPE DE DOWNLOAD */
type EnvelopeMove struct {
	Body BodyMove `xml:"Body"`
}
type BodyMove struct {
	ManipulateFileResponse ManipulateFileResponse `xml:"ManipulateFileResponse>return"`
}
type ManipulateFileResponse struct {
	ErrorCode int    `xml:"errorCode"`
	ErrorMsg  string `xml:"errorMsg"`
}

/* FIM DO ENVELOPE DE DOWNLOAD */

/* COMECO DO ENVELOPE DE DOWNLOAD */
type EnvelopeDownload struct {
	Body BodyDownload `xml:"Body"`
}
type BodyDownload struct {
	DownloadResponse DownloadResponse `xml:"DownloadResponse>return"`
}
type DownloadResponse struct {
	FileName      string `xml:"fileName"`
	FileType      string `xml:"fileType"`
	Server        string `xml:"server"`
	Port          int    `xml:"port"`
	SaveAs        string `xml:"saveAs"`
	ObjId         string `xml:"objId"`
	QwareUserName string `xml:"qwareUserName"`
	VersionId     int    `xml:"versionId"`
	QwareUserPass string `xml:"qwareUserPass"`
	ErrorCode     int    `xml:"errorCode"`
	ErrorMsg      string `xml:"errorMsg"`
}

/* FIM DO ENVELOPE DE DOWNLOAD */

/* COMECO DO ENVELOPE DE MYGROUPS */
type EnvelopeMyGroups struct {
	Body BodyMyGroups `xml:"Body"`
}
type BodyMyGroups struct {
	MyGroupsResponses []MyGroupsResponse `xml:"MyGroupsResponse>return"`
}
type MyGroupsResponse struct {
	GroupId                 int       `xml:"groupId"`
	Apelido                 string    `xml:"apelido"`
	Nome                    string    `xml:"nome"`
	Description             string    `xml:"description"`
	OwnerCreator            string    `xml:"ownerCreator"`
	CreationDate            time.Time `xml:"creationDate"`
	InputDateLimit          time.Time `xml:"inputDateLimit"`
	OutputDateLimit         time.Time `xml:"outputDateLimit"`
	SizeInBytes             string    `xml:"sizeInBytes"`
	DaysLimitDiscart        int       `xml:"daysLimitDiscart"`
	AceptVersion            string    `xml:"aceptVersion"`
	Status                  int       `xml:"status"`
	SubordinateGroup        int       `xml:"subordinateGroup"`
	ManagerGroupId          int       `xml:"managerGroupId"`
	Suporte1Id              int       `xml:"suporte1Id"`
	AreaId                  int       `xml:"areaId"`
	Suporte2Id              int       `xml:"suporte2Id"`
	FileSystemStorageDomain int       `xml:"fileSystemStorageDomain"`
	NotificationTypeDomain  int       `xml:"notificationTypeDomain"`
}

/* FIM DO ENVELOPE DE MYGROUPS */

/* COMECO DO ENVELOPE DE LISTFILES */
type EnvelopeListFiles struct {
	Body BodyListFiles `xml:"Body"`
}
type BodyListFiles struct {
	ListFilesResponse []ListFilesResponse `xml:"ListFilesResponse>return"`
}
type ListFilesResponse struct {
	Fileid            int       `xml:"fileid"`
	FileName          string    `xml:"fileName"`
	ExtensionName     string    `xml:"extensionName"`
	OwnerMemberId     int       `xml:"ownerMemberId"`
	OwnerMemberName   string    `xml:"ownerMemberName"`
	ErrorCode         int       `xml:"errorCode"`
	ErrorMsg          string    `xml:"errorMsg"`
	DataLimitValidate time.Time `xml:"dataLimitValidate"`
	CreationDate      time.Time `xml:"creationDate"`
	HashName          string    `xml:"hashName"`
	TotalRecords      int       `xml:"totalRecords"`
	IconeName         string    `xml:"iconeName"`
}

/* FIM DO ENVELOPE DE LISTFILES */
/* COMEÇO DO ENVELOPE DE LOGIN */
type EnvelopeLogin struct {
	Body BodyLogin `xml:"Body"`
}
type BodyLogin struct {
	LoginResponse LoginResponse `xml:"LoginResponse"`
}
type LoginResponse struct {
	LoginDTO LoginDTO `xml:"return"`
}
type LoginDTO struct {
	MemberId          int       `xml:"memberId"`
	QwarePasswordEnc  string    `xml:"qwarePasswordEnc"`
	MemberName        string    `xml:"memberName"`
	MemberEmail       string    `xml:"memberEmail"`
	AreaCodePhone1    int       `xml:"areaCodePhone1"`
	NumberPhone1      string    `xml:"numberPhone1"`
	PhisicalAddress   string    `xml:"phisicalAddress"`
	ZipCode           string    `xml:"zipCode"`
	LoginCpfId        string    `xml:"loginCpfId"`
	DataCreation      time.Time `xml:"dataCreation"`
	StatusMemberDomFk int       `xml:"statusMemberDomFk"`
	AreaIdFk          int       `xml:"areaIdFk"`
	QwareUser         string    `xml:"qwareUser"`
	ErrorCode         int       `xml:"errorCode"`
	LoginKey          string    `xml:"loginKey"`
	ErrorMsg          string    `xml:"errorMsg"`
}

/* FIM DO ENVELOPE DE LOGIN */

type FileVersionRetorno struct {
	Path          string
	FileVersionId int
}

type ClientSoap struct {
	dnsServer  string
	rootConfig string
	loginKey   string
}

func New(username, password, dnsServer, rootConfig string) (*ClientSoap, error) {

	loginKey, erro := Login(username, password, dnsServer, rootConfig)

	if erro != nil || loginKey == "" {
		err := errors.New(fmt.Sprintf("Houve um erro tetando realizar o login: %+v", erro))
		return &ClientSoap{}, err
	}

	return &ClientSoap{
		dnsServer:  dnsServer,
		rootConfig: rootConfig,
		loginKey:   loginKey,
	}, nil

}

func Login(username, password, dnsServer, rootConfig string) (string, error) {

	r := EnvelopeLogin{}

	tagsName := map[string]string{
		"login":  username,
		"esenha": password,
	}

	bh, err := populateXML(rootConfig, "Login.xml", tagsName)

	b, err := doRequest(dnsServer+"/qwcfpWebService/Login?wsdl", bh)
	if err != nil {
		return "", err
	}

	err = xml.Unmarshal(b, &r)

	if err != nil {
		return "", err
	}

	loginKey := r.Body.LoginResponse.LoginDTO.LoginKey

	if len(loginKey) > 0 {
		return loginKey, nil
	} else {
		err := errors.New(fmt.Sprintf("Error code:%+v\nError message:%+v", r.Body.LoginResponse.LoginDTO.ErrorCode, r.Body.LoginResponse.LoginDTO.ErrorMsg))
		return "", err
	}

}

func (cs *ClientSoap) GetGroup(groupName string) (int, error) {

	mb := EnvelopeMyGroups{}

	groupFile := 0

	tagsName := map[string]string{
		"loginKey": cs.loginKey,
	}

	bh, err := populateXML(cs.rootConfig, "MyGroups.xml", tagsName)

	if err != nil {
		return groupFile, err
	}

	b, err := doRequest(cs.dnsServer+"/qwcfpWebService/MyGroups?wsdl", bh)
	if err != nil {
		return groupFile, err
	}

	err = xml.Unmarshal(b, &mb)

	if err != nil {
		return groupFile, err
	}

	for i := 0; i < len(mb.Body.MyGroupsResponses); i++ {
		myGroup := mb.Body.MyGroupsResponses[i]

		if myGroup.Apelido == groupName {
			groupFile = myGroup.GroupId
		}
	}

	if groupFile == 0 {
		err := errors.New(fmt.Sprintf("O grupo %s nao foi encontrado", groupName))
		return groupFile, err
	}

	return groupFile, nil

}

func (cs *ClientSoap) CreateNewGroup(subordinatedGroupId int, apelido string, groupName string) (ManagerGroupResponse, error) {

	mo := EnvelopeNewGroup{}

	tagsName := map[string]string{
		"subordinatedGroupId": strconv.Itoa(subordinatedGroupId),
		"apelido":             apelido,
		"groupName":           groupName,
		"loginKey":            cs.loginKey,
	}

	bh, err := populateXML(cs.rootConfig, "ManagerGroup.xml", tagsName)

	if err != nil {
		return ManagerGroupResponse{}, err
	}

	b, err := doRequest(cs.dnsServer+"/qwcfpWebService/ManagerGroup?wsdl", bh)
	if err != nil {
		return ManagerGroupResponse{}, err
	}

	err = xml.Unmarshal(b, &mo)

	if err != nil {
		return ManagerGroupResponse{}, err
	}

	if mo.Body.ManagerGroupResponse.ErrorCode != 0 {
		errE := errors.New(fmt.Sprintf("%d: %s\n", mo.Body.ManagerGroupResponse.ErrorCode, mo.Body.ManagerGroupResponse.ErrorMsg))
		return ManagerGroupResponse{}, errE
	}

	return mo.Body.ManagerGroupResponse, nil

	//fmt.Printf("Mas que porra e essa: %+v\n\n\n\n", mo.Body.ManagerGroupResponse[0])

}

func (cs *ClientSoap) MoveFile(fileid, groupid int) (ManipulateFileResponse, error) {

	mk := EnvelopeMove{}

	tagsName := map[string]string{
		"loginKey":      cs.loginKey,
		"idFileVersion": strconv.Itoa(fileid),
		"groupTo":       strconv.Itoa(groupid),
	}

	bh, err := populateXML(cs.rootConfig, "ManipulateFile.xml", tagsName)

	b, err := doRequest(cs.dnsServer+"/qwcfpWebService/ManipulateFile?wsdl", bh)
	if err != nil {
		return ManipulateFileResponse{}, err
	}

	/*stringhss := CToGoString(b[:])
	fmt.Println(stringhss)*/

	err = xml.Unmarshal(b, &mk)

	if err != nil {
		return ManipulateFileResponse{}, err
	}

	if mk.Body.ManipulateFileResponse.ErrorCode != 0 {
		errE := errors.New(fmt.Sprintf("%d: %s\n", mk.Body.ManipulateFileResponse.ErrorCode, mk.Body.ManipulateFileResponse.ErrorMsg))
		return ManipulateFileResponse{}, errE
	}

	return mk.Body.ManipulateFileResponse, nil

}

/*
groupName - Alias do Grupo
dnsServer - DNS do servidor do QWCFP sem barra no final
rootConfig -  Caminho dos XMLS de configuracao
*/
func (cs *ClientSoap) GetFilesFromQWCFP(groupName string, pager, records int) ([]DownloadResponse, error) {

	if len(cs.loginKey) > 0 {
		//Listando todos os arquivos do grupo
		listFilesArray, errListFiles := cs.ListFiles(groupName, pager, records)

		if errListFiles != nil {
			return nil, errListFiles
		}

		//	fmt.Printf("Encontrados %d arquivos no grupo %s\n", len(listFilesArray), groupName)

		/*	var names = []string{}
			var filesid = []int{}*/

		if len(listFilesArray) > 0 {

			var downArray = []DownloadResponse{}

			for i := 0; i < len(listFilesArray); i++ {

				FileNameEx := listFilesArray[i].FileName
				// names = append(names, FileNameEx)

				FileNameId := listFilesArray[i].Fileid

				// filesid = append(filesid, FileNameId)

				//Download do arquivo

				versions, errorHapp := cs.ListVersions(FileNameId)

				if errorHapp != nil {
					fmt.Printf("Ocorreu um erro pegando a versao do arquivo %+v", errorHapp)
					continue
				}

				//fmt.Printf("%+v", versions)

				for k := 0; k < len(versions); k++ {

					down, errorHapp := cs.Download(groupName, FileNameEx, versions[k])

					if errorHapp != nil {
						fmt.Printf("Ocorreu um erro fazendo download do arquivo %s, no grupo %s, na versao %+v, \n\n Erro:\n %+v", FileNameEx, groupName, versions[k], errorHapp)
						continue
					}

					if down.ErrorCode != 0 {
						errorHapp = errors.New(fmt.Sprintf("%d: %s\n", down.ErrorCode, down.ErrorMsg))
						fmt.Printf("Ocorreu um erro fazendo download do arquivo %s, no grupo %s, na versao %+v, \n\n Erro:\n %+v", FileNameEx, groupName, versions[k], errorHapp)
						continue
					}

					if down != (DownloadResponse{}) {

						downArray = append(downArray, down)
					}

				}

			}

			/*fmt.Println("Quantidade de arquivos %d", len(names))
			fmt.Println("Quantidade de versoes%d", len(downArray))

			var retorno = []FileVersionRetorno{}

			for i := 0; i < len(names); i++ {
				for u := 0; u < len(downArray); u++ {

					fvr := FileVersionRetorno{
						FileName:      names[i],
						Path:          downArray[u].FileName,
						FileVersionId: downArray[u].VersionId,
						FileId:        filesid[i],
					}

					retorno = append(retorno, fvr)
				}
			}*/

			return downArray, nil

		} else {
			err := errors.New("Group empty")
			return nil, err
		}
	} else {
		err := errors.New("LoginKey empty")
		return nil, err
	}

}

func (cs *ClientSoap) ListVersions(fileid int) ([]int, error) {

	mj := EnvelopeListVersions{}

	tagsName := map[string]string{
		"loginKey": cs.loginKey,
		"fileId":   strconv.Itoa(fileid),
	}

	bh, err := populateXML(cs.rootConfig, "ListVersions.xml", tagsName)

	b, err := doRequest(cs.dnsServer+"/qwcfpWebService/ListVersions?wsdl", bh)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(b, &mj)

	if err != nil {
		return nil, err
	}

	/*stringhss := CToGoString(b[:])
	fmt.Println(stringhss)*/

	var versions = []int{}

	for j := 0; j < len(mj.Body.ListVersionsResponse); j++ {
		versions = append(versions, mj.Body.ListVersionsResponse[j].Version)
	}

	return versions, nil

}

func (cs *ClientSoap) ListFiles(groupName string, pager, records int) ([]ListFilesResponse, error) {

	mx := EnvelopeListFiles{}

	tagsName := map[string]string{
		"loginKey": cs.loginKey,
		"group":    groupName,
		"pager":    strconv.Itoa(pager),
		"records":  strconv.Itoa(records),
	}

	bh, err := populateXML(cs.rootConfig, "ListFiles.xml", tagsName)

	b, err := doRequest(cs.dnsServer+"/qwcfpWebService/ListFiles?wsdl", bh)
	if err != nil {
		return nil, err
	}

	/*	stringhss := CToGoString(b[:])
		fmt.Println(stringhss)
	*/
	err = xml.Unmarshal(b, &mx)

	if err != nil {
		return nil, err
	}

	return mx.Body.ListFilesResponse, nil

}

func (cs *ClientSoap) Download(groupName string, FileNameEx string, version int) (DownloadResponse, error) {

	mr := EnvelopeDownload{}

	tagsName := map[string]string{
		"loginKey":      cs.loginKey,
		"fileName":      FileNameEx,
		"versionNumber": strconv.Itoa(version),
		"infoGroup":     groupName,
	}

	bh, err := populateXML(cs.rootConfig, "Download.xml", tagsName)

	bvg, err := doRequest(cs.dnsServer+"/qwcfpWebService/Download?wsdl", bh)
	if err != nil {
		return DownloadResponse{}, err
	}

	err = xml.Unmarshal(bvg, &mr)

	if err != nil {
		return DownloadResponse{}, err
	}

	return mr.Body.DownloadResponse, nil
}

func doRequest(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req.ContentLength = int64(len(body))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func parseFileToArraY(fileName string) ([]string, error) {

	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	reader := bufio.NewReader(xmlFile)
	buffer := bytes.NewBuffer(make([]byte, 0))

	var chunk []byte
	var eol bool
	var str_array []string

	for {
		if chunk, eol, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(chunk)
		if !eol {
			str_array = append(str_array, buffer.String())
			buffer.Reset()
		}
	}

	if err == io.EOF {
		err = nil
	}

	return str_array, err
}

func populateXML(rootConfig string, fileName string, tagsName map[string]string) ([]byte, error) {

	str_arrayEx, err := parseFileToArraY(rootConfig + fileName)

	if err != nil {
		return nil, err
	}

	var xml []string

	for j := 0; j < len(str_arrayEx); j++ {
		line := str_arrayEx[j]

		line = strings.TrimSpace(line)

		for k, v := range tagsName {
			if line == "<"+k+"></"+k+">" {
				line = "<" + k + ">" + v + "</" + k + ">"
			}
		}

		xml = append(xml, line)
	}

	var bk = []byte{}

	for u := 0; u < len(xml); u++ {
		b := []byte(xml[u])
		for op := 0; op < len(b); op++ {
			bk = append(bk, b[op])
		}
	}

	return bk, nil
}
