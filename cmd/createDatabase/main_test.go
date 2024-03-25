package main

import "testing"
import "github.com/philhanna/cwcomp"

func Test_main(t *testing.T) {
	config := cwcomp.Configuration
	config.DATABASE.NAME = "/tmp/sample.db"
	cwcomp.Configuration = config
	main()
}
