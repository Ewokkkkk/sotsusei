CREATE TABLE foods (
    food_id SERIAL PRIMARY KEY,
    food_name varchar(30),
    food_name_ruby varchar(30)
);

CREATE TABLE permissions (
	permission_id varchar(1) PRIMARY KEY,
	permission_name varchar(10)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    permission_id varchar(1),
    user_name varchar(30),
    user_name_ruby varchar(30),
    password varchar(128),
    birth_year int,
    birth_month int,
    birth_day int,
    email varchar(100) UNIQUE,
    zipcode varchar(7),
    prefecture_address varchar(50),
    city_address varchar(50),
    street_address varchar(80),
    phone_number varchar(32),
    created_date timestamp,
    changed_date timestamp,
    foreign key (permission_id) references permissions on update no action
);

CREATE TABLE shipping_addresses (
	shipping_id SERIAL PRIMARY KEY,
	user_id int,
	shipping_prefecture varchar(50),
	shipping_street varchar(80),
	created_date timestamp,
	changed_date timestamp,
	foreign key (user_id) references users on update no action
);

CREATE TABLE sellers (
    seller_id SERIAL PRIMARY KEY,
    permission_id varchar(1),
    seller_name varchar(30),
    seller_name_ruby varchar(30),
    password varchar(32),
    birth_year int,
    birth_month int,
    birth_day int,
    email varchar(100),
    zipcode varchar(7),
    prefecture_address varchar(50),
    street_address varchar(80),
    phone_number varchar(32),
    created_date timestamp,
    changed_date timestamp,
    foreign key (permission_id) references permissions on update no action
);

CREATE TABLE products(
	product_id SERIAL PRIMARY KEY,
	food_id int,
	seller_id int,
	product_name varchar(30),
	product_locality varchar(30),
	product_img1 varchar(100),
	product_img2 varchar(100),
	product_img3 varchar(100),
	product_price int,
	product_stock int,
	product_description varchar(500),
	deadline timestamp,
	created_date timestamp,
	updated_date timestamp,
	foreign key (food_id) references foods on update no action,
	foreign key (seller_id) references sellers on update no action
);

CREATE TABLE purchases(
	purchase_id SERIAL PRIMARY KEY,
	user_id int,
	purchase_zipcode varchar(7),
	purchase_prefecture varchar(50),
	purchase_street varchar(80),
	total_price int,
	postage int,
	order_date timestamp,
	shipping_data timestamp,
	shipping_status varchar(15),
	foreign key (user_id) references users on update no action
);

CREATE TABLE orders(
	order_id SERIAL PRIMARY KEY,
	purchase_id int,
	product_id int,
	quantity int,
	subtotal int,
	foreign key (product_id) references products on update no action,
	foreign key (purchase_id) references purchases on update no action
);

INSERT INTO public.sellers VALUES (1, '1', 'test', 'test', 'tes', 1111, 11, 11, 'TEST', '1111111', '1TEST', 'TEST', '11111111111', NULL, NULL);

INSERT INTO foods
  (food_name, food_name_ruby)
VALUES
  ('?????????', '????????????'),
  ('????????????', '????????????'),
  ('??????????????????', '??????????????????'),
  ('????????????', '????????????'),
  ('?????????', '?????????'),
  ('????????????????????????', '????????????????????????'),
  ('?????????', '?????????'),
  ('???', '??????'),
  ('??????', '??????'),
  ('??????', '??????'),
  ('??????', '????????????'),
  ('???????????????', '???????????????'),
  ('????????????', '????????????'),
  ('?????????', '?????????'),

  ('???', '??????'),
  ('??????', '??????'),
  ('?????????', '?????????'),
  ('????????????', '????????????'),
  ('??????????????????', '??????????????????'),
  ('?????????', '?????????'),
  ('?????????', '?????????'),
  ('????????????', '????????????'),
  ('????????????', '????????????'),
  ('????????????', '????????????'),
  ('?????????', '?????????'),
  ('?????????', '????????????'),	
  ('?????????', '?????????'),

  ('???????????????', '???????????????'),
  ('??????', '????????????'),
  ('??????????????????', '??????????????????'),
  ('??????', '???????????????'),
  ('??????', '???????????????'),
  ('??????', '????????????'),
  ('???????????????', '???????????????'),
  ('?????????', '?????????'),
  ('??????', '??????'),
  ('?????????', '?????????'),
  ('?????????', '????????????'),

  ('????????????', '????????????'),
  ('?????????', '????????????'),
  ('??????', '????????????'),
  ('?????????', '?????????'),
  ('??????????????????', '??????????????????'),
  ('??????????????????', '??????????????????'),
  ('?????????', '???????????????'),
  ('??????????????????', '??????????????????'),
  ('?????????', '?????????'),

  ('???', '??????'),
  ('??????', '??????'),
  ('?????????', '????????????'),
  ('??????', '??????'),
  ('????????????', '????????????'),
  ('????????????', '????????????'),
  ('??????', '??????'),

  ('??????', '????????????'),
  ('?????????', '?????????'),
  ('??????????????????', '??????????????????'),
  ('?????????', '?????????'),
  ('?????????', '?????????'),
  ('????????????', '????????????'),
  ('??????', '??????'),
  ('?????????', '?????????'),
  ('????????????', '????????????'),
  ('??????', '??????'),
  ('?????????', '?????????'),
  ('??????????????????', '??????????????????'),
  ('?????????', '?????????'),
  ('???????????????', '??????????????????'),

  ('??????', '????????????'),
  ('????????????', '????????????'),
  ('?????????', '?????????'),
  ('?????????', '?????????'),
  ('??????', '?????????'),
  ('????????????', '????????????'),
  ('?????????', '?????????'),
  ('???', '??????'),
  ('?????????', '?????????'),
  ('???????????????', '???????????????'),

  ('????????????', '????????????'),
  ('??????', '????????????'),
  ('??????', '??????'),
  ('?????????', '?????????'),

  ('?????????', '?????????'),
  ('???????????????', '???????????????'),
  ('?????????', '?????????'),
  ('?????????', '?????????'),
  ('?????????', '?????????'),
  ('????????????', '????????????'),

  ('?????????', '?????????'),
  ('?????????', '?????????');

INSERT INTO public.products VALUES (1, 52, 1, '????????????2????????????(1???150g)', '?????????????????????', '/assets/img/ninzin.jpg', '/assets/img/ninzin.jpg', NULL, 50, 10, '??????????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (2, 49, 1, '??????2????????????', '?????????????????????', '/assets/img/nasu.jpg', '/assets/img/nasu.jpg', NULL, 80, 9, '????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (3, 23, 1, '????????????3????????????', '?????????????????????', '/assets/img/kyu-ri.jpg', '/assets/img/kyu-ri.jpg', NULL, 70, 7, '??????????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (4, 63, 1, '????????????5????????????', '?????????????????????', '/assets/img/pi-man.jpg', '/assets/img/pi-man.jpg', NULL, 150, 5, '??????????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (5, 40, 1, '?????????2???', '?????????????????????', '/assets/img/tamanegi.jpg', '/assets/img/tamanegi.jpg', NULL, 80, 10, '??????????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (6, 47, 1, '?????????3????????????', '?????????????????????', '/assets/img/tomato.jpg', '/assets/img/tomato.jpg', NULL, 150, 7, '???????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
INSERT INTO public.products VALUES (7, 41, 1, '????????????1???250g', '?????????????????????', '/assets/img/daikon.jpg', '/assets/img/daikon.jpg', NULL, 160, 3, '????????????????????????????????????????????????????????????', '2021-12-30 00:00:00', NULL, NULL);
