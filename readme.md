

## MENV

Tool to manage environment file. It encrypt your environment file and create MenvFile which you can push to your remote repository.
You can retrive you file using `menv update`.

#### Why and what problem its solves ?

The sole purpose of this tool to creates an encrypted Menvfile containing your environment configuration and metadata.
It allows you to securely store and retrieve your environment file even if you delete your local project.
This ensures easy restoration of environment files and configurations without relying on local project files.


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
