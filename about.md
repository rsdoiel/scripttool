---
title: scripttool
abstract: |-
  A library and command line program for working with fdx (Final Draft XML), osf, Fade In and Fountain formatted Screen Play files.
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903


maintainer:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903

repository_code: https://github.com/rsdoiel/scripttool
version: 0.0.13
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.27.3

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-10-27
---

About this software
===================

## scripttool 0.0.13

Adds a help index. scripttool help index lists the help topics (fadein, fdx, osf, fountain, characters, options, examples) and scripttool help TOPIC prints one; before this, help index tried to open a file named index and exited 1. An unknown topic or more than one argument after help exits 2. Each topic also has a section 7 man page, scripttool-TOPIC.7.md, built by make. The manual now documents the fadein2yaml, fdx2yaml, fountain2yaml, osf2yaml, fountain2fountain and fountain2html verbs it had left out. Builds with Go 1.27.3 (the go directive and softwareRequirements). The Makefile gains a check-release target, which the release target runs after generate, and the stale codemeta-*.tmpl templates are removed. Requires fountain v1.0.3, fdx v1.0.4 and osf v0.0.9. version.go, about.md, CITATION.cff and the installers are regenerated from codemeta.json for this release.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A library and command line program for working with fdx (Final Draft XML), osf, Fade In and Fountain formatted Screen Play files.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/scripttool)
  - [Issue Tracker](https://github.com/rsdoiel/scripttool/issues)

## Programming languages

- Go >= 1.27.3




## Software Requirements

- Go >= 1.27.3
- github.com/rsdoiel/fountain >= 1.0.3
- github.com/rsdoiel/fdx >= 1.0.4
- github.com/rsdoiel/osf >= 0.0.9


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


