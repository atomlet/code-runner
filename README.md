**Code Runner** is a command-line tool that executes specific commands based on file extensions. It reads commands related to file types from a configuration file, sets environment variables for each file, and runs the commands via the Windows command prompt.

## Features

- **Read Configuration File**: The program reads a config.json file that maps file extensions to corresponding commands.
- **Set Environment Variables**: The program automatically sets environment variables based on the file path, such as file path, file name, and file directory.
- **Execute Commands**: Based on the file extension, the program retrieves the corresponding command from the configuration file and executes it via the Windows command prompt.
- **Customizable Commands**: You can easily modify the config.json file to define different commands for different file extensions.

## Installation

1. Clone this repository:

```sh
git clone https://github.com/yourusername/code-runner.git
cd code-runner
```

2. Build the program:

```sh
go build
```

## Usage

1. Create a config.json file in the root directory of the project, using the following format:

```json
{
  ".c": "cd {DIR} && gcc {FILE_PATH} -o {FILE_NAME} && {FILE_NAME}"
}
```

2. Run the program with the file path as an argument:

```sh
code-runner example.c
```

3. The program will look up the corresponding command based on the file's extension and execute it in the Windows command prompt.

## Environment Variables

The program automatically sets the following environment variables for use by the command:

- `FILE_PATH`: The absolute path of the file
- `FILE_NAME`: The file name (excluding the extension)
- `DIR`: The directory of the file

## Example

Suppose you have a Go file example.go, and the config.json file is configured with the Go command as follows:

```json
{
  ".go": "go run {FILE_PATH}"
}
```

You can execute it by running:

```sh
code-runner example.go
```