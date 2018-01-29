# Buffanetes

Buffanetes is a [Kubernetes](https://kubernetes.io) deployment tool for 
[Buffalo](https://gobuffalo.io) web applications.

Buffanetes assumes you have a standard Buffalo application set up. From there, it starts
by generating the things you'll need to deploy your application to Kubernetes.

From there, buffanetes it provides an easy to use command line interface (CLI) that you
can use to deploy all the pieces of your application to a Kubernetes cluster - from
the main webapp to grifts to database migrations.

Buffanetes isn't a patchwork of Kubernetes utilities - it's a holistic tool that deploys
your application to your Kubernetes cluster, without making you learn and write all the
Kubernetes fundamentals (Kubernetes manifests, Helm charts, and so on...)

# Your Workflow

Buffanetes doesn't change your development workflow. You still use the `buffalo` CLI to
do all your local development. You'll use the Buffanetes CLI (called `buffnet`) when you're ready 
to deploy your application to your Kubernetes cluster.

Let's look at a simple workflow.

When you first run `buffalo new myapp`, you'll get your brand new app. Thanks Buffalo!

You simply `cd` into the `myapp` directory and run `buffnet init` to set up your app for
Kubernetes.

Then, after you build some stuff into your app, you're ready to deploy it. You'll run
`buffnet deploy web` for that.

Now you're cooking! Your webapp is built, deployed, and ready to go. Enjoy!

# Documentation

You can do much more with Buffanetes, and this section holds all the documentation for all 
that awesome functionality. Let's get started!

## The Buffanetes Configuration

You saw `buffnet deploy web` above. To make that happen, Buffanetes looked for a 
[TOML](https://github.com/toml-lang/toml) file called in the `.buffanetes` directory.

The initial config that `buffnet new myapp` generated is called `.buffanetes/web.toml`, and
it's pretty simple. It tells Buffanetes a few simple things about your application:

- What organization your application lives in, and what the application name is
- What command to run to build the application for development
- Where to `docker push` the application for development
- How to route requests to the application development

## The `buffnet deploy web` Command

You saw the `buffnet deploy web` command above. The command deploys to your development 
environment by default, but if you want to deploy to another environment, you can pass 
the `--env` flag to do it. For example, to deploy to `production`, run this:

```console
buffnet deploy web --env production
```

You'll need to have production configuration in your `.buffanetes/web.toml` file to do this.

## Grifts

But wait, there's more! Buffanetes takes care of your grifts too. After you put a grift
into the `grifts` directory, Buffanetes can pick them up and generate config for them. 
Run `buffnet detect grifts` for that. Just like for the `web.toml`, it generates development
config for them and puts it in `.buffanetes/grifts.toml`. If you add another grift, run 
`buffnet detect grifts` again and you're all set.

When you're ready to run a grift, run `buffnet deploy grift <name>`. Your grift will run 
in development by default, but just like `buffnet deploy web`, you pass the `--env` 
flag to override that.

## Database Migrations

You guessed it, Buffanetes can also run database migrations for you. By now, you probably
know what's coming. `buffnet deploy migrations up` packages up your database migrations
and runs all of your `*.up.fizz` migrations. And as you expect, it runs them in development
by default and you can override the default with `--env`.

You can pass `buffnet deploy migrations down` to run all of the `*.down.fizz` migrations, but
be careful! It'll prompt you before you run them so that you don't accidentally wipe out your
database.

If you're _absolutely sure_ that you want to run the migrations and don't want to see the
prompt, pass the `--i-am-sure` flag.