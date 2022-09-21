package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func main1() int {
	main()
	return 0
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"demo": main1,
	}))
}

func TestSelf(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
		Setup: func(env *testscript.Env) error {
			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get working directory: %w", err)
			}
			env.Setenv("DEMO", wd)
			return nil
		},
		UpdateScripts:       os.Getenv("DEMO_UPDATE") == "1",
		RequireExplicitExec: true,
	})
}
