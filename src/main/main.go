package main

import (
	"flag"
	"gomandelbrot/src/debug"
	"gomandelbrot/src/render"
	_ "runtime/pprof" // so that debugging is enabled asap

	"github.com/gopxl/pixel/pixelgl"
)

var (
	memProfileFilePath string
	cpuProfileFilePath string
)

func init() {
	flag.StringVar(&memProfileFilePath, "mem", "", "write memory profile to the specified file")
	flag.StringVar(&cpuProfileFilePath, "cpu", "", "write CPU profile to the specified file")
	flag.Parse()
}

func main() {
	// open/override pprof files
	files := debug.InitPProf(cpuProfileFilePath, memProfileFilePath)
	// close files on termination
	defer debug.StopPProf(files)
	// make sure pixel runs on main thread
	pixelgl.Run(render.Run)
}
