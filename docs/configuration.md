# Buffanetes Configuration

Buffanetes is completely driven by its configuration files. All the configs are written
in [TOML](https://github.com/toml-lang/toml). As you read before in the 
[directory structure doc](./directory_structure.md) document, all the configuration files
go under the `.buffanetes` directory.

# What is Configuration?

Buffanetes goes by some conventions to configure your server, grifts and migrations
to run properly in Kubernetes. This means that it sets up all the pieces to be deployed to 
Kubernetes (i.e. `Pod`s, `Job`s, etc...), but it also injects the app with the proper environment
variables.

You can override those environment variables if you want, or add more too. There are three ways to
do it:

## Hard-Coded Environment Variables

Here's how you hard-code an environment variable:

```toml
[env.hardcode]
name = MY_ENV
value = abcd
```

When you do this, your app running in Kubernetes will see `MY_ENV = abcd` in its environment.

## Environment Variables from your Machine

Here's how you tell Buffanetes to grab an environment variable's value from the host it's running
on, and expose it to the running app:

```toml
[env.local]
name = MY_SECOND_ENV
host = HOST_ENV_VAR_NAME
```

## Environment Variables from a Secret

Here's how you tell Buffanetes to configure the running app to:

1. Get a value from a secret that's already in Kubernetes
2. Inject it into the running app as an environment variable

```toml
[env.secret]
name = MY_ENV_FROM_SECRET
secret-name = topsecret
secret-key = topsecret1
```

# Configuring Grifts

Configuring Grifts is pretty simple
