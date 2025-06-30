
   ```
 ██████╗  ██╗  ██╗    ███████╗████████╗ █████╗ ██╗     ██╗  ██╗███████╗██████╗ 
██╔════╝ ██║  ██║    ██╔════╝╚══██╔══╝██╔══██╗██║     ██║ ██╔╝██╔════╝██╔══██╗
██║  ███╗███████║    ███████╗   ██║   ███████║██║     █████╔╝ █████╗  ██████╔╝
██║   ██║██╔══██║    ╚════██║   ██║   ██╔══██║██║     ██╔═██╗ ██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║    ███████║   ██║   ██║  ██║███████╗██║  ██╗███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝    ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝                                                                           
```
# GitHub User Activity CLI

A simple command-line tool built in Go that fetches and displays recent activity for any GitHub user using the GitHub API.

## Features

- Fetch recent public activity for any GitHub user
- Display formatted information for different event types:
  - **CreateEvent**: Shows repository creation details
  - **PushEvent**: Shows branch, commit SHA, and commit message
  - **ReleaseEvent**: Shows release information and tag details
  - **Other Events**: Shows basic event information
- Clean, readable output format
- Fast and lightweight

## Prerequisites

- Go 1.16 or later
- Internet connection to access GitHub API

## Dependencies

This project uses the following external dependency:

- `github.com/tidwall/gjson` - Fast JSON parser for Go
- `github.com/fatih/color` - Color Formatting for CLI
## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd Github-User-Activity
```

2. Install dependencies:
```bash
go mod init Github-User-Activity
go get github.com/tidwall/gjson
```

3. Build the application:
```bash
go build -o github-activity
```

## Usage

Run the application with a GitHub username as an argument:

```bash
./github-activity <username>
```
default action is to mask emails, to show email use --show-emails flag
```
./github-activity --show-emails <username>
```

### Examples

```bash

# Check activity for user "A-noob-in-Coding"
./github-activity A-noob-in-Coding
```

## Sample Output

```
=== GitHub PushEvent ===
Event ID     : 12345678901
Type         : PushEvent
Actor        : username
Repository   : username/repo-name
Branch       : refs/heads/main
Commit SHA   : abc123def456...
Commit Msg   : Fix bug in user authentication
Pushed At    : 2024-01-15T10:30:00Z

=== GitHub CreateEvent ===
Event ID     : 12345678902
Type         : CreateEvent
Actor        : username
Repository   : username/new-repo
Create Description : A new project for learning Go
Pushed At    : 2024-01-14T15:45:00Z
```

## Project Structure

```
Github-User-Activity/
├── main.go           # Main application entry point
├── utils/
│   └── utils.go      # JSON processing utilities
├── go.mod            # Go module file
├── go.sum            # Go dependencies checksum
└── README.md         # This file
```

## API Information

This tool uses the GitHub Events API:
- **Endpoint**: `https://api.github.com/users/{username}/events`
- **Documentation**: [GitHub Events API](https://docs.github.com/en/rest/activity/events)

## Supported Event Types

The application provides detailed information for the following GitHub event types:

| Event Type | Information Displayed |
|------------|----------------------|
| PushEvent | Branch, Commit SHA, Commit Message |
| CreateEvent | Description |
| ReleaseEvent | Release Body, Tag Name |
| Others | Basic event information |

## Error Handling

The application handles the following error cases:

- **Invalid arguments**: Shows usage information if username is not provided
- **Network errors**: Displays error message if GitHub API is unreachable
- **User not found**: Shows error if the specified user doesn't exist
- **API errors**: Handles non-200 HTTP responses from GitHub API

## Limitations

- Only displays public events (private repository events are not accessible)
- Limited to the most recent events returned by GitHub API (typically up to 300 events)

## [Project URL](https://roadmap.sh/projects/github-user-activity)

