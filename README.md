# CRUD Music

A simple CRUD Music with **Golang** and **Gin** Framework

## Requirements

- **Go** 1.19+
- **MySQL** 5.7+

## Installation

1. Copy file `.env.example` to `.env`

```env
DB_DRIVER="mysql"
DB_USER=YOUR_DATABASE_USERNAME
DB_PASSWORD=YOUR_DATABASE_PASSWORD
DB_HOST=localhost
DB_PORT=3306
DB_NAME=YOUR_DATABASE_NAME
```

2. Install Golang Package

```bash
go get
```

3. Run this SQL to create new Music table

```sql
CREATE TABLE IF NOT EXISTS `musics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `artist` varchar(50) DEFAULT NULL,
  `is_published` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
```

4. Run App
   - with `Makefile`
     ```bash
     make run
     ```
   - or you can run this
     ```bash
     go run main.go
     ```

## Support

For support, just `star` this repository :)
