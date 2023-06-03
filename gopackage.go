package gengrpc

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/golistic/xos"
	"golang.org/x/tools/go/packages"
)

type GoPackage struct {
	packageName   string
	pathInPackage string
	contractsPath string
}

func NewGoPackage(packageName, pathInPackage string) (*GoPackage, error) {
	gp := &GoPackage{
		packageName:   packageName,
		pathInPackage: pathInPackage,
	}

	if err := gp.pathProtoContracts(); err != nil {
		return nil, err
	}

	return gp, nil
}

var _ sourcer = (*GoPackage)(nil)

func (gp *GoPackage) pathProtoContracts() error {
	// Load the package information
	cfg := &packages.Config{
		Mode: packages.NeedFiles,
		Dir:  ".",
		Env:  os.Environ(),
	}

	pkgs, err := packages.Load(cfg, gp.packageName)
	if err != nil || len(pkgs) == 0 {
		return fmt.Errorf("could not load package '%s' (%w)", gp.packageName, err)
	}

	if len(pkgs) > 1 {
		return fmt.Errorf("found multiple '%s' packages (not possible)", gp.packageName)
	}

	pkg := pkgs[0]

	if len(pkg.Errors) != 0 {
		err := fmt.Errorf(strings.Replace(pkg.Errors[0].Error(), "\n", " ", -1))
		return fmt.Errorf("source package '%s' not available (%w)", gp.packageName, err)
	}

	pkgPath := path.Dir(pkg.GoFiles[0])

	p := filepath.Join(pkgPath, gp.pathInPackage)
	if !xos.IsDir(p) {
		return fmt.Errorf("'%s' in package '%s' is not directory", gp.pathInPackage, gp.packageName)
	}

	gp.contractsPath = p

	return nil
}

func (gp *GoPackage) ContractPath() string {
	return gp.contractsPath
}

func (gp *GoPackage) Contracts() ([]string, error) {
	var files []string

	err := filepath.Walk(gp.contractsPath, func(filepath string, info fs.FileInfo, err error) error {
		if path.Ext(filepath) == ".proto" {
			files = append(files, strings.Replace(filepath, gp.contractsPath+"/", "", -1))
		}
		return nil
	})

	return files, err
}
