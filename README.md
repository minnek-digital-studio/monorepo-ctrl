# Monorepo Control

Monorepo Control is a command-line tool that simplifies the process of managing monorepos by automating tasks, such as running pre-commit hooks for modified packages. It's designed to work with Husky and other Git hooks tools.

## Features

- Automatically detects modified packages and runs pre-commit hooks.
- Customizable configuration to specify workspaces, extensions, and commands.
- Easy to integrate with existing Git hooks management tools.

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

4. Add the compiled binary to your `PATH`:

   ```
   export PATH=$PATH:/path/to/monorepo-ctrl
   ```

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

   Replace `scopeName` with the name you defined in the `mnk-config.json` file.

## Contributing

Contributions are welcome! If you find a bug, have a feature request, or want to improve the documentation, please create an issue or submit a pull request.

## License

Monorepo Control is released under the [MIT License](https://opensource.org/licenses/MIT).