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

The samples above work for many folks, but you can do more with Buffanetes.

Please visit the [documentation](./docs) for complete docs, examples and more.
