# Change log for cwcomp
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning].
The format is based on [Keep a Changelog].
	
## [Unreleased]

- Renamed `grid*` to `puzzle*`

## [v0.6.0] - 2023-05-27
Added feature to create SVG from a graph

## [v0.5.0] - 2023-05-21
Development version with a complete `model` package.
- Added database support for saving and reloading grids
- Added tool for creating initial database
- Added support for suggesting words that match a pattern
- Added support for finding the constraints imposed by crossing words

## [v0.4.0] - 2023-05-16
Development version with a number of changes to the data model:
- Modified the `WordNumber` structure to hold just the sequence number and starting point.
- Created the `Word` type to hold the starting point, direction, and clue.
This is easier to work with than storing across and down words separately
and always requiring a type switch.
- Created a grid `words` slice to hold all words
- Changed the word number map to just a `wordNumbers` slice of word number pointers
- Added methods to lookup words
- Kept test coverage up to 100%

## [v0.3.0] - 2023-05-14
- Added across and down clues to WordNumber
- Added regular expression package

## [v0.2.0] - 2023-05-14
Development version with black cell undo/redo

## [v0.1.0] - 2023-05-10
Development version with complete grid package

[Semantic Versioning]: http://semver.org
[Keep a Changelog]: http://keepachangelog.com
[Unreleased]: https://github.com/philhanna/cwcomp/compare/v0.6.0..HEAD
[v0.6.0]: https://github.com/philhanna/cwcomp/compare/v0.5.0..v0.6.0
[v0.5.0]: https://github.com/philhanna/cwcomp/compare/v0.4.0..v0.5.0
[v0.4.0]: https://github.com/philhanna/cwcomp/compare/v0.3.0..v0.4.0
[v0.3.0]: https://github.com/philhanna/cwcomp/compare/v0.2.0..v0.3.0
[v0.2.0]: https://github.com/philhanna/cwcomp/compare/v0.1.0..v0.2.0
[v0.1.0]: https://github.com/philhanna/cwcomp/compare/4e55e9e..v0.1.0
