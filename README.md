<h1 align="center"> BLUECCPRINT </h1>

A CLI utility to create a new C++ projects with a predefined structure inspired by Melkeydev's go-blueprint built in Go.

## Setup

### Requirements

* go
* make

Clone the project

```bash
  $ git clone https://github.com/ErlanRG/bluecpprint
```

Go to the project directory

```bash
  $ cd bluecpprint
```

Build and install

```bash
  $ make && make install
```

The binary will be installed in `$HOME/.local/bin/`

## Usage

Create a project

```bash
  $ bluecpprint <project_name>
```

## Structure

Bluecpprint generates a simple project structure with some extra features:

```
my_project
├── .git
├── bin
├── include
├── src/
│   └── main.cpp
├── .clang-format
├── .gitignore
└── Makefile
└── README.md
```

* Git repository automatically created.
* Microsoft's clang-format configuration.
* Makefile with build, clean and run scripts right of the bat

## Acknowledgements

 - [go-blueprint](https://github.com/Melkeydev/go-blueprint) - Melkeydev
 - [readme.so](https://github.com/octokatherine/readme.so) - octokatherine

## Feedback

If you have any feedback, please reach out to me at erlanrangel@gmail.com

## Contributing

Pull requests are more than welcome.
For major changes, please submit a new issue.

## License

[MIT](https://choosealicense.com/licenses/mit/)

