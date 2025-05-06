# Go TIcketing App
This is a RESTful API built using Golang with Gin frameworks, GORM as ORM for MySQL. Features check event, and bookinh ticket for events and concerts. 

### Apps Business Process 
* User registration, login, and update with email and password authorization
* See available events based on time constraint, artist, etc.
* Perform booking, finish booking, and cancel booking
* Check ticket booking history

### Technical Features
* User registration with SHA-256 password hashing
* User login with JWT-based authentication
* Middleware for protected routes
* User role authorization using Middleware
* MySQL database connection using .env
* Goroutine, Sync, and Mutex for booking transactions

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)   ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
