package env

import (
	"fmt"
	"os"
	"testing"

	sdk "github.com/hashicorp/sentinel-sdk"
	framework "github.com/hashicorp/sentinel-sdk/framework"
	plugintesting "github.com/hashicorp/sentinel-sdk/testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	plugintesting.Clean()
	os.Exit(exitCode)
}

func TestImport_impl(t *testing.T) {
	var _ sdk.Import = New()
}

func TestRoot_impl(t *testing.T) {
	var _ framework.Root = new(root)
}

func TestImport(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Name   string
		Source string
	}{
		{
			"read_good",
			fmt.Sprintf(`main = subject.read("%s/../testdata/foo.json") is "{\"foo\":\"bar\",\"baz\":\"qux\"}"`, cwd),
		},
		{
			"list_good",
			fmt.Sprintf(`main = subject.list("%s/../testdata") is [{"directory": false, "name": "foo.json"}]`, cwd),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			plugintesting.TestImport(t, plugintesting.TestImportCase{
				Source: tc.Source,
			})
		})
	}
}
