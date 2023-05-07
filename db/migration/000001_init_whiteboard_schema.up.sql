-- authors table
CREATE TABLE "authors" (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  created_at timestamp not null default (now()),
  updated_at timestamp not null default (now())
);

-- whiteboard table
CREATE TABLE whiteboard (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  created_by INT,
  created_at timestamp not null default (now()),
  updated_at timestamp not null default (now()),
  FOREIGN KEY (created_by) REFERENCES authors(id)
);

-- stroke table
CREATE TABLE stroke (
  id INT PRIMARY KEY,
  whiteboard_id INT,
  color VARCHAR(255),
  line_width INT,
  created_by INT,
  created_at timestamp not null default (now()),
  updated_at timestamp not null default (now()),
  FOREIGN KEY (whiteboard_id) REFERENCES whiteboard(id),
  FOREIGN KEY (created_by) REFERENCES authors(id)
);

-- point table
CREATE TABLE point (
  id INT PRIMARY KEY,
  stroke_id INT,
  x_coordinate FLOAT,
  y_coordinate FLOAT,
  created_at timestamp not null default (now()),
  FOREIGN KEY (stroke_id) REFERENCES stroke(id)
);
