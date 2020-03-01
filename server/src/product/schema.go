// Product Schema
//
// Establishes associative relationship between food item and ticket
//
package product

const Schema = `
CREATE TABLE IF NOT EXISTS product (
	id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	name varchar(255) NOT NULL,
	price float NOT NULL DEFAULT 0.00 CHECK (price > 0),
	description varchar
);

CREATE TABLE IF NOT EXISTS ticket_order (
	id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	course integer,
	quantity integer NOT NULL DEFAULT 1,
	ticket_id integer REFERENCES ticket,
	product_id integer REFERENCES product
);
`