CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES roles(id), -- Assign a single role to each user
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id),
    role_id UUID REFERENCES roles(id),
    PRIMARY KEY (user_id, role_id)
);


func (h *handler) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        h.handlerResponse(c, "Invalid input", http.StatusBadRequest, err.Error())
        return
    }

    // Assign a default role (e.g., "user") if none is provided
    if user.RoleID == "" {
        defaultRoleID := "uuid-for-default-role"
        user.RoleID = defaultRoleID
    }

    // Save the user with the assigned role
    id, err := h.strg.User().Create(c.Request.Context(), &user)
    if err != nil {
        h.handlerResponse(c, "Error creating user", http.StatusInternalServerError, err.Error())
        return
    }

    h.handlerResponse(c, "User created successfully", http.StatusCreated, gin.H{"id": id})
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole := c.GetString("role") // Assume role is extracted from JWT or session

        if userRole != requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
            c.Abort()
            return
        }

        c.Next()
    }
}
