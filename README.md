# Welcome to dirf 2.0

## what is dirf

- Dirf is simple command line utility to make multiple,nested, recursive directory and files inside them

## Usage

- Paths can be passed as arguments
  `$ dirf a/b+c/d+e a/b/c`
  -Paths can be passed from standard input
  `$ dirf --noArgs`
- Here d and e are files and rest are directories.
- Below command will create a nested directory like this
  `$ dirf a/b+c/d+e `
  <img src="./screenshots/screenshot_1.png" alt="tree view of created directory">

## Build

- Run
  `go build -o dirf`
- Move executable in the user binary folder for global uses.

### Example

- `a/` will create a directory
- `b` will create a file.
- `a/b` will create a file `b` inside directory `a`.
- `a/b+c/` will create a directory `b` and `c` inside directory `a`
- `a/b+c` will create a file `b` and `c`inside directory `a`
- `a/$.component.jsx+$.styles.css ---> a/a.component.jsx a/a.styles.css` $ will replaced by parent name

### Whats' new

- Rewritten in go
- can pass multiple paths
- can use $ expansion
- have help menu (-h,--help)

## Development

- This is improved version of previous dirf which is written in python.
- Try to cleanup code
- I will try to add more features.
