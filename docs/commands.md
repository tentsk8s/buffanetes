# Buffanetes CLI Commands

The Buffanetes CLI is called `buffnet`. This section introduces the commands that the 
Buffanetes CLI supports.

# `buffnet deploy web`

This command deploys your web application, according to your `.buffanetes/web.toml` config.
It defaults to deploying to your development environment, but you can pass the `--env` flag 
to override the environment. For example, to deploy to `production`, run this:

```console
buffnet deploy web --env production
```

You'll need to have production configuration in your `.buffanetes/web.toml` file to do this.

# `buffnet deploy grift`

Buffanetes takes care of your grifts too. After you put a grift into the `grifts` directory, Buffanetes can pick them up and generate config for them. Run `buffnet detect grifts` for that.

If you add another grift, just run `buffnet detect grifts` again and you're all set.

Just like the `web.toml` file tells Buffanetes how to deploy your web application, 
`grifts.toml` tells if how to run your grifts.

When you're ready to run a grift, run `buffnet deploy grift <grift name>`. Your grift will run 
in development by default, but just like `buffnet deploy web`, pass the `--env` flag to 
override that.

# Database Migrations

You guessed it, Buffanetes can also run database migrations for you. By now, you can probably
guess how it works.

`buffnet deploy migrations up` packages up your database migrations and runs all of your 
`*.up.fizz` migrations. And just like the other commands, it runs them in development by 
default and you can override the default with `--env`.

You can pass `buffnet deploy migrations down` to run all of the `*.down.fizz` migrations, but
be careful! The command will prompt you before you run them so that you don't accidentally 
wipe out your database.

If you're _absolutely sure_ that you want to run the migrations and don't want to see the
prompt, pass the `--i-am-sure` flag.
