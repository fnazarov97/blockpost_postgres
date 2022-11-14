BEGIN;

INSERT INTO author (id, firstname, lastname) VALUES
 ('81d41ef3-7d71-4631-87c5-c8c087f464cb','Farhod', 'Nazarov') ON CONFLICT DO NOTHING;
INSERT INTO author (id, firstname, lastname) VALUES
 ('d0d9c394-f22a-4ead-b71d-9e5a825565a6','Jone', 'Doe') ON CONFLICT DO NOTHING;

INSERT INTO article (id, title, body, author_id) VALUES
 ('e05f4dca-c92c-478c-8820-df5705190e8c','string1', 'text1', 'd0d9c394-f22a-4ead-b71d-9e5a825565a6') ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES
 ('c9ded5e8-94db-4a19-9a29-86858878f346','string2', 'text2', '81d41ef3-7d71-4631-87c5-c8c087f464cb') ON CONFLICT DO NOTHING;

 COMMIT;