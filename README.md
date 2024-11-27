<div align="left">
    <img src="https://avatars.githubusercontent.com/u/179362694?s=200&v=4" width="40%" align="left" style="margin-right: 15px"/>
    <div style="display: inline-block;">
        <h2 style="display: inline-block; vertical-align: middle; margin-top: 0;">GO-AUTH-SERVICE</h2>
        <p>
        This project is an Authentication Microservice built using Golang, designed to handle registration, login, and OAuth2 authentication processes using Google and GitHub as third party authorization providers. This project focuses on security, performance, and flexibility to be integrated with other microservices.
</p>
        <p>
	<img src="https://img.shields.io/github/license/tuxedo-labs/go-auth-service?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/tuxedo-labs/go-auth-service?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/tuxedo-labs/go-auth-service?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/tuxedo-labs/go-auth-service?style=default&color=0080ff" alt="repo-language-count">
</p>
        <p><!-- default option, no dependency badges. -->
</p>
        <p>
</p>
    </div>
</div>
<br clear="left"/>

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

<code>❯ REPLACE-ME</code>

---

##  Features

<code>❯ REPLACE-ME</code>

---

##  Project Structure

```sh
└── go-auth-service/
    ├── .github
    │   └── workflows
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── config
    │   └── config.go
    ├── docker-compose.yml
    ├── docs
    │   └── ROUTE.md
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── handler
    │   ├── middleware
    │   ├── models
    │   ├── repository
    │   └── service
    ├── main.go
    └── pkg
        ├── client
        ├── provider
        └── utils
```


###  Project Index
<details open>
	<summary><b><code>GO-AUTH-SERVICE/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/main.go'>main.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/go.mod'>go.mod</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/go.sum'>go.sum</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/Makefile'>Makefile</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/docker-compose.yml'>docker-compose.yml</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/Dockerfile'>Dockerfile</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- .github Submodule -->
		<summary><b>.github</b></summary>
		<blockquote>
			<details>
				<summary><b>workflows</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/.github/workflows/ci.yml'>ci.yml</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- config Submodule -->
		<summary><b>config</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/config/config.go'>config.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- pkg Submodule -->
		<summary><b>pkg</b></summary>
		<blockquote>
			<details>
				<summary><b>provider</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/pkg/provider/google.go'>google.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/pkg/provider/github.go'>github.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>client</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/pkg/client/client.go'>client.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>utils</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/pkg/utils/jwt.go'>jwt.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/pkg/utils/hash.go'>hash.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- internal Submodule -->
		<summary><b>internal</b></summary>
		<blockquote>
			<details>
				<summary><b>models</b></summary>
				<blockquote>
					<details>
						<summary><b>entity</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/models/entity/user.go'>user.go</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
					<details>
						<summary><b>request</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/models/request/user.go'>user.go</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<details>
				<summary><b>middleware</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/middleware/middleware.go'>middleware.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>repository</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/repository/auth_repository.go'>auth_repository.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>service</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/service/auth_service.go'>auth_service.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>handler</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/tuxedo-labs/go-auth-service/blob/master/internal/handler/auth_handler.go'>auth_handler.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with go-auth-service, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules
- **Container Runtime:** Docker


###  Installation

Install go-auth-service using one of the following methods:

**Build from source:**

1. Clone the go-auth-service repository:
```sh
❯ git clone https://github.com/tuxedo-labs/go-auth-service
```

2. Navigate to the project directory:
```sh
❯ cd go-auth-service
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker build -t tuxedo-labs/go-auth-service .
```




###  Usage
Run go-auth-service using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker run -it {image_name}
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```


---
##  Project Roadmap

- [X] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

##  Contributing

- **💬 [Join the Discussions](https://github.com/tuxedo-labs/go-auth-service/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/tuxedo-labs/go-auth-service/issues)**: Submit bugs found or log feature requests for the `go-auth-service` project.
- **💡 [Submit Pull Requests](https://github.com/tuxedo-labs/go-auth-service/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/tuxedo-labs/go-auth-service
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/tuxedo-labs/go-auth-service/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=tuxedo-labs/go-auth-service">
   </a>
</p>
</details>

---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

---
