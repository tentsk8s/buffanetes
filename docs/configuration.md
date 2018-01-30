# Buffanetes Configuration

Buffanetes is completely driven by its configuration files. All the configs are written
in [TOML](https://github.com/toml-lang/toml). As you read before in the 
[directory structure doc](./directory_structure.md) document, all the configuration files
go under the `.buffanetes` directory.

# Buffanetes Metadata

The first configuration file that Buffanetes looks for is `meta.toml`. Buffanetes almost
always looks for this file before it does anything else. It tells Buffanetes what your
app is called and how to build and push Docker images.

Here's an example:

```toml
# this tells buffanetes your app's name. it's used all over the place
name = myapp
# this tells buffanetes the environment that your app can be deployed in. environments are used
environments = [development, staging, production]
version = 1
```

# Buffanetes Web Configuration

If you try to deploy your web application, Buffanetes will look for a `web.toml` file for
guidance on what to do.

It holds details on how your web application should be deployed. Here's an example:

```toml
# this route tells buffanetes to route http://buffanetes.io to port 8080 (the default for Buffalo). in the future, Buffanetes might support multiple ports!
[routes.development]
    host = dev.buffanetes.io

```