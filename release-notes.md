---
title: scripttool releases
---

Releases
========

v0.0.11, 2026-10-20, fountain v1.0.3
------------------------------------

Requires `fountain` v1.0.3, which types a boneyard and reads a scene number written at the end of a scene heading (`INT. KITCHEN - DAY #12A#`). The `go` directive is now 1.26.4. Conversions are unchanged: a boneyard stays a visible paragraph in FDX and OSF, and a scene number stays in the heading text. `fdx` v1.0.4 maps the number to Final Draft's `Number` attribute; `scripttool` still requires `fdx` v1.0.2 and `osf` v0.0.7, so that arrives when `go.mod` is moved to the new `fdx`.

v0.0.7, 2022-08-07, project refresh
-----------------------------------

This release is mostly a code re-organization due to changes in how Go 1.18 supports modules. The packages osf, fountain, and fdx are now part of the scripttool. The modules are largely stable. It's been years since I've changed them.  The scripttool itself has bee restructure. A major change is the command line syntax. I am now using a `scripttool [-help|-version|-license] VERB [OPTIONS] [ADDITIONAL_PARAMETERS]`.  Each  "verb" can have its own set of options and  they come after the "verb" before it expected additional parameters.  

The build process for source code and website have been significantly updated. I'm using Pandoc and pdtk to build the website. The old Makefile and Python scripts have been replace by a two new Makefile (i.e. Makefile, website.mak). 

In the `scripttool` command the "characters" verb will list the characters found in a fountain formatted screenplay. By default it is order of appearance but there is now an option `-alpha` for alphabetical sorts. 

