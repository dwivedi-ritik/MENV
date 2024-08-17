

## MENV

Tool to manage environment file. It encrypt your environment file and create MenvFile which you can push to your remote repository.
You can retrive you file using `menv update`.

#### Uses

```
A tool to manage your enviroment files

Commands:
  init     Initialize the menv file
  update   Update your Menvfile with environment file changes
  generate Generate your environment file from Menvfile


Options:
        init     -f      Name of environment file
        generate -y      Yes for overridden message

Examples:
  menv init
  menv init -f config.json
  menv update
  menv generate

```

#### Installation

- Run `make` to build binary, move this binary to your env path

or for Linux or Mac follow these steps

- `git clone git@github.com:dwivedi-ritik/menv.git`
- `cd menv`
- `chmod +x ./install.sh`
- `./install.sh`
- restart your shell
