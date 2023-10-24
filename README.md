# A few comments
- Time to run averages around 15 seconds on my M1 macbook pro. This should leave plenty of room on slower machines.
- There is still alot of fine tuning that can be done to increase performance.
- There is a lot of dead code that I didn't prune. I left it for visibility, so I can demonstrate my process as I worked through the problem.
- This is only my second project in go. So my understanding of the language was constantly evolving throughout this project. Leading to some inconsistencies.
  
# Build and execute
This was built using go version 1.21.2

run `go build .`
This will output an executable called vehicle-routing-problem

The executable takes a single argument for location of a problem file

example `./vehicle-routing-problem ./problems/problem1.txt`

# Visualizations
- Visualizations can increase runtime and should not be used when benchmarking the program runtime.
- Use the `--driver-route-file` to define a file location for driver schedules.
- Use `--driver-count` to define how many drivers per driver route file. Default will include all drivers in single file. If multiple files are required a suffix will be added to the file name.
- A file called `annealing_graph.png` and `annealing_temp.png` will be created, this cannot be changed at the moment.
