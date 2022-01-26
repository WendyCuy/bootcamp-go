/* --SQL 1 MOVIES DB ---

--- Mostrar todos los registros de la tabla movies---*/
SELECT * FROM movies;

/* --Mostrar el nombre, apellido y rating de todos los actores--*/
SELECT first_name, last_name, rating FROM actors;

/* Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla
como el campo estén en español*/
SELECT title AS titulo FROM series AS series;

/* Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5*/
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

/* Mostrar el título de las peliculas, el rating y los premios de las peliculas con un rating mayor a 
7.5 y con más de dos premios. */
SELECT title, rating, awards FROM movies WHERE rating > 7.5 and awards > 2;

/* Mostrar el titulo de las peliculas y el rating en forma ascendente*/
SELECT title, rating FROM movies ORDER BY title ASC, rating ASC;

/*Mostrar los títulos de las primeras tres peliculas en la base de datos.*/
SELECT title FROM movies limit 3;

/*Mostrar el top 5 de las peliculas con mayor rating*/
SELECT MAX(rating) FROM movies;

/*Mostrar las top 5 a 10 de las peliculas con mayor rating*/
SELECT rating FROM movies WHERE rating between 5 and 10;

/*Listar los primeros 10 actores (sería la página 1). */
SELECT 
/*Luego usar offset para traer la página 3.*/
/*Hacer lo mismo para la página 5*/
/*Mostrar el titulo y rating de todas las peliculas cuyo titulo sea de Toy Story.*/
/*Mostrar a todos los actores cuyos nombres empiecen con Sam.*/
/*Mostrar el titulo de las peliculas que salieron entre el 2004 y 2008*/
/*Traer el titulo de las peliculas con el rating mayor a 3, con más de 1 premio y con fecha de
lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.*/
/*Traer el top 3 a partir del registro 10 de la consulta anterior*/



