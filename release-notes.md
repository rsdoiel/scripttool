---
title: scripttool releases
---

Releases
========

v0.0.13, 2026-10-27, help index
-------------------------------

Adds a help index. `scripttool help index` lists the help topics (`fadein`, `fdx`, `osf`, `fountain`, `characters`, `options`, `examples`) and `scripttool help TOPIC` prints one. Before this, `help index` tried to open a file named `index` and exited 1. An unknown topic, or more than one argument after `help`, exits 2. Each topic also has a section 7 man page, `scripttool-TOPIC.7.md`, built by `make`.

The manual now documents the `fadein2yaml`, `fdx2yaml`, `fountain2yaml`, `osf2yaml`, `fountain2fountain` and `fountain2html` verbs it had left out.

Builds with Go 1.27.3 (the `go` directive and `softwareRequirements`). The Makefile gains a `check-release` target, which the `release` target runs after `generate`, and the stale `codemeta-*.tmpl` templates are removed. Requires `fountain` v1.0.3, `fdx` v1.0.4 and `osf` v0.0.9. `version.go`, `about.md`, `CITATION.cff` and the installers are regenerated from `codemeta.json` for this release.

v0.0.12, 2026-10-20, release text corrected
-------------------------------------------

Corrects v0.0.11, whose notes said `scripttool` still required `fdx` v1.0.2 and `osf` v0.0.7. It requires `fdx` v1.0.4 and `osf` v0.0.9 (and `fountain` v1.0.3), so a Fountain scene number (`INT. KITCHEN - DAY #12A#`) becomes Final Draft's `Number` attribute in `fountain2fdx` and comes back in `fdx2fountain`. A boneyard stays a visible paragraph in FDX and OSF, and OSF keeps no per-scene number. `version.go`, `about.md`, `CITATION.cff` and the installers are regenerated from `codemeta.json` for this release.

v0.0.11, 2026-10-20, fountain v1.0.3
------------------------------------

Requires `fountain` v1.0.3, `fdx` v1.0.4 and `osf` v0.0.9. The `go` directive is now 1.26.4. `fountain` v1.0.3 types a boneyard and reads a scene number written at the end of a scene heading (`INT. KITCHEN - DAY #12A#`), and `fdx` v1.0.4 writes that number as Final Draft's `Number` attribute and reads it back, so a scene number survives `fountain2fdx` and `fdx2fountain`. A boneyard stays a visible paragraph in FDX and OSF, and OSF keeps no per-scene number.

v0.0.7, 2022-08-07, project refresh
-----------------------------------

This release is mostly a code re-organization due to changes in how Go 1.18 supports modules. The packages osf, fountain, and fdx are now part of the scripttool. The modules are largely stable. It's been years since I've changed them.  The scripttool itself has bee restructure. A major change is the command line syntax. I am now using a `scripttool [-help|-version|-license] VERB [OPTIONS] [ADDITIONAL_PARAMETERS]`.  Each  "verb" can have its own set of options and  they come after the "verb" before it expected additional parameters.  

The build process for source code and website have been significantly updated. I'm using Pandoc and pdtk to build the website. The old Makefile and Python scripts have been replace by a two new Makefile (i.e. Makefile, website.mak). 

In the `scripttool` command the "characters" verb will list the characters found in a fountain formatted screenplay. By default it is order of appearance but there is now an option `-alpha` for alphabetical sorts. 

