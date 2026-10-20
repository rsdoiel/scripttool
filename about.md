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
version: 0.0.12
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.26.4

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-10-20
---

About this software
===================

## scripttool 0.0.12

Corrects the release text of v0.0.11, which said scripttool still required fdx v1.0.2 and osf v0.0.7: it requires fdx v1.0.4 and osf v0.0.9, so a Fountain scene number (INT. KITCHEN - DAY #12A#) becomes Final Draft's Number attribute in fountain2fdx and comes back in fdx2fountain. Requires fountain v1.0.3; the go directive is 1.26.4. A boneyard stays a visible paragraph in FDX and OSF, and OSF keeps no per-scene number. version.go, about.md, CITATION.cff and the installers are regenerated from codemeta.json for this release.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A library and command line program for working with fdx (Final Draft XML), osf, Fade In and Fountain formatted Screen Play files.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/scripttool)
  - [Issue Tracker](https://github.com/rsdoiel/scripttool/issues)

## Programming languages

- Go >= 1.26.4




## Software Requirements

- Go >= 1.26.4
- github.com/rsdoiel/fountain >= 1.0.3
- github.com/rsdoiel/fdx >= 1.0.4
- github.com/rsdoiel/osf >= 0.0.9


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


