CloudSpend/
├── cmd/                 
│   ├── main.go                # Entry point, initializes the app
├── docker-compose.yml         # Docker services for PostgreSQL & pgAdmin
└── internal/                  
    ├── config/                
    │   ├── config.go          # Loads environment variables & configurations
    ├── database/              
    │   ├── db.go              # Database connection using GORM
    │   ├── migrate.go         # Database migration logic
    ├── handlers/              
    │   ├── transaction_handler.go  # Handles transaction-related API requests
    │   ├── user_handler.go         # Handles user-related API requests
    ├── models/                
    │   ├── transaction.go     # Defines the Transaction struct
    │   ├── user.go            # Defines the User struct
    ├── repository/            
    │   ├── transaction_repo.go  # Implements DB logic for transactions
    │   ├── user_repo.go         # Implements DB logic for users
    ├── server/                
    │   ├── server.go          # Initializes HTTP server with mux router
