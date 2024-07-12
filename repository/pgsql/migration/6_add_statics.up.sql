-- 6_add_statics.up.sql

-- Добавление нового поля subject в таблицу ad_campaigns
ALTER TABLE "cpms"
    ADD COLUMN "old_position" INTEGER;

-- Создание таблицы

-- Создание индексов
