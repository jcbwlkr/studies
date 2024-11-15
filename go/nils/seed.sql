DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    name VARCHAR(36) NOT NULL DEFAULT "",
    attributes JSON DEFAULT NULL
) Engine=InnoDB;

INSERT INTO `products` (`name`, `attributes`) VALUES
    ('shirt', '{"size": "large", "color": "black"}'),
    ('pants', 'null'),
    ('undefined', NULL);
