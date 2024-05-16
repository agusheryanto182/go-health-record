CREATE TABLE IF NOT EXISTS records(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    identity_number BIGINT NOT NULL REFERENCES patients(identity_number),
    symptoms VARCHAR(2000) NOT NULL,
    medications VARCHAR(2000) NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
) ;

CREATE INDEX IF NOT EXISTS idx_records_identity_number ON records(identity_number) ;
CREATE INDEX IF NOT EXISTS idx_records_created_by ON records(created_by) ;
CREATE INDEX IF NOT EXISTS idx_records_created_at ON records(created_at) ;    