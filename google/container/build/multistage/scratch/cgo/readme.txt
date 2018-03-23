A Docker multi-stage build to package a Go executable in a dedicated run container, based on the scratch image, to make it really minimal.
This time we disable cgo to create a "fat" binary. Fat is relative though - it goes from 2M to 6M.
This still fails though because we're trying to use https without any ca-certificates.
