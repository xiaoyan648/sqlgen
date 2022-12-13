package parser

import (
	"encoding/json"
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/tools/go/packages"
)

// InterfacePath interface path
type InterfacePath struct {
	Name     string
	FullName string
	Files    []string
	Package  string
}

// GetInterfacePath get interface's directory path and all files it contains
func GetInterfacePath(v interface{}) (paths []*InterfacePath, err error) {
	value := reflect.ValueOf(v)
	k1 := value.Kind()
	_ = k1
	if value.Kind() != reflect.Func && value.Kind() != reflect.Slice {
		err = fmt.Errorf("model param is not function or string slice:%s", value.String())
		return
	}

	if value.Kind() == reflect.Slice {
		cfg, _ := v.([]string)
		if len(cfg) != 2 {
			err = fmt.Errorf("model param string slice need len 2:%s", value.String())
			return
		}
		paths = parsePackage(cfg[0], cfg[1])
		return
	}

	for i := 0; i < value.Type().NumIn(); i++ {
		var path InterfacePath
		arg := value.Type().In(i)
		path.FullName = arg.String()

		// keep the last model
		for _, n := range strings.Split(arg.String(), ".") {
			path.Name = n
		}

		ctx := build.Default
		var p *build.Package

		if strings.Split(arg.String(), ".")[0] == "main" {
			_, file, _, _ := runtime.Caller(3)
			p, err = ctx.ImportDir(filepath.Dir(file), build.ImportComment)
		} else {
			p, err = ctx.Import(arg.PkgPath(), "", build.ImportComment)
		}

		if err != nil {
			return
		}

		for _, file := range p.GoFiles {
			goFile := fmt.Sprintf("%s/%s", p.Dir, file)
			if fileExists(goFile) {
				path.Files = append(path.Files, goFile)
			}
		}

		if len(path.Files) == 0 {
			err = fmt.Errorf("interface file not found:%s", value.String())
			return
		}

		paths = append(paths, &path)
	}

	return
}

func parsePackage(srcLocaltion, fileName string) (paths []*InterfacePath) {
	absDir, _ := filepath.Abs(srcLocaltion)
	cfg := &packages.Config{
		Dir: absDir,
		// nolint: staticcheck
		Mode:  packages.LoadSyntax,
		Tests: false,
		// BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}

	for _, pkg := range pkgs {
		paths = append(paths, &InterfacePath{
			Name:     fileName,
			FullName: pkg.Name + "." + fileName,
			Files:    pkg.GoFiles,
			// Package:  pkg.PkgPath,
		})
		data, _ := json.Marshal(pkg)
		fmt.Printf("%+v", string(data))
	}

	return
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetModelMethod get diy methods
func GetModelMethod(v interface{}, skip int) (method *DIYMethods, err error) {
	method = new(DIYMethods)

	// get diy method info by input value, must input a function or a struct
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Func:
		fullPath := runtime.FuncForPC(value.Pointer()).Name()
		err = method.parserPath(fullPath)
		if err != nil {
			return nil, err
		}
	case reflect.Struct:
		method.pkgPath = value.Type().PkgPath()
		method.BaseStructType = value.Type().Name()
	default:
		return nil, fmt.Errorf("method param must be a function or struct")
	}

	var p *build.Package

	// if struct in main file
	ctx := build.Default
	if method.pkgPath == "main" {
		_, file, _, _ := runtime.Caller(skip)
		p, err = ctx.ImportDir(filepath.Dir(file), build.ImportComment)
	} else {
		p, err = ctx.Import(method.pkgPath, "", build.ImportComment)
	}
	if err != nil {
		return nil, fmt.Errorf("diy method dir not found:%s.%s %w", method.pkgPath, method.MethodName, err)
	}

	for _, file := range p.GoFiles {
		goFile := p.Dir + "/" + file
		if fileExists(goFile) {
			method.pkgFiles = append(method.pkgFiles, goFile)
		}
	}
	if len(method.pkgFiles) == 0 {
		return nil, fmt.Errorf("diy method file not found:%s.%s", method.pkgPath, method.MethodName)
	}

	// read files got methods
	return method, method.LoadMethods()
}
