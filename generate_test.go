package gengrpc_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golistic/xos"
	"github.com/golistic/xt"

	_ "github.com/golistic/gengrpctestdata" // just load dependency

	"github.com/golistic/gengrpc"
)

const testModule = "github.com/golistic/gengrpctestdata"
const testSourcePackage = "github.com/golistic/gengrpctestdata"
const testPathInSource = "proto"
const testResultsDir = "_testresults"

func TestGenerate(t *testing.T) {
	t.Run("generate without source-path set", func(t *testing.T) {
		dest := filepath.Join(testResultsDir, "difi2923")
		_ = os.RemoveAll(dest)

		source, err := gengrpc.NewGoPackage(testSourcePackage, testPathInSource)
		xt.OK(t, err)

		xt.OK(t, gengrpc.Generate(testModule, source, dest))
		xt.Assert(t, xos.IsRegularFile(filepath.Join(dest, "echo", "echo.pb.go")))
	})

	t.Run("without path in source", func(t *testing.T) {
		dest := filepath.Join(testResultsDir, "wid82kdi")
		_ = os.RemoveAll(dest)

		source, err := gengrpc.NewGoPackage(testSourcePackage, "")
		xt.OK(t, err)

		xt.OK(t, gengrpc.Generate(testModule, source, dest))
		xt.Assert(t, xos.IsRegularFile(filepath.Join(dest, "echo", "echo.pb.go")))
	})

	t.Run("incorrect path in source provided", func(t *testing.T) {
		_, err := gengrpc.NewGoPackage(testSourcePackage, "something/that/doesnot/exists")
		xt.KO(t, err)
		xt.Assert(t, strings.Contains(err.Error(), "is not directory"))
	})
}
