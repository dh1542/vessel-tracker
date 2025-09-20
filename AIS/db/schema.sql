CREATE TABLE position_reports (
                                  mmsi BIGINT PRIMARY KEY NOT NULL,
                                  ship_name VARCHAR NOT NULL,
                                  latitude DOUBLE PRECISION NOT NULL,
                                  longitude DOUBLE PRECISION NOT NULL,
                                  cog INTEGER NOT NULL,
                                  sog INTEGER NOT NULL,
                                  true_heading INTEGER NOT NULL,
                                  navigational_status INTEGER NOT NULL,
                                  position_accuracy BOOLEAN NOT NULL,
                                  communication_state BIGINT NOT NULL,
                                  rate_of_turn INTEGER NOT NULL,
                                  special_manoeuvre_indicator INTEGER NOT NULL,
                                  repeat_indicator INTEGER NOT NULL,
                                  message_id INTEGER NOT NULL,
                                  valid BOOLEAN NOT NULL,
                                  time_utc TIMESTAMP NOT NULL,
                                    destination VARCHAR NOT NULL

);

