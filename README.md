# Pokedex Backend

Welcome to the repository for the Pokedex backend. This backend is written in GO as a Tech Demo for CS732. 

This is a simple backend that allows an admin (Professor Oak) to add new Pokemons and remove old Pokemon.

Below is how to get started!

# Getting Started

Ensure you have GO installed. Installation can be found [here](https://go.dev/dl/)
Ensure you have Git installed. Installation can be found [here](https://git-scm.com/downloads)

Next, clone the repository.
```
git clone https://github.com/VoxelFS/pokedex-backend/
```

Navigate to the directory.
```
cd pokedex-backend
```

Create a `.env` file at the root directory with the following attributes"
```
BINARY=
MONGO_DB=
MONGO_DB_USERNAME=
MONGO_DB_PASSWORD=
```

Next, install the dependencies.
```
go mod tidy
```

Finally, run the code using the Makefile.
```
make restart
```

# Project Structure

Below is an overview of the project structure:

```
.
├── cmd
│   └── api
│       └── main.go # The main entry point of the backend.
├── internal 
│   ├── db # Handle connections to the database.
│   ├── handlers # This modules contains HTTP handlers, which define how the application responds to requests.
│   ├── middleware # Handles authentication and authorisation.
│   └── services # Handles business logic, like processing data or interacting with the DB to insert, delete or get data.
└── pkg 
    ├── utils # Utility functions that are commonly used throughout the application.
    └── write_response # Used for helping format and send HTTP responses across the application
```
