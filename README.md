# eShop

# Online Electronics Shop

## Project Description  
Online Electronics Shop is an e-commerce platform aimed at providing a user-friendly way for younger audiences to purchase electronics such as TVs, phones, PCs, laptops, and more. The project incorporates CRUD functionality for managing users and ensures seamless navigation for visitors.

## Team Members  
- Ismail  
- Adilkhan  
- Ualikhan  

## Screenshot of the Main Page  
![image](https://github.com/user-attachments/assets/23c4003f-10af-407d-ba46-fd5b0b0c8f3a)


## How to Start the Project  

### Prerequisites  
1. Install [Go](https://golang.org/doc/install).  
2. Install PostgreSQL and ensure it's running on your system.  
3. Install `gorm` and `pq` libraries for Go:  
   ```bash
   go get -u gorm.io/gorm
   go get -u github.com/lib/pq
Steps
1) Clone the repository:
git clone <repository_url>
cd <repository_folder>

2) Set up the PostgreSQL database:
Create a database named advprog.
Update the dsn string in the main function to match your database credentials.
3) Run the Go server:
go run main.go
4) Open the browser and navigate to:
http://localhost:8080 for the main page.
http://localhost:8080/about for the About page.

## Tools and Resources Used
Programming Language: Go
Database: PostgreSQL
Libraries:
1)GORM for database operations
2)pq as a PostgreSQL driver
Static Files: HTML, CSS (in the static folder)
## Tools:
Intellij IDEA for development
