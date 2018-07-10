Developer Developer Developer
=============================

## How it works


This repo is made up of nested directories. Each successive directory devlops the idea a step further.

Each directory contains a brief `readme.txt` to explain what's going on at that level. There's an `_.txt` containing the commands you need to run that level.

Other than that, you'll find a `main.go` and usually a `Dockerfile`.

For levels that rely on the `scratch` image, there's also a `filesystem.sh` that exports the container filesystem for you into a directory called `export`. To get a quick listing of the filesystem, try a `ls -lRh export`.

## Links
[The slides for the talk](https://carb.onl/ddd) are available online, along with an [accompanying blog post](https://medium.com/@davidcarboni/simplicity-wins-4fc54efae1f0). [The built image can be found on Docker Hub](https://hub.docker.com/r/davidcarboni/go-scratch-nonprivileged-multistage/) and [the complete repo](https://github.com/davidcarboni/go-scratch-nonprivileged-multistage) is available on Github. It's made up of four files, including the readme, which I think is deliciously simple.
