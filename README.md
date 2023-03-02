# Simple Security Checker

A simple security checker for URLs that need to be closed (More 200+ common URLs). This tool allows you to check a list of URLs for a specific response code (e.g. 200 OK) and logs any errors or successful responses. This is useful for monitoring open URLs that need to be closed to maintain security.

## Detailed Description

This security checker is designed to help identify and monitor URLs that should not be publicly accessible. By default, the tool checks a list of URLs for a specific response code (200 OK), but this can be easily modified in the code to check for other response codes as well.

The tool is useful for monitoring open URLs that need to be closed in order to maintain security. For example, if your website has an admin panel at `/admin` that should only be accessible to authorized users, you can use this tool to check that the admin panel is not accessible to the public.

To use the tool, simply clone the repository to your local machine, install Go if you haven't already, and run `go run main.go` from the project directory. You can then modify the list of URLs to check by editing the `config.json` file.

Overall, this tool is a simple yet effective way to monitor your website's security and ensure that sensitive URLs are not publicly accessible.

## Usage

1. Clone the repository to your local machine.

```sh
git clone https://github.com/sv222/simple-security-checker.git
```

2. Install Go if you haven't already: `https://golang.org/doc/install`
3. Open a terminal and navigate to the project directory.
4. Run `go run main.go` to start the security checker.
5. To modify the list of URLs to check, edit the `config.json` file in the project directory.

### URL Slugs

More then 200 URL slugs are provided in the `config.json` file:

- /login
- /admin
- /dashboard
- /settings
- /users
....
etc.

## Build and Usage as Compiled CLI

To build the project as a standalone CLI tool, follow these steps:

1. Clone the repository to your local machine.
2. Install Go if you haven't already: `https://golang.org/doc/install`
3. Open a terminal and navigate to the project directory.
4. Run the command `go build -o security-checker main.go` to build the executable file.
5. Run the executable file with `./security-checker`.

You can also specify a custom config file with the `-config` flag, e.g. `./security-checker -config=myconfig.json`. By default, the tool looks for a `config.json` file in the same directory as the executable.

Overall, building the tool as a standalone CLI provides a convenient way to run the security checker without having to install Go on your machine or use Docker. This is useful if you need to run the tool on a machine that doesn't have Go installed, or if you want to distribute the tool to other users who may not be familiar with Go.

## Building and Usage with Docker

To build and run the security checker with Docker, follow these steps:

1. Clone the repository to your local machine.
2. Install Docker if you haven't already: `https://docs.docker.com/get-docker/`
3. Open a terminal and navigate to the project directory.
4. Build the Docker image with the following command: `docker build -t security-checker .`
5. Run the Docker container with the following command: `docker run -v $(pwd)/config.json:/app/config.json security-checker`
6. To modify the list of URLs to check, edit the `config.json` file on your local machine.

The Docker container runs the security checker using the `config.json` file mounted as a volume, so any changes you make to the file on your local machine will be reflected in the Docker container.

By using Docker, you can easily build and deploy the security checker in any environment without needing to worry about installing Go or managing dependencies.

## Contribution

If you would like to contribute to the Simple Security Checker, you can fork the project on GitHub and create a pull request with your changes. Contributions are always welcome, and we appreciate your help in making the program better.

## License

The Simple Security Checker is licensed under the MIT License.
