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
version: 0.0.11
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

## scripttool 0.0.11

Requires fountain v1.0.3, fdx v1.0.4 and osf v0.0.9; the go directive is now 1.26.4. Fountain v1.0.3 parses scene numbers (INT. KITCHEN - DAY #12A#) and boneyards, and fdx v1.0.4 writes a scene number as Final Draft's Number attribute and reads it back, so a number survives fountain2fdx and fdx2fountain. A boneyard stays a visible paragraph in FDX and OSF, and OSF keeps no per-scene number.

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


