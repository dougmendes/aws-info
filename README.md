# AWS Info

This project is an example of how to retrieve various information from AWS, such as IAM users, their attached policies, and associated groups, using the AWS SDK for Go.

## Project Structure

```plaintext
aws-info/
│
├── config/
│   └── config.go
│
├── services/
│   ├── iam_service.go
│
└── main.go
```

## Prerequisites

Before you begin, ensure you have the following installed:

- Go
- AWS CLI configured with your credentials

## How to run

1 - Clone the repository to your local directory:
```
git clone https://github.com/your-username/aws-info.git
cd aws-info
```
2 - Ensure that Go is installed and configured correctly on your machine.

3 - Update the Go modules:
```
go mod tidy
```

4 - Run the application:
```
go run main.go
```

## Key Functions

### Configuration

The LoadAWSConfig function in config/config.go loads the AWS SDK configuration with the specified region.

### IAM Services

The functions in services/iam_service.go interact with the AWS IAM service to list users, their attached policies, and associated groups:

- ListUsers: Lists all IAM users.
- ListAttachedUserPolicies: Lists the policies attached to a specific user.
- ListGroupsForUser: Lists the groups to which a specific user belongs.
- ListAttachedGroupPolicies: Lists the policies attached to a specific group.

### Adding Other AWS Services
To add functions for other AWS services, create a new file in the services directory, e.g., other_service.go, and implement the desired functions. Update main.go to call these new functions as needed.