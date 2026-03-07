DROP INDEX IF EXISTS idx_tasks_status;
DROP TABLE IF EXISTS tasks;

-- Apaga o tipo Enum que criamos
DROP TYPE IF EXISTS task_status;