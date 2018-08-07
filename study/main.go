package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/golang/glog"
)

func main() {
	// exec backup cmd
	// path, err := exec.LookPath("mysqldump")
	// if err != nil {
	// 	log.Fatal("LookPath: ", err)
	// }
	cmd := exec.Command("bash", "-c", "/usr/bin/mysqldump -h mariadb-0.mariadb -uroot -proot -A | gzip > /tmp/backup.sql.gz")

	glog.Infof("%s", cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
