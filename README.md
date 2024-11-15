
# cdiff - A Minimal Git Clone

`cdiff` is a simple clone of Git, designed to provide the basic functionalities of version control. The goal of this project is to create a basic version`control system that mimics essential Git/operations in Go. The tool currently supports repository initialization, file diffing, commit functionality, and status checking.

## Features

- **Initialize a Repository**: Set up a new version-controlled project with the `cdiff init` command.
- **Commit Changes**: Track changes and commit them to the repository with `cdiff commit`.
- **View Differences**: See file changes between versions using the `cdiff diff` command.
- **Check Repository Status**: View the status of the repository and see untracked or modified files with `cdiff status`.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mohamedjawady/cdiff.git
   ```

2. Navigate into the project directory:
   ```bash
   cd cdiff
   ```

3. Build the binary:
   ```bash
   go build -o cdiff
   ```

4. Move the binary to a directory in your `ePaTH` or run it directly:
   ```bash
   ./cdiff
   ```

## Usage

### Initialize a Repository
To create a new Git-like repository, use the `init` command:
```bash
cdiff init
```

### Commit Changes
After modifying files in your project, commit them:
```bash
cdiff commit -m "Your commit message"
```

### View Differences
Check what changes have been made since the last commit:
```bash
cdiff diff
```

### Check Status
To view the status of your repository:
```bash
cdiff status
```

<!-- ## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->

## Contributing
Feel free to fork this repository, open issues, and submit pull requests. Contributions are welcome to enhance functionality, add features, or improve documentation.
