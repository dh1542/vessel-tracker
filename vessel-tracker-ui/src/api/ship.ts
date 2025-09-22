import { Ship, Bounds, SQLNonNullString } from "@/types";

export const fetchShips = async (bounds: Bounds): Promise<Ship[]> => {
  const url = `/api/${bounds.minLatitude}/${bounds.maxLatitude}/${bounds.minLongitude}/${bounds.maxLongitude}`;
  console.log("Fetching ships from: " + url);
  const response = await fetch(url);
  const json = await response.json();
  console.log("Response JSON: ", json);

  const ships: Ship[] = [];

  if (json != null) {
    json.map((shipData: any) => {
      const ship: Ship = {
        mmsi: shipData.Mmsi,
        name: shipData.ShipName,
        position: {
          latitude: shipData.Latitude,
          longitude: shipData.Longitude,
        },
        heading: shipData.TrueHeading,
        destination: shipData.Destination.String,
      };
      ships.push(ship);
      console.log("Fetched ship: ", ship);
    });
  }
  console.log("Total ships fetched: ", ships.length);
  return ships;
};
