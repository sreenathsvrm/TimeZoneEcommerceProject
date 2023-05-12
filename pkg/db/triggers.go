package db

var (
	cartTotalPriceUpdate = `CREATE OR REPLACE FUNCTION update_cart_total() 
	RETURNS TRIGGER AS $$ 
	BEGIN 
	IF (TG_OP = 'DELETE') THEN 
		UPDATE carts c
			SET total_price = (
				SELECT COALESCE ( SUM (p.prize * ci.qty) , 0)::bigint
				FROM cart_items ci INNER JOIN products p ON ci.product_id = p.id 
				WHERE ci.cart_id = OLD.cart_id  
			)
		WHERE c.id = OLD.cart_id;	 
		RETURN NEW; 
	ELSE 
		UPDATE carts c 
			SET total_price = (
				SELECT COALESCE (SUM (p.prize * ci.qty),0)::bigint 
				FROM cart_items ci INNER JOIN products p ON ci.product_id = p.id 
				WHERE ci.cart_id = NEW.cart_id 
			)
		WHERE c.id = NEW.cart_id;
	END IF; 
	RETURN NEW; 
	END;
	$$ LANGUAGE plpgsql;`

	cartTotalPriceUpateTrigger = `CREATE OR REPLACE TRIGGER update_cart_total 
	AFTER INSERT OR UPDATE OR DELETE ON cart_items 
	FOR EACH ROW EXECUTE FUNCTION update_cart_total();`
)
