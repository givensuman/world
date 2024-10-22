<img src="./assets/theworld_balatro.png" align="center" alt="The World tarot card" />

### world

#### a declarative package manager for Arch, btw

### Installation

TODO

### Usage

Your system packages are documented in `WORLD_DIR`, by default `$HOME/.world`. This directory has the following structure:

```
WORLD_DIR
    packages
    groups
        cli_tools
    backups
        01_01_2001_12_00_00
```

`world` can be followed by a number of commands, whose naming is meant to reflect natural language:

| command   | aliases           | description                                                                      |
| --------- | ----------------- | -------------------------------------------------------------------------------- |
| add       | a, ad             | Adds a package to the `package` file and installs it to the system               |
| remove    | r, re, rem        | Removes the package from the `package` file and uninstalls it from system        |
| install   | i, inst           | Installs a package to the system without saving it to the `package` file         |
| uninstall | u, un, ui, uninst | Uninstalls a package from the system without removing it from the `package` file |

### Motivation

Honestly, obsessive compulsive disorder. Not being sure what packages are installed in my system and which ones I actually care about annoys me, and it's particularly pervasive on Arch when using standard package managers like `pacman` or `yay`.

This program is inspired by the package manager of Alpine Linux `pkg`, which stores a list of all the packages installed on the system in a single text file, called `world`. It aims to have similar, if not broader, functionality and to facilitate package management on Arch and Arch-based distros with intuitive commands meant to emulate natural language.

### What it is

A frontend for `libalpm`, much like `pacman`, `yay`, or `paru`, but with a quirk of maintaining system packages in declarative, safe-to-edit text files.

### What it's not

Any kind of immutability to your system. Arch is rolling release and this does not change that. `world` does not track package versions (at least not yet) because doing so would make maintaining your package files hell.

### Pros

View, add, or remove the packages on your system the same way you would edit any text file. Use version control systems to keep a history of your system, and even remotely clone\* your package list.

### Cons

Primarily that alternatives like `yay`, `paru`, and especially `pacman` are older, better written, better maintained, better tested.

### Is this a JoJo reference?

Sure.

<img src="./assets/theworld_jojo.jpg" alt="JoJo's Bizarre Adventure">

### License

[MIT]("/LICENSE") ~
