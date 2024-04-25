# Step 1 - Basics

## • Retrieve **all** the information contained in the `artists` table.

```
SELECT * FROM artists;
```

## • Retrieve **only** `name` and `genre` from the table `artists`.

```
SELECT name, genre FROM artists;
```

## • Retrieve the list of all `artists` of `genre` `hip-hop/rap`.

```
SELECT * FROM artists WHERE genre = 'hip-hop/rap';
```

&nbsp;

# Step 2 - Relations

## • Retrieve `name` from `artists` and `musics`.

> You must specify the name of your result column with artists_names and musics_names.

```
SELECT artists.name AS artists_names, musics.name AS musics_names
FROM participations
JOIN artists ON participations.artistid = artists.id
JOIN musics on participations.musicid = musics.id
```

## • Retrieve all `artist` who singed in the music `We Are The World`.

> Those artists must be sorted in `descending` order according to their number of fans.

```
SELECT artists.* FROM artists
JOIN participations ON artists.id = participations.artistid
JOIN musics on participations.musicid = musics.id
WHERE musics.id = 1
ORDER by artists.nbfans DESC;
```

## • Retrieve all the `musics` from `Booba`.

> They must be sorted in `alphabetical` order.

```
SELECT musics.* FROM musics
JOIN participations on musics.id = participations.musicid
JOIN artists ON participations.artistid = artists.id
WHERE artists.id = 8
ORDER BY name ASC
```

&nbsp;

# Step 3 - CRUD

## • Add a new `artist` with his `id` set to `100`.

```
INSERT INTO artists (id, name, genre, nbfans)
VALUES (100, 'Bigflo et Oli', 'rap', 2000000);
```

## • Delete all musics that have the `Gold` `certification`.

> ⚠️ artists and musics are linked using a relationship table, you will maybe need to do 2 queries to delete records.

```
DELETE FROM participations
WHERE musicid IN (SELECT id from musics
                  WHERE certification = 'Gold');
```

```
DELETE FROM musics
WHERE certification = 'Gold';
```

## • Add the music `Take What You Want` to the `artists` you previously created.

```
INSERT INTO participations (artistid, musicid)
VALUES (100, 10);
```

&nbsp;

# Step 4 - Good counts make good friends

## • Count the number of artists

```
SELECT COUNT(artists) FROM artists
```

## • Count the number of artists in each genre.

```
SELECT genre, COUNT(artists)
FROM artists
GROUP BY genre;
```

## • Count the number of musics sorted by their certification and displayed in ascending order.

```
SELECT certification, COUNT(id)
FROM musics
GROUP BY certification
ORDER BY count(id) ASC;
```

## • Count the number of musics of each artists, sorted by their certification and displayed in ascending order.

```
SELECT artists.name, musics.certification, COUNT(musics.id) FROM artists
JOIN participations ON artists.id = participations.artistid
JOIN musics ON participations.musicid = musics.id
GROUP BY artists.name, musics.certification
ORDER BY COUNT(musics.id) ASC;
```

&nbsp;

# Step 5 - Rap Game

## • Write a query that retrieve all the rappers in the database, sorted in descending order by their fans' number.

```
SELECT * FROM artists
WHERE LOWER(genre::text) LIKE '%' || LOWER('rap') || '%'
ORDER BY nbfans DESC;
```
