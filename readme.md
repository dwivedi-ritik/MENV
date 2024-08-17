

## MENV

Tool to manage environment file. It encrypt your environment file and create MenvFile which you can push to your remote repository.
You can retrive you file using `menv update`.

#### Uses

- menv init -f environment-file
- menv update

if no option passed, its picks .env or .env.local file.


#### Installation

- Run `make` to build binary, move this binary to your env path

or for Linux or Mac follow these steps

- `git clone git@github.com:dwivedi-ritik/menv.git`
- `cd menv`
- `chmod +x ./install.sh`
- `./install.sh`
- restart your shell
