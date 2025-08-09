Installation for development of **scripttool**
===========================================

**scripttool** A library and command line program for working with fdx (Final Draft XML), osf, Fade In and Fountain formatted Screen Play files.

Quick install with curl or irm
------------------------------

There is an experimental installer.sh script that can be run with the following command to install latest table release. This may work for macOS, Linux and if you’re using Windows with the Unix subsystem. This would be run from your shell (e.g. Terminal on macOS).

~~~shell
curl https://rsdoiel.github.io/scripttool/installer.sh | sh
~~~

This will install the programs included in scripttool in your `$HOME/bin` directory.

If you are running Windows 10 or 11 use the Powershell command below.

~~~ps1
irm https://rsdoiel.github.io/scripttool/installer.ps1 | iex
~~~

### If your are running macOS or Windows

You may get security warnings if you are using macOS or Windows. See the notes for the specific operating system you're using to fix issues.

- [INSTALL_NOTES_macOS.md](INSTALL_NOTES_macOS.md)
- [INSTALL_NOTES_Windows.md](INSTALL_NOTES_Windows.md)

Installing from source
----------------------

### Required software


### Steps

1. git clone https://github.com/rsdoiel/scripttool
2. Change directory into the `scripttool` directory
3. Make to build, test and install

~~~shell
git clone https://github.com/rsdoiel/scripttool
cd scripttool
make
make test
make install
~~~

