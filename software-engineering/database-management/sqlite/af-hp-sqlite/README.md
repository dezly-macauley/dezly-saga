If you have just cloned this repo, run this command
`uv sync`

This is will install all of the required dependencies

_______________________________________________________________________________

You will need to have `direnv` installed on your system

Add the following lines to your .zshrc
```bash

eval "$(direnv hook zsh)"
export DIRENV_LOG_FORMAT=""

```

Use this to load a litecli configuration for a specific database

```
litecli --liteclirc $litecli_config_path database_file.sqlite
```
