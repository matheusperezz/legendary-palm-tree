create table cars
(
    id                  serial primary key,
    marca               varchar,
    modelo              varchar,
    motor               varchar,
    ano                 varchar,
    potencia_em_cavalos varchar
);

INSERT INTO cars(marca, modelo, motor, ano, potencia_em_cavalos) VALUES
                                               ('Porsche', '911 Turbo S', 'Flat 6, Traseiro, Turbo, 3.7L', '2022', '650'),
                                               ('Ferrari', '488 Pista', 'V8, Central, Turbo, 3.9L', '2020', '720'),
                                               ('Lamborghini', 'Aventador SVJ', 'V12, Central, Aspirado, 6.5L', '2021', '770'),
                                               ('McLaren', 'Artura', 'V6, Central, Turbo, 3.0L', '2023', '585');