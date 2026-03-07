-- Habilita a extensão de UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Cria o tipo ENUM no Postgres com os valores iniciais permitidos
CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'DONE');

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    -- Substituímos o VARCHAR(50) pelo nosso novo tipo ENUM
    status task_status NOT NULL DEFAULT 'TODO',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Cria o índice no status para buscas rápidas
CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);