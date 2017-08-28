//Package fileurl contains tools for dealing with urls that refer to files.
package fileurl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

//IsFileURL returns whether a URL is a link to a file.
func IsFileURL(path string) bool {
	match, _ := regexp.Match(`https?:\/\/.+\..+\/.+\.[a-zA-Z0-9]+`, []byte(path))
	return match
	//starts with http:// or https://
	//any characters
	//a domain (foo.whatever.net)
	//then a slash
	//some filename
	//and then a period and file extension
	//eg: http://foo.bar.baz/someFolder/hello.tsv
}

//Copy the data located at the URL to the destination.
func Copy(url string, dst io.Writer) (n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("http.Get: %v", err)
	}
	defer resp.Body.Close()
	return io.Copy(dst, resp.Body)

}

//DownloadTemp downloads the url to a temporary file starting with prefix in the current working directory.
func DownloadTemp(url string, prefix string) (file *os.File, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err = ioutil.TempFile(cwd, prefix)
	if err != nil {
		return nil, err
	}
	_, err = Copy(url, file)
	if err != nil {
		file.Close()
	}
	return file, err
}

//Download downloads the url to the file specified by path. This will truncate an existing file!
func Download(url string, filepath string) (file *os.File, err error) {
	file, err = os.Create(filepath)
	if err != nil {
		return nil, err
	}
	_, err = Copy(url, file)
	if err != nil {
		file.Close()
	}
	return file, err
}

//DownloadClose downloads the url to the file specified by path, then closes it. This will truncate an existing file!
func DownloadClose(url string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = Copy(url, file)
	return err
}
