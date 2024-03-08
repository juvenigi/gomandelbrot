package debug

import (
	"log"
	"os"
	"runtime/pprof"
)

// PprofFiles a small dictionary of file handles for pprof to write profiles to
type PprofFiles struct {
	memFile *os.File
	cpuFile *os.File
}

// InitPProf initialise pprof and open pprof files based on paths specified.
func InitPProf(cpuProfPath string, memProfPath string) *PprofFiles {
	var err error
	var files PprofFiles
	if cpuProfPath != "" {
		files.cpuFile, err = os.Create(cpuProfPath)
		err = pprof.StartCPUProfile(files.cpuFile)
	}
	if memProfPath != "" {
		files.memFile, err = os.Create(memProfPath)
	}
	if err != nil {
		log.Fatal("Error while starting profiling: ", err.Error())
	}
	return &files
}

// StopPProf stops cpuprofile and dumps data for pprof files that are open
func StopPProf(files *PprofFiles) {
	var err error
	if files.memFile != nil {
		err = pprof.WriteHeapProfile(files.memFile)
		err = files.memFile.Close()
	}
	if files.cpuFile != nil {
		pprof.StopCPUProfile()
		err = files.cpuFile.Close()
	}
	if err != nil {
		log.Fatal("Error while stopping profiling: ", err.Error())
	}
}
