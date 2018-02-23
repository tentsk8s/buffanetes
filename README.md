# Buffanetes

Buffanetes is a deployment tool that brings [Kubernetes](https://kubernetes.io) to your 
[Buffalo](https://gobuffalo.io) web applications.

Buffanetes aims to take all of the Kubernetes work off of your shoulders, so you can focus on
building your app and let Buffanetes take care of everything else for you.

It lets you deploy your webapp or run your grifts and migrations all with a single CLI
and a bit of TOML configuration.

# Requirements

Aside from the Buffalo CLI and its dependencies (Go, Webpack, etc...), you'll need to have 
[Docker](https://docker.com) installed on your machine and [Helm](https://helm.sh) installed
on the Kubernetes cluster. Someone else can do the Helm part for you.

# Assumptions

Buffanetes assumes that you have a standard Buffalo application set up. From there, it will take
care of packaging up your app, generating all the Kubernetes "stuff" that you need to deploy
it, and then actually deploy it to your Kubernetes cluster.

# Your Workflow

First thing's first, Buffanetes doesn't change your development workflow. You still use the 
`buffalo` CLI to do all your local development. You use the Buffanetes CLI (called `buffnet`) 
when you're ready to deploy your code to your Kubernetes cluster.

Let's look at a simple workflow, starting with no Buffalo app.

```console
$ buffalo new myapp # first, create a new Buffalo app with sensible defaults
$ cd myapp # let's go into the new app
$ buffnet init # now set up the Buffanetes config files
$ buffalo dev # do your normal edit-test workflow
$ buffnet deploy web # time to deploy your webapp!
```

Now you're cooking! You created a Buffalo app, built it, and deployed it to Kubernetes.

Enjoy!

# Documentation

The simple workflow above works for many folks, but Buffanetes unlocks more for you.

Please visit the [documentation](./docs) for complete docs, examples and more.
