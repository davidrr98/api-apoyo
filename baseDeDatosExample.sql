CREATE TABLE periodo (
   id_periodo serial primary key not null,
   nombre varchar(6) not null,
   estado varchar(14) not null,
   fecha_inicio date not null,
   fecha_fin date   
);
