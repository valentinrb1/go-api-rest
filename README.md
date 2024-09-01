# GO API REST using Gin

## Date - 2023

The program creates two web services using the Gin framework: a user service and a processing service. These services run on different ports and have different paths and handlers associated with them. The program also sets up log files and redirects standard output and error output to these files.

### Authors:
- **Robledo, Valentín**

## Summary
In this project, a server was implemented that has a user service and a processing service.

The user service has two functions, one to create a user and another to obtain a list of users registered in the system.

On the other hand, the processing service has a function to send information about a process and another to obtain a list of processes.

## How to clone this repository?
You can clone this repository to any directory you want using the following command:

```console
git clone https://github.com/valentinrb1/go-api-rest.git
```

## How to use?
First, we move to the project folder.

```console
cd go-api-rest
```

Within the project directory we can use the functions of the *make* command.

*make build*, if we want to compile the project.
```console
make build
```

If we want to run the server, we can use the *make install* command that compiles the project, checks dependencies and sets the *Nginx* and *Systemd* configurations.
```console
make install
```

To run the server:

```console
lab6
```

---
## Operation
As mentioned above, the server has two services, one for users and one for processing.

### User service
The client can make two types of requests to the user service, one to create a user and another to obtain a list of users registered in the system.

These *endpoints* cannot be accessed by an unauthenticated user, therefore, the client must first make an authentication *request*. Then, the server will return a *token* that must be used in the *requests* of the user *endpoints*. To achieve this, the Jason Web Token (JWT) protocol was used.

The *requests* are the following:
- POST dashboard.com/api/users/login
- POST dashboard.com/api/users/createuser
- GET dashboard.com/api/users/listall

### Processing service
The processing service has two *endpoints*, one used by the sensors to send information about a processing and one used by the client to obtain the information about these processings.

The *requests* are the following:
- POST sensors.com/api/processings/submit
- GET sensors.com/api/processings/summary

## Configuration
The files in the *config* directory are used to configure the server.

- *lab6.conf*: *Nginx* configuration file. Allows us to configure the server's port and domain.
- *lab6.service*: *Systemd* configuration file. Allows us to add the service to the system.

## Testing
The *Postman* software was used to perform the *testing*. It allows us to make *requests* to the server's *endpoints*.
