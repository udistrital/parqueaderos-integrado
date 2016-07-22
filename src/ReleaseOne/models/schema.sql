CREATE TABLE `propietario` (
	    `uid` INT(10) NOT NULL AUTO_INCREMENT,
	    `documento` VARCHAR(64) NULL DEFAULT NULL,
	    `primer_nombre` VARCHAR(64) NULL DEFAULT NULL,
	    `segundo_nombre` VARCHAR(64) NULL DEFAULT NULL,
	    `primer_apellido` VARCHAR(64) NULL DEFAULT NULL,
	    `segundo_apellido` VARCHAR(64) NULL DEFAULT NULL,
	    PRIMARY KEY (`uid`)
);
CREATE TABLE `vehiculo` (
	      `uid` INT(10) NOT NULL AUTO_INCREMENT,
	      `id_nfc` INT(4) NOT NULL,
	      `placa` VARCHAR(10) NULL DEFAULT NULL,
       	      `id_propietario` INT(10) NOT NULL,
	      PRIMARY KEY (`uid`)
);
