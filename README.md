# **A2SV Portal Project**  

The **A2SV Portal** is a web application designed to track students' progress and manage coding contests. It provides tools for monitoring performance, organizing contests, and fostering collaboration within the **A2SV community**.  

## **Setup Instructions**  

### **1. Clone the Repository**  
```sh
git clone https://github.com/Elizabethyonas/A2SV-Portal-Project.git
cd A2SV-Portal-Project
```

### **2. Install Dependencies**  
```sh
go mod tidy
```

### **3. Configure Environment Variables**  
Create a `.env` file inside the `backend` directory and add the required variables:  

```ini
DB_HOST=your_database_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
JWT_SECRET=your_jwt_secret
```

### **4. Run the Application**  
```sh
cd backend
go run main.go
