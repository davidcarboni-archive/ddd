A Docker multi-stage build to package a Go executable in a dedicated run container, based on the scratch image, to make it really minimal.
This time we add ca-certificates so we can make the https call.
However, we're the root user, which is more privilege than we need to run the executable.
