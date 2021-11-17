use mr_president;

insert into space_owners (name) value ('珈琲物語');

insert into spaces (
    owner_id,
    headline,
    access,
    number_of_visitors,
    main_customers_sex,
    min_main_customers_age,
    max_main_customers_age,
    price,
    website_url,
    coordinate
  ) value (
    1,
    '〜自家製焙煎珈琲〜珈琲物語★テレビで紹介されました★若い女性に大人気',
    '北千住駅徒歩10分',
    500,
    '2',
    20,
    30,
    3000,
    'https://coffee-story.co.jp/',
    Point(35.7516667, 139.7307955)
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

insert into space_displayers (space_id, image_url, description) value (
    1,
    'https://cdn.pixabay.com/photo/2020/07/12/07/47/bee-5396362_1280.jpg',
    '棚 170 * 180'
  );

insert into space_displayers (space_id, image_url, description) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/road.jpg',
    '棚 170 * 180'
  );

insert into space_displayers (space_id, image_url, description) value (
    1,
    'https://cdn.vuetifyjs.com/images/cards/plane.jpg',
    '棚 170 * 180'
  );