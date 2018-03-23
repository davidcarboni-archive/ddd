Developer Developer Developer
=============================

Demo repo to step through the repo: https://github.com/davidcarboni/go-scratch-nonprivileged-multistage

## How it works


This repo is made up of nested directories. Each successive directory devlops the idea a step further.

Each directory contains a brief `readme.txt` to explain what's going on at that level. There's an `_.txt` containing the commands you need to run that level.

Other than that, you'll find a `main.go` and usually a `Dockerfile`.

For levels that rely on the `scratch` image, there's also a `filesystem.sh` that exports the container filesystem for you into a directory called `export`. To get a quick listing of the filesystem, try a `ls -lRh export`.
