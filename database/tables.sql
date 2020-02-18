CREATE TABLE daylevels(
 id serial PRIMARY KEY,
 focus integer NOT NULL,
 fischio_orecchie integer NOT NULL,
 power_energy integer NOT NULL,
 dormito integer NOT NULL,
 PR  integer NOT NULL,
 ansia  integer NOT NULL,
 arrabiato integer NOT NULL,
 irritato integer NOT NULL,
 depresso  integer NOT NULL,
 cinque_tibetani BOOLEAN NOT NULL,
 meditazione BOOLEAN NOT NULL,
 createdOn TIMESTAMP default current_timestamp
);

CREATE TABLE meditations(
 id serial PRIMARY KEY,
 meditation text NOT NULL,
 timesused integer NOT NULL,
 createdOn TIMESTAMP default current_timestamp
);

CREATE TABLE actionsmed(
 id serial PRIMARY KEY,
 action text NOT NULL,
 idmed integer NOT NULL,
 idusr integer NOT NULL,
 createdOn TIMESTAMP default current_timestamp
);

CREATE TABLE users (
id serial PRIMARY KEY,
first_name text NOT NULL,
last_name text NOT NULL,
user_name VARCHAR(50) UNIQUE,
user_pwd text NOT NULL,
idrole integer NOT NULL
);

CREATE TABLE roles (
id serial PRIMARY KEY,
role text NOT NULL
);

ALTER TABLE actionsmed
    ADD CONSTRAINT fk_actions_meditation FOREIGN KEY (idmed) REFERENCES meditations (id),
    ADD CONSTRAINT fk_actions_users FOREIGN KEY (idusr) REFERENCES users (id);


ALTER TABLE users
    ADD CONSTRAINT fk_users_roles FOREIGN KEY (idrole) REFERENCES roles (id);

ALTER TABLE meditations
  ADD COLUMN pref_month VARCHAR(15),
  ADD COLUMN pref_day integer;

ALTER TABLE users
  ADD COLUMN user_pwd text;

ALTER TABLE users
    ALTER COLUMN user_name TYPE VARCHAR(50),
    ADD UNIQUE (user_name);


INSERT INTO roles (role)
    VALUES ('user');


##UPDATE roles SET role = 'user';


GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO ml;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO ml;


ALTER TABLE daylevels
    ADD COLUMN uid integer,
    ADD CONSTRAINT fk_daylevels_users FOREIGN KEY (uid) REFERENCES users (id);


UPDATE daylevels SET uid = 1;


ALTER TABLE daylevels
    ALTER COLUMN uid TYPE integer;
    
ALTER TABLE daylevels
    ALTER COLUMN uid SET NOT NULL;


INSERT INTO daylevels (
	focus, 
	fischio_orecchie,
	power_energy,
	dormito,
	pr,
	ansia,
	arrabiato,
	irritato,
	depresso,
	cinque_tibetani,
	meditazione,
	uid) 
	VALUES (1,2,3,4,5,6,7,8,9,false,false,2);

CREATE TABLE attivitafisica (
    id serial PRIMARY KEY,
    passi integer NOT NULL,
    corsa integer NOT NULL,
    bici integer NOT NULL,
    sci integer NOT NULL,
    uid integer NOT NULL,
    createdOn TIMESTAMP default current_timestamp
);

ALTER TABLE  attivitafisica
ADD CONSTRAINT fk_attfisica_users FOREIGN KEY (uid) REFERENCES users (id);

CREATE TABLE misurevitali (
    id serial PRIMARY KEY,
    peso integer NOT NULL,
    girovita integer NOT NULL,
    uid integer NOT NULL,
    createdOn TIMESTAMP default current_timestamp
);

ALTER TABLE misurevitali
   ADD CONSTRAINT fk_daylevels_users FOREIGN KEY (uid) REFERENCES users (id);


ALTER TABLE misurevitali
    ADD COLUMN grassocorpo decimal,
    ADD COLUMN IMC decimal,
    ADD COLUMN pulse decimal;

ALTER TABLE misurevitali
    ALTER COLUMN peso TYPE decimal,
    ALTER COLUMN girovita TYPE decimal;

 
// run every time you create a table
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO ml;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO ml;


// run every time you create a table
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mlpsuser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO mlpsuser;

SELECT * FROM users;

// read the id (users) and replace the uid below

INSERT INTO misurevitali (
	peso, 
	girovita,
    grassocorpo,
    imc,
    pulse,
	uid) 
	VALUES (75,95,25,25,65,2);

CREATE TABLE sleepsmart (
    id SERIAL PRIMARY KEY,
    hsleep TIME NOT NULL,
    hwake TIME NOT NULL,
    hasleep DECIMAL,
    nrCycles DECIMAL,
    cyclesInterrupted BOOLEAN,
    alcool BOOLEAN NOT NULL,
    coffee INTEGER NOT NULL,
    lastCoffe TIME NOT NULL,
    diner INTEGER NOT NULL,
    prepSleep INTEGER NOT NULL,
    blueLight INTEGER NOT NULL,
    melatoninLevel INTEGER NOT NULL,
    uid INTEGER NOT NULL,
    createdOn TIMESTAMP default current_timestamp
);

ALTER TABLE sleepsmart
   ADD CONSTRAINT fk_daylevels_users FOREIGN KEY (uid) REFERENCES users (id);

INSERT INTO sleepsmart (
	hsleep,
    hwake,
    hasleep,
    nrCycles,
    cyclesInterrupted,
    alcool,
    coffee,
    lastCoffe,
    diner,
    prepSleep,
    blueLight,
    melatoninLevel,
	uid) 
	VALUES ('22:00:00','05:00:00',7,4,true,true,2,'16:00:00',1,1,1,1,6);

ALTER TABLE sleepsmart
ADD COLUMN dataDormire date;

ALTER TABLE sleepsmart
ALTER COLUMN dataDormire SET NOT NULL;

ALTER TABLE sleepsmart
ADD COLUMN datasveglia date;

ALTER TABLE sleepsmart
ALTER COLUMN datasveglia SET NOT NULL;


INSERT INTO sleepsmart (
	hsleep,
    hwake,
    hasleep,
    nrCycles,
    cyclesInterrupted,
    alcool,
    coffee,
    lastCoffe,
    diner,
    prepSleep,
    blueLight,
    melatoninLevel,
	uid) 
	VALUES ('22:00:00','05:00:00',7,4,true,true,2,'16:00:00',1,1,1,1,6);


    // on ubuntu 130 (PROD) userid=2
    INSERT INTO sleepsmart (datadormire, hsleep, datasveglia, hwake, hasleep, nrCycles,cyclesInterrupted,
	alcool,coffee,lastCoffe,diner,prepSleep,blueLight,melatoninLevel,uid) 
	VALUES
    ('2019-07-19','22:00:00','2019-07-20','05:00:00',7,4,true,true,2,'12:00:00',1,1,1,1,2);


    // run every time you create a table

    // esempio su Ubuntu (user diverso :: ml)
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO ml;
    GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO ml;