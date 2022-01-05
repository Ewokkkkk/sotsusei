CREATE TABLE food_data (
	food_id SERIAL PRIMARY KEY,
	food_name varchar(30),
	food_name_ruby varchar(30)
);

CREATE TABLE permission_data(
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
	ciry_address varchar(50),
	street_address varchar(80),
	phone_number varchar(32),
	created_date timestamp,
	changed_date timestamp,
	foreign key (permission_id) references permission_data on update no action
);

CREATE TABLE shipping_address (
	shipping_id SERIAL PRIMARY KEY,
	user_id int,
	shipping_prefecture varchar(50),
	shipping_street varchar(80),
	created_date timestamp,
	changed_date timestamp,
	foreign key (user_id) references users on update no action
);

CREATE TABLE seller_data (
	seller_id SERIAL PRIMARY KEY,
	permission_id varchar(1),
	seller_name varchar(30),
	seller_name_ruby varchar(30),
	seller_name varchar(30),
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
	foreign key (permission_id) references permission_data on update no action
);

CREATE TABLE product_data(
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
	foreign key (food_id) references food_data on update no action,
	foreign key (seller_id) references seller_data on update no action
);

CREATE TABLE purchase(
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

CREATE TABLE order_data(
	order_id SERIAL PRIMARY KEY,
	purchase_id int,
	product_id int,
	quantity int,
	subtotal int,
	foreign key (product_id) references product_data on update no action,
	foreign key (purchase_id) references purchase on update no action
);