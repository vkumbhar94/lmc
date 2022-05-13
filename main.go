package main

import (
	"bytes"
	"fmt"
	"github.com/vkumbhar94/lmc-util/pkg/conv"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	con := &conv.OldArgusConf{}
	err = yaml.Unmarshal(file, con)
	if err != nil {
		fmt.Println(err)
		return
	}
	newc := con.ToNewArgusConf()
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2) // this is what you're looking for
	err = yamlEncoder.Encode(&newc)
	if err != nil {
		fmt.Println(err)
		return
	}
	//b, err := yaml.Marshal(newc)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = ioutil.WriteFile(os.Args[1]+".new", b, fs.ModePerm)
	err = ioutil.WriteFile(os.Args[1]+".new", b.Bytes(), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}
