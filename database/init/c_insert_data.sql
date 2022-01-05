use main_database;

-- space1
insert into space_owners (name) value ('珈琲物語');

insert into spaces (
    owner_id,
    headline,
    -- access,
    -- weekly_visitors,
    -- main_customers_sex,
    -- min_main_customers_age,
    -- max_main_customers_age,
    daily_price,
    -- website_url,
    coordinate
  ) value (
    1,
    '〜自家製焙煎珈琲〜珈琲物語★テレビで紹介されました★若い女性に大人気',
    -- '北千住駅徒歩10分',
    -- 500,
    -- '2',
    -- 20,
    -- 30,
    3000,
    -- 'https://coffee-story.co.jp/',
    ST_PointFromText('Point(35.7516667 139.7307955)', 4326)
  );

insert into space_images (space_id, image_url) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/cooking.png'
  );

insert into space_images (space_id, image_url) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/sunshine.jpg'
  );

insert into space_images (space_id, image_url) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/house.jpg'
  );

insert into space_displays (space_id, image_url, description) value (
    1,
    'https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/road.jpg',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/plane.jpg',
    '棚 170 * 180'
  );

-- space2
insert into space_owners (name) value ('LOTUS');

insert into spaces (
    owner_id,
    headline,
    -- access,
    -- weekly_visitors,
    main_customers_sex,
    min_main_customers_age,
    max_main_customers_age,
    daily_price,
    website_url,
    coordinate
  ) value (
    2,
    '~まるごとピーチ~LOTUS★テレビで紹介されました★若い女性に大人気',
    -- '表参道駅徒歩10分',
    -- 1000,
    '2',
    20,
    30,
    4000,
    'https://www.instagram.com/lotus_omotesando/?hl=ja',
    ST_PointFromText('Point(35.6684671 139.7100234)', 4326)
  );

insert into space_images (space_id, image_url) value (
    2,
    'https://images.unsplash.com/photo-1551963831-b3b1ca40c98e'
  );

insert into space_images (space_id, image_url) value (
    2,
    'https://images.unsplash.com/photo-1551782450-a2132b4ba21d'
  );

insert into space_images (space_id, image_url) value (
    2,
    'https://images.unsplash.com/photo-1522770179533-24471fcdba45'
  );

insert into space_displays (space_id, image_url, description) value (
    2,
    'https://images.unsplash.com/photo-1444418776041-9c7e33cc5a9c',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    2,
    'https://images.unsplash.com/photo-1533827432537-70133748f5c8',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    2,
    'https://images.unsplash.com/photo-1558642452-9d2a7deb7f62',
    '棚 170 * 180'
  );

-- space3
insert into space_owners (name) value ('星野リゾート');

insert into spaces (
    owner_id,
    headline,
    access,
    weekly_visitors,
    -- main_customers_sex,
    -- min_main_customers_age,
    -- max_main_customers_age,
    daily_price,
    website_url,
    coordinate
  ) value (
    3,
    '東京の中心で現代に蘇る日本旅館',
    '東京駅徒歩10分',
    400,
    -- '1',
    -- 30,
    -- 40,
    30000,
    'https://hoshinoya.com/tokyo/',
    ST_PointFromText('Point(35.6877316 139.7628476)', 4326)
  );

insert into space_images (space_id, image_url) value (
    3,
    'https://images.unsplash.com/photo-1516802273409-68526ee1bdd6'
  );

insert into space_images (space_id, image_url) value (
    3,
    'https://images.unsplash.com/photo-1518756131217-31eb79b20e8f'
  );

insert into space_images (space_id, image_url) value (
    3,
    'https://images.unsplash.com/photo-1597645587822-e99fa5d45d25'
  );

insert into space_displays (space_id, image_url, description) value (
    3,
    'https://images.unsplash.com/photo-1567306301408-9b74779a11af',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    3,
    'https://images.unsplash.com/photo-1471357674240-e1a485acb3e1',
    '棚 170 * 180'
  );

insert into space_displays (space_id, image_url, description) value (
    3,
    'https://images.unsplash.com/photo-1589118949245-7d38baf380d6',
    '棚 170 * 180'
  );