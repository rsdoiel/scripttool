---
title: scripttool, a program for converting screen play formats
---

About the test data
===================

The testing data for scripttool and its sub-packages fountain, fdx and
osf are held in common in this repository's _testdata_ directory.

The [fountain](https://fountain.io) website has some good files 
for reviewing formatting difference between fountain, fdx and 
PDF. Some are referenced in the test programs but they are 
optional.  It is not clear to me the licensing arrangements for 
the text so I have not included them in this repository. You can 
use the *bash* script `get-optional-testdata.bash` to retrieve 
them from the fountain.io website.

If you want to include them in the test sequence go to the 
fountain website and download them and place them in the _testdata_ 
directory.  When you run `go test` they will be found and included 
in the basic test process.

Optional test FDX files
-----------------------

+ [Big Fish](https://fountain.io/_downloads/Big%20Fish.fountain)
    + [fdx](https://fountain.io/_downloads/Big%20Fish.fdx)
    + [pdf](https://fountain.io/_downloads/Big%20Fish.pdf) 
+ [Brick & Steel](https://fountain.io/_downloads/Brick%20&%20Steel.fountain)
    + [fdx](https://fountain.io/_downloads/Brick%20&%20Steel.fdx)
    + [pdf](https://fountain.io/_downloads/Brick%20&%20Steel.pdf)
+ [Birthday Card](https://fountain.io/_downloads/The%20Last%20Birthday%20Card.fountain)
    + [fdx](https://fountain.io/_downloads/The%20Last%20Birthday%20Card.fdx)
    + [pdf](https://fountain.io/_downloads/The%20Last%20Birthday%20Card.pdf)

