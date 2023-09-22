# Spellchecker Git Commit Hook
![](https://github.com/DevitoDbug/ruth/blob/master/resources/screenshot.jpg?raw=true)

This is a simple command-line tool that checks the spelling in a Git commit message using the apilayer Spellchecker API and automatically commits the changes if there are no spelling errors in the commit message. If there are spelling errors, it displays the mistakes and suggestions to the user.

## Prerequisites
Before using this tool, make sure you have the following prerequisites installed:

* Go programming language (https://golang.org/doc/install)
* Git (https://git-scm.com/downloads)

## Installation
1. Clone this repository to your local machine:

2. Change into the project directory:

3. Build the Go executable:

4. Move the executable to a directory in your PATH (e.g., /usr/local/bin/):

## Usage
To use this tool, follow these steps:

1. Ensure you are inside a Git repository where you want to enforce spelling checks for commit messages.

2. When you make a commit, use the __*ruth*__ command followed by your commit message. For example:
   + __*ruth*__ fix buggs in the code

3. Replace "fix buggs in the code" with your actual commit message.

- The tool will check the spelling of your commit message using the apilayer Spellchecker API.
- If there are no spelling errors, it will automatically add and commit your changes
- If there are spelling errors, it will display the mistakes and suggestions for correction, and the commit will not be made. You can then edit your commit message accordingly.

## Configuration
You can configure the tool by modifying the following lines in the main function of the code:

- **url**:This is the API endpoint for the apilayer Spellchecker API. You can replace it with your own endpoint if needed.

- **apikey**: Replace the API key in the req.Header.Set("apikey", "your-api-key") line with your apilayer API key.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
Thanks to apilayer for providing the Spellchecker API used in this tool.
This tool is inspired by the need to maintain consistent and error-free commit messages in Git repositories.



