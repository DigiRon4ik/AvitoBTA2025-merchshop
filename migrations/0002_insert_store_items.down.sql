-- Удаление всех записей, добавленных в таблицу store
DELETE
FROM store
WHERE slug IN ('t-shirt', 'cup', 'book', 'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');
