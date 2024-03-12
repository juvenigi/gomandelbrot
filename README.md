# gomandelbrot

This project is a Mandelbrot set visualizer implemented in [Go](https://go.dev) using the [Pixel](https://github.com/gopxl/pixel) game development library. 

⚒ The current implementation focuses on rendering the set in a fixed window size of 600x720, covering the complex plane from -2-1.5i to 2+1.5i. More functionality will be added pretty soon™

## Planned Enhancements

- **Interactive Controls**: Future updates will include mouse and keyboard controls for interactive exploration of the Mandelbrot set. Users will be able to zoom in, pan, and navigate through different regions of the set.
- **Performance Improvements**: Ongoing efforts will focus on optimizing the performance of the visualization, ensuring smooth rendering even for larger and more detailed Mandelbrot sets.

## Usage

The program supports optional command-line flags for profiling purposes:

    -mem: Specify a file to write memory profiles.
    -cpu: Specify a file to write CPU profiles.

both profiles can be run using default profiler:
```bash
go tool pprof <your file>
```

## How to Run

Ensure you have Go (1.22 or later) installed on your system. To run the program, execute the following command in the terminal:

```bash
export CGO_ENABLED=1 # make sure CGO is enabled for 'go-gl/gl' dependency to work (part of pixel)
go run main.go -mem=memprofile.prof -cpu=cprofile.prof
```
This will generate memory and CPU profiles in the specified files.

## Dependencies

- [Pixel](https://github.com/gopxl/pixel): A 2D game development library in Go.

## Contribution

Feel free to use this code in your projects. Adding some attribution would be nice :)


Happy exploring the beauty of fractals! 🌌
