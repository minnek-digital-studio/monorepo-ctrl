# Monorepo Control

Monorepo Control is a command-line tool that simplifies the process of managing monorepos by automating tasks, such as running pre-commit hooks for modified packages. It's designed to work with Husky and other Git hooks tools.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Add Monorepo Control to your PATH](#add-monorepo-control-to-your-path)
  - [Standalone usage](#standalone-usage)
- [Usage](#usage)
  - [Integration with pre-commit hook](#integration-with-pre-commit-hook)
    - [Global installation (added to PATH)](#global-installation-added-to-path)
    - [Standalone installation](#standalone-installation)
- [Contributing](#contributing)
- [License](#license)
## Features

- Automatically detects modified packages and runs pre-commit hooks.
- Customizable configuration to specify workspaces, extensions, and commands.
- Easy to integrate with existing Git hooks management tools.
- Supports passing a custom configuration file path.
- Can be used in "standalone" mode without adding it to the PATH.

## Installation

To install Monorepo Control, follow these steps:

1. Clone the repository:

   ```
   git clone git@github.com:minnek-digital-studio/monorepo-ctrl.git
   ```

2. Change to the repository directory:

   ```
   cd monorepo-ctrl
   ```

3. Build the project:

   ```
   go build -o monorepo-ctrl ./cmd
   ```

### Add Monorepo Control to your PATH

To use Monorepo Control globally, add the compiled binary to your `PATH`:

```
export PATH=$PATH:/path/to/monorepo-ctrl
```

### Standalone usage

For standalone usage, you can place the compiled binary and the configuration file in the `.husky` directory and use it without adding it to the PATH.

## Usage

1. Create a `mnk-config.json` file in the root of your monorepo with the following structure:

   ```json
   {
     "monorepo-ctrl": {
       "global": {
         "workspaces": ["workspaces", "packages"],
         "extensions": [".js", ".ts", ".tsx"]
       },
       "configs": [
         {
           "name": "scopeName",
           "commands": ["npm run lint", "npm run test"]
         }
       ]
     }
   }
   ```

   - Replace `scopeName` with the desired name for your scope.
   - Customize the `workspaces`, `extensions`, and `commands` as needed.

2. Run Monorepo Control with the desired scope:

   ```
   monorepo-ctrl scopeName
   ```

   Replace `scopeName` with the name you defined in the `mnk-config.json` file. A good example for a scope name is the type of hook you want to run, such as `commit` or `push`.

   If you want to use a custom configuration file path, use the `--config` or `-c` option:

   ```
   monorepo-ctrl scopeName --config path/to/config/file.json
   ```

   If Monorepo Control is not in your PATH and you have to add the binary and the `mnk-config.json` file to the `.husky` directory, you can run it in standalone mode without adding it to the PATH.

### Integration with pre-commit hook

#### Global installation (added to PATH)

To use Monorepo Control in your pre-commit hook when it is installed globally, create a `.husky/pre-commit` file with the following content:

```bash
#!/bin/bash

# Execute the monorepo-ctrl with the "commit" argument and custom configuration file
monorepo-ctrl scopeName
```

This will ensure that Monorepo Control is executed with the specified scope and configuration file during the pre-commit process.

#### Standalone installation

To use Monorepo Control in your pre-commit hook in standalone mode, create a `.husky/pre-commit` file with the following content:

```bash
#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Give execute permission to the monorepo-ctrl file
chmod +x $SCRIPT_DIR/monorepo-ctrl

# Execute the monorepo-ctrl file with the "commit" argument and custom configuration file
$SCRIPT_DIR/monorepo-ctrl scopeName -c "$SCRIPT_DIR/mnk-config.json"
```

This will ensure that Monorepo Control is executed with the specified scope and configuration file during the pre-commit process when used in standalone mode, without adding it to the PATH.

## Contributing

Contributions are welcome! If you find a bug, have a feature request, or want to improve the documentation, please create an issue or submit a pull request.

## License

Monorepo Control is released under the [MIT License](https://opensource.org/licenses/MIT).