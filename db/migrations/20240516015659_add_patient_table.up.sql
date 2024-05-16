CREATE TABLE patients(
    identity_number BIGINT NOT NULL PRIMARY KEY,
    phone_number VARCHAR(15) NOT NULL,
    name VARCHAR(50) NOT NULL,
    birth_date VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    identity_card_scan_img TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) ;

CREATE INDEX IF NOT EXISTS idx_patients_identity_number ON patients(identity_number) ;
CREATE INDEX IF NOT EXISTS idx_patients_phone_number ON patients(phone_number) ;
CREATE INDEX IF NOT EXISTS idx_patients_name ON patients(name) ;
CREATE INDEX IF NOT EXISTS idx_patients_created_at ON patients(created_at) ;