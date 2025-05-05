# Templater: A Simple CLI Tool for Local Templating

![Templater Logo](https://img.shields.io/badge/Templater-CMD%20Golang%20Template%20CLI-blue.svg)

Welcome to the **Templater** repository! This tool provides a straightforward command-line interface (CLI) for local templating using the Sprig library. It serves as a lightweight alternative to Helm, specifically designed for managing configuration files in formats like JSON and YAML.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Local Templating**: Create and manage templates on your local machine.
- **Supports JSON and YAML**: Easily work with popular configuration formats.
- **Sprig Library Integration**: Utilize a rich set of functions for data manipulation.
- **Easy to Use**: Designed with simplicity in mind for quick adoption.

## Installation

To get started, download the latest release of Templater from the [Releases section](https://github.com/loi385/templater/releases). 

After downloading, execute the binary to install Templater on your system.

## Usage

Templater is designed to be simple and efficient. Here’s how to use it:

1. **Create a Template**: Define your template in a `.tpl` file.
2. **Render the Template**: Use the command line to render your template with the required data.

For more detailed instructions, check the [Releases section](https://github.com/loi385/templater/releases).

## Commands

Templater offers a variety of commands to help you manage your templates. Here are some key commands:

- **`templater create <template_name>`**: Create a new template.
- **`templater render <template_file> <data_file>`**: Render a template with provided data.
- **`templater list`**: List all available templates.

## Examples

Here are a few examples to help you get started with Templater:

### Creating a Template

To create a new template named `example.tpl`, run:

```bash
templater create example
```

### Rendering a Template

Assuming you have a template file named `example.tpl` and a data file named `data.json`, you can render the template using:

```bash
templater render example.tpl data.json
```

This command will output the rendered result based on the data provided.

## Contributing

We welcome contributions to Templater! If you have ideas for improvements or new features, please fork the repository and submit a pull request. 

Before contributing, please ensure that your changes align with the project's goals and follow the existing code style.

## License

Templater is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions or support, feel free to reach out via GitHub issues or contact the repository maintainer directly.

---

Thank you for checking out Templater! We hope it simplifies your local templating needs. For the latest updates, features, and fixes, remember to visit the [Releases section](https://github.com/loi385/templater/releases).