create table if not exists Drink (
  id integer primary key autoincrement,
  `name` varchar not null,
  price smallint unsigned not null,
  stock smallint unsigned not null,
  created_at datetime not null,
  updated_at datetime not null
);

insert into Drink (`name`, price, stock, created_at, updated_at) values ('黒くてしゅわしゅわしたやつ', 120, 30, datetime('now'), datetime('now'));
insert into Drink (`name`, price, stock, created_at, updated_at) values ('おいしいお茶', 100, 20, datetime('now'), datetime('now'));
insert into Drink (`name`, price, stock, created_at, updated_at) values ('運動後に飲むとうまいお茶', 130, 10, datetime('now'), datetime('now'));
insert into Drink (`name`, price, stock, created_at, updated_at) values ('黄色いシュワッとしたやつ', 100, 10, datetime('now'), datetime('now'));
