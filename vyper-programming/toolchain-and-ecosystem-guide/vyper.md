First make sure that you have `uv` installed

```
which uv
```

/usr/bin/uv

mkdir `name-of-project`

cd `name-of-project`

uv init

This will create a new Python project with the latest version of uv

If there is no .gitignore file after you have created this, 
then create your own:

```
# Python bytecode
__pycache__/
*.py[cod]

# Virtual environments
.env/
.venv/

# Distribution/Packaging
*.egg-info/
dist/
build/
```
