# Go Ticketing App
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
* Docker for containerization deployment

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)   ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)

# Models

## 🧑‍💼 User Model

Represents a user who can book tickets for events.

- `Name` - Full name of the user *(required)*
- `Email` - Unique email address *(required)*
- `Password` - Hashed password
- `Role` - Either `admin` or `user` *(default: user)*
- `Tickets` - List of tickets the user has booked (relation)

## 🎫 Ticket Model

Represents a ticket booked by a user for a specific event.

- `UserID` - Foreign key to the `User`
- `EventID` - Foreign key to the `Event`
- `Type` - Ticket type *(e.g., VIP, Regular)* *(required)*
- `Status` - Current status of the ticket:  
  - `tersedia` *(default)*  
  - `habis`  
  - `dibatalkan`  
- `User` - Belongs to a user
- `Event` - Belongs to an event

## 📅 Event Model

Represents an event that users can book tickets for.

- `Name` - Unique name of the event *(required)*
- `Description` - Details about the event
- `Location` - Venue/location of the event
- `StartTime` - Start date and time *(required)*
- `EndTime` - End date and time *(required)*
- `Capacity` - Maximum number of tickets available *(required)*
- `Price` - Cost of a single ticket *(required)*
- `Status` - Event status:
  - `aktif` *(default)*  
  - `berlangsung`  
  - `selesai`  
- `Tickets` - List of tickets associated with the event (relation)

---
