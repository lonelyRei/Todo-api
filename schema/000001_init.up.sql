create table users
(
    id serial not null primary key,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

create table todo_lists
(
    id serial not null primary key,
    title varchar(255) not null,
    description varchar(255)
);

create table users_lists
(
    id serial not null primary key,
    user_id int not null,
    list_id int not null,
    foreign KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

create table todo_items
(
    id serial not null primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

create table lists_items
(
    id serial not null primary key,
    item_id int not null,
    list_id int not null,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);