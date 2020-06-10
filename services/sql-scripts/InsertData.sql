INSERT INTO Test.memberships
  ( name, description)
VALUES
  ('basic', 'Rent 4 articles a week. No simultaneously. No delivery.'), 
  ('silver', 'Rent 6 articles a week. Maximal 2 simultaneously. No delivery.'), 
  ('gold', 'Rent 10 articles a week. No maximal. Delivery');

INSERT INTO Test.access_levels
  ( id, role)
VALUES
  (11, 'client'), 
  (12, 'sales'), 
  (13, 'chief_sells');

INSERT INTO Test.users( name, password, membership_id,access_levels_id) VALUES ("Matias", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", 3, 13);