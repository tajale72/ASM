
CREATE TABLE finance (
	user_id serial PRIMARY KEY,
	discovercardbalance  VARCHAR(50),
    debitCardbalance          VARCHAR(50),
    stockmarket_buyingpower   VARCHAR(50),
	rent                      VARCHAR(50),
	utilities                 VARCHAR(50),
	phonebill                 VARCHAR(50),
	carinsurance              VARCHAR(50),
	walmartGrocery            VARCHAR(50),
  total_net_worth            VARCHAR(50),
    date VARCHAR(50)
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);



//Sample INSERT query in accunts table 
INSERT INTO accounts (user_id ,discovercardbalance,debitcardbalance,stockmarket_buyingpower,rent,utilities,phonebill,carinsurance,walmartgrocery , date)VALUES ('1', '2', '3', '4' , '5', '6', '7', '8', '9' ,'10');


SELECT * FROM accounts ORDER BY is_notice desc, thread_id desc, id;


Update and select
//Update accounts SET rent = '700' Where(accounts.user_id = select accounts.user_id from accounts INNER JOIN users a ON a.id = accounts.user_id); 

Update accounts a  SET rent = '800' FROM accounts INNER JOIN users u  ON u.id = accounts.user_id; 