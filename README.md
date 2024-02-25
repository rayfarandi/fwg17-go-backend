# Coffee Shop Backend


Welcome to the Coffee Shop Backend Web Project! This repository contains the back-end source code for the Online Coffee Shop web application. With Golang and Gin framework structure.

Built using

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Google Chrome](https://img.shields.io/badge/Google%20Chrome-4285F4?style=for-the-badge&logo=GoogleChrome&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)

## Getting Started

To run the project locally, follow these simple steps:

1. Clone this repository
```sh
git clone https://github.com/rayfarandi/fwg17-go-backend
cd fwg17-go-backend
```

2. Open in VSCode
```sh
code .
```

3. install all the dependencies
```
go mod tidy
```

4. run the project
```
go run .
```

## Technologies Used
- Gin: This project leverages the efficiency and flexibility of Gin, a fast and lightweight web framework for Golang, to ensure the development of robust and scalable server-side applications.
- Golang: This project is built on Go, harnessing its efficient concurrency model and performance characteristics to ensure the development of scalable and high-performance server-side applications.
  
## project Structure
The project structure is organized as follows:

-src/: contains the source code of the project
  - /controllers: containing functions responsible for managing data input and output
  - /lib: containing reusable functions for specific tasks
  - /middleware: containing functions executed in the order of request
  - /models: containing queries to the database or business logic
  - /router: contains endpoint paths
  - /service: containing a response struct
    
-uploads: containing uploaded files

-main/ : main file in the application

## Coffee Shop - Frontend Repository
https://github.com/rayfarandi/fwg17-beginner-frontend

## Contributing

Contributions are always welcome!

## Authors

- [@rayfarandi](https://github.com/rayfarandi)

## Feedback

If you have any feedback, please reach out to us at rayfarandi1994@gmail.com
