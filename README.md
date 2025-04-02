# A2SV Portal Project

The A2SV Portal is an Application designed to track the progress of students and manage coding contests. It provides tools for monitoring performance, organizing contests, and fostering collaboration within the A2SV community.

## Setup

1. Clone the repository:
        git clone https://github.com/Elizabethyonas/A2SV-Portal-Project.git
    cd A2SV-Portal-Project
    

2. Install dependencies:
        go mod tidy
    

3. Set up environment variables:
    Create a .env file in the backend directory and configure the required variables:
        DB_HOST=your_database_host
    DB_USER=your_database_user
    DB_PASSWORD=your_database_password
    JWT_SECRET=your_jwt_secret
    

4. Run the application:
        cd backend
    go run main.go