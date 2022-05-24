-- table holds all of the items possible in all of the warehouses
CREATE TABLE Inventory (
    name VARCHAR(255) NOT NULL,
    price DECIMAL NOT NULL,
    description VARCHAR(255) NOT NULL,
    PRIMARY KEY(name)
);

-- each table is a warehouse in a particular city
CREATE TABLE Houston (
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    city VARCHAR(255) DEFAULT 'Houston',
    FOREIGN KEY(name) REFERENCES Inventory(name)
);

CREATE TABLE Austin (
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    city VARCHAR(255) DEFAULT 'Austin',
    FOREIGN KEY(name) REFERENCES Inventory(name)
);

CREATE TABLE Seattle (
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    city VARCHAR(255) DEFAULT 'Seattle',
    FOREIGN KEY(name) REFERENCES Inventory(name)
);

CREATE TABLE SanFrancisco (
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    city VARCHAR(255) DEFAULT 'San Francisco',
    FOREIGN KEY(name) REFERENCES Inventory(name)
);

CREATE TABLE Chicago (
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    city VARCHAR(255) DEFAULT 'Chicago',
    FOREIGN KEY(name) REFERENCES Inventory(name)
);

-- all of the items our company sells
INSERT INTO Inventory(name, price, description) VALUES('bread', 1.00, 'fresh bread');
INSERT INTO Inventory(name, price, description) VALUES('milk', 2.00, 'fresh milk');
INSERT INTO Inventory(name, price, description) VALUES('cheese', 3.00, 'fresh cheese');
INSERT INTO Inventory(name, price, description) VALUES('eggs', 4.00, 'fresh eggs');
INSERT INTO Inventory(name, price, description) VALUES('sugar', 5.00, 'fresh sugar');
INSERT INTO Inventory(name, price, description) VALUES('butter', 6.00, 'fresh butter');
INSERT INTO Inventory(name, price, description) VALUES('flour', 7.00, 'fresh flour');
INSERT INTO Inventory(name, price, description) VALUES('salt', 8.00, 'fresh salt');
INSERT INTO Inventory(name, price, description) VALUES('pepper', 9.00, 'fresh pepper');
INSERT INTO Inventory(name, price, description) VALUES('soda', 10.00, 'fresh soda');
INSERT INTO Inventory(name, price, description) VALUES('water', 11.00, 'fresh water');
INSERT INTO Inventory(name, price, description) VALUES('chocolate', 12.00, 'fresh chocolate');
INSERT INTO Inventory(name, price, description) VALUES('candy', 13.00, 'fresh candy');
INSERT INTO Inventory(name, price, description) VALUES('juice', 14.00, 'fresh juice');
INSERT INTO Inventory(name, price, description) VALUES('sausage', 15.00, 'fresh sausage');
INSERT INTO Inventory(name, price, description) VALUES('chips', 16.00, 'fresh chips');
INSERT INTO Inventory(name, price, description) VALUES('candy bar', 17.00, 'fresh candy bar');
INSERT INTO Inventory(name, price, description) VALUES('chocolate bar', 18.00, 'fresh chocolate bar');
INSERT INTO Inventory(name, price, description) VALUES('chocolate chip cookie', 19.00, 'fresh chocolate chip cookie');

-- all of the items in Houston
INSERT INTO Houston(name, quantity) VALUES('bread', 100000);
INSERT INTO Houston(name, quantity) VALUES('cheese', 100000);
INSERT INTO Houston(name, quantity) VALUES('eggs', 100000);
INSERT INTO Houston(name, quantity) VALUES('sugar', 100000);
INSERT INTO Houston(name, quantity) VALUES('butter', 100000);
INSERT INTO Houston(name, quantity) VALUES('flour', 100000);
INSERT INTO Houston(name, quantity) VALUES('salt', 100000);
INSERT INTO Houston(name, quantity) VALUES('pepper', 100000);

-- all the items in Austin
INSERT INTO Austin(name, quantity) VALUES('bread', 10000);
INSERT INTO Austin(name, quantity) VALUES('cheese', 10000);
INSERT INTO Austin(name, quantity) VALUES('eggs', 10000);
INSERT INTO Austin(name, quantity) VALUES('sugar', 10000);
INSERT INTO Austin(name, quantity) VALUES('butter', 10000);
INSERT INTO Austin(name, quantity) VALUES('flour', 10000);
INSERT INTO Austin(name, quantity) VALUES('salt', 10000);
INSERT INTO Austin(name, quantity) VALUES('pepper', 10000);
INSERT INTO Austin(name, quantity) VALUES('juice', 10000);

-- all the items in Seattle
INSERT INTO Seattle(name, quantity) VALUES('bread', 10000);
INSERT INTO Seattle(name, quantity) VALUES('cheese', 10000);
INSERT INTO Seattle(name, quantity) VALUES('eggs', 10000);
INSERT INTO Seattle(name, quantity) VALUES('sugar', 10000);
INSERT INTO Seattle(name, quantity) VALUES('butter', 10000);
INSERT INTO Seattle(name, quantity) VALUES('flour', 10000);
INSERT INTO Seattle(name, quantity) VALUES('salt', 10000);
INSERT INTO Seattle(name, quantity) VALUES('pepper', 10000);

-- all the items in San Francisco
INSERT INTO SanFrancisco(name, quantity) VALUES('bread', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('cheese', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('eggs', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('sugar', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('butter', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('flour', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('salt', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('pepper', 10000);
INSERT INTO SanFrancisco(name, quantity) VALUES('soda', 10000);

-- all the items in Chicago
INSERT INTO Chicago(name, quantity) VALUES('bread', 10000);
INSERT INTO Chicago(name, quantity) VALUES('cheese', 10000);
INSERT INTO Chicago(name, quantity) VALUES('eggs', 10000);
INSERT INTO Chicago(name, quantity) VALUES('sugar', 10000);
INSERT INTO Chicago(name, quantity) VALUES('butter', 10000);
INSERT INTO Chicago(name, quantity) VALUES('flour', 10000);
INSERT INTO Chicago(name, quantity) VALUES('salt', 10000);
INSERT INTO Chicago(name, quantity) VALUES('pepper', 10000);
INSERT INTO Chicago(name, quantity) VALUES('soda', 10000);
INSERT INTO Chicago(name, quantity) VALUES('water', 10000);

