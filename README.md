{
    "title": "scripttool, a program for converting screen play formats"
}
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

# scripttool

A tool for working with screenplay file formats (e.g. fdx, fountain)

## Converting formats

Convert from *screenplay.fdx* to *screenplay.fountain*

```shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
```

Convert from *screenplay.fadein* to *screenplay.fountain*

```shell
    scripttool fadein2fountain screenplay.fadein screenplay.fountain
```

Convert from *screenplay.fountain* to *screenplay.osf*

```shell
    scripttool fountain2osf screenplay.foutain screenplay.osf
```


## Script Reports

This is still in the planning stage. List the characters in *screenplay.fdx*

```shell
    scripttool characters screenplay.fdx
```

