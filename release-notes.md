---
title: scripttool releases
---

Releases
========

v0.0.7, 2022-08-07
------------------

This release is mostly a code re-organization due to changes in how Go 1.18 supports modules. The packages osf, fountain, and fdx are now part of the scripttool. The modules are largely stable. It's been years since I've changed them.  The scripttool itself has bee restructure. A major change is the command line syntax. I am now using a `scripttool [-help|-version|-license] VERB [OPTIONS] [ADDITIONAL_PARAMETERS]`.  Each  "verb" can have its own set of options and  they come after the "verb" before it expected additional parameters.  

The build process for source code and website have been significantly updated. I'm using Pandoc and pdtk to build the website. The old Makefile and Python scripts have been replace by a two new Makefile (i.e. Makefile, website.mak). 

In the `scripttool` command the "characters" verb will list the characters found in a fountain formatted screenplay. By default it is order of appearance but there is now an option `-alpha` for alphabetical sorts. 

