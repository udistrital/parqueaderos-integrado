Descripcion de Cambios.

Se modifica el cogido github.com/astaxie/beego/orm/db.go del Framework de "beego" para el caso de este proyecto en particular desactivando la caracteristica de tratamineto de datos tipo timestamp (time.Time) que gestionaba la adicion del componente de localtime.

Se modifica el codigo de github.com/beego/bee/g_addcode.go adicionando la capacidad al framework de conversion de datos tipo Geometry de la base de datos postgresa su equivalente en Go (Geometry-string). Es necesario reconstruir el binario de bee.
