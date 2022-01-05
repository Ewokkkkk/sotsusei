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
    foreign key (permission_id) references permission_data on update no action
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
	foreign key (food_id) references food_data on update no action,
	foreign key (seller_id) references seller_data on update no action
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
	foreign key (product_id) references product_data on update no action,
	foreign key (purchase_id) references purchase on update no action
);

INSERT INTO foods
  (food_name, food_name_ruby)
VALUES
  ('青じそ', 'アオジソ'),
  ('あしたば', 'アシタバ'),
  ('アスパラガス', 'アスパラガス'),
  ('アボガド', 'アボガド'),
  ('アロエ', 'アロエ'),
  ('アーティチョーク', 'アーティチョーク'),
  ('いちご', 'イチゴ'),
  ('梅', 'ウメ'),
  ('うり', 'ウリ'),
  ('うど', 'ウド'),
  ('枝豆', 'エダマメ'),
  ('エンダイブ', 'エンダイブ'),
  ('えんどう', 'エンドウ'),
  ('オクラ', 'オクラ'),

  ('柿', 'カキ'),
  ('かぶ', 'カブ'),
  ('かぼす', 'カボス'),
  ('かぼちゃ', 'カボチャ'),
  ('カリフラワー', 'カリフラワー'),
  ('キウイ', 'キウイ'),
  ('きのこ', 'キノコ'),
  ('キャベツ', 'キャベツ'),
  ('きゅうり', 'キュウリ'),
  ('きんかん', 'キンカン'),
  ('ゴーヤ', 'ゴーヤ'),
  ('小松菜', 'コマツナ'),	
  ('ごぼう', 'ゴボウ'),

  ('さつまいも', 'サツマイモ'),
  ('里芋', 'サトイモ'),
  ('さやいんげん', 'サヤインゲン'),
  ('山椒', 'サンショウ'),
  ('春菊', 'シュンギク'),
  ('生姜', 'ショウガ'),
  ('じゃがいも', 'ジャガイモ'),
  ('すだち', 'スダチ'),
  ('セリ', 'セリ'),
  ('セロリ', 'セロリ'),
  ('そら豆', 'ソラマメ'),

  ('たけのこ', 'タケノコ'),
  ('玉ねぎ', 'タマネギ'),
  ('大根', 'ダイコン'),
  ('チコリ', 'チコリ'),
  ('チンゲンサイ', 'チンゲンサイ'),
  ('つるむらさき', 'ツルムラサキ'),
  ('唐辛子', 'トウガラシ'),
  ('とうもろこし', 'トウモロコシ'),
  ('トマト', 'トマト'),

  ('梨', 'ナシ'),
  ('なす', 'ナス'),
  ('菜の花', 'ナノハナ'),
  ('ニラ', 'ニラ'),
  ('にんじん', 'ニンジン'),
  ('ニンニク', 'ニンニク'),
  ('ねぎ', 'ネギ'),

  ('白菜', 'ハクサイ'),
  ('バナナ', 'バナナ'),
  ('パイナップル', 'パイナップル'),
  ('バジル', 'バジル'),
  ('パセリ', 'パセリ'),
  ('パプリカ', 'パプリカ'),
  ('びわ', 'ビワ'),
  ('ビーツ', 'ビーツ'),
  ('ピーマン', 'ピーマン'),
  ('ふき', 'フキ'),
  ('ぶどう', 'ブドウ'),
  ('ブロッコリー', 'ブロッコリー'),
  ('ベリー', 'ベリー'),
  ('ほうれん草', 'ホウレンソウ'),

  ('松茸', 'マツタケ'),
  ('マンゴー', 'マンゴー'),
  ('三つ葉', 'ミツバ'),
  ('壬生菜', 'ミブナ'),
  ('水菜', 'ミズナ'),
  ('みょうが', 'ミョウガ'),
  ('メロン', 'メロン'),
  ('桃', 'モモ'),
  ('もやし', 'モヤシ'),
  ('モロヘイヤ', 'モロヘイヤ'),

  ('ヤーコン', 'ヤーコン'),
  ('山芋', 'ヤマイモ'),
  ('ゆず', 'ユズ'),
  ('ゆり根', 'ユリネ'),

  ('ライム', 'ライム'),
  ('らっきょう', 'ラッキョウ'),
  ('りんご', 'リンゴ'),
  ('レタス', 'レタス'),
  ('レモン', 'レモン'),
  ('れんこん', 'レンコン'),

  ('わさび', 'ワサビ'),
  ('わらび', 'ワラビ');