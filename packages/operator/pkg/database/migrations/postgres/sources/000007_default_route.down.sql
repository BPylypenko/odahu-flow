BEGIN;
alter table odahu_operator_route drop column is_default;
COMMIT;