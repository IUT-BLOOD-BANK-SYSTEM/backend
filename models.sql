

CREATE TABLE blood_types (
    id UUID PRIMARY KEY,
    bloods_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT  
 NOW()
);
CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    date_of_birth VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    blood_type_id UUID REFERENCES blood_types("id"),
    diseases VARCHAR(100) NOT NULL,
    passport_number VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT  
 NOW()
);

CREATE TABLE regions (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT  
 NOW()
);

CREATE TABLE hospitals (
    id UUID PRIMARY KEY,                    
    name VARCHAR(100) NOT NULL,             
    region_id UUID REFERENCES regions("id"), 
    address VARCHAR(100) NOT NULL,         
    phone_number VARCHAR(100) NOT NULL,    
    head_nurse_id UUID,                    
    created_at TIMESTAMP DEFAULT NOW(),     
    updated_at TIMESTAMP DEFAULT NOW(),     
);
CREATE TABLE head_nurses (
    id UUID PRIMARY KEY,                   
    first_name VARCHAR(100) NOT NULL,      
    second_name VARCHAR(100) NOT NULL,     
    date_of_birth VARCHAR(100) NOT NULL,   
    phone_number VARCHAR(100) NOT NULL,    
    address VARCHAR(100) NOT NULL,         
    email VARCHAR(100) UNIQUE NOT NULL,   
    password VARCHAR(255) NOT NULL,        
    hospital_id UUID,                     
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW(),    
    FOREIGN KEY (hospital_id) REFERENCES hospitals(id)  
);




CREATE TABLE doctors (
    id UUID PRIMARY KEY,                   
    first_name VARCHAR(100) NOT NULL,      
    second_name VARCHAR(100) NOT NULL,     
    hospital_id UUID,                     
    specialization VARCHAR(100) NOT NULL,   
    phone_number VARCHAR(100) NOT NULL,    
    address VARCHAR(100) NOT NULL,         
    email VARCHAR(100) UNIQUE NOT NULL,   
    password VARCHAR(255) NOT NULL,        
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW(),    
    FOREIGN KEY (hospital_id) REFERENCES hospitals(id)  
);


CREATE TABLE blood_donations (
    id UUID PRIMARY KEY,                   
    donor_id UUID REFERENCES users("id"),       
    hospital_id UUID REFERENCES hospitals("id"),      
    blood_type_id UUID REFERENCES blood_types("id"),      
    quantity NUMERIC NOT NULL,   
    donation_date DATE NOT NULL,    
    status VARCHAR(55) DEFAULT 'in_proccess',   
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE blood_requests (
    id UUID PRIMARY KEY,                   
    doctor_id UUID REFERENCES doctors("id"),       
    hospital_id UUID REFERENCES hospitals("id"),      
    blood_type_id UUID REFERENCES blood_types("id"),      
    quantity NUMERIC NOT NULL,   
    status VARCHAR(55) DEFAULT 'in_proccess',   
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE blood_inventories (
    id UUID PRIMARY KEY,                   
    hospital_id UUID REFERENCES hospitals("id"),      
    blood_type_id UUID REFERENCES blood_types("id"),      
    quantity NUMERIC NOT NULL,   
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW()
);




CREATE TABLE appointments (
    id UUID PRIMARY KEY,                   
    donor_id UUID REFERENCES users("id"),       
    hospital_id UUID REFERENCES hospitals("id"),      
    appointment_date DATE NOT NULL,    
    status VARCHAR(55) DEFAULT 'in_proccess',   
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW()
);




CREATE TABLE blood_acceptances (
    id UUID PRIMARY KEY,                   
    doctor_id UUID REFERENCES doctors("id"),       
    patient_id UUID REFERENCES users("id"),      
    hospital_id UUID REFERENCES hospitals("id"),     
    acceptance_date  DATE NOT NULL,  
    status VARCHAR(55) DEFAULT 'checking',   
    created_at TIMESTAMP DEFAULT NOW(),    
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE sis_admins (
    id UUID PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT  
 NOW()
);