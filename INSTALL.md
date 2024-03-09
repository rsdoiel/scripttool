---
title: scripttool installation
---

Installation
============

*scripttool* is a command line program run from a shell like Bash. You can find compiled version in the [releases](https://github.com/rsdoiel/scripttool/releases/latest) 

## Quick install via curl

This project can be installed via curl and sh for POSIX compatible systems.
Enter the following in your shell.

~~~
curl https://rsdoiel.github.io/scripttool/installer.sh | sh
~~~

Compiled version
----------------

This is generalized instructions for a release. 

Compiled versions are available for Mac OS (Intel and M1/M2 processor, macOS-x86_64, macOS-arm64), 
Linux (Intel process, Linux-x86_64), Windows (Intel processor, Windows-x86_64), 
Raspberry Pi OS (arm7 32bit processor, RaspberryPiOS-arm7) and Pine64 (arm64 processor, Linux-aarch64)

VERSION\_NUMBER is a [semantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on GitHub and click latest release

>    https://github.com/rsdoiel/scripttool/releases/latest


| Platform       | Zip Filename                                    |
|----------------|-------------------------------------------------|
| Windows        | scripttool-VERSION_NUMBER-Windows-x86_64.zip    |
| Mac OS (Intel) | scripttool-VERSION_NUMBER-macOS-x86_64.zip      |
| Mac OS (M1,M2) | scripttool-VERSION_NUMBER-macOS-arm64.zip       |
| Linux/Intel    | scripttool-VERSION_NUMBER-Linux-x86_64.zip      |
| Raspberry Pi   | scripttool-VERSION_NUMBER-RaspberryPiOS-arm7.zip |
| Pine64         | scripttool-VERSION_NUMBER-Linux-aarch64.zip     |


The basic recipe
----------------

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.  
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. export PATH="$HOME/bin:$PATH")
+ Test


### Mac OS

1. Download the zip file
2. Unzip the zip file
3. Copy the executable to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip scripttool-*-macOS-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scripttool -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executable to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip scripttool-*-Windows-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scripttool -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executable to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip scripttool-*-Linux-x86_64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scripttool -version
```


### Raspberry Pi OS

Released version is for a Raspberry Pi 3 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executable to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip scripttool-*-RaspberryPiOS-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scripttool -version
```


Compiling from source
---------------------

_scripttool_ is "go gettable".  Use the "go get" command to download the dependent packages
as well as _scripttool_'s source code.

```shell
    go get -u github.com/rsdoiel/scripttool/...
```

Or clone the repository and then compile

```shell
    cd
    git clone https://github.com/rsdoiel/scripttool src/github.com/rsdoiel/scripttool
    cd src/github.com/rsdoiel/scripttool
    make
    make test
    make install
```


