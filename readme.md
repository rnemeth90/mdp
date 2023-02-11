# mdp [![build-release-binary](https://github.com/rnemeth90/mdp/actions/workflows/build.yaml/badge.svg)](https://github.com/rnemeth90/mdp/actions/workflows/build.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/rnemeth90/mdp/)](https://goreportcard.com/report/github.com/rnemeth90/mdp/)
## Description
mdp is a simple utility for generating HTML previews for Markdown files. 

## Getting Started
gopher $ mdp
mdp

Usage:
  mdp --f myfile.md
  mdp --f myfile.md --p

Options:
      --f string   the markdown file to preview
      --p          preview the file
 
### Dependencies
*  to build yourself, you must have Go v1.16+ installed

### Installing
Download the latest release [here](https://github.com/rnemeth90/mdp/releases)
 
## Help
If you need help, submit an issue to me

## To Do
- [x] Add ability to preview a file in browser
- [x] Create preview html file in temp directory
- [x] Pause while the file is being previewed
- [x] update run function to use io.writer, then update tests
- [ ] auto update the preview file
- [x] Remove preview file after exiting application
 
## Version History
*  1.0.0
    * Initial Release

## License
This project is licensed under the MIT License - see the LICENSE.md file for details
