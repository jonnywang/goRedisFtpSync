package ftpSync

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type TConfig struct {
	ListenServer      string `xml:"listen"`
	FtpServerAddress  string `xml:"ftp>address"`
	FtpServerUser     string `xml:"ftp>user"`
	FtpServerPassword string `xml:"ftp>password"`
	DependCommand     string `xml:"depend_command"`
}

var GConfig *TConfig

func ParseXmlConfig(path string) (*TConfig, error) {
	if len(path) == 0 {
		return nil, errors.New("not found configure xml file")
	}

	n, err := GetFileSize(path)
	if err != nil || n == 0 {
		return nil, errors.New("not found configure xml file")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	GConfig = &TConfig{}

	data := make([]byte, n)

	m, err := f.Read(data)
	if err != nil {
		return nil, err
	}

	if int64(m) != n {
		return nil, errors.New(fmt.Sprintf("expect read configure xml file size %d but result is %d", n, m))
	}

	err = xml.Unmarshal(data, &GConfig)
	if err != nil {
		return nil, err
	}

	Logger.Printf("read config %+v", GConfig)

	return GConfig, nil
}
