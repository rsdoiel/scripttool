
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/scripttool)

scripttool
==========

A tool for working with screenplay file formats (e.g. fdx, fountain, osf, FadeIn)

Converting to fountain format
-----------------------------

Convert from *screenplay.fdx* to *screenplay.fountain*

~~~shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
~~~

Convert from *screenplay.fadein* to *screenplay.fountain*

~~~shell
    scripttool fadein2fountain screenplay.fadein screenplay.fountain
~~~

Convert from *screenplay.osf* to *screenplay.fountain*

~~~shell
    scripttool osf2fountain screenplay.osf screenplay.fountain
~~~


Working with fountain files
---------------------------

Pretty print fountain files

~~~shell
    scripttool fountainfmt screenplay.fountain
~~~

Render a fountain file as JSON

~~~shell
    scripttool fountain2json screenplay.fountain
~~~


Convert from fountain format
----------------------------

Convert from *screenplay.fountain* to *screenplay.fdx*

~~~shell
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
~~~

Convert from *screenplay.fountain* to *screenplay.fadein*

~~~shell
    scripttool fountain2fadein screenplay.fountain screenplay.fadein
~~~

Convert from *screenplay.fountain* to *screenplay.osf*

~~~shell
    scripttool fountain2osf screenplay.fountain screenplay.osf
~~~



Script Reports
--------------

NOTE: Currently only fountain documents are supported for reporting. Currently the character list report is implemented as a proof of concept.

List the characters in *screenplay.fountain*

~~~shell
    scripttool characters screenplay.fountain
~~~

