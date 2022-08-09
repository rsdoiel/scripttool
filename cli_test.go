package scripttool

import (
	"io/fs"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	testdataBase = "testdata"
	testoutBase  = "testout"
)

func TestRunScripttool(t *testing.T) {
	var verb string
	appName := os.Args[0]
	// Run the JSON command
	if _, err := os.Stat(testoutBase); err == nil {
		os.RemoveAll(testoutBase)
	}
	os.MkdirAll(testoutBase, 0775)
	fSys := os.DirFS(testdataBase)
	fs.WalkDir(fSys, ".", func(p string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			inName := info.Name()
			inExt := path.Ext(inName)
			outName := strings.TrimSuffix(inName, inExt) + ".json"

			//fmt.Printf("DEBUG pth: %q fName: %q, ext: %q\n", p, inName, inExt)
			switch inExt {
			case ".txt":
				verb = "fountain2json"
			case ".fountain":
				verb = "fountain2json"
			case ".fdx":
				verb = "fdx2json"
			case ".fadein":
				verb = "fadein2json"
			case ".osf":
				verb = "osf2json"
			default:
				// We're only doing
				//fmt.Printf("Skipping %q, unsupported to file extension %q\n", inName, inExt)
				return nil
			}
			// We'll just rely on our standard VERB INPUT_NAME OUTPUT_NAME
			in := os.Stdin
			out := os.Stdout
			eout := os.Stderr
			inputFName := path.Join(testdataBase, inName)
			outputFName := path.Join(testoutBase, outName)
			args := []string{appName, verb, inputFName, outputFName}
			if err := RunScripttool(in, out, eout, args); err != nil {
				t.Errorf("unexpected error RunScripttool(in, out, eout, %+v) -> %s", args, err)
			}
		}
		return nil
	})
}
