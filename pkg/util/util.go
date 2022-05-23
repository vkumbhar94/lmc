package util

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vkumbhar94/lmc/pkg/icon"
	"gopkg.in/yaml.v2"
	"net"
	"os"
)

var Debug = false
var Quiet = false

func SetDebug(b bool) {
	Debug = b
}
func SetQuiet(b bool) {
	Quiet = b
}
func Flatten(items []interface{}) ([]string, error) {
	return doFlatten([]string{}, items)
}

func doFlatten(result []string, items interface{}) ([]string, error) {
	var err error

	switch v := items.(type) {
	case string:
		result = append(result, v)
	case []string:
		result = append(result, v...)
	case []interface{}:
		for _, item := range v {
			result, err = doFlatten(result, item)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New(fmt.Sprintf("Flatten does not support %T", v))
	}

	return result, err
}

func StringSliceContains(slice []string, s string) bool {
	for _, element := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func GetRandomPort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	defer listener.Close()
	if err != nil {
		return 0, err
	}

	return listener.Addr().(*net.TCPAddr).Port, nil
}

func printlnConfig(prefix string, cmd *cobra.Command, out interface{}) {
	marshal, err := yaml.Marshal(out)
	if err != nil {
		cmd.PrintErrln(err)
	}
	cmd.Printf("\n%s\n%v\n", prefix, string(marshal))
}
func PrintlnDebug(cmd *cobra.Command, msg string) {
	if Debug {
		cmd.Println(msg)
	}

}
func PrintlnDebugConfig(cmd *cobra.Command, prefix string, out interface{}) {
	if Debug {
		printlnConfig(prefix, cmd, out)
	}
}

func printf(cmd *cobra.Command, msg string, ic icon.Icon) {
	cmd.Printf("\n%s %s\n", ic, msg)
}
func PrintlnSuccess(cmd *cobra.Command, msg string) {
	if !Quiet {
		printf(cmd, msg, icon.SuccessTick)
	}
}
func PrintlnFailed(cmd *cobra.Command, msg string) {
	printf(cmd, msg, icon.FailedCross)
}
func PrintlnRunning(cmd *cobra.Command, msg string) {
	if !Quiet {
		printf(cmd, msg, icon.RoundStar)
	}
}
