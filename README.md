# GoGlide: A Docker Builder Image for Go + Glide

To use the regular versions:

```
FROM technosophos/goglide:1.5-1

# Now start working...
COPY . wherever
WORKDIR wherever

# Glide is in $PATH, so you can use it wherever
RUN glide up

# ...
```

To use the `-onbuild` versions create a Dockerfile like this:

```
FROM technosophos/goglide:1.5-1-onbuild
```

Yeah, that's really it.
