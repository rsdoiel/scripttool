
# Action Items

## Bugs

## Next

+ [ ] Review XML output from ToXML() after a FromFountain() call, see where I need to add mapping for Text elements and embedded styling
+ [ ] validate that I am producing fdx files that Final Draft, FadeIn and Trelby can read

## Someday, Maybe

+ [ ] Write and fdx2html using [scrippets](https://fountain.io/scrippets) approach
+ [ ] Left/Right alignment should be respected based based on Paragraph Type
+ [ ] Plaintext formatting needs to be pickup and respected from whole FinalDraft document (e.g. respect definitions, Layout, etc)
+ [ ] Screen Headers and Footers can have Text, Dynamic, SceneProperties in any order, right now converting back to XML renders them in fixed order because they are ignored when rendering and plaintext

## Completed

+ [x] Add ParseFile() to fdx.go
+ [x] String() Paragraph needs to handle trailing new lines based on Paragraph Type
+ [x] write tests that validate the source FDX content in _testout_ version
