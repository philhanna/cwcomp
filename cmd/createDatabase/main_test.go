package main

import "testing"
import "github.com/philhanna/cwcomp"

func Test_main(t *testing.T) {
	config := cwcomp.GetConfiguration()
	config.DATABASE.NAME = "/tmp/sample.db"
	cwcomp.GetConfiguration = func() *cwcomp.Configuration {
		return config
	}
	main()
}
