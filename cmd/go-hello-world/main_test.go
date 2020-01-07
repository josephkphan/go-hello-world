package main

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

var _ = DescribeTable("Fake Test",
	func() {
		err := PrettyPrint([]string{"fake-test"})

		Expect(err).ToNot(HaveOccurred())
	},
	Entry("No Resources"),
)
