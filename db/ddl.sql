create table if not exists drinks (
  id integer primary key autoincrement,
  `name` varchar not null comment '名前',
  price smallint unsigned not null comment '価格',
  stock smallint unsigned not null comment '在庫',
  created_at datetime not null,
  updated_at datetime not null
);

insert into drinks (`name`, price, stock, created_at, updated_at) values ('黒くてしゅわしゅわしたやつ', 120, 30, datetime('now'), datetime('now'));
insert into drinks (`name`, price, stock, created_at, updated_at) values ('おいしいお茶', 100, 20, datetime('now'), datetime('now'));
insert into drinks (`name`, price, stock, created_at, updated_at) values ('運動後に飲むとうまいお茶', 130, 10, datetime('now'), datetime('now'));
insert into drinks (`name`, price, stock, created_at, updated_at) values ('黄色いシュワッとしたやつ', 100, 10, datetime('now'), datetime('now'));
