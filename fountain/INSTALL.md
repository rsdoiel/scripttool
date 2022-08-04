{
	"title": "fountain, a Go package for working with Fountain screenplay format"
}


Installation
============

*fountain2json*, *fountainfmt*, *fountain2html* are a command line programs run from a shell like Bash. You can find compiled version in the [releases](https://github.com/rsdoiel/fountain/releases/latest) 

Compiled version
----------------

This is generalized instructions for a release. 

Compiled versions are available for Mac OS X (x86 and M1 processors, macos-amd64, macos-arm64), 
Linux (x86 processor, linux-amd64), Windows (x86 processor, windows-amd64), 
Rapsberry Pi OS (arm7 processor, raspbian-arm7) and Pine64 (arm64 processor, linux-arm64)


VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/rsdoiel/fountain/releases/latest


| Platform    | Zip Filename                           |
|-------------|----------------------------------------|
| Windows     | fountain-VERSION_NUMBER-windows-amd64.zip |
| Mac OS X    | fountain-VERSION_NUMBER-macos-amd64.zip  |
| Mac OS X    | fountain-VERSION_NUMBER-macos-arm64.zip  |
| Linux/Intel | fountain-VERSION_NUMBER-linux-amd64.zip   |
| Raspbery Pi | fountain-VERSION_NUMBER-raspbian-arm7.zip |
| Pine64      | fountain-VERSION_NUMBER-linux-arm64.zip   |


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


### Mac OS X

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip fountain-*-macos-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    fountain -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip fountain-*-windows-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    fountain -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip fountain-*-linux-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    fountain -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip fountain-*-raspbian-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    fountain -version
```


Compiling from source
---------------------

The _fountain_ command line programs are "go gettable".  Use the "go get" command to download the dependant packages
as well as source code.

```shell
    go get -u github.com/rsdoiel/fountain/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/rsdoiel/fountain src/github.com/rsdoiel/fountain
    cd src/github.com/rsdoiel/fountain
    make
    make test
    make install
```


