1. Install Go: Ensure that Go is installed on your machine. You can download and install Go from the official website: https://golang.org/dl/


2. Set up your project: Create a new directory for your project and navigate to it in your terminal.


3. Initialize Go module: Run the following command to initialize a Go module in your project directory:
    ```
    go mod init <module-name>
    ```
    Replace <module-name> with the name of your project.


4. Install dependencies: Install the required packages by running the following commands:
    ```
    go get -u github.com/gin-gonic/gin
    go get -u github.com/jackc/pgx
    go get -u github.com/jackc/pgx/pgxpool
    go get -u github.com/golang-migrate/migrate/v4
    ```


5. Set up your database: Make sure you have a PostgreSQL database set up and running. Take note of the database connection details (host, port, username, password, database name).


6. Create a .env file: In your project's root directory, create a new file named .env and open it in a text editor.

    Add database credentials to .env file: Inside the .env file, add the following lines and replace the placeholders with your actual database credentials:
    ```
    dbUrl = "postgres://login:password@host:5432/dbname"
    ```


7. Use the Migrate package to manage your database migrations. You can use instructions from official github: https://github.com/golang-migrate/migrate


8. Clone or download your Golang project from your version control system (e.g., Git).


9. Build and run the project: In the terminal, navigate to the root directory of your project. Run the following command to build and run your project:
    ```
    go run .
    ```
    This command assumes that the main Go file is in the root directory and named main.go.



10. Endpoints:

    To save CSV data in the database, make a POST request to the endpoint 
    ```
    http://localhost:<port>/promotions/csv
    ```
    To retrieve records by ID, make a GET request to the endpoint 
    ```
    http://localhost:<port>/promotions/:id
    ```
    You can use a tool like cURL or Postman to send requests

