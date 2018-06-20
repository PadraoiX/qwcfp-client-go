package soap

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
	"strings"
	"time"
)

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

var (
	r  EnvelopeLogin
	mx EnvelopeListFiles
	mr EnvelopeDownload
	mb EnvelopeMyGroups
	mk EnvelopeMove
)

type FileVersionRetorno struct {
	FileName      string
	Path          string
	FileVersionId int
	FileId        int
}

/*
groupName - Alias do Grupo
dnsServer - DNS do servidor do QWCFP sem barra no final
rootConfig -  Caminho dos XMLS de configuracao
*/
func GetFilesFromQWCFP(loginKey string, groupName string, dnsServer string, rootConfig string) ([]FileVersionRetorno, error) {

	if len(loginKey) > 0 {
		//Listando todos os arquivos do grupo
		listFilesArray, errListFiles := ListFiles(groupName, loginKey, dnsServer, rootConfig)

		if errListFiles != nil {
			return nil, errListFiles
		}

		var names = []string{}
		var filesid = []int{}

		if len(listFilesArray) > 0 {

			var downArray = []DownloadResponse{}

			for i := 0; i < len(listFilesArray); i++ {

				FileNameEx := listFilesArray[i].FileName
				FileNameId := listFilesArray[i].Fileid

				names = append(names, FileNameEx)
				filesid = append(filesid, FileNameId)

				//Download do arquivo
				down, errDow := Download(groupName, loginKey, FileNameEx, dnsServer, rootConfig)

				if errDow != nil {
					return nil, errDow
				}

				if down.ErrorCode != 0 {
					err := errors.New(fmt.Sprintf("%d: %s\n", down.ErrorCode, down.ErrorMsg))
					return nil, err
				}

				downArray = append(downArray, down)
			}

			var retorno = []FileVersionRetorno{}

			for u := 0; u < len(downArray); u++ {

				fvr := FileVersionRetorno{
					FileName:      names[u],
					Path:          downArray[u].FileName,
					FileVersionId: downArray[u].VersionId,
					FileId:        filesid[u],
				}

				retorno = append(retorno, fvr)
			}

			return retorno, nil

		} else {
			err := errors.New("Group empty!")
			return nil, err
		}
	} else {
		err := errors.New("LoginKey empty!")
		return nil, err
	}

}

func GetGroup(groupName string, loginKey string, dnsServer string, rootConfig string) (int, error) {

	str_array, err := parseFileToArraY(rootConfig + "MyGroups.xml")

	var groupFile = 0

	if err != nil {
		return groupFile, err
	}

	var listFilesXML []string

	for i := 0; i < len(str_array); i++ {
		line := str_array[i]

		line = strings.TrimSpace(line)

		if line == "<loginKey></loginKey>" {
			line = "<loginKey>" + loginKey + "</loginKey>"
		}

		listFilesXML = append(listFilesXML, line)
	}

	var bh = []byte{}

	for i := 0; i < len(listFilesXML); i++ {
		b := []byte(listFilesXML[i])
		for j := 0; j < len(b); j++ {
			bh = append(bh, b[j])
		}
	}

	b, err := doRequest(dnsServer+"/qwcfpWebService/MyGroups?wsdl", bh)
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

	return groupFile, nil

}

func MoveFile(fileid string, groupid string, loginKey string, dnsServer string, rootConfig string) (bool, error) {

	str_array, err := parseFileToArraY(rootConfig + "ManipulateFile.xml")

	var sucesso = false

	if err != nil {
		return sucesso, err
	}

	var listFilesXML []string

	for i := 0; i < len(str_array); i++ {
		line := str_array[i]

		line = strings.TrimSpace(line)

		if line == "<loginKey></loginKey>" {
			line = "<loginKey>" + loginKey + "</loginKey>"
		}

		if line == "<idFileVersion></idFileVersion>" {
			line = "<idFileVersion>" + fileid + "</idFileVersion>"
		}

		if line == "<groupTo></groupTo>" {
			line = "<groupTo>" + groupid + "</groupTo>"
		}

		listFilesXML = append(listFilesXML, line)
	}

	var bh = []byte{}

	//fmt.Println(listFilesXML)

	for i := 0; i < len(listFilesXML); i++ {
		b := []byte(listFilesXML[i])
		for j := 0; j < len(b); j++ {
			bh = append(bh, b[j])
		}
	}

	b, err := doRequest(dnsServer+"/qwcfpWebService/ManipulateFile?wsdl", bh)
	if err != nil {
		return sucesso, err
	}

	/*stringhss := CToGoString(b[:])
	fmt.Println(stringhss)*/

	err = xml.Unmarshal(b, &mk)

	if err != nil {
		return sucesso, err
	}

	//fmt.Printf("%d: %s\n", mk.Body.ManipulateFileResponse.ErrorCode, mk.Body.ManipulateFileResponse.ErrorMsg)

	return mk.Body.ManipulateFileResponse.ErrorCode == 0, nil

}

func Login(dnsServer string, rootConfig string) (string, error) {

	str_array, err := parseFileToArraY(rootConfig + "Login.xml")

	if err != nil {
		return "", err
	}

	var bs = []byte{}

	for i := 0; i < len(str_array); i++ {
		b := []byte(str_array[i])
		for j := 0; j < len(b); j++ {
			bs = append(bs, b[j])
		}
	}

	b, err := doRequest(dnsServer+"/qwcfpWebService/Login?wsdl", bs)
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

func ListFiles(groupName string, loginKey string, dnsServer string, rootConfig string) ([]ListFilesResponse, error) {

	str_array, err := parseFileToArraY(rootConfig + "ListFiles.xml")

	if err != nil {
		return nil, err
	}

	var listFilesXML []string

	for i := 0; i < len(str_array); i++ {
		line := str_array[i]

		line = strings.TrimSpace(line)

		if line == "<loginKey></loginKey>" {
			line = "<loginKey>" + loginKey + "</loginKey>"
		}

		if line == "<group></group>" {
			line = "<group>" + groupName + "</group>"
		}

		listFilesXML = append(listFilesXML, line)
	}

	var bh = []byte{}

	for i := 0; i < len(listFilesXML); i++ {
		b := []byte(listFilesXML[i])
		for j := 0; j < len(b); j++ {
			bh = append(bh, b[j])
		}
	}

	b, err := doRequest(dnsServer+"/qwcfpWebService/ListFiles?wsdl", bh)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(b, &mx)

	if err != nil {
		return nil, err
	}

	return mx.Body.ListFilesResponse, nil

}

func Download(groupName string, loginKey string, FileNameEx string, dnsServer string, rootConfig string) (DownloadResponse, error) {
	str_arrayEx, err := parseFileToArraY(rootConfig + "Download.xml")

	if err != nil {
		return DownloadResponse{}, err
	}

	var downloadXML []string

	for j := 0; j < len(str_arrayEx); j++ {
		line := str_arrayEx[j]

		line = strings.TrimSpace(line)

		if line == "<loginKey></loginKey>" {
			line = "<loginKey>" + loginKey + "</loginKey>"
		}

		if line == "<fileName></fileName>" {
			line = "<fileName>" + FileNameEx + "</fileName>"
		}

		if line == "<infoGroup></infoGroup>" {
			line = "<infoGroup>" + groupName + "</infoGroup>"
		}

		downloadXML = append(downloadXML, line)
	}

	var bk = []byte{}

	for u := 0; u < len(downloadXML); u++ {
		b := []byte(downloadXML[u])
		for op := 0; op < len(b); op++ {
			bk = append(bk, b[op])
		}
	}

	bvg, err := doRequest(dnsServer+"/qwcfpWebService/Download?wsdl", bk)
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
