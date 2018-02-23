# Buffanetes CLI Commands

The Buffanetes CLI is called `buffnet`. This section introduces the commands that the 
Buffanetes CLI supports.

# `buffnet deploy web`

This command builds your web application (via `docker build`) and then deploys it according to your 
`.buffanetes/web.toml` config file.

# `buffnet deploy grift <name>`

Buffanetes takes care of your grifts too. This command builds your grifts (via `docker build`) and 
then deploys the grift `name` to your cluster. It will watch the grift run and show you the logs
as well.

Just like the `web.toml` file tells Buffanetes how to run your web application, 
`grifts.toml` tells if how to run your grifts.

# `buffanet deploy migrations`

You guessed it, Buffanetes can also run database migrations for you.

This command packages up your database migrations and runs `buffalo db migrate up` in your 
cluster. There is currently no way to run `buffalo db migrate down`.
