### Lab6 Sistemas Operativos II
## Ingeniería en Computación - FCEFyN - UNC

## Introduction
Embedded systems are often accessed remotely. There are different techniques to do so, one widely used way is the RESTful APIs. These provide a defined and robust interface for remote communication and manipulation of the embedded system. Defined for a Client-Server scheme, they are used in all verticals of the technology industry, from IoT applications to multiplayer games.

## Objective
The objective of this practical work is for the student to have an end-to-end vision of a basic implementation of a RESTful API on an embedded system working with applications that resemble professional life, through different language stacks.
The student must implement it by interacting with all layers of the process. From functional testing (high level) to the C and Go code of the service (low level).

## Development
### Requirements
To carry out this practical work, a computer with GNU/Linux _kernel_ is necessary, since we will use [SystemD][sysD] to implement the management of our services.

### Development
There are N IoT sensors connected to a server written in Go which receives from them a series of parameters that are listed below. The server has a _user service_ and a _processing service_. Each service must expose a _REST API_ with _Media Type_ `application/json` [about mediatype] for all its functionalities and on the user side, only allow the operation to authenticated users.
The service must have a [nginx][ngnx] configured in front to be able to direct the _request_ to the corresponding service.

The web server must *only* respond to `dashboard.com` for the user service and `sensors.com` for the processing service. It must return _404 Not Found_ for any other non-existent _path_ with a message of choice in JSON format.
For simplicity, we will only use _HTTP_, but clarifying that this has *serious security problems*.

Next, we will detail the two services to be created and the functionalities of each one.

### User Service
This service will be responsible for creating users and listing them. These users must be able to _log in_ via _SSH_ after their creation. Only authenticated users will be able to access them. And the tasks of each endpoint using [JWT][JWT]

#### POST /api/users
Endpoints for creating users in the operating system:

```C
    GET http://{{server}}/api/users/login
```

```C
    POST http://{{server}}/api/users/createuser
```
Request
```C
        curl --request POST \
            --url http:// {server}}/api/users \
            -u USER:SECRET \
            --header 'accept: application/json' \
            --header 'content-type: application/json' \
            --data '{"username": "myuser", "password": "mypassword"}' \
            --authentification jwt
```
Respuesta
```C

        {
            "id": 142,
            "username": "myuser",
            "created_at": "2019-06-22 02:19:59"
        }

```


#### GET /api/users
Endpoint to get all operating system users and their IDs.
```C
    GET http://{{server}}/api/users/listall
```
Request
```C
    curl --request GET \
        --url http://{{server}}/api/users \
        -u USER:SECRET \
        --header 'accept: application/json' \
        --header 'content-type: application/json' \
        --authentification jwt
```
Respuesta
```C
    {
      "data": [
          {
              "user_id": 2,
              "username": "user1",
          },
          {
              "user_id": 1,
              "username": "user2"
          },
          ...
      ]
    }
```

### Processing Service
Must list a _Media Type_ , `application/json`. With information about: processing, free memory, swap, etc.

#### POST /processing/submit
```C
    POST http://{{server}}/processing/submit
```
Request

```C
    curl --request POST \
        --url http://{{server}}/contador/increment \
        -u USER:SECRET \
        --header 'accept: application/json' \
        --header 'content-type: application/json'
```


#### GET /processing/summary
This endpoint allows you to know the value of all sensors on average.
```C
    GET http://{{server}}/processing/summary
```
Request
```C
    curl --request GET \
        --url http://{{server}}/contador/value \
        -u USER:SECRET \
        --header 'accept: application/json' \
        --header 'content-type: application/json'
```

This endpoint has no logging requirements.

## Delivery
The source files must be provided, as well as any other files associated with the compilation, "Makefile" project files and the correctly documented code, all in the repository, where the Student must demonstrate progress week by week through _commits_.

A report must also be delivered, a _How to_ type guide, explaining step by step what was done (it can be a _Markdown_). The report must also contain the design of the solution with a detailed explanation of it. It must be assumed that the compilation tests will be carried out on a computer that has the typical console tools for program development (Example: gcc, make), and there are NO "GUI" tools for compiling them (Example: eclipse).

The makefile install must copy the systemd configuration files so that they can then be enabled and executed by command line.
The script should copy the necessary files for the Nginx systemd service so that they can then be enabled and executed from the command line.
The services must pass a battery of tests written in _postman_ provided. TBD.

### Correctness Criteria
- Split the code into modules judiciously.
- Code style.
- Error handling
- The code must not contain staticcheck errors.


## Evaluation
This practical work is individual and must be submitted before Friday, June 2, 2023 at 11:55 p.m. through the LEV. It will be corrected and then a date must be coordinated for the oral defense of the same.

## References and help
- [Systrem D ](https://systemd.io/)
- [System D en Freedesktop](https://www.freedesktop.org/wiki/Software/systemd/)
- [nginx](https://docs.nginx.com/)
- [Kore Web PLataform](https://kore.io/)

[sysD]: https://www.freedesktop.org/wiki/Software/systemd/
[ngnx]: https://docs.nginx.com/
[ulfi]: https://github.com/babelouest/ulfius
[logrotate]: https://en.wikipedia.org/wiki/Log_rotation
[mediatype]: https://en.wikipedia.org/wiki/Media_type
[JWT]: https://jwt.io
