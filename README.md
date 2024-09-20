# Mail Client 

This is a simple email client application with a RESTful API backend written in Go. 
It supports multiple email providers, with current implementation focused on ProtonMail

**!!This is a work in progress**

## Goal

The overall goal of this project is to make a feature rich and extensible email client backend that can be used to aggregate all email services in one frontend. 

## Future Features

* RESTful API for email operations
* Support for multiple email providers through SMTP and IMAP protocols and specific Google, ProtonMail, and Microsoft Exchange implementations
* JWT-based authentication
* Email listing, viewing, and sending

## Prerequisites

* Go 1.16 or higher

## Installation:

1. Clone the repository:

```bash
git clone https://github.com/gituser12981u2/mailclient.git
```

```bash
cd mailclient
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create a `config.json` file in the project root with the following structure:

```json
{
    "server_address": ":8080",
    "provider:" "proton",
    "provider_config": {
        "username": "your_protonmail_username",
        "password": "your_protonmail_password"
    },
    "jwt_secret": "your_jwt_secret"
}
```

## Running the Application 

To start the server, run:

```bash
go run main.go
```

## API Endpoints

* GET `/api/v1/emails`: List emails
* GET `/api/v1/emails/:id`: Get a specific email 
* POST `/api/v1/emails`: Send a new email 

## Project Structure

* main.go: Entry point of the application 
* `internal/`: Contains the core application code
    * `api/`: API-related code including route setup
    * `config/`: Configuration loading and management
    * `models/`: Data models
    * `providers/`: Email provider implementations
    * `services/`: Business logic services

## Contributing

Contributions are welcome. Please feel free to sumbit a Pull Request. 

## License

This project is licensed under the MIT License.
